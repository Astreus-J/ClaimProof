// Package chain wraps go-ethereum's ethclient and abigen-generated contract
// bindings for both Ethereum Sepolia and Creditcoin CC3 Testnet, and submits
// signed transactions to ClaimVault on behalf of the worker's wallet.
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

// CreditcoinTestnetChainID is Creditcoin CC3 Testnet's EVM chain ID.
const CreditcoinTestnetChainID = 102031

// Client wraps a Creditcoin RPC connection, the worker's signing key, and
// the ClaimVault contract binding for submitting claims.
type Client struct {
	eth        *ethclient.Client
	signingKey *ecdsa.PrivateKey
	chainID    *big.Int
	claimVault *ClaimVault
}

// New connects to the given RPC endpoint and binds ClaimVault at
// claimVaultAddress, signing transactions with signingKey for chainID.
func New(
	ctx context.Context,
	rpcURL string,
	signingKey *ecdsa.PrivateKey,
	chainID *big.Int,
	claimVaultAddress common.Address,
) (*Client, error) {
	eth, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("chain: dial %s: %w", rpcURL, err)
	}

	claimVault, err := NewClaimVault(claimVaultAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("chain: bind ClaimVault at %s: %w", claimVaultAddress, err)
	}

	return &Client{eth: eth, signingKey: signingKey, chainID: chainID, claimVault: claimVault}, nil
}

// Close releases the underlying RPC connection.
func (c *Client) Close() {
	c.eth.Close()
}

// SubmitClaim calls ClaimVault.submitClaim with the given Attestcoin proof
// and the AI claims agent's suggested payout, signed by the worker's key.
func (c *Client) SubmitClaim(
	ctx context.Context,
	chainKey uint64,
	blockHeight uint64,
	encodedTransaction []byte,
	merkleRoot [32]byte,
	siblings []INativeQueryVerifierMerkleProofEntry,
	lowerEndpointDigest [32]byte,
	continuityRoots [][32]byte,
	suggestedPayout *big.Int,
) (*types.Transaction, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(c.signingKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("chain: create transactor: %w", err)
	}
	opts.Context = ctx

	tx, err := c.claimVault.SubmitClaim(
		opts, chainKey, blockHeight, encodedTransaction, merkleRoot, siblings, lowerEndpointDigest, continuityRoots, suggestedPayout,
	)
	if err != nil {
		return nil, fmt.Errorf("chain: submit claim: %w", err)
	}
	return tx, nil
}

// Order is a registered protection order read from ClaimVault.
type Order struct {
	Buyer            common.Address
	ProtectionAmount *big.Int
	Claimed          bool
}

// GetOrder reads a registered order from ClaimVault.
func (c *Client) GetOrder(ctx context.Context, orderID *big.Int) (Order, error) {
	result, err := c.claimVault.Orders(&bind.CallOpts{Context: ctx}, orderID)
	if err != nil {
		return Order{}, fmt.Errorf("chain: read order %s: %w", orderID, err)
	}
	return Order{Buyer: result.Buyer, ProtectionAmount: result.ProtectionAmount, Claimed: result.Claimed}, nil
}

// WaitMined blocks until tx is mined and returns its receipt.
func (c *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, c.eth, tx)
	if err != nil {
		return nil, fmt.Errorf("chain: wait mined: %w", err)
	}
	return receipt, nil
}
