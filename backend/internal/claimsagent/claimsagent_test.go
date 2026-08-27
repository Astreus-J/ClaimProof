package claimsagent

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubLLMClient struct {
	response string
	err      error
}

func (s *stubLLMClient) Complete(ctx context.Context, prompt string) (string, error) {
	return s.response, s.err
}

func TestNew_StoresPolicyCap(t *testing.T) {
	llm := &stubLLMClient{}
	a := New(llm, big.NewInt(500))

	assert.Equal(t, big.NewInt(500), a.policyCap)
	assert.Equal(t, llm, a.llm)
}

func TestSuggestPayout_FullRefundWhenLLMSuggestsHundredPercent(t *testing.T) {
	llm := &stubLLMClient{response: `{"suggestedPercentage": 100, "reasoning": "package never arrived"}`}
	a := New(llm, big.NewInt(10_000))

	amount, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "package never arrived",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1_000), amount)
}

func TestSuggestPayout_PartialRefundHonored(t *testing.T) {
	llm := &stubLLMClient{response: `{"suggestedPercentage": 30, "reasoning": "minor SLA breach, package arrived 1 day late"}`}
	a := New(llm, big.NewInt(10_000))

	amount, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "1 day late",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(300), amount)
}

func TestSuggestPayout_CappedAtPolicyCapEvenWhenLLMSuggestsMore(t *testing.T) {
	// The policy cap is the real enforcement mechanism: an off-chain sanity
	// check independent of (and in addition to) ClaimVault's own on-chain cap.
	llm := &stubLLMClient{response: `{"suggestedPercentage": 100}`}
	a := New(llm, big.NewInt(500))

	amount, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "never arrived",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(500), amount)
}

func TestSuggestPayout_ClampsOutOfRangePercentageFromLLM(t *testing.T) {
	llm := &stubLLMClient{response: `{"suggestedPercentage": 250}`}
	a := New(llm, big.NewInt(10_000))

	amount, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "x",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1_000), amount, "a percentage above 100 must clamp to 100, not overpay")
}

func TestSuggestPayout_ClampsNegativePercentageFromLLM(t *testing.T) {
	llm := &stubLLMClient{response: `{"suggestedPercentage": -50}`}
	a := New(llm, big.NewInt(10_000))

	amount, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "x",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(0), amount)
}

func TestSuggestPayout_FallsBackToFullAmountOnUnparsableLLMResponse(t *testing.T) {
	llm := &stubLLMClient{response: "sorry, I cannot help with that"}
	a := New(llm, big.NewInt(10_000))

	amount, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "x",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1_000), amount, "an unparsable AI response must not deny a verified claim")
}

func TestSuggestPayout_HandlesMarkdownFencedJSON(t *testing.T) {
	llm := &stubLLMClient{response: "```json\n{\"suggestedPercentage\": 40}\n```"}
	a := New(llm, big.NewInt(10_000))

	amount, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "x",
	})

	require.NoError(t, err)
	assert.Equal(t, big.NewInt(400), amount)
}

func TestSuggestPayout_ReturnsErrorWhenLLMCallFails(t *testing.T) {
	llm := &stubLLMClient{err: errors.New("network error")}
	a := New(llm, big.NewInt(10_000))

	amount, err := a.SuggestPayout(context.Background(), ClaimContext{
		OrderID: 1, ProtectionAmount: big.NewInt(1_000), FailureDescription: "x",
	})

	assert.Nil(t, amount)
	assert.Error(t, err)
}
