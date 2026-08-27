// Command api runs the ClaimProof storefront API: it accepts "buy
// protection" requests from the frontend and, acting as the trusted
// operator (see docs/architecture.md's sequence diagram — "Store", not the
// buyer's own wallet, signs both writes), creates the shipment on Ethereum
// Sepolia and registers the order on Creditcoin's ClaimVault.
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Astreus-J/ClaimProof/backend/internal/chain"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, logger); err != nil {
		logger.Error("api exited with error", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, logger *slog.Logger) error {
	cfg, err := loadConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	sepoliaClient, err := chain.NewSepoliaClient(
		ctx, cfg.sepoliaRPCURL, cfg.operatorPrivateKey, big.NewInt(chain.SepoliaTestnetChainID), cfg.deliveryTrackerAddress,
	)
	if err != nil {
		return fmt.Errorf("create Sepolia client: %w", err)
	}
	defer sepoliaClient.Close()

	creditcoinClient, err := chain.New(
		ctx, cfg.creditcoinRPCURL, cfg.operatorPrivateKey, big.NewInt(chain.CreditcoinTestnetChainID), cfg.claimVaultAddress,
	)
	if err != nil {
		return fmt.Errorf("create Creditcoin client: %w", err)
	}
	defer creditcoinClient.Close()

	handler := &orderHandler{sepolia: sepoliaClient, creditcoin: creditcoinClient, logger: logger}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/orders", handler.handleCreateOrder)

	server := &http.Server{
		Addr:    cfg.listenAddr,
		Handler: withCORS(cfg.corsOrigin, mux),
	}

	serverErr := make(chan error, 1)
	go func() {
		logger.Info("api started", "addr", cfg.listenAddr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		return fmt.Errorf("server failed: %w", err)
	case <-ctx.Done():
		logger.Info("shutdown signal received")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("server shutdown: %w", err)
		}
		logger.Info("api shutting down")
		return nil
	}
}

// withCORS allows the frontend's dev server to call this API directly from
// the browser. This service serves no cookies/credentials, so a single
// configurable allowed origin (default "*" for local/demo use) is
// sufficient — see backend/.env.example, API_CORS_ORIGIN.
func withCORS(origin string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
