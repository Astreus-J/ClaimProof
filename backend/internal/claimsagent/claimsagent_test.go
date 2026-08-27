package claimsagent

import (
	"context"
	"errors"
	"math/big"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubLLMClient struct {
	response       string
	err            error
	capturedSystem string
	capturedUser   string
}

func (s *stubLLMClient) Complete(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	s.capturedSystem = systemPrompt
	s.capturedUser = userPrompt
	return s.response, s.err
}

// withOrderID substitutes a literal orderIdConfirmation value into a
// response template like `{"orderIdConfirmation": %d, "suggestedPercentage": 100}`.
func withOrderID(template string, orderID uint64) string {
	return strings.Replace(template, "%d", strconv.FormatUint(orderID, 10), 1)
}

func TestNew_StoresPolicyCap(t *testing.T) {
	llm := &stubLLMClient{}
	a := New(llm, big.NewInt(500))

	assert.Equal(t, big.NewInt(500), a.policyCap)
	assert.Equal(t, llm, a.llm)
}

func TestSuggestPayout_SendsSeparateSystemPromptWithAntiHallucinationInstructions(t *testing.T) {
	llm := &stubLLMClient{response: withOrderID(`{"orderIdConfirmation": %d, "suggestedPercentage": 100}`, 1)}
	a := New(llm, big.NewInt(10_000))

	_, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "package never arrived",
	})

	require.NoError(t, err)
	assert.NotEmpty(t, llm.capturedSystem, "a system prompt must be sent, separate from the user prompt")
	assert.Contains(t, strings.ToLower(llm.capturedSystem), "never invent")
	assert.Contains(t, strings.ToLower(llm.capturedSystem), "orderidconfirmation")
}

func TestSuggestPayout_NeverExposesTheMonetaryAmountToTheLLM(t *testing.T) {
	// The LLM only ever judges severity as a percentage; the actual wei
	// amount is computed in Go afterward. This keeps the LLM out of the loop
	// on real financial figures entirely, narrowing what it could hallucinate.
	llm := &stubLLMClient{response: withOrderID(`{"orderIdConfirmation": %d, "suggestedPercentage": 100}`, 7)}
	a := New(llm, big.NewInt(10_000))

	_, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 7, ProtectionAmount: big.NewInt(123_456_789), FailureDescription: "package never arrived",
	})

	require.NoError(t, err)
	assert.NotContains(t, llm.capturedSystem, "123456789")
	assert.NotContains(t, llm.capturedUser, "123456789")
}

func TestSuggestPayout_FullRefundWhenLLMSuggestsHundredPercent(t *testing.T) {
	llm := &stubLLMClient{response: withOrderID(`{"orderIdConfirmation": %d, "suggestedPercentage": 100, "reasoning": "package never arrived"}`, 1)}
	a := New(llm, big.NewInt(10_000))

	suggestion, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "package never arrived",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1_000), suggestion.AmountWei)
}

func TestSuggestPayout_ReturnsTheLLMsOwnReasoningVerbatim(t *testing.T) {
	llm := &stubLLMClient{response: withOrderID(`{"orderIdConfirmation": %d, "suggestedPercentage": 30, "reasoning": "minor, one-day SLA breach"}`, 1)}
	a := New(llm, big.NewInt(10_000))

	suggestion, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "1 day late",
	})

	require.NoError(t, err)
	assert.Equal(t, "minor, one-day SLA breach", suggestion.Reasoning)
}

func TestSuggestPayout_PartialRefundHonored(t *testing.T) {
	llm := &stubLLMClient{response: withOrderID(`{"orderIdConfirmation": %d, "suggestedPercentage": 30, "reasoning": "minor SLA breach"}`, 1)}
	a := New(llm, big.NewInt(10_000))

	suggestion, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "1 day late",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(300), suggestion.AmountWei)
}

func TestSuggestPayout_CappedAtPolicyCapEvenWhenLLMSuggestsMore(t *testing.T) {
	// The policy cap is the real enforcement mechanism: an off-chain sanity
	// check independent of (and in addition to) ClaimVault's own on-chain cap.
	llm := &stubLLMClient{response: withOrderID(`{"orderIdConfirmation": %d, "suggestedPercentage": 100}`, 1)}
	a := New(llm, big.NewInt(500))

	suggestion, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "never arrived",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(500), suggestion.AmountWei)
}

func TestSuggestPayout_ClampsOutOfRangePercentageFromLLM(t *testing.T) {
	llm := &stubLLMClient{response: withOrderID(`{"orderIdConfirmation": %d, "suggestedPercentage": 250}`, 1)}
	a := New(llm, big.NewInt(10_000))

	suggestion, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "x",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1_000), suggestion.AmountWei, "a percentage above 100 must clamp to 100, not overpay")
}

func TestSuggestPayout_ClampsNegativePercentageFromLLM(t *testing.T) {
	llm := &stubLLMClient{response: withOrderID(`{"orderIdConfirmation": %d, "suggestedPercentage": -50}`, 1)}
	a := New(llm, big.NewInt(10_000))

	suggestion, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "x",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(0), suggestion.AmountWei)
}

func TestSuggestPayout_FallsBackToFullAmountOnUnparsableLLMResponse(t *testing.T) {
	llm := &stubLLMClient{response: "sorry, I cannot help with that"}
	a := New(llm, big.NewInt(10_000))

	suggestion, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "x",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1_000), suggestion.AmountWei, "an unparsable AI response must not deny a verified claim")
	assert.Contains(t, suggestion.Reasoning, "could not be parsed")
}

func TestSuggestPayout_HandlesMarkdownFencedJSON(t *testing.T) {
	llm := &stubLLMClient{response: "```json\n" + withOrderID(`{"orderIdConfirmation": %d, "suggestedPercentage": 40}`, 1) + "\n```"}
	a := New(llm, big.NewInt(10_000))

	suggestion, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "x",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(400), suggestion.AmountWei)
}

func TestSuggestPayout_ReturnsErrorWhenLLMCallFails(t *testing.T) {
	llm := &stubLLMClient{err: errors.New("network error")}
	a := New(llm, big.NewInt(10_000))

	suggestion, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "x",
	})

	assert.Nil(t, suggestion)
	assert.Error(t, err)
}

// --- Layer 2: post-response verification (hallucination / cross-talk detection) ---

func TestSuggestPayout_FallsBackToFullAmountWhenOrderIdConfirmationMismatches(t *testing.T) {
	// The LLM answered about a *different* order than the one it was asked
	// about — a hallucination/cross-talk signal. The suggested percentage
	// must be discarded entirely, not merely distrusted-but-used.
	llm := &stubLLMClient{response: withOrderID(`{"orderIdConfirmation": %d, "suggestedPercentage": 10}`, 999)}
	a := New(llm, big.NewInt(10_000))

	suggestion, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "x",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1_000), suggestion.AmountWei, "an order-ID mismatch must fall back to a full refund, ignoring the mismatched percentage")
	assert.Contains(t, suggestion.Reasoning, "did not match")
}

func TestSuggestPayout_FallsBackToFullAmountWhenOrderIdConfirmationMissing(t *testing.T) {
	llm := &stubLLMClient{response: `{"suggestedPercentage": 10}`}
	a := New(llm, big.NewInt(10_000))

	suggestion, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "x",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1_000), suggestion.AmountWei)
}
