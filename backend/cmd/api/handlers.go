package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// sepoliaCreator is the subset of *chain.SepoliaClient the order handler
// needs — an interface so it can be faked in tests without a live RPC.
type sepoliaCreator interface {
	CreateShipment(ctx context.Context, orderID *big.Int, buyer common.Address, slaSeconds *big.Int) (*types.Transaction, error)
	WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)
}

// creditcoinRegistrar is the subset of *chain.Client the order handler
// needs — an interface so it can be faked in tests without a live RPC.
type creditcoinRegistrar interface {
	RegisterOrder(ctx context.Context, orderID *big.Int, buyer common.Address, protectionAmount *big.Int) (*types.Transaction, error)
	WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)
}

type orderHandler struct {
	sepolia    sepoliaCreator
	creditcoin creditcoinRegistrar
	logger     *slog.Logger
}

type createOrderRequest struct {
	OrderID             uint64 `json:"orderId"`
	Buyer               string `json:"buyer"`
	ProtectionAmountWei string `json:"protectionAmountWei"`
	SLASeconds          uint64 `json:"slaSeconds"`
}

type createOrderResponse struct {
	CreateShipmentTxHash string `json:"createShipmentTxHash"`
	RegisterOrderTxHash  string `json:"registerOrderTxHash"`
}

type errorResponse struct {
	Error string `json:"error"`
}

// handleCreateOrder is the storefront's "buy protection" entry point. It
// mirrors docs/architecture.md's sequence diagram: the *operator* (this
// service), not the buyer's own wallet, signs both createShipment (Sepolia)
// and registerOrder (Creditcoin) — the buyer's wallet is used only to
// identify the address a payout should later go to.
func (h *orderHandler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var req createOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if !common.IsHexAddress(req.Buyer) {
		writeError(w, http.StatusBadRequest, "buyer is not a valid address")
		return
	}
	buyer := common.HexToAddress(req.Buyer)

	protectionAmount, ok := new(big.Int).SetString(req.ProtectionAmountWei, 10)
	if !ok || protectionAmount.Sign() <= 0 {
		writeError(w, http.StatusBadRequest, "protectionAmountWei must be a positive base-10 integer")
		return
	}

	if req.SLASeconds == 0 {
		writeError(w, http.StatusBadRequest, "slaSeconds must be greater than zero")
		return
	}

	orderID := new(big.Int).SetUint64(req.OrderID)
	slaSeconds := new(big.Int).SetUint64(req.SLASeconds)
	ctx := r.Context()
	logger := h.logger.With("orderId", req.OrderID)

	shipmentTx, err := h.sepolia.CreateShipment(ctx, orderID, buyer, slaSeconds)
	if err != nil {
		logger.Error("failed to create shipment", "error", err)
		writeError(w, http.StatusBadGateway, "failed to create shipment on Sepolia")
		return
	}
	if _, err := h.sepolia.WaitMined(ctx, shipmentTx); err != nil {
		logger.Error("shipment transaction failed to mine", "error", err)
		writeError(w, http.StatusBadGateway, "shipment transaction failed to mine")
		return
	}
	logger.Info("shipment created", "txHash", shipmentTx.Hash().Hex())

	registerTx, err := h.creditcoin.RegisterOrder(ctx, orderID, buyer, protectionAmount)
	if err != nil {
		logger.Error("failed to register order", "error", err)
		writeError(w, http.StatusBadGateway, "shipment created, but failed to register the order — contact support")
		return
	}
	if _, err := h.creditcoin.WaitMined(ctx, registerTx); err != nil {
		logger.Error("register order transaction failed to mine", "error", err)
		writeError(w, http.StatusBadGateway, "shipment created, but the order registration transaction failed to mine")
		return
	}
	logger.Info("order registered", "txHash", registerTx.Hash().Hex())

	writeJSON(w, http.StatusCreated, createOrderResponse{
		CreateShipmentTxHash: shipmentTx.Hash().Hex(),
		RegisterOrderTxHash:  registerTx.Hash().Hex(),
	})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Error: message})
}
