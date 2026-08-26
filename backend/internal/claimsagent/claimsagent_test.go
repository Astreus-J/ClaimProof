package claimsagent

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
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
	a := New(llm, 500)

	assert.Equal(t, uint64(500), a.policyCap)
	assert.Equal(t, llm, a.llm)
}

func TestSuggestPayout_NotYetImplemented(t *testing.T) {
	a := New(&stubLLMClient{}, 500)

	amount, err := a.SuggestPayout(context.Background(), ClaimContext{OrderID: 1})

	assert.Equal(t, uint64(0), amount)
	assert.Error(t, err)
}
