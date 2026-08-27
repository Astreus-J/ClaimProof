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

func TestGeminiClient_Complete_SendsSystemInstructionSeparateFromUserContent(t *testing.T) {
	var capturedBody geminiRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1beta/models/gemini-2.0-flash:generateContent", r.URL.Path)
		assert.Equal(t, "test-api-key", r.URL.Query().Get("key"))
		require.NoError(t, json.NewDecoder(r.Body).Decode(&capturedBody))

		_ = json.NewEncoder(w).Encode(geminiResponse{
			Candidates: []geminiCandidate{
				{Content: geminiContent{Parts: []geminiPart{{Text: `{"suggestedPercentage": 80}`}}}},
			},
		})
	}))
	defer server.Close()

	client := NewGeminiClient("test-api-key", "gemini-2.0-flash", nil)
	client.baseURL = server.URL

	text, err := client.Complete(context.Background(), "you are a claims adjuster", "evaluate this claim")

	require.NoError(t, err)
	assert.Equal(t, `{"suggestedPercentage": 80}`, text)

	require.NotNil(t, capturedBody.SystemInstruction)
	require.Len(t, capturedBody.SystemInstruction.Parts, 1)
	assert.Equal(t, "you are a claims adjuster", capturedBody.SystemInstruction.Parts[0].Text)

	require.Len(t, capturedBody.Contents, 1)
	require.Len(t, capturedBody.Contents[0].Parts, 1)
	assert.Equal(t, "evaluate this claim", capturedBody.Contents[0].Parts[0].Text)

	require.NotNil(t, capturedBody.GenerationConfig)
	assert.Equal(t, "application/json", capturedBody.GenerationConfig.ResponseMimeType)
}

func TestGeminiClient_Complete_ReturnsErrorOnNonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	client := NewGeminiClient("bad-key", "gemini-2.0-flash", nil)
	client.baseURL = server.URL

	text, err := client.Complete(context.Background(), "system", "x")

	assert.Empty(t, text)
	assert.Error(t, err)
	assert.NotContains(t, err.Error(), "bad-key", "the API key must never leak into error messages")
}

func TestGeminiClient_Complete_ReturnsErrorWhenNoCandidates(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(geminiResponse{Candidates: nil})
	}))
	defer server.Close()

	client := NewGeminiClient("test-api-key", "gemini-2.0-flash", nil)
	client.baseURL = server.URL

	text, err := client.Complete(context.Background(), "system", "x")

	assert.Empty(t, text)
	assert.Error(t, err)
}

func TestNewGeminiClient_DefaultsModelAndHTTPClient(t *testing.T) {
	client := NewGeminiClient("test-api-key", "", nil)

	assert.Equal(t, defaultGeminiModel, client.model)
	assert.Equal(t, http.DefaultClient, client.httpClient)
}
