package listener

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_StoresConfiguration(t *testing.T) {
	l := New("wss://example.invalid", "0xabc")

	assert.Equal(t, "wss://example.invalid", l.rpcURL)
	assert.Equal(t, "0xabc", l.contractAddress)
}

func TestListen_NotYetImplemented(t *testing.T) {
	l := New("wss://example.invalid", "0xabc")

	ch, err := l.Listen(context.Background())

	assert.Nil(t, ch)
	assert.Error(t, err)
}
