package main

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func silentLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

// useFastBackoffForTest keeps these tests fast without waiting out the real
// production backoff schedule.
func useFastBackoffForTest(t *testing.T) {
	t.Helper()
	original := retryBaseBackoff
	retryBaseBackoff = time.Millisecond
	t.Cleanup(func() { retryBaseBackoff = original })
}

func TestRetryWithBackoff_SucceedsOnFirstAttempt(t *testing.T) {
	useFastBackoffForTest(t)
	calls := 0
	err := retryWithBackoff(context.Background(), silentLogger(), "op", func() error {
		calls++
		return nil
	})

	require.NoError(t, err)
	assert.Equal(t, 1, calls)
}

func TestRetryWithBackoff_RetriesUntilSuccess(t *testing.T) {
	useFastBackoffForTest(t)
	calls := 0
	err := retryWithBackoff(context.Background(), silentLogger(), "op", func() error {
		calls++
		if calls < 3 {
			return errors.New("transient failure")
		}
		return nil
	})

	require.NoError(t, err)
	assert.Equal(t, 3, calls)
}

func TestRetryWithBackoff_ReturnsWrappedErrorAfterExhaustingAttempts(t *testing.T) {
	useFastBackoffForTest(t)
	calls := 0
	err := retryWithBackoff(context.Background(), silentLogger(), "op", func() error {
		calls++
		return errors.New("persistent failure")
	})

	require.Error(t, err)
	assert.Equal(t, maxRetryAttempts, calls)
	assert.ErrorContains(t, err, "op")
	assert.ErrorContains(t, err, "persistent failure")
}

func TestRetryWithBackoff_ReturnsContextErrorWhenCanceledBetweenAttempts(t *testing.T) {
	useFastBackoffForTest(t)
	ctx, cancel := context.WithCancel(context.Background())
	calls := 0
	err := retryWithBackoff(ctx, silentLogger(), "op", func() error {
		calls++
		cancel()
		return errors.New("failure")
	})

	require.Error(t, err)
	assert.ErrorIs(t, err, context.Canceled)
	assert.Equal(t, 1, calls)
}
