package listener

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Astreus-J/ClaimProof/backend/internal/chain"
)

func TestNew_ConnectsAndBindsFiltererWithoutError(t *testing.T) {
	// ethclient.DialContext over HTTP(S) does not open a connection eagerly
	// (JSON-RPC over HTTP is stateless), so this succeeds without a live node.
	l, err := New(context.Background(), "https://sepolia.example.invalid", common.HexToAddress("0x1"))
	require.NoError(t, err)
	require.NotNil(t, l)
	defer l.Close()
}

func TestNew_ReturnsErrorOnMalformedURL(t *testing.T) {
	l, err := New(context.Background(), "not-a-url", common.HexToAddress("0x1"))

	assert.Nil(t, l)
	assert.Error(t, err)
}

func TestToDeliveryFailedEvent_ConvertsAllFields(t *testing.T) {
	buyer := common.HexToAddress("0xD8108A4C6384866691b32c618892F0385CfC7a62")
	txHash := common.HexToHash("0xabc")
	raw := &chain.DeliveryTrackerMockDeliveryFailed{
		OrderId:   big.NewInt(42),
		Buyer:     buyer,
		Timestamp: big.NewInt(1_798_000_000),
		Raw:       types.Log{TxHash: txHash, BlockNumber: 11573061},
	}

	event := toDeliveryFailedEvent(raw)

	assert.Equal(t, uint64(42), event.OrderID)
	assert.Equal(t, buyer, event.Buyer)
	assert.Equal(t, uint64(1_798_000_000), event.Timestamp)
	assert.Equal(t, txHash, event.TxHash)
	assert.Equal(t, uint64(11573061), event.BlockNumber)
}
