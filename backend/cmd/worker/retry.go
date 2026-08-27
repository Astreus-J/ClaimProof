package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

const maxRetryAttempts = 5

// retryBaseBackoff is a var (not a const) so tests can shrink it to keep
// retry tests fast without changing production behavior.
var retryBaseBackoff = 2 * time.Second

const maxRetryBackoff = 30 * time.Second

// retryWithBackoff retries fn with exponential backoff on RPC/HTTP failure,
// up to maxRetryAttempts, or until ctx is done.
func retryWithBackoff(ctx context.Context, logger *slog.Logger, operation string, fn func() error) error {
	backoff := retryBaseBackoff

	var lastErr error
	for attempt := 1; attempt <= maxRetryAttempts; attempt++ {
		if err := fn(); err != nil {
			lastErr = err
			logger.Warn("operation failed, retrying", "operation", operation, "attempt", attempt, "error", err)

			if attempt == maxRetryAttempts {
				break
			}

			select {
			case <-ctx.Done():
				return fmt.Errorf("%s: %w", operation, ctx.Err())
			case <-time.After(backoff):
			}

			backoff *= 2
			if backoff > maxRetryBackoff {
				backoff = maxRetryBackoff
			}
			continue
		}
		return nil
	}

	return fmt.Errorf("%s: exhausted %d attempts: %w", operation, maxRetryAttempts, lastErr)
}
