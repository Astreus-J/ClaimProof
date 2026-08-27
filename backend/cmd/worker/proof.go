package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/Astreus-J/ClaimProof/backend/internal/chain"
	"github.com/Astreus-J/ClaimProof/backend/internal/proofbuilder"
)

// submitClaimArgs holds a Prover API proof decoded into the argument shapes
// ClaimVault.submitClaim expects.
type submitClaimArgs struct {
	encodedTransaction  []byte
	merkleRoot          [32]byte
	siblings            []chain.INativeQueryVerifierMerkleProofEntry
	lowerEndpointDigest [32]byte
	continuityRoots     [][32]byte
}

// toSubmitClaimArgs converts a Proof from the Prover API's hex-string
// fields into the fixed-size byte arrays ClaimVault.submitClaim requires.
func toSubmitClaimArgs(p *proofbuilder.Proof) (*submitClaimArgs, error) {
	encodedTransaction, err := decodeHexBytes(p.TxBytes)
	if err != nil {
		return nil, fmt.Errorf("decode txBytes: %w", err)
	}

	merkleRoot, err := decodeHex32(p.MerkleProof.Root)
	if err != nil {
		return nil, fmt.Errorf("decode merkle root: %w", err)
	}

	siblings := make([]chain.INativeQueryVerifierMerkleProofEntry, len(p.MerkleProof.Siblings))
	for i, s := range p.MerkleProof.Siblings {
		hash, err := decodeHex32(s.Hash)
		if err != nil {
			return nil, fmt.Errorf("decode sibling %d hash: %w", i, err)
		}
		siblings[i] = chain.INativeQueryVerifierMerkleProofEntry{Hash: hash, IsLeft: s.IsLeft}
	}

	lowerEndpointDigest, err := decodeHex32(p.ContinuityProof.LowerEndpointDigest)
	if err != nil {
		return nil, fmt.Errorf("decode lower endpoint digest: %w", err)
	}

	continuityRoots := make([][32]byte, len(p.ContinuityProof.Roots))
	for i, r := range p.ContinuityProof.Roots {
		root, err := decodeHex32(r)
		if err != nil {
			return nil, fmt.Errorf("decode continuity root %d: %w", i, err)
		}
		continuityRoots[i] = root
	}

	return &submitClaimArgs{
		encodedTransaction:  encodedTransaction,
		merkleRoot:          merkleRoot,
		siblings:            siblings,
		lowerEndpointDigest: lowerEndpointDigest,
		continuityRoots:     continuityRoots,
	}, nil
}

func decodeHexBytes(s string) ([]byte, error) {
	return hex.DecodeString(strings.TrimPrefix(s, "0x"))
}

func decodeHex32(s string) ([32]byte, error) {
	var out [32]byte
	b, err := decodeHexBytes(s)
	if err != nil {
		return out, err
	}
	if len(b) != 32 {
		return out, fmt.Errorf("expected 32 bytes, got %d", len(b))
	}
	copy(out[:], b)
	return out, nil
}
