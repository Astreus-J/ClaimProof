package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setValidEnv(t *testing.T) {
	t.Helper()
	t.Setenv("SEPOLIA_WSS_URL", "wss://sepolia.example.invalid")
	t.Setenv("CREDITCOIN_TESTNET_RPC_URL", "https://rpc.cc3-testnet.creditcoin.network")
	t.Setenv("DELIVERY_TRACKER_MOCK_ADDRESS", "0x5c293e0C72E52fAca66befbEd2a65552431Ce46d")
	t.Setenv("CLAIM_VAULT_ADDRESS", "0xd6f0680F366d2de5849ab00Ff2Ca48aa1D030bCd")
	t.Setenv("ATTESTCOIN_PROVER_URL", "https://prover.cc3-testnet.creditcoin.network")
	t.Setenv("WORKER_PRIVATE_KEY", "3aebd232ed3f5a75fe376864b3d45f0518a7cbb0b68c3dfac2eda0c88df6c28d")
	t.Setenv("LLM_PROVIDER", "gemini")
	t.Setenv("LLM_API_KEY", "test-key")
	t.Setenv("POLICY_CAP_WEI", "1000000000000000000")
}

func TestLoadConfig_ParsesAllFieldsFromValidEnv(t *testing.T) {
	setValidEnv(t)

	cfg, err := loadConfig()

	require.NoError(t, err)
	assert.Equal(t, "wss://sepolia.example.invalid", cfg.sepoliaWSSURL)
	assert.Equal(t, "https://rpc.cc3-testnet.creditcoin.network", cfg.creditcoinRPCURL)
	assert.Equal(t, "0x5c293e0C72E52fAca66befbEd2a65552431Ce46d", cfg.deliveryTrackerAddress.Hex())
	assert.Equal(t, "0xd6f0680F366d2de5849ab00Ff2Ca48aa1D030bCd", cfg.claimVaultAddress.Hex())
	assert.Equal(t, "https://prover.cc3-testnet.creditcoin.network", cfg.attestcoinProverURL)
	assert.NotNil(t, cfg.workerPrivateKey)
	assert.Equal(t, "gemini", cfg.llmProvider)
	assert.Equal(t, "test-key", cfg.llmAPIKey)
	assert.Empty(t, cfg.llmModel)
	assert.Equal(t, "1000000000000000000", cfg.policyCapWei.String())
}

func TestLoadConfig_AcceptsPrivateKeyWith0xPrefix(t *testing.T) {
	setValidEnv(t)
	t.Setenv("WORKER_PRIVATE_KEY", "0x3aebd232ed3f5a75fe376864b3d45f0518a7cbb0b68c3dfac2eda0c88df6c28d")

	cfg, err := loadConfig()

	require.NoError(t, err)
	assert.NotNil(t, cfg.workerPrivateKey)
}

func TestLoadConfig_LLMModelIsOptionalAndOverridable(t *testing.T) {
	setValidEnv(t)
	t.Setenv("LLM_MODEL", "gemini-3.6-flash")

	cfg, err := loadConfig()

	require.NoError(t, err)
	assert.Equal(t, "gemini-3.6-flash", cfg.llmModel)
}

func TestLoadConfig_ReturnsErrorWhenRequiredVarMissing(t *testing.T) {
	requiredVars := []string{
		"SEPOLIA_WSS_URL",
		"CREDITCOIN_TESTNET_RPC_URL",
		"DELIVERY_TRACKER_MOCK_ADDRESS",
		"CLAIM_VAULT_ADDRESS",
		"ATTESTCOIN_PROVER_URL",
		"WORKER_PRIVATE_KEY",
		"LLM_PROVIDER",
		"LLM_API_KEY",
		"POLICY_CAP_WEI",
	}

	for _, missing := range requiredVars {
		t.Run(missing, func(t *testing.T) {
			setValidEnv(t)
			t.Setenv(missing, "")

			cfg, err := loadConfig()

			assert.Nil(t, cfg)
			assert.Error(t, err)
		})
	}
}

func TestLoadConfig_ReturnsErrorOnInvalidContractAddress(t *testing.T) {
	setValidEnv(t)
	t.Setenv("CLAIM_VAULT_ADDRESS", "not-an-address")

	cfg, err := loadConfig()

	assert.Nil(t, cfg)
	assert.Error(t, err)
}

func TestLoadConfig_ReturnsErrorOnInvalidPrivateKey(t *testing.T) {
	setValidEnv(t)
	t.Setenv("WORKER_PRIVATE_KEY", "not-a-key")

	cfg, err := loadConfig()

	assert.Nil(t, cfg)
	assert.Error(t, err)
}

func TestLoadConfig_ReturnsErrorOnInvalidPolicyCap(t *testing.T) {
	setValidEnv(t)
	t.Setenv("POLICY_CAP_WEI", "not-a-number")

	cfg, err := loadConfig()

	assert.Nil(t, cfg)
	assert.Error(t, err)
}
