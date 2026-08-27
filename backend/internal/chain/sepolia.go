package chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SepoliaTestnetChainID is Ethereum Sepolia's EVM chain ID.
const SepoliaTestnetChainID = 11155111

// SepoliaClient wraps a Sepolia RPC connection, the storefront operator's
// signing key, and the DeliveryTrackerMock contract binding for registering
// new shipments on the buyer's behalf.
type SepoliaClient struct {
	eth                 *ethclient.Client
	signingKey          *ecdsa.PrivateKey
	chainID             *big.Int
	deliveryTrackerMock *DeliveryTrackerMock
}

// NewSepoliaClient connects to the given RPC endpoint and binds
// DeliveryTrackerMock at deliveryTrackerMockAddress, signing transactions
// with signingKey for chainID.
func NewSepoliaClient(
	ctx context.Context,
	rpcURL string,
	signingKey *ecdsa.PrivateKey,
	chainID *big.Int,
	deliveryTrackerMockAddress common.Address,
) (*SepoliaClient, error) {
	eth, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("chain: dial %s: %w", rpcURL, err)
	}

	deliveryTrackerMock, err := NewDeliveryTrackerMock(deliveryTrackerMockAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("chain: bind DeliveryTrackerMock at %s: %w", deliveryTrackerMockAddress, err)
	}

	return &SepoliaClient{eth: eth, signingKey: signingKey, chainID: chainID, deliveryTrackerMock: deliveryTrackerMock}, nil
}

// Close releases the underlying RPC connection.
func (c *SepoliaClient) Close() {
	c.eth.Close()
}

// CreateShipment calls DeliveryTrackerMock.createShipment, signed by the
// storefront operator's key. The contract itself is permissionless (see
// docs/THREAT_MODEL.md, T6) — the operator key here only pays gas, it
// grants no special authority on this side.
func (c *SepoliaClient) CreateShipment(ctx context.Context, orderID *big.Int, buyer common.Address, slaSeconds *big.Int) (*types.Transaction, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(c.signingKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("chain: create transactor: %w", err)
	}
	opts.Context = ctx

	tx, err := c.deliveryTrackerMock.CreateShipment(opts, orderID, buyer, slaSeconds)
	if err != nil {
		return nil, fmt.Errorf("chain: create shipment for order %s: %w", orderID, err)
	}
	return tx, nil
}

// WaitMined blocks until tx is mined and returns its receipt.
func (c *SepoliaClient) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, c.eth, tx)
	if err != nil {
		return nil, fmt.Errorf("chain: wait mined: %w", err)
	}
	return receipt, nil
}
