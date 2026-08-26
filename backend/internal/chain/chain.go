// Package chain wraps go-ethereum's ethclient and abigen-generated contract
// bindings for both Ethereum Sepolia and Creditcoin CC3 Testnet, and submits
// signed transactions to ClaimVault on behalf of the worker's wallet.
package chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Client wraps an Ethereum-compatible RPC connection and the worker's
// signing key for submitting transactions.
type Client struct {
	eth        *ethclient.Client
	signingKey *ecdsa.PrivateKey
}

// New connects to the given RPC endpoint and returns a Client that signs
// transactions with signingKey.
func New(ctx context.Context, rpcURL string, signingKey *ecdsa.PrivateKey) (*Client, error) {
	eth, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("chain: dial %s: %w", rpcURL, err)
	}
	return &Client{eth: eth, signingKey: signingKey}, nil
}

// Close releases the underlying RPC connection.
func (c *Client) Close() {
	c.eth.Close()
}

// SubmitClaim calls ClaimVault.submitClaim with the given Attestcoin proof
// data, signed by the worker's key. Implemented in Sprint 2, once
// ClaimVault.submitClaim exists and abigen bindings are generated for it.
func (c *Client) SubmitClaim(ctx context.Context) error {
	return fmt.Errorf("chain: SubmitClaim not implemented until Sprint 2")
}
