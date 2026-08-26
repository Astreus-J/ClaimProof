package proofbuilder

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_DefaultsToStandardHTTPClientWhenNilGiven(t *testing.T) {
	c := New("https://prover.example.invalid", nil)

	assert.Equal(t, http.DefaultClient, c.httpClient)
}

func TestNew_UsesProvidedHTTPClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	custom := server.Client()
	c := New(server.URL, custom)

	assert.Equal(t, custom, c.httpClient)
	assert.Equal(t, server.URL, c.baseURL)
}

func TestGetProof_NotYetImplemented(t *testing.T) {
	c := New("https://prover.example.invalid", nil)

	proof, err := c.GetProof(context.Background(), "0xdeadbeef")

	assert.Nil(t, proof)
	assert.Error(t, err)
}

func TestWaitUntilHeightAttested_NotYetImplemented(t *testing.T) {
	c := New("https://prover.example.invalid", nil)

	err := c.WaitUntilHeightAttested(context.Background(), 1, 100)

	assert.Error(t, err)
}
