package claimsagent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	defaultAnthropicModel   = "claude-sonnet-5"
	defaultAnthropicBaseURL = "https://api.anthropic.com"
	anthropicAPIVersion     = "2023-06-01"
	anthropicMaxTokens      = 1024
)

type anthropicMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type anthropicRequest struct {
	Model     string             `json:"model"`
	MaxTokens int                `json:"max_tokens"`
	System    string             `json:"system,omitempty"`
	Messages  []anthropicMessage `json:"messages"`
}

type anthropicContentBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type anthropicResponse struct {
	Content []anthropicContentBlock `json:"content"`
}

// AnthropicClient implements LLMClient against Anthropic's Messages API.
type AnthropicClient struct {
	apiKey     string
	model      string
	httpClient *http.Client
	baseURL    string
}

// NewAnthropicClient creates an AnthropicClient for the given API key. If
// model is empty, "claude-sonnet-5" is used. If httpClient is nil,
// http.DefaultClient is used.
func NewAnthropicClient(apiKey, model string, httpClient *http.Client) *AnthropicClient {
	if model == "" {
		model = defaultAnthropicModel
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &AnthropicClient{apiKey: apiKey, model: model, httpClient: httpClient, baseURL: defaultAnthropicBaseURL}
}

// Complete sends systemPrompt as Anthropic's top-level system field and
// userPrompt as the single user message, and returns the first text
// content block.
func (a *AnthropicClient) Complete(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	reqBody, err := json.Marshal(anthropicRequest{
		Model:     a.model,
		MaxTokens: anthropicMaxTokens,
		System:    systemPrompt,
		Messages:  []anthropicMessage{{Role: "user", Content: userPrompt}},
	})
	if err != nil {
		return "", fmt.Errorf("claimsagent: encode Anthropic request: %w", err)
	}

	url := fmt.Sprintf("%s/v1/messages", a.baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("claimsagent: build Anthropic request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", a.apiKey)
	req.Header.Set("anthropic-version", anthropicAPIVersion)

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("claimsagent: call Anthropic model %s: %w", a.model, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("claimsagent: Anthropic model %s returned status %d", a.model, resp.StatusCode)
	}

	var parsed anthropicResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return "", fmt.Errorf("claimsagent: decode Anthropic response: %w", err)
	}

	if len(parsed.Content) == 0 {
		return "", fmt.Errorf("claimsagent: Anthropic model %s returned no content blocks", a.model)
	}
	return parsed.Content[0].Text, nil
}
