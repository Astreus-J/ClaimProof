// Package listener watches Ethereum Sepolia for DeliveryFailed events
// emitted by DeliveryTrackerMock and publishes them on a channel for the
// worker to consume.
package listener

import (
	"context"
	"fmt"
)

// DeliveryFailedEvent is a decoded DeliveryFailed log from DeliveryTrackerMock.
type DeliveryFailedEvent struct {
	OrderID   uint64
	Buyer     string
	Timestamp uint64
}

// Listener watches Ethereum Sepolia for DeliveryFailed events emitted by the
// contract at ContractAddress.
type Listener struct {
	rpcURL          string
	contractAddress string
}

// New creates a Listener for the given Sepolia WSS RPC endpoint and
// DeliveryTrackerMock contract address.
func New(rpcURL, contractAddress string) *Listener {
	return &Listener{rpcURL: rpcURL, contractAddress: contractAddress}
}

// Listen subscribes to DeliveryFailed events via go-ethereum's ethclient and
// sends decoded events on the returned channel until ctx is canceled.
// Implemented in Sprint 3.
func (l *Listener) Listen(ctx context.Context) (<-chan DeliveryFailedEvent, error) {
	return nil, fmt.Errorf("listener: Listen not implemented until Sprint 3")
}
