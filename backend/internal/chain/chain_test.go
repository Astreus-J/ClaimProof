package chain

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testSigningKey(t *testing.T) *ecdsa.PrivateKey {
	t.Helper()
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	return key
}

func TestNew_ConnectsAndBindsClaimVaultWithoutError(t *testing.T) {
	// ethclient.DialContext over HTTP(S) does not open a connection eagerly
	// (JSON-RPC over HTTP is stateless), so this succeeds without a live node.
	key := testSigningKey(t)

	c, err := New(context.Background(), "https://creditcoin.example.invalid", key, big.NewInt(102031), common.HexToAddress("0x1"))
	require.NoError(t, err)
	require.NotNil(t, c)
	defer c.Close()
}

func TestNew_ReturnsErrorOnMalformedURL(t *testing.T) {
	key := testSigningKey(t)

	c, err := New(context.Background(), "not-a-url", key, big.NewInt(102031), common.HexToAddress("0x1"))

	assert.Nil(t, c)
	assert.Error(t, err)
}
