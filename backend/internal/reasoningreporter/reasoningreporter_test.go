package reasoningreporter

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReport_SendsOrderIDReasoningAndPayoutToTheExpectedPath(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody map[string]string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	c := New(server.URL, nil)
	err := c.Report(context.Background(), 42, "full refund, no severity detail given", "1000000000000000000")

	require.NoError(t, err)
	assert.Equal(t, http.MethodPost, gotMethod)
	assert.Equal(t, "/api/claims/42", gotPath)
	assert.Equal(t, "full refund, no severity detail given", gotBody["reasoning"])
	assert.Equal(t, "1000000000000000000", gotBody["suggestedPayoutWei"])
}

func TestReport_ReturnsErrorOnServerFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	c := New(server.URL, nil)
	err := c.Report(context.Background(), 42, "x", "1000")

	assert.Error(t, err)
}

func TestReport_ReturnsErrorWhenContextCanceled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	c := New(server.URL, nil)
	err := c.Report(ctx, 42, "x", "1000")

	assert.Error(t, err)
}
