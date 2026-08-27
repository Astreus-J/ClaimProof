package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setValidEnv(t *testing.T) {
	t.Helper()
	t.Setenv("SEPOLIA_RPC_URL", "https://sepolia.example.invalid")
	t.Setenv("CREDITCOIN_TESTNET_RPC_URL", "https://rpc.cc3-testnet.creditcoin.network")
	t.Setenv("DELIVERY_TRACKER_MOCK_ADDRESS", "0x5c293e0C72E52fAca66befbEd2a65552431Ce46d")
	t.Setenv("CLAIM_VAULT_ADDRESS", "0xd6f0680F366d2de5849ab00Ff2Ca48aa1D030bCd")
	t.Setenv("WORKER_PRIVATE_KEY", "3aebd232ed3f5a75fe376864b3d45f0518a7cbb0b68c3dfac2eda0c88df6c28d")
}

func TestLoadConfig_ParsesAllFieldsFromValidEnvAndAppliesDefaults(t *testing.T) {
	setValidEnv(t)

	cfg, err := loadConfig()

	require.NoError(t, err)
	assert.Equal(t, ":8080", cfg.listenAddr, "API_LISTEN_ADDR should default when unset")
	assert.Equal(t, "*", cfg.corsOrigin, "API_CORS_ORIGIN should default when unset")
	assert.Equal(t, "https://sepolia.example.invalid", cfg.sepoliaRPCURL)
	assert.Equal(t, "https://rpc.cc3-testnet.creditcoin.network", cfg.creditcoinRPCURL)
	assert.Equal(t, "0x5c293e0C72E52fAca66befbEd2a65552431Ce46d", cfg.deliveryTrackerAddress.Hex())
	assert.Equal(t, "0xd6f0680F366d2de5849ab00Ff2Ca48aa1D030bCd", cfg.claimVaultAddress.Hex())
	assert.NotNil(t, cfg.operatorPrivateKey)
}

func TestLoadConfig_RespectsExplicitListenAddrAndCorsOrigin(t *testing.T) {
	setValidEnv(t)
	t.Setenv("API_LISTEN_ADDR", ":9090")
	t.Setenv("API_CORS_ORIGIN", "http://localhost:3000")

	cfg, err := loadConfig()

	require.NoError(t, err)
	assert.Equal(t, ":9090", cfg.listenAddr)
	assert.Equal(t, "http://localhost:3000", cfg.corsOrigin)
}

func TestLoadConfig_ReturnsErrorWhenRequiredVarMissing(t *testing.T) {
	requiredVars := []string{
		"SEPOLIA_RPC_URL",
		"CREDITCOIN_TESTNET_RPC_URL",
		"DELIVERY_TRACKER_MOCK_ADDRESS",
		"CLAIM_VAULT_ADDRESS",
		"WORKER_PRIVATE_KEY",
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

func TestLoadConfig_ReturnsErrorOnInvalidAddress(t *testing.T) {
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
