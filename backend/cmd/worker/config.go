package main

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// config holds the worker's runtime configuration, read from environment
// variables per 12-factor conventions — see backend/.env.example.
type config struct {
	sepoliaWSSURL          string
	creditcoinRPCURL       string
	deliveryTrackerAddress common.Address
	claimVaultAddress      common.Address
	attestcoinProverURL    string
	workerPrivateKey       *ecdsa.PrivateKey
	llmProvider            string
	llmAPIKey              string
	llmModel               string
	policyCapWei           *big.Int
}

func loadConfig() (*config, error) {
	sepoliaWSSURL, err := requireEnv("SEPOLIA_WSS_URL")
	if err != nil {
		return nil, err
	}

	creditcoinRPCURL, err := requireEnv("CREDITCOIN_TESTNET_RPC_URL")
	if err != nil {
		return nil, err
	}

	deliveryTrackerAddress, err := requireEnvAddress("DELIVERY_TRACKER_MOCK_ADDRESS")
	if err != nil {
		return nil, err
	}

	claimVaultAddress, err := requireEnvAddress("CLAIM_VAULT_ADDRESS")
	if err != nil {
		return nil, err
	}

	attestcoinProverURL, err := requireEnv("ATTESTCOIN_PROVER_URL")
	if err != nil {
		return nil, err
	}

	workerPrivateKeyHex, err := requireEnv("WORKER_PRIVATE_KEY")
	if err != nil {
		return nil, err
	}
	workerPrivateKey, err := crypto.HexToECDSA(strings.TrimPrefix(workerPrivateKeyHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("parse WORKER_PRIVATE_KEY: %w", err)
	}

	llmProvider, err := requireEnv("LLM_PROVIDER")
	if err != nil {
		return nil, err
	}

	llmAPIKey, err := requireEnv("LLM_API_KEY")
	if err != nil {
		return nil, err
	}

	// LLM_MODEL is optional — an empty value falls back to that provider's
	// own default model.
	llmModel := os.Getenv("LLM_MODEL")

	policyCapWeiStr, err := requireEnv("POLICY_CAP_WEI")
	if err != nil {
		return nil, err
	}
	policyCapWei, ok := new(big.Int).SetString(policyCapWeiStr, 10)
	if !ok {
		return nil, fmt.Errorf("POLICY_CAP_WEI %q is not a valid base-10 integer", policyCapWeiStr)
	}

	return &config{
		sepoliaWSSURL:          sepoliaWSSURL,
		creditcoinRPCURL:       creditcoinRPCURL,
		deliveryTrackerAddress: deliveryTrackerAddress,
		claimVaultAddress:      claimVaultAddress,
		attestcoinProverURL:    attestcoinProverURL,
		workerPrivateKey:       workerPrivateKey,
		llmProvider:            llmProvider,
		llmAPIKey:              llmAPIKey,
		llmModel:               llmModel,
		policyCapWei:           policyCapWei,
	}, nil
}

func requireEnv(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return "", fmt.Errorf("%s is required", key)
	}
	return v, nil
}

func requireEnvAddress(key string) (common.Address, error) {
	v, err := requireEnv(key)
	if err != nil {
		return common.Address{}, err
	}
	if !common.IsHexAddress(v) {
		return common.Address{}, fmt.Errorf("%s %q is not a valid address", key, v)
	}
	return common.HexToAddress(v), nil
}
