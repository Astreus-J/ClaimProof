package chain

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew_ConnectsOverHTTPWithoutError(t *testing.T) {
	// ethclient.DialContext over HTTP(S) does not open a connection eagerly
	// (JSON-RPC over HTTP is stateless), so this succeeds without a live node.
	c, err := New(context.Background(), "https://sepolia.example.invalid", nil)
	require.NoError(t, err)
	require.NotNil(t, c)
	defer c.Close()
}

func TestNew_ReturnsErrorOnMalformedURL(t *testing.T) {
	c, err := New(context.Background(), "not-a-url", nil)

	assert.Nil(t, c)
	assert.Error(t, err)
}

func TestSubmitClaim_NotYetImplemented(t *testing.T) {
	c, err := New(context.Background(), "https://sepolia.example.invalid", nil)
	require.NoError(t, err)
	defer c.Close()

	err = c.SubmitClaim(context.Background())

	assert.Error(t, err)
}
