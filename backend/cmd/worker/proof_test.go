package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Astreus-J/ClaimProof/backend/internal/proofbuilder"
)

func TestToSubmitClaimArgs_DecodesAllFields(t *testing.T) {
	proof := &proofbuilder.Proof{
		TxBytes: "0x1234abcd",
		MerkleProof: proofbuilder.MerkleProof{
			Root: "0x" + "11" + repeat("00", 31),
			Siblings: []proofbuilder.MerkleProofEntry{
				{Hash: "0x" + "22" + repeat("00", 31), IsLeft: true},
				{Hash: "0x" + "33" + repeat("00", 31), IsLeft: false},
			},
		},
		ContinuityProof: proofbuilder.ContinuityProof{
			LowerEndpointDigest: "0x" + "44" + repeat("00", 31),
			Roots: []string{
				"0x" + "55" + repeat("00", 31),
				"0x" + "66" + repeat("00", 31),
			},
		},
	}

	args, err := toSubmitClaimArgs(proof)

	require.NoError(t, err)
	assert.Equal(t, []byte{0x12, 0x34, 0xab, 0xcd}, args.encodedTransaction)
	assert.Equal(t, byte(0x11), args.merkleRoot[0])
	require.Len(t, args.siblings, 2)
	assert.Equal(t, byte(0x22), args.siblings[0].Hash[0])
	assert.True(t, args.siblings[0].IsLeft)
	assert.Equal(t, byte(0x33), args.siblings[1].Hash[0])
	assert.False(t, args.siblings[1].IsLeft)
	assert.Equal(t, byte(0x44), args.lowerEndpointDigest[0])
	require.Len(t, args.continuityRoots, 2)
	assert.Equal(t, byte(0x55), args.continuityRoots[0][0])
	assert.Equal(t, byte(0x66), args.continuityRoots[1][0])
}

func TestToSubmitClaimArgs_ReturnsErrorOnMalformedTxBytes(t *testing.T) {
	proof := &proofbuilder.Proof{TxBytes: "not-hex"}

	args, err := toSubmitClaimArgs(proof)

	assert.Nil(t, args)
	assert.Error(t, err)
}

func TestToSubmitClaimArgs_ReturnsErrorOnWrongLengthMerkleRoot(t *testing.T) {
	proof := &proofbuilder.Proof{
		TxBytes:     "0x1234",
		MerkleProof: proofbuilder.MerkleProof{Root: "0xabcd"}, // only 2 bytes, not 32
	}

	args, err := toSubmitClaimArgs(proof)

	assert.Nil(t, args)
	assert.Error(t, err)
}

func repeat(s string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}
