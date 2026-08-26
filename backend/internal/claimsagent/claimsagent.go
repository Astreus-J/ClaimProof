// Package claimsagent asks an LLM to suggest a payout amount for a claim,
// bounded by an on-chain policy cap. The suggestion is advisory only — per
// CLAUDE.md's core invariant, the AI never authorizes payment alone; only
// ClaimVault.submitClaim's Attestcoin reverification does.
package claimsagent

import (
	"context"
	"fmt"
)

// ClaimContext is the order/claim data given to the LLM to reason about.
type ClaimContext struct {
	OrderID            uint64
	ProtectionAmount   uint64
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
	policyCap uint64
}

// New creates an Agent backed by the given LLM client and policy cap (wei).
func New(llm LLMClient, policyCap uint64) *Agent {
	return &Agent{llm: llm, policyCap: policyCap}
}

// SuggestPayout returns a suggested payout amount for the given claim,
// never exceeding the configured policy cap. Implemented in Sprint 3.
func (a *Agent) SuggestPayout(ctx context.Context, claim ClaimContext) (uint64, error) {
	return 0, fmt.Errorf("claimsagent: SuggestPayout not implemented until Sprint 3")
}
