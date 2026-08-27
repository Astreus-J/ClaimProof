package proofbuilder

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testChainKey = 1

func TestNew_DefaultsToStandardHTTPClientAndPollIntervalWhenZeroValuesGiven(t *testing.T) {
	c := New("https://prover.example.invalid", testChainKey, nil, 0)

	assert.Equal(t, http.DefaultClient, c.httpClient)
	assert.Equal(t, defaultPollInterval, c.pollInterval)
}

func TestGetProof_DecodesFullResponseShape(t *testing.T) {
	generatedAt := time.Date(2026, 8, 26, 20, 8, 29, 0, time.UTC)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/proof-by-tx/1/0xdeadbeef", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Proof{
			ChainKey:     1,
			HeaderNumber: 11573061,
			TxIndex:      225,
			TxHash:       "0xdeadbeef",
			TxBytes:      "0x1234",
			ContinuityProof: ContinuityProof{
				LowerEndpointDigest: "0xaaa",
				Roots:               []string{"0xbbb", "0xccc"},
			},
			MerkleProof: MerkleProof{
				Root: "0xbbb",
				Siblings: []MerkleProofEntry{
					{Hash: "0xddd", IsLeft: true},
				},
			},
			Cached:      true,
			GeneratedAt: generatedAt,
		})
	}))
	defer server.Close()

	c := New(server.URL, testChainKey, nil, 0)
	proof, err := c.GetProof(context.Background(), "0xdeadbeef")

	require.NoError(t, err)
	require.NotNil(t, proof)
	assert.Equal(t, uint64(11573061), proof.HeaderNumber)
	assert.Equal(t, uint64(225), proof.TxIndex)
	assert.Equal(t, "0xaaa", proof.ContinuityProof.LowerEndpointDigest)
	assert.Equal(t, []string{"0xbbb", "0xccc"}, proof.ContinuityProof.Roots)
	assert.Equal(t, "0xbbb", proof.MerkleProof.Root)
	assert.Len(t, proof.MerkleProof.Siblings, 1)
	assert.True(t, proof.MerkleProof.Siblings[0].IsLeft)
	assert.True(t, proof.Cached)
	assert.True(t, generatedAt.Equal(proof.GeneratedAt))
}

func TestGetProof_ReturnsErrorOnNonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	c := New(server.URL, testChainKey, nil, 0)
	proof, err := c.GetProof(context.Background(), "0xdeadbeef")

	assert.Nil(t, proof)
	assert.Error(t, err)
}

func TestGetProof_ReturnsErrorOnMalformedJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("not json"))
	}))
	defer server.Close()

	c := New(server.URL, testChainKey, nil, 0)
	proof, err := c.GetProof(context.Background(), "0xdeadbeef")

	assert.Nil(t, proof)
	assert.Error(t, err)
}

func TestWaitUntilHeightAttested_ReturnsOnceTargetHeightReached(t *testing.T) {
	heights := []uint64{10, 15, 20}
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/attested-height/1", r.URL.Path)
		h := heights[callCount]
		if callCount < len(heights)-1 {
			callCount++
		}
		_ = json.NewEncoder(w).Encode(map[string]uint64{"attestedHeight": h})
	}))
	defer server.Close()

	c := New(server.URL, testChainKey, nil, 5*time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := c.WaitUntilHeightAttested(ctx, 20)

	require.NoError(t, err)
	assert.Equal(t, len(heights)-1, callCount)
}

func TestWaitUntilHeightAttested_ReturnsContextErrorWhenCanceled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]uint64{"attestedHeight": 1})
	}))
	defer server.Close()

	c := New(server.URL, testChainKey, nil, 5*time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	err := c.WaitUntilHeightAttested(ctx, 999999)

	require.Error(t, err)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestWaitUntilHeightAttested_ReturnsErrorOnServerFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	c := New(server.URL, testChainKey, nil, 5*time.Millisecond)

	err := c.WaitUntilHeightAttested(context.Background(), 1)

	assert.Error(t, err)
}
