// Package reasoningreporter reports the AI claims agent's payout reasoning
// to cmd/api, so the frontend has something to fetch and display. This is
// purely informational — the actual claim/payout truth always lives
// on-chain (see internal/claimsagent's two-layer anti-hallucination guard);
// losing a report here (a stopped api, a dropped request) must never affect
// claim submission, so callers should treat Report's error as best-effort.
package reasoningreporter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client reports payout reasoning to cmd/api's POST /api/claims/{orderId}.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// New creates a Client targeting cmd/api at baseURL. If httpClient is nil,
// http.DefaultClient is used.
func New(baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{baseURL: baseURL, httpClient: httpClient}
}

type reportBody struct {
	Reasoning          string `json:"reasoning"`
	SuggestedPayoutWei string `json:"suggestedPayoutWei"`
}

// Report sends the AI's reasoning and suggested payout for orderID to
// cmd/api. It has no financial effect and no bearing on payout authorization
// — see the package doc — so a failure here should be logged and ignored by
// callers, not treated as fatal to claim processing.
func (c *Client) Report(ctx context.Context, orderID uint64, reasoning, suggestedPayoutWei string) error {
	body, err := json.Marshal(reportBody{Reasoning: reasoning, SuggestedPayoutWei: suggestedPayoutWei})
	if err != nil {
		return fmt.Errorf("reasoningreporter: encode body: %w", err)
	}

	url := fmt.Sprintf("%s/api/claims/%d", c.baseURL, orderID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("reasoningreporter: build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("reasoningreporter: request %s: %w", url, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("reasoningreporter: unexpected status %d from %s", resp.StatusCode, url)
	}
	return nil
}
