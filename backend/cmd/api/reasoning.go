package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
)

// claimReasoning is the AI claims agent's advisory output for one order —
// purely informational (see internal/reasoningreporter's package doc): the
// real claim/payout truth always lives on-chain. This service's in-memory
// store is intentionally not persisted; losing it on restart only means the
// dashboard temporarily has nothing to show for older claims, not any loss
// of on-chain state.
type claimReasoning struct {
	Reasoning          string `json:"reasoning"`
	SuggestedPayoutWei string `json:"suggestedPayoutWei"`
}

type reasoningStore struct {
	mu   sync.RWMutex
	data map[uint64]claimReasoning
}

func newReasoningStore() *reasoningStore {
	return &reasoningStore{data: make(map[uint64]claimReasoning)}
}

func (s *reasoningStore) set(orderID uint64, r claimReasoning) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[orderID] = r
}

func (s *reasoningStore) get(orderID uint64) (claimReasoning, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	r, ok := s.data[orderID]
	return r, ok
}

// claimReasoningHandler serves the worker's write and the frontend's read
// of a claim's AI reasoning.
//
// The POST side has no access control — like DeliveryTrackerMock's
// reportDeliveryFailure, this is a deliberate simplification for the
// hackathon MVP scope. Unlike that function, it is not even on the
// payout-critical path: this store never influences a payout amount or
// authorization, so the worst case of abuse is misleading dashboard copy,
// not financial impact.
type claimReasoningHandler struct {
	store  *reasoningStore
	logger *slog.Logger
}

func (h *claimReasoningHandler) handleSetReasoning(w http.ResponseWriter, r *http.Request) {
	orderID, ok := parseOrderIDPathValue(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "orderId in path must be a non-negative integer")
		return
	}

	var body claimReasoning
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if body.Reasoning == "" {
		writeError(w, http.StatusBadRequest, "reasoning must not be empty")
		return
	}

	h.store.set(orderID, body)
	w.WriteHeader(http.StatusNoContent)
}

func (h *claimReasoningHandler) handleGetReasoning(w http.ResponseWriter, r *http.Request) {
	orderID, ok := parseOrderIDPathValue(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "orderId in path must be a non-negative integer")
		return
	}

	reasoning, found := h.store.get(orderID)
	if !found {
		writeError(w, http.StatusNotFound, "no AI reasoning recorded yet for this order")
		return
	}
	writeJSON(w, http.StatusOK, reasoning)
}

func parseOrderIDPathValue(r *http.Request) (uint64, bool) {
	orderID, err := strconv.ParseUint(r.PathValue("orderId"), 10, 64)
	if err != nil {
		return 0, false
	}
	return orderID, true
}
