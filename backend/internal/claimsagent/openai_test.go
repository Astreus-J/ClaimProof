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

func TestOpenAIClient_Complete_SendsSystemAndUserMessagesWithJSONMode(t *testing.T) {
	var capturedBody openAIRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/chat/completions", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		require.NoError(t, json.NewDecoder(r.Body).Decode(&capturedBody))

		_ = json.NewEncoder(w).Encode(openAIResponse{
			Choices: []openAIChoice{
				{Message: openAIMessage{Role: "assistant", Content: `{"suggestedPercentage": 80}`}},
			},
		})
	}))
	defer server.Close()

	client := NewOpenAIClient("test-api-key", "gpt-5-mini", nil)
	client.baseURL = server.URL

	text, err := client.Complete(context.Background(), "you are a claims adjuster", "evaluate this claim")

	require.NoError(t, err)
	assert.Equal(t, `{"suggestedPercentage": 80}`, text)

	require.Len(t, capturedBody.Messages, 2)
	assert.Equal(t, "system", capturedBody.Messages[0].Role)
	assert.Equal(t, "you are a claims adjuster", capturedBody.Messages[0].Content)
	assert.Equal(t, "user", capturedBody.Messages[1].Role)
	assert.Equal(t, "evaluate this claim", capturedBody.Messages[1].Content)
	require.NotNil(t, capturedBody.ResponseFormat)
	assert.Equal(t, "json_object", capturedBody.ResponseFormat.Type)
}

func TestOpenAIClient_Complete_ReturnsErrorOnNonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	client := NewOpenAIClient("bad-key", "gpt-5-mini", nil)
	client.baseURL = server.URL

	text, err := client.Complete(context.Background(), "system", "x")

	assert.Empty(t, text)
	assert.Error(t, err)
	assert.NotContains(t, err.Error(), "bad-key")
}

func TestOpenAIClient_Complete_ReturnsErrorWhenNoChoices(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(openAIResponse{Choices: nil})
	}))
	defer server.Close()

	client := NewOpenAIClient("test-api-key", "gpt-5-mini", nil)
	client.baseURL = server.URL

	text, err := client.Complete(context.Background(), "system", "x")

	assert.Empty(t, text)
	assert.Error(t, err)
}

func TestNewOpenAIClient_DefaultsModelAndHTTPClient(t *testing.T) {
	client := NewOpenAIClient("test-api-key", "", nil)

	assert.Equal(t, defaultOpenAIModel, client.model)
	assert.Equal(t, http.DefaultClient, client.httpClient)
}
