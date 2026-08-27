// Package listener watches Ethereum Sepolia for DeliveryFailed events
// emitted by DeliveryTrackerMock and publishes them on a channel for the
// worker to consume.
package listener

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Astreus-J/ClaimProof/backend/internal/chain"
)

// DeliveryFailedEvent is a decoded DeliveryFailed log from DeliveryTrackerMock.
type DeliveryFailedEvent struct {
	OrderID     uint64
	Buyer       common.Address
	Timestamp   uint64
	TxHash      common.Hash
	BlockNumber uint64
}

// Listener watches Ethereum Sepolia for DeliveryFailed events emitted by a
// DeliveryTrackerMock instance.
type Listener struct {
	client   *ethclient.Client
	filterer *chain.DeliveryTrackerMockFilterer
}

// New connects to the given Sepolia RPC endpoint (must be WSS for Listen's
// live subscription to work) and binds DeliveryTrackerMock's event filterer
// at contractAddress.
func New(ctx context.Context, rpcURL string, contractAddress common.Address) (*Listener, error) {
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("listener: dial %s: %w", rpcURL, err)
	}

	filterer, err := chain.NewDeliveryTrackerMockFilterer(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("listener: bind DeliveryTrackerMock at %s: %w", contractAddress, err)
	}

	return &Listener{client: client, filterer: filterer}, nil
}

// Close releases the underlying RPC connection.
func (l *Listener) Close() {
	l.client.Close()
}

// Listen subscribes to DeliveryFailed events and sends decoded events on the
// returned channel until ctx is canceled or the subscription errors.
func (l *Listener) Listen(ctx context.Context) (<-chan DeliveryFailedEvent, error) {
	raw := make(chan *chain.DeliveryTrackerMockDeliveryFailed)
	sub, err := l.filterer.WatchDeliveryFailed(&bind.WatchOpts{Context: ctx}, raw, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("listener: watch DeliveryFailed: %w", err)
	}

	events := make(chan DeliveryFailedEvent)
	go func() {
		defer close(events)
		defer sub.Unsubscribe()
		for {
			select {
			case <-ctx.Done():
				return
			case <-sub.Err():
				return
			case log := <-raw:
				select {
				case events <- toDeliveryFailedEvent(log):
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return events, nil
}

func toDeliveryFailedEvent(log *chain.DeliveryTrackerMockDeliveryFailed) DeliveryFailedEvent {
	return DeliveryFailedEvent{
		OrderID:     log.OrderId.Uint64(),
		Buyer:       log.Buyer,
		Timestamp:   log.Timestamp.Uint64(),
		TxHash:      log.Raw.TxHash,
		BlockNumber: log.Raw.BlockNumber,
	}
}
