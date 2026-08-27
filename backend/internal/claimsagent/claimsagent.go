// Package claimsagent asks an LLM to suggest a payout amount for a claim,
// bounded by an on-chain policy cap. The suggestion is advisory only — per
// CLAUDE.md's core invariant, the AI never authorizes payment alone; only
// ClaimVault.submitClaim's Attestcoin reverification does. The policy cap
// enforced here is an off-chain sanity check independent of (and in
// addition to) ClaimVault's own on-chain payoutCap.
package claimsagent

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

// ClaimContext is the order/claim data given to the LLM to reason about.
type ClaimContext struct {
	OrderID            uint64
	ProtectionAmount   *big.Int
	FailureDescription string
}

// LLMClient is the interface claimsagent depends on, so the LLM call can be
// mocked in tests without a real API key.
type LLMClient interface {
	Complete(ctx context.Context, prompt string) (string, error)
}

// Agent suggests a payout amount for a claim, never exceeding policyCap.
type Agent struct {
	llm       LLMClient
	policyCap *big.Int
}

// New creates an Agent backed by the given LLM client and policy cap (wei).
func New(llm LLMClient, policyCap *big.Int) *Agent {
	return &Agent{llm: llm, policyCap: policyCap}
}

type llmSuggestion struct {
	SuggestedPercentage int64  `json:"suggestedPercentage"`
	Reasoning           string `json:"reasoning"`
}

// SuggestPayout asks the LLM what fraction of the order's protectionAmount
// to refund, and returns that fraction of protectionAmount, clamped to
// policyCap. If the LLM's response can't be parsed, it falls back to a full
// refund suggestion — an off-chain parsing hiccup must never itself deny a
// claim that Attestcoin has already verified actually happened; the cap
// still bounds the result either way.
func (a *Agent) SuggestPayout(ctx context.Context, claim ClaimContext) (*big.Int, error) {
	response, err := a.llm.Complete(ctx, buildPrompt(claim))
	if err != nil {
		return nil, fmt.Errorf("claimsagent: LLM completion: %w", err)
	}

	percentage := parseSuggestedPercentage(response)

	suggested := new(big.Int).Mul(claim.ProtectionAmount, big.NewInt(percentage))
	suggested.Div(suggested, big.NewInt(100))

	if suggested.Cmp(a.policyCap) > 0 {
		return new(big.Int).Set(a.policyCap), nil
	}
	return suggested, nil
}

func buildPrompt(claim ClaimContext) string {
	return fmt.Sprintf(
		`You are a claims adjuster for a parametric delivery-protection product. A delivery failure has already been cryptographically verified on-chain — your only job is to recommend what fraction of the buyer's protection amount to refund, based on the severity described below.

Order ID: %d
Failure description: %s

Respond with ONLY a JSON object, no other text: {"suggestedPercentage": <integer 0-100>, "reasoning": "<one sentence>"}`,
		claim.OrderID, claim.FailureDescription,
	)
}

// parseSuggestedPercentage extracts a 0-100 percentage from the LLM's
// response, clamping out-of-range values and defaulting to 100 (full
// refund) if the response isn't valid JSON.
func parseSuggestedPercentage(response string) int64 {
	trimmed := strings.TrimSpace(response)
	trimmed = strings.TrimPrefix(trimmed, "```json")
	trimmed = strings.TrimPrefix(trimmed, "```")
	trimmed = strings.TrimSuffix(trimmed, "```")
	trimmed = strings.TrimSpace(trimmed)

	var parsed llmSuggestion
	if err := json.Unmarshal([]byte(trimmed), &parsed); err != nil {
		return 100
	}

	switch {
	case parsed.SuggestedPercentage < 0:
		return 0
	case parsed.SuggestedPercentage > 100:
		return 100
	default:
		return parsed.SuggestedPercentage
	}
}
