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
	"time"
)

// MerkleProofEntry is one sibling hash in a Merkle inclusion proof.
type MerkleProofEntry struct {
	Hash   string `json:"hash"`
	IsLeft bool   `json:"isLeft"`
}

// MerkleProof is the Merkle inclusion proof for one transaction within its
// block.
type MerkleProof struct {
	Root     string             `json:"root"`
	Siblings []MerkleProofEntry `json:"siblings"`
}

// ContinuityProof chains block digests so a single attested checkpoint
// cryptographically anchors a whole range of blocks.
type ContinuityProof struct {
	LowerEndpointDigest string   `json:"lowerEndpointDigest"`
	Roots               []string `json:"roots"`
}

// Proof holds the data `ClaimVault.submitClaim` needs to call the Attestcoin
// precompile's verifyAndEmit(). This is the exact response shape of
// `GET /api/v1/proof-by-tx/{chainKey}/{transactionHash}` on the Prover API,
// confirmed against the official TypeScript SDK's source — see
// docs/ATTESTCOIN_INTEGRATION.md.
type Proof struct {
	ChainKey        uint64          `json:"chainKey"`
	HeaderNumber    uint64          `json:"headerNumber"`
	TxIndex         uint64          `json:"txIndex"`
	TxHash          string          `json:"txHash"`
	TxBytes         string          `json:"txBytes"`
	ContinuityProof ContinuityProof `json:"continuityProof"`
	MerkleProof     MerkleProof     `json:"merkleProof"`
	Cached          bool            `json:"cached"`
	GeneratedAt     time.Time       `json:"generatedAt"`
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

// WaitUntilHeightAttested polls `GET /api/v1/attested-height/{chainKey}`
// until its `attestedHeight` reaches blockHeight — the Prover API has no
// push/webhook mechanism, so this is a client-side poll loop, matching the
// official SDK's own implementation. Implemented in Sprint 3.
func (c *Client) WaitUntilHeightAttested(ctx context.Context, chainKey, blockHeight uint64) error {
	return fmt.Errorf("proofbuilder: WaitUntilHeightAttested not implemented until Sprint 3")
}

// GetProof fetches the inclusion and continuity proof for a transaction
// hash via `GET /api/v1/proof-by-tx/{chainKey}/{transactionHash}`.
// Implemented in Sprint 3.
func (c *Client) GetProof(ctx context.Context, txHash string) (*Proof, error) {
	return nil, fmt.Errorf("proofbuilder: GetProof not implemented until Sprint 3")
}
