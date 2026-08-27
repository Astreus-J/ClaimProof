package claimsagent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	defaultGeminiModel   = "gemini-3.6-flash"
	defaultGeminiBaseURL = "https://generativelanguage.googleapis.com"
)

type geminiPart struct {
	Text string `json:"text"`
}

type geminiContent struct {
	Parts []geminiPart `json:"parts"`
}

type geminiRequest struct {
	Contents []geminiContent `json:"contents"`
}

type geminiCandidate struct {
	Content geminiContent `json:"content"`
}

type geminiResponse struct {
	Candidates []geminiCandidate `json:"candidates"`
}

// GeminiClient implements LLMClient against Google's Gemini API.
type GeminiClient struct {
	apiKey     string
	model      string
	httpClient *http.Client
	baseURL    string
}

// NewGeminiClient creates a GeminiClient for the given API key. If model is
// empty, "gemini-2.0-flash" is used. If httpClient is nil,
// http.DefaultClient is used.
func NewGeminiClient(apiKey, model string, httpClient *http.Client) *GeminiClient {
	if model == "" {
		model = defaultGeminiModel
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &GeminiClient{apiKey: apiKey, model: model, httpClient: httpClient, baseURL: defaultGeminiBaseURL}
}

// Complete sends prompt to Gemini's generateContent endpoint and returns the
// first candidate's text.
func (g *GeminiClient) Complete(ctx context.Context, prompt string) (string, error) {
	reqBody, err := json.Marshal(geminiRequest{Contents: []geminiContent{{Parts: []geminiPart{{Text: prompt}}}}})
	if err != nil {
		return "", fmt.Errorf("claimsagent: encode Gemini request: %w", err)
	}

	url := fmt.Sprintf("%s/v1beta/models/%s:generateContent?key=%s", g.baseURL, g.model, g.apiKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("claimsagent: build Gemini request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("claimsagent: call Gemini model %s: %w", g.model, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("claimsagent: Gemini model %s returned status %d", g.model, resp.StatusCode)
	}

	var parsed geminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return "", fmt.Errorf("claimsagent: decode Gemini response: %w", err)
	}

	if len(parsed.Candidates) == 0 || len(parsed.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("claimsagent: Gemini model %s returned no candidates", g.model)
	}
	return parsed.Candidates[0].Content.Parts[0].Text, nil
}
