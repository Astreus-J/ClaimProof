package main

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// config holds the API's runtime configuration, read from environment
// variables per 12-factor conventions — see backend/.env.example.
type config struct {
	listenAddr             string
	corsOrigin             string
	sepoliaRPCURL          string
	creditcoinRPCURL       string
	deliveryTrackerAddress common.Address
	claimVaultAddress      common.Address
	operatorPrivateKey     *ecdsa.PrivateKey
}

func loadConfig() (*config, error) {
	listenAddr := os.Getenv("API_LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":8080"
	}

	corsOrigin := os.Getenv("API_CORS_ORIGIN")
	if corsOrigin == "" {
		corsOrigin = "*"
	}

	sepoliaRPCURL, err := requireEnv("SEPOLIA_RPC_URL")
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

	workerPrivateKeyHex, err := requireEnv("WORKER_PRIVATE_KEY")
	if err != nil {
		return nil, err
	}
	operatorPrivateKey, err := crypto.HexToECDSA(strings.TrimPrefix(workerPrivateKeyHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("parse WORKER_PRIVATE_KEY: %w", err)
	}

	return &config{
		listenAddr:             listenAddr,
		corsOrigin:             corsOrigin,
		sepoliaRPCURL:          sepoliaRPCURL,
		creditcoinRPCURL:       creditcoinRPCURL,
		deliveryTrackerAddress: deliveryTrackerAddress,
		claimVaultAddress:      claimVaultAddress,
		operatorPrivateKey:     operatorPrivateKey,
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
