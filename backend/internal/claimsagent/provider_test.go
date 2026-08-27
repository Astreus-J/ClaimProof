package claimsagent

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLLMClient_Gemini(t *testing.T) {
	client, err := NewLLMClient("gemini", "key", "", nil)

	require.NoError(t, err)
	_, ok := client.(*GeminiClient)
	assert.True(t, ok)
}

func TestNewLLMClient_OpenAI(t *testing.T) {
	client, err := NewLLMClient("openai", "key", "", nil)

	require.NoError(t, err)
	_, ok := client.(*OpenAIClient)
	assert.True(t, ok)
}

func TestNewLLMClient_Anthropic(t *testing.T) {
	client, err := NewLLMClient("anthropic", "key", "", nil)

	require.NoError(t, err)
	_, ok := client.(*AnthropicClient)
	assert.True(t, ok)
}

func TestNewLLMClient_IsCaseInsensitive(t *testing.T) {
	client, err := NewLLMClient("GEMINI", "key", "", nil)

	require.NoError(t, err)
	_, ok := client.(*GeminiClient)
	assert.True(t, ok)
}

func TestNewLLMClient_ReturnsErrorOnUnknownProvider(t *testing.T) {
	client, err := NewLLMClient("some-unknown-provider", "key", "", nil)

	assert.Nil(t, client)
	assert.Error(t, err)
}
