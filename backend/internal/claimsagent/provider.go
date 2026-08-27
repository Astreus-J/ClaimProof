package claimsagent

import (
	"fmt"
	"net/http"
	"strings"
)

// NewLLMClient constructs the LLMClient for the named provider
// ("gemini", "openai", or "anthropic", case-insensitive). model may be
// empty to use that provider's default. This is the project's single
// swap point for changing AI providers — internal/claimsagent's own logic
// (prompting, policy-cap enforcement, hallucination checks) never needs to
// change when the provider does.
func NewLLMClient(provider, apiKey, model string, httpClient *http.Client) (LLMClient, error) {
	switch strings.ToLower(provider) {
	case "gemini":
		return NewGeminiClient(apiKey, model, httpClient), nil
	case "openai":
		return NewOpenAIClient(apiKey, model, httpClient), nil
	case "anthropic":
		return NewAnthropicClient(apiKey, model, httpClient), nil
	default:
		return nil, fmt.Errorf("claimsagent: unknown LLM provider %q (want one of: gemini, openai, anthropic)", provider)
	}
}
