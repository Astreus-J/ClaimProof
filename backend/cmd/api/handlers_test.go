package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func silentLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

func fakeTx() *types.Transaction {
	to := common.HexToAddress("0x1")
	return types.NewTx(&types.LegacyTx{Nonce: 0, GasPrice: big.NewInt(1), Gas: 21000, To: &to, Value: big.NewInt(0)})
}

type fakeSepoliaCreator struct {
	createErr error
	waitErr   error
}

func (f *fakeSepoliaCreator) CreateShipment(ctx context.Context, orderID *big.Int, buyer common.Address, slaSeconds *big.Int) (*types.Transaction, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if f.createErr != nil {
		return nil, f.createErr
	}
	return fakeTx(), nil
}

func (f *fakeSepoliaCreator) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if f.waitErr != nil {
		return nil, f.waitErr
	}
	return &types.Receipt{Status: 1}, nil
}

type fakeCreditcoinRegistrar struct {
	registerErr error
	waitErr     error
}

func (f *fakeCreditcoinRegistrar) RegisterOrder(ctx context.Context, orderID *big.Int, buyer common.Address, protectionAmount *big.Int) (*types.Transaction, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if f.registerErr != nil {
		return nil, f.registerErr
	}
	return fakeTx(), nil
}

func (f *fakeCreditcoinRegistrar) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if f.waitErr != nil {
		return nil, f.waitErr
	}
	return &types.Receipt{Status: 1}, nil
}

func doRequest(t *testing.T, handler http.Handler, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/api/orders", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec
}

func validBody() string {
	return `{"orderId": 43, "buyer": "0xD8108A4C6384866691b32c618892F0385CfC7a62", "protectionAmountWei": "30000000000000000", "slaSeconds": 86400}`
}

func TestHandleCreateOrder_Success(t *testing.T) {
	h := &orderHandler{
		sepolia:    &fakeSepoliaCreator{},
		creditcoin: &fakeCreditcoinRegistrar{},
		logger:     silentLogger(),
	}

	rec := doRequest(t, http.HandlerFunc(h.handleCreateOrder), validBody())

	require.Equal(t, http.StatusCreated, rec.Code)
	var resp createOrderResponse
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.NotEmpty(t, resp.CreateShipmentTxHash)
	assert.NotEmpty(t, resp.RegisterOrderTxHash)
}

func TestHandleCreateOrder_InvalidJSON(t *testing.T) {
	h := &orderHandler{sepolia: &fakeSepoliaCreator{}, creditcoin: &fakeCreditcoinRegistrar{}, logger: silentLogger()}

	rec := doRequest(t, http.HandlerFunc(h.handleCreateOrder), "not json")

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleCreateOrder_InvalidBuyerAddress(t *testing.T) {
	h := &orderHandler{sepolia: &fakeSepoliaCreator{}, creditcoin: &fakeCreditcoinRegistrar{}, logger: silentLogger()}

	rec := doRequest(t, http.HandlerFunc(h.handleCreateOrder), `{"orderId": 1, "buyer": "not-an-address", "protectionAmountWei": "1000", "slaSeconds": 60}`)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleCreateOrder_NonNumericProtectionAmount(t *testing.T) {
	h := &orderHandler{sepolia: &fakeSepoliaCreator{}, creditcoin: &fakeCreditcoinRegistrar{}, logger: silentLogger()}

	rec := doRequest(t, http.HandlerFunc(h.handleCreateOrder), `{"orderId": 1, "buyer": "0xD8108A4C6384866691b32c618892F0385CfC7a62", "protectionAmountWei": "not-a-number", "slaSeconds": 60}`)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleCreateOrder_ZeroProtectionAmount(t *testing.T) {
	h := &orderHandler{sepolia: &fakeSepoliaCreator{}, creditcoin: &fakeCreditcoinRegistrar{}, logger: silentLogger()}

	rec := doRequest(t, http.HandlerFunc(h.handleCreateOrder), `{"orderId": 1, "buyer": "0xD8108A4C6384866691b32c618892F0385CfC7a62", "protectionAmountWei": "0", "slaSeconds": 60}`)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleCreateOrder_ZeroSLASeconds(t *testing.T) {
	h := &orderHandler{sepolia: &fakeSepoliaCreator{}, creditcoin: &fakeCreditcoinRegistrar{}, logger: silentLogger()}

	rec := doRequest(t, http.HandlerFunc(h.handleCreateOrder), `{"orderId": 1, "buyer": "0xD8108A4C6384866691b32c618892F0385CfC7a62", "protectionAmountWei": "1000", "slaSeconds": 0}`)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleCreateOrder_ReturnsBadGatewayWhenCreateShipmentFails(t *testing.T) {
	h := &orderHandler{
		sepolia:    &fakeSepoliaCreator{createErr: errors.New("rpc down")},
		creditcoin: &fakeCreditcoinRegistrar{},
		logger:     silentLogger(),
	}

	rec := doRequest(t, http.HandlerFunc(h.handleCreateOrder), validBody())

	assert.Equal(t, http.StatusBadGateway, rec.Code)
}

func TestHandleCreateOrder_ReturnsBadGatewayWhenShipmentTxFailsToMine(t *testing.T) {
	h := &orderHandler{
		sepolia:    &fakeSepoliaCreator{waitErr: errors.New("timeout")},
		creditcoin: &fakeCreditcoinRegistrar{},
		logger:     silentLogger(),
	}

	rec := doRequest(t, http.HandlerFunc(h.handleCreateOrder), validBody())

	assert.Equal(t, http.StatusBadGateway, rec.Code)
}

func TestHandleCreateOrder_ReturnsBadGatewayWhenRegisterOrderFails(t *testing.T) {
	h := &orderHandler{
		sepolia:    &fakeSepoliaCreator{},
		creditcoin: &fakeCreditcoinRegistrar{registerErr: errors.New("not worker")},
		logger:     silentLogger(),
	}

	rec := doRequest(t, http.HandlerFunc(h.handleCreateOrder), validBody())

	assert.Equal(t, http.StatusBadGateway, rec.Code)
	assert.Contains(t, rec.Body.String(), "shipment created")
}

func TestHandleCreateOrder_SurvivesClientDisconnect(t *testing.T) {
	h := &orderHandler{
		sepolia:    &fakeSepoliaCreator{},
		creditcoin: &fakeCreditcoinRegistrar{},
		logger:     silentLogger(),
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // simulate the buyer's browser disconnecting before the on-chain submissions finish

	req := httptest.NewRequest(http.MethodPost, "/api/orders", bytes.NewBufferString(validBody())).WithContext(ctx)
	rec := httptest.NewRecorder()
	h.handleCreateOrder(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code, "a disconnected client must not abort in-flight chain submissions")
}

func TestHandleCreateOrder_ReturnsBadGatewayWhenRegisterOrderTxFailsToMine(t *testing.T) {
	h := &orderHandler{
		sepolia:    &fakeSepoliaCreator{},
		creditcoin: &fakeCreditcoinRegistrar{waitErr: errors.New("timeout")},
		logger:     silentLogger(),
	}

	rec := doRequest(t, http.HandlerFunc(h.handleCreateOrder), validBody())

	assert.Equal(t, http.StatusBadGateway, rec.Code)
}
