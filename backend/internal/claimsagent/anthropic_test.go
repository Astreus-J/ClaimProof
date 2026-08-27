package claimsagent

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnthropicClient_Complete_SendsSystemFieldSeparateFromMessages(t *testing.T) {
	var capturedBody anthropicRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/messages", r.URL.Path)
		assert.Equal(t, "test-api-key", r.Header.Get("x-api-key"))
		assert.NotEmpty(t, r.Header.Get("anthropic-version"))
		require.NoError(t, json.NewDecoder(r.Body).Decode(&capturedBody))

		_ = json.NewEncoder(w).Encode(anthropicResponse{
			Content: []anthropicContentBlock{{Type: "text", Text: `{"suggestedPercentage": 80}`}},
		})
	}))
	defer server.Close()

	client := NewAnthropicClient("test-api-key", "claude-sonnet-5", nil)
	client.baseURL = server.URL

	text, err := client.Complete(context.Background(), "you are a claims adjuster", "evaluate this claim")

	require.NoError(t, err)
	assert.Equal(t, `{"suggestedPercentage": 80}`, text)

	assert.Equal(t, "you are a claims adjuster", capturedBody.System)
	require.Len(t, capturedBody.Messages, 1)
	assert.Equal(t, "user", capturedBody.Messages[0].Role)
	assert.Equal(t, "evaluate this claim", capturedBody.Messages[0].Content)
}

func TestAnthropicClient_Complete_ReturnsErrorOnNonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	client := NewAnthropicClient("bad-key", "claude-sonnet-5", nil)
	client.baseURL = server.URL

	text, err := client.Complete(context.Background(), "system", "x")

	assert.Empty(t, text)
	assert.Error(t, err)
	assert.NotContains(t, err.Error(), "bad-key")
}

func TestAnthropicClient_Complete_ReturnsErrorWhenNoContentBlocks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(anthropicResponse{Content: nil})
	}))
	defer server.Close()

	client := NewAnthropicClient("test-api-key", "claude-sonnet-5", nil)
	client.baseURL = server.URL

	text, err := client.Complete(context.Background(), "system", "x")

	assert.Empty(t, text)
	assert.Error(t, err)
}

func TestNewAnthropicClient_DefaultsModelAndHTTPClient(t *testing.T) {
	client := NewAnthropicClient("test-api-key", "", nil)

	assert.Equal(t, defaultAnthropicModel, client.model)
	assert.Equal(t, http.DefaultClient, client.httpClient)
}
