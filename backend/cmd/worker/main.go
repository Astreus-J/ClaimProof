// Command worker runs the ClaimProof off-chain worker: it watches Ethereum
// Sepolia for DeliveryFailed events, obtains an Attestcoin proof, asks the
// AI claims agent for a suggested payout, and submits the claim to
// ClaimVault on Creditcoin. See docs/architecture.md for the full flow.
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, logger); err != nil {
		logger.Error("worker exited with error", "error", err)
		os.Exit(1)
	}
}

// run wires the worker's components and blocks until ctx is canceled by
// SIGINT/SIGTERM. Component wiring (listener, proofbuilder, claimsagent,
// chain) is added in Sprint 3, once each package's real implementation
// lands — this skeleton only proves the process starts up and shuts down
// cleanly.
func run(ctx context.Context, logger *slog.Logger) error {
	logger.Info("worker starting")
	<-ctx.Done()
	logger.Info("worker shutting down")
	return nil
}
