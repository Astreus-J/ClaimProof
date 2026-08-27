package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newReasoningMux() http.Handler {
	h := &claimReasoningHandler{store: newReasoningStore(), logger: silentLogger()}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/claims/{orderId}", h.handleSetReasoning)
	mux.HandleFunc("GET /api/claims/{orderId}", h.handleGetReasoning)
	return mux
}

func TestClaimReasoning_GetReturnsNotFoundBeforeAnyReportIsPosted(t *testing.T) {
	mux := newReasoningMux()

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/claims/42", nil))

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestClaimReasoning_PostThenGetRoundTrips(t *testing.T) {
	mux := newReasoningMux()

	body, _ := json.Marshal(map[string]string{
		"reasoning":          "full refund, no severity detail given",
		"suggestedPayoutWei": "1000000000000000000",
	})
	postRec := httptest.NewRecorder()
	mux.ServeHTTP(postRec, httptest.NewRequest(http.MethodPost, "/api/claims/42", bytes.NewReader(body)))
	require.Equal(t, http.StatusNoContent, postRec.Code)

	getRec := httptest.NewRecorder()
	mux.ServeHTTP(getRec, httptest.NewRequest(http.MethodGet, "/api/claims/42", nil))
	require.Equal(t, http.StatusOK, getRec.Code)

	var got claimReasoning
	require.NoError(t, json.NewDecoder(getRec.Body).Decode(&got))
	assert.Equal(t, "full refund, no severity detail given", got.Reasoning)
	assert.Equal(t, "1000000000000000000", got.SuggestedPayoutWei)
}

func TestClaimReasoning_GetForADifferentOrderIdIsNotFound(t *testing.T) {
	mux := newReasoningMux()

	body, _ := json.Marshal(map[string]string{"reasoning": "x", "suggestedPayoutWei": "1"})
	postRec := httptest.NewRecorder()
	mux.ServeHTTP(postRec, httptest.NewRequest(http.MethodPost, "/api/claims/42", bytes.NewReader(body)))
	require.Equal(t, http.StatusNoContent, postRec.Code)

	getRec := httptest.NewRecorder()
	mux.ServeHTTP(getRec, httptest.NewRequest(http.MethodGet, "/api/claims/43", nil))
	assert.Equal(t, http.StatusNotFound, getRec.Code)
}

func TestClaimReasoning_PostRejectsInvalidJSON(t *testing.T) {
	mux := newReasoningMux()

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/claims/42", bytes.NewBufferString("not json")))

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestClaimReasoning_PostRejectsEmptyReasoning(t *testing.T) {
	mux := newReasoningMux()

	body, _ := json.Marshal(map[string]string{"reasoning": "", "suggestedPayoutWei": "1"})
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/claims/42", bytes.NewReader(body)))

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
