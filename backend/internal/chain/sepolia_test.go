package chain

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSepoliaClient_ConnectsAndBindsDeliveryTrackerMockWithoutError(t *testing.T) {
	key := testSigningKey(t)

	c, err := NewSepoliaClient(context.Background(), "https://sepolia.example.invalid", key, big.NewInt(SepoliaTestnetChainID), common.HexToAddress("0x1"))
	require.NoError(t, err)
	require.NotNil(t, c)
	defer c.Close()
}

func TestNewSepoliaClient_ReturnsErrorOnMalformedURL(t *testing.T) {
	key := testSigningKey(t)

	c, err := NewSepoliaClient(context.Background(), "not-a-url", key, big.NewInt(SepoliaTestnetChainID), common.HexToAddress("0x1"))

	assert.Nil(t, c)
	assert.Error(t, err)
}
