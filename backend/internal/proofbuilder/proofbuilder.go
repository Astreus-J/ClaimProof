// Package proofbuilder talks to the Attestcoin Prover REST API
// (https://prover.cc3-testnet.creditcoin.network) to obtain Merkle
// inclusion and continuity proofs for a source-chain transaction. There is
// no official Go SDK for Attestcoin, so this package implements the HTTP
// client directly — see docs/ATTESTCOIN_INTEGRATION.md for the
// request/response shape, confirmed against the official TypeScript SDK's
// source (`@gluwa/usc-sdk`).
package proofbuilder

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// defaultPollInterval matches the official SDK's default polling interval
// for WaitUntilHeightAttested.
const defaultPollInterval = 15 * time.Second

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
// `GET /api/v1/proof-by-tx/{chainKey}/{transactionHash}` on the Prover API.
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

type attestedHeightResponse struct {
	AttestedHeight *uint64 `json:"attestedHeight"`
}

// Client fetches proofs from the Attestcoin Prover REST API over HTTP, for
// a single fixed source chain identified by chainKey.
type Client struct {
	baseURL      string
	chainKey     uint64
	httpClient   *http.Client
	pollInterval time.Duration
}

// New creates a Client for the given Prover API base URL and source chain
// key. If httpClient is nil, http.DefaultClient is used. If pollInterval is
// <= 0, a 15-second default (matching the official SDK) is used.
func New(baseURL string, chainKey uint64, httpClient *http.Client, pollInterval time.Duration) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	if pollInterval <= 0 {
		pollInterval = defaultPollInterval
	}
	return &Client{baseURL: baseURL, chainKey: chainKey, httpClient: httpClient, pollInterval: pollInterval}
}

// WaitUntilHeightAttested polls `GET /api/v1/attested-height/{chainKey}`
// until its `attestedHeight` reaches blockHeight, or ctx is done. The
// Prover API has no push/webhook mechanism, so this is a client-side poll
// loop, matching the official SDK's own implementation.
func (c *Client) WaitUntilHeightAttested(ctx context.Context, blockHeight uint64) error {
	ticker := time.NewTicker(c.pollInterval)
	defer ticker.Stop()

	for {
		height, err := c.attestedHeight(ctx)
		if err != nil {
			return fmt.Errorf("proofbuilder: check attested height for chain key %d: %w", c.chainKey, err)
		}
		if height != nil && *height >= blockHeight {
			return nil
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("proofbuilder: waiting for height %d to be attested: %w", blockHeight, ctx.Err())
		case <-ticker.C:
		}
	}
}

func (c *Client) attestedHeight(ctx context.Context) (*uint64, error) {
	url := fmt.Sprintf("%s/api/v1/attested-height/%d", c.baseURL, c.chainKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request %s: %w", url, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d from %s", resp.StatusCode, url)
	}

	var parsed attestedHeightResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, fmt.Errorf("decode response from %s: %w", url, err)
	}
	return parsed.AttestedHeight, nil
}

// GetProof fetches the inclusion and continuity proof for a transaction
// hash via `GET /api/v1/proof-by-tx/{chainKey}/{transactionHash}`. Callers
// should ensure the containing block is attested first, via
// WaitUntilHeightAttested — requesting a proof too early returns an error.
func (c *Client) GetProof(ctx context.Context, txHash string) (*Proof, error) {
	url := fmt.Sprintf("%s/api/v1/proof-by-tx/%d/%s", c.baseURL, c.chainKey, txHash)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("proofbuilder: build request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("proofbuilder: request %s: %w", url, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("proofbuilder: unexpected status %d from %s", resp.StatusCode, url)
	}

	var proof Proof
	if err := json.NewDecoder(resp.Body).Decode(&proof); err != nil {
		return nil, fmt.Errorf("proofbuilder: decode response from %s: %w", url, err)
	}
	return &proof, nil
}
