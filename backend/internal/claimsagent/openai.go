package claimsagent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	defaultOpenAIModel   = "gpt-5-mini"
	defaultOpenAIBaseURL = "https://api.openai.com"
)

type openAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIResponseFormat struct {
	Type string `json:"type"`
}

type openAIRequest struct {
	Model          string                `json:"model"`
	Messages       []openAIMessage       `json:"messages"`
	ResponseFormat *openAIResponseFormat `json:"response_format,omitempty"`
}

type openAIChoice struct {
	Message openAIMessage `json:"message"`
}

type openAIResponse struct {
	Choices []openAIChoice `json:"choices"`
}

// OpenAIClient implements LLMClient against OpenAI's Chat Completions API.
type OpenAIClient struct {
	apiKey     string
	model      string
	httpClient *http.Client
	baseURL    string
}

// NewOpenAIClient creates an OpenAIClient for the given API key. If model is
// empty, "gpt-5-mini" is used. If httpClient is nil, http.DefaultClient is
// used.
func NewOpenAIClient(apiKey, model string, httpClient *http.Client) *OpenAIClient {
	if model == "" {
		model = defaultOpenAIModel
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &OpenAIClient{apiKey: apiKey, model: model, httpClient: httpClient, baseURL: defaultOpenAIBaseURL}
}

// Complete sends systemPrompt and userPrompt as the system and user chat
// messages, requesting JSON-object output, and returns the first choice's
// message content.
func (o *OpenAIClient) Complete(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	reqBody, err := json.Marshal(openAIRequest{
		Model: o.model,
		Messages: []openAIMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
		ResponseFormat: &openAIResponseFormat{Type: "json_object"},
	})
	if err != nil {
		return "", fmt.Errorf("claimsagent: encode OpenAI request: %w", err)
	}

	url := fmt.Sprintf("%s/v1/chat/completions", o.baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("claimsagent: build OpenAI request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.apiKey)

	resp, err := o.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("claimsagent: call OpenAI model %s: %w", o.model, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("claimsagent: OpenAI model %s returned status %d", o.model, resp.StatusCode)
	}

	var parsed openAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return "", fmt.Errorf("claimsagent: decode OpenAI response: %w", err)
	}

	if len(parsed.Choices) == 0 {
		return "", fmt.Errorf("claimsagent: OpenAI model %s returned no choices", o.model)
	}
	return parsed.Choices[0].Message.Content, nil
}
