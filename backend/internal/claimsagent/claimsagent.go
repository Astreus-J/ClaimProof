// Package claimsagent asks an LLM to suggest a payout amount for a claim,
// bounded by an on-chain policy cap. The suggestion is advisory only — per
// CLAUDE.md's core invariant, the AI never authorizes payment alone; only
// ClaimVault.submitClaim's Attestcoin reverification does.
//
// Two independent layers guard against the LLM hallucinating or inventing
// data, rather than relying on prompting alone:
//
//  1. Grounding at generation time: a system prompt (kept separate from the
//     per-claim user prompt) explicitly forbids inventing or assuming facts
//     not given, and the LLM is never told the claim's real monetary value —
//     it only ever judges severity as a 0-100 percentage, so it has nothing
//     concrete to hallucinate about the payout amount itself.
//  2. Verification after generation: the LLM must echo back the exact order
//     ID it was given (orderIdConfirmation). If that echo doesn't match the
//     order actually being processed — a sign the response is about a
//     different claim, or fabricated — the suggested percentage is
//     discarded outright and a safe full-refund default is used instead.
//     The policy cap enforced here (and ClaimVault's own on-chain payoutCap)
//     independently bound whatever number survives both layers.
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
// mocked in tests without a real API key. systemPrompt and userPrompt are
// sent as distinct roles wherever the underlying provider supports it.
type LLMClient interface {
	Complete(ctx context.Context, systemPrompt, userPrompt string) (string, error)
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
	OrderIDConfirmation *uint64 `json:"orderIdConfirmation"`
	SuggestedPercentage int64   `json:"suggestedPercentage"`
	Reasoning           string  `json:"reasoning"`
}

const systemPrompt = `You are a claims adjuster for a parametric delivery-protection product. A delivery failure has already been cryptographically verified on-chain by a separate system — that is settled fact, not something for you to judge.

Your ONLY job is to recommend what fraction (0-100%) of the buyer's protection amount to refund, based on the severity described in the failure description you are given.

Strict rules:
- Never invent, assume, or infer any fact not explicitly present in the user message. Do not guess at order details, dates, amounts, or circumstances that were not stated.
- You are never told the claim's monetary value — do not estimate or reference one.
- If the failure description does not give you enough detail to judge severity, default to suggestedPercentage: 100 rather than guessing.
- You must echo back the exact Order ID given to you, unchanged, as orderIdConfirmation — this is used to verify your response corresponds to the correct claim.
- Respond with ONLY a single JSON object, no other text and no markdown code fences: {"orderIdConfirmation": <integer, exactly the Order ID given>, "suggestedPercentage": <integer 0-100>, "reasoning": "<one sentence>"}`

// SuggestPayout asks the LLM what fraction of the order's protectionAmount
// to refund, and returns that fraction of protectionAmount, clamped to
// policyCap. If the LLM's response can't be parsed or fails the order-ID
// cross-check, it falls back to a full-refund suggestion — an off-chain
// parsing hiccup or hallucination must never itself deny a claim that
// Attestcoin has already verified actually happened; the cap still bounds
// the result either way.
func (a *Agent) SuggestPayout(ctx context.Context, claim ClaimContext) (*big.Int, error) {
	userPrompt := fmt.Sprintf("Order ID: %d\nFailure description: %s", claim.OrderID, claim.FailureDescription)

	response, err := a.llm.Complete(ctx, systemPrompt, userPrompt)
	if err != nil {
		return nil, fmt.Errorf("claimsagent: LLM completion: %w", err)
	}

	percentage := parseSuggestedPercentage(response, claim.OrderID)

	suggested := new(big.Int).Mul(claim.ProtectionAmount, big.NewInt(percentage))
	suggested.Div(suggested, big.NewInt(100))

	if suggested.Cmp(a.policyCap) > 0 {
		return new(big.Int).Set(a.policyCap), nil
	}
	return suggested, nil
}

// parseSuggestedPercentage extracts a 0-100 percentage from the LLM's
// response. It returns the safe default of 100 (full refund) if the
// response isn't valid JSON, or if the echoed orderIdConfirmation doesn't
// match expectedOrderID — the latter is treated as a hallucination/cross-talk
// signal and the suggested percentage is discarded outright, not merely
// distrusted-but-used.
func parseSuggestedPercentage(response string, expectedOrderID uint64) int64 {
	const fullRefundFallback = 100

	trimmed := strings.TrimSpace(response)
	trimmed = strings.TrimPrefix(trimmed, "```json")
	trimmed = strings.TrimPrefix(trimmed, "```")
	trimmed = strings.TrimSuffix(trimmed, "```")
	trimmed = strings.TrimSpace(trimmed)

	var parsed llmSuggestion
	if err := json.Unmarshal([]byte(trimmed), &parsed); err != nil {
		return fullRefundFallback
	}

	if parsed.OrderIDConfirmation == nil || *parsed.OrderIDConfirmation != expectedOrderID {
		return fullRefundFallback
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
