// Package proofbuilder talks to the Attestcoin Prover REST API
// (https://prover.cc3-testnet.creditcoin.network) to obtain Merkle
// inclusion and continuity proofs for a source-chain transaction. There is
// no official Go SDK for Attestcoin, so this package implements the HTTP
// client directly — see docs/ATTESTCOIN_INTEGRATION.md for the
// request/response shape as it is discovered.
package proofbuilder

import (
	"context"
	"fmt"
	"net/http"
)

// Proof holds the data `ClaimVault.submitClaim` needs to call the Attestcoin
// precompile's verifyAndEmit(), as returned by the Prover API for a single
// attested transaction.
type Proof struct {
	ChainKey        uint64
	BlockHeight     uint64
	EncodedTx       []byte
	MerkleProof     []byte
	ContinuityProof []byte
}

// Client fetches proofs from the Attestcoin Prover REST API over HTTP.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// New creates a Client for the given Prover API base URL. If httpClient is
// nil, http.DefaultClient is used.
func New(baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{baseURL: baseURL, httpClient: httpClient}
}

// WaitUntilHeightAttested blocks until the given source-chain block height
// has been attested by the Attestcoin protocol. Implemented in Sprint 3.
func (c *Client) WaitUntilHeightAttested(ctx context.Context, chainKey, blockHeight uint64) error {
	return fmt.Errorf("proofbuilder: WaitUntilHeightAttested not implemented until Sprint 3")
}

// GetProof fetches the inclusion and continuity proof for a transaction
// hash. Implemented in Sprint 3.
func (c *Client) GetProof(ctx context.Context, txHash string) (*Proof, error) {
	return nil, fmt.Errorf("proofbuilder: GetProof not implemented until Sprint 3")
}
