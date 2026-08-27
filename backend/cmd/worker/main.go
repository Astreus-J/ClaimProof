// Command worker runs the ClaimProof off-chain worker: it watches Ethereum
// Sepolia for DeliveryFailed events, obtains an Attestcoin proof, asks the
// AI claims agent for a suggested payout, and submits the claim to
// ClaimVault on Creditcoin. See docs/architecture.md for the full flow.
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Astreus-J/ClaimProof/backend/internal/chain"
	"github.com/Astreus-J/ClaimProof/backend/internal/claimsagent"
	"github.com/Astreus-J/ClaimProof/backend/internal/listener"
	"github.com/Astreus-J/ClaimProof/backend/internal/proofbuilder"
)

// sepoliaChainKey is Attestcoin's chain key identifying Ethereum Sepolia as
// a source chain — see docs/ATTESTCOIN_INTEGRATION.md.
const sepoliaChainKey = 1

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, logger); err != nil {
		logger.Error("worker exited with error", "error", err)
		os.Exit(1)
	}
}

// run wires the worker's components and processes DeliveryFailed events
// until ctx is canceled by SIGINT/SIGTERM. On shutdown, it stops accepting
// new events but waits for any claim already in flight to finish, so the
// worker never leaves a claim half-submitted.
func run(ctx context.Context, logger *slog.Logger) error {
	cfg, err := loadConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	l, err := listener.New(ctx, cfg.sepoliaWSSURL, cfg.deliveryTrackerAddress)
	if err != nil {
		return fmt.Errorf("create listener: %w", err)
	}
	defer l.Close()

	pb := proofbuilder.New(cfg.attestcoinProverURL, sepoliaChainKey, nil, 0)

	llm, err := claimsagent.NewLLMClient(cfg.llmProvider, cfg.llmAPIKey, cfg.llmModel, nil)
	if err != nil {
		return fmt.Errorf("create LLM client: %w", err)
	}
	agent := claimsagent.New(llm, cfg.policyCapWei)

	chainClient, err := chain.New(ctx, cfg.creditcoinRPCURL, cfg.workerPrivateKey, big.NewInt(chain.CreditcoinTestnetChainID), cfg.claimVaultAddress)
	if err != nil {
		return fmt.Errorf("create chain client: %w", err)
	}
	defer chainClient.Close()

	events, err := l.Listen(ctx)
	if err != nil {
		return fmt.Errorf("start listening: %w", err)
	}

	logger.Info("worker started, listening for DeliveryFailed events")

	var inFlight sync.WaitGroup
	for {
		select {
		case <-ctx.Done():
			logger.Info("shutdown signal received, waiting for in-flight claims to finish")
			inFlight.Wait()
			logger.Info("worker shutting down")
			return nil
		case event, ok := <-events:
			if !ok {
				logger.Warn("event subscription closed, shutting down")
				inFlight.Wait()
				return errors.New("event subscription closed unexpectedly")
			}
			inFlight.Add(1)
			go func() {
				defer inFlight.Done()
				processClaim(ctx, logger, pb, agent, chainClient, event)
			}()
		}
	}
}

// processClaim carries one DeliveryFailed event from detection through to a
// mined (or failed) claim submission. Errors are logged and the claim is
// abandoned — the event's on-chain trigger remains available for a manual
// resubmission, since submitClaim's anti-replay key is per-proof, not
// per-attempt.
func processClaim(
	ctx context.Context,
	logger *slog.Logger,
	pb *proofbuilder.Client,
	agent *claimsagent.Agent,
	chainClient *chain.Client,
	event listener.DeliveryFailedEvent,
) {
	logger = logger.With("orderId", event.OrderID, "txHash", event.TxHash.Hex())
	logger.Info("delivery failure detected")

	orderID := new(big.Int).SetUint64(event.OrderID)
	order, err := chainClient.GetOrder(ctx, orderID)
	if err != nil {
		logger.Error("failed to read order", "error", err)
		return
	}
	if order.Claimed {
		logger.Info("order already claimed, skipping")
		return
	}

	if err := retryWithBackoff(ctx, logger, "wait for attestation", func() error {
		return pb.WaitUntilHeightAttested(ctx, event.BlockNumber)
	}); err != nil {
		logger.Error("failed waiting for attestation", "error", err)
		return
	}

	var proof *proofbuilder.Proof
	if err := retryWithBackoff(ctx, logger, "get proof", func() error {
		var err error
		proof, err = pb.GetProof(ctx, event.TxHash.Hex())
		return err
	}); err != nil {
		logger.Error("failed to get proof", "error", err)
		return
	}

	args, err := toSubmitClaimArgs(proof)
	if err != nil {
		logger.Error("failed to decode proof", "error", err)
		return
	}

	suggestion, err := agent.SuggestPayout(ctx, claimsagent.ClaimContext{
		OrderID:            event.OrderID,
		ProtectionAmount:   order.ProtectionAmount,
		FailureDescription: "delivery not confirmed before the SLA deadline",
	})
	if err != nil {
		logger.Error("failed to get AI payout suggestion", "error", err)
		return
	}
	logger.Info("AI suggested payout", "suggestedPayoutWei", suggestion.AmountWei.String(), "reasoning", suggestion.Reasoning)

	var tx *types.Transaction
	if err := retryWithBackoff(ctx, logger, "submit claim", func() error {
		submitted, err := chainClient.SubmitClaim(
			ctx, sepoliaChainKey, event.BlockNumber, args.encodedTransaction, args.merkleRoot,
			args.siblings, args.lowerEndpointDigest, args.continuityRoots, suggestion.AmountWei,
		)
		if err != nil {
			return err
		}
		tx = submitted
		return nil
	}); err != nil {
		logger.Error("failed to submit claim", "error", err)
		return
	}
	logger.Info("claim submitted", "claimTxHash", tx.Hash().Hex())

	receipt, err := chainClient.WaitMined(ctx, tx)
	if err != nil {
		logger.Error("failed waiting for claim to be mined", "error", err)
		return
	}
	logger.Info("claim mined", "status", receipt.Status, "blockNumber", receipt.BlockNumber.Uint64())
}
