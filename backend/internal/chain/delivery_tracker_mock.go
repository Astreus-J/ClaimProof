// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chain

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = time.Tick
	_ = context.Background
)

// DeliveryTrackerMockMetaData contains all meta data concerning the DeliveryTrackerMock contract.
var DeliveryTrackerMockMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"confirmDelivery\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createShipment\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slaSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportDeliveryFailure\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shipments\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slaDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDeliveryTrackerMock.ShipmentStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DeliveryConfirmed\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeliveryFailed\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShipmentCreated\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slaDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ShipmentAlreadyExists\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ShipmentNotFound\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ShipmentNotPending\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SlaNotExpired\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slaDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f80fd5b5061051f8061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c80632ac08a931461004e578063613e9883146100a85780639cb2da19146100bd578063fd84cb97146100d0575b5f80fd5b61008f61005c366004610410565b5f6020819052908152604090208054600182015460028301546003909301546001600160a01b0390921692909160ff1684565b60405161009f949392919061043b565b60405180910390f35b6100bb6100b6366004610410565b6100e3565b005b6100bb6100cb366004610484565b6101e7565b6100bb6100de366004610410565b61034e565b5f81815260208190526040902080546001600160a01b031661012057604051632139528960e21b8152600481018390526024015b60405180910390fd5b5f600382015460ff16600281111561013a5761013a610427565b1461015b57604051636407a63b60e01b815260048101839052602401610117565b8060020154421015610190576002810154604051631bfaf79960e21b8152610117918491600401918252602082015260400190565b60038101805460ff1916600217905580546040514281526001600160a01b039091169083907fe4ce927580aec56981151f3af94fdeab8c1fb94f1d388862dbaf4c6fcbdec4dc906020015b60405180910390a35050565b5f838152602081905260409020546001600160a01b03161561021f57604051633686d35160e01b815260048101849052602401610117565b6001600160a01b0382166102465760405163d92e233d60e01b815260040160405180910390fd5b5f61025182426104c4565b90506040518060800160405280846001600160a01b031681526020018581526020018281526020015f600281111561028b5761028b610427565b90525f8581526020818152604091829020835181546001600160a01b0319166001600160a01b039091161781559083015160018083019190915591830151600280830191909155606084015160038301805493949193909260ff19909116919084908111156102fc576102fc610427565b0217905550905050826001600160a01b0316847fedf892d41ca950980cd2f16b337b699f8ef55ad673f3c2810af49b6a1c20004d8360405161034091815260200190565b60405180910390a350505050565b5f81815260208190526040902080546001600160a01b031661038657604051632139528960e21b815260048101839052602401610117565b5f600382015460ff1660028111156103a0576103a0610427565b146103c157604051636407a63b60e01b815260048101839052602401610117565b60038101805460ff1916600117905580546040514281526001600160a01b039091169083907f2f4a952729db32131cb3d658cc231127e0ce7e50f83508ece9a49c985f2384d0906020016101db565b5f60208284031215610420575f80fd5b5035919050565b634e487b7160e01b5f52602160045260245ffd5b6001600160a01b03851681526020810184905260408101839052608081016003831061047557634e487b7160e01b5f52602160045260245ffd5b82606083015295945050505050565b5f805f60608486031215610496575f80fd5b8335925060208401356001600160a01b03811681146104b3575f80fd5b929592945050506040919091013590565b808201808211156104e357634e487b7160e01b5f52601160045260245ffd5b9291505056fea264697066735822122000870e19691775ace29c2e52d558cfd6112dec822e17b6348bf1249ea73664a864736f6c634300081a0033",
}

// DeliveryTrackerMockABI is the input ABI used to generate the binding from.
// Deprecated: Use DeliveryTrackerMockMetaData.ABI instead.
var DeliveryTrackerMockABI = DeliveryTrackerMockMetaData.ABI

// DeliveryTrackerMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DeliveryTrackerMockMetaData.Bin instead.
var DeliveryTrackerMockBin = DeliveryTrackerMockMetaData.Bin

// DeployDeliveryTrackerMock deploys a new Ethereum contract, binding an instance of DeliveryTrackerMock to it.
func DeployDeliveryTrackerMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DeliveryTrackerMock, error) {
	parsed, err := DeliveryTrackerMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DeliveryTrackerMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DeliveryTrackerMock{DeliveryTrackerMockCaller: DeliveryTrackerMockCaller{contract: contract}, DeliveryTrackerMockTransactor: DeliveryTrackerMockTransactor{contract: contract}, DeliveryTrackerMockFilterer: DeliveryTrackerMockFilterer{contract: contract}}, nil
}

// DeliveryTrackerMock is an auto generated Go binding around an Ethereum contract.
type DeliveryTrackerMock struct {
	DeliveryTrackerMockCaller     // Read-only binding to the contract
	DeliveryTrackerMockTransactor // Write-only binding to the contract
	DeliveryTrackerMockFilterer   // Log filterer for contract events
}

// DeliveryTrackerMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeliveryTrackerMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeliveryTrackerMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeliveryTrackerMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeliveryTrackerMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeliveryTrackerMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeliveryTrackerMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeliveryTrackerMockSession struct {
	Contract     *DeliveryTrackerMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DeliveryTrackerMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeliveryTrackerMockCallerSession struct {
	Contract *DeliveryTrackerMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// DeliveryTrackerMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeliveryTrackerMockTransactorSession struct {
	Contract     *DeliveryTrackerMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DeliveryTrackerMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeliveryTrackerMockRaw struct {
	Contract *DeliveryTrackerMock // Generic contract binding to access the raw methods on
}

// DeliveryTrackerMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeliveryTrackerMockCallerRaw struct {
	Contract *DeliveryTrackerMockCaller // Generic read-only contract binding to access the raw methods on
}

// DeliveryTrackerMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeliveryTrackerMockTransactorRaw struct {
	Contract *DeliveryTrackerMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeliveryTrackerMock creates a new instance of DeliveryTrackerMock, bound to a specific deployed contract.
func NewDeliveryTrackerMock(address common.Address, backend bind.ContractBackend) (*DeliveryTrackerMock, error) {
	contract, err := bindDeliveryTrackerMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeliveryTrackerMock{DeliveryTrackerMockCaller: DeliveryTrackerMockCaller{contract: contract}, DeliveryTrackerMockTransactor: DeliveryTrackerMockTransactor{contract: contract}, DeliveryTrackerMockFilterer: DeliveryTrackerMockFilterer{contract: contract}}, nil
}

// NewDeliveryTrackerMockCaller creates a new read-only instance of DeliveryTrackerMock, bound to a specific deployed contract.
func NewDeliveryTrackerMockCaller(address common.Address, caller bind.ContractCaller) (*DeliveryTrackerMockCaller, error) {
	contract, err := bindDeliveryTrackerMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeliveryTrackerMockCaller{contract: contract}, nil
}

// NewDeliveryTrackerMockTransactor creates a new write-only instance of DeliveryTrackerMock, bound to a specific deployed contract.
func NewDeliveryTrackerMockTransactor(address common.Address, transactor bind.ContractTransactor) (*DeliveryTrackerMockTransactor, error) {
	contract, err := bindDeliveryTrackerMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeliveryTrackerMockTransactor{contract: contract}, nil
}

// NewDeliveryTrackerMockFilterer creates a new log filterer instance of DeliveryTrackerMock, bound to a specific deployed contract.
func NewDeliveryTrackerMockFilterer(address common.Address, filterer bind.ContractFilterer) (*DeliveryTrackerMockFilterer, error) {
	contract, err := bindDeliveryTrackerMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeliveryTrackerMockFilterer{contract: contract}, nil
}

// bindDeliveryTrackerMock binds a generic wrapper to an already deployed contract.
func bindDeliveryTrackerMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeliveryTrackerMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeliveryTrackerMock *DeliveryTrackerMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeliveryTrackerMock.Contract.DeliveryTrackerMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeliveryTrackerMock *DeliveryTrackerMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeliveryTrackerMock.Contract.DeliveryTrackerMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeliveryTrackerMock *DeliveryTrackerMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeliveryTrackerMock.Contract.DeliveryTrackerMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeliveryTrackerMock *DeliveryTrackerMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeliveryTrackerMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeliveryTrackerMock *DeliveryTrackerMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeliveryTrackerMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeliveryTrackerMock *DeliveryTrackerMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeliveryTrackerMock.Contract.contract.Transact(opts, method, params...)
}

// Shipments is a free data retrieval call binding the contract method 0x2ac08a93.
//
// Solidity: function shipments(uint256 ) view returns(address buyer, uint256 orderId, uint256 slaDeadline, uint8 status)
func (_DeliveryTrackerMock *DeliveryTrackerMockCaller) Shipments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Buyer       common.Address
	OrderId     *big.Int
	SlaDeadline *big.Int
	Status      uint8
}, error) {
	var out []interface{}
	err := _DeliveryTrackerMock.contract.Call(opts, &out, "shipments", arg0)

	outstruct := new(struct {
		Buyer       common.Address
		OrderId     *big.Int
		SlaDeadline *big.Int
		Status      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Buyer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.OrderId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SlaDeadline = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// Shipments is a free data retrieval call binding the contract method 0x2ac08a93.
//
// Solidity: function shipments(uint256 ) view returns(address buyer, uint256 orderId, uint256 slaDeadline, uint8 status)
func (_DeliveryTrackerMock *DeliveryTrackerMockSession) Shipments(arg0 *big.Int) (struct {
	Buyer       common.Address
	OrderId     *big.Int
	SlaDeadline *big.Int
	Status      uint8
}, error) {
	return _DeliveryTrackerMock.Contract.Shipments(&_DeliveryTrackerMock.CallOpts, arg0)
}

// Shipments is a free data retrieval call binding the contract method 0x2ac08a93.
//
// Solidity: function shipments(uint256 ) view returns(address buyer, uint256 orderId, uint256 slaDeadline, uint8 status)
func (_DeliveryTrackerMock *DeliveryTrackerMockCallerSession) Shipments(arg0 *big.Int) (struct {
	Buyer       common.Address
	OrderId     *big.Int
	SlaDeadline *big.Int
	Status      uint8
}, error) {
	return _DeliveryTrackerMock.Contract.Shipments(&_DeliveryTrackerMock.CallOpts, arg0)
}

// ConfirmDelivery is a paid mutator transaction binding the contract method 0xfd84cb97.
//
// Solidity: function confirmDelivery(uint256 orderId) returns()
func (_DeliveryTrackerMock *DeliveryTrackerMockTransactor) ConfirmDelivery(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _DeliveryTrackerMock.contract.Transact(opts, "confirmDelivery", orderId)
}

// ConfirmDelivery is a paid mutator transaction binding the contract method 0xfd84cb97.
//
// Solidity: function confirmDelivery(uint256 orderId) returns()
func (_DeliveryTrackerMock *DeliveryTrackerMockSession) ConfirmDelivery(orderId *big.Int) (*types.Transaction, error) {
	return _DeliveryTrackerMock.Contract.ConfirmDelivery(&_DeliveryTrackerMock.TransactOpts, orderId)
}

// ConfirmDelivery is a paid mutator transaction binding the contract method 0xfd84cb97.
//
// Solidity: function confirmDelivery(uint256 orderId) returns()
func (_DeliveryTrackerMock *DeliveryTrackerMockTransactorSession) ConfirmDelivery(orderId *big.Int) (*types.Transaction, error) {
	return _DeliveryTrackerMock.Contract.ConfirmDelivery(&_DeliveryTrackerMock.TransactOpts, orderId)
}

// CreateShipment is a paid mutator transaction binding the contract method 0x9cb2da19.
//
// Solidity: function createShipment(uint256 orderId, address buyer, uint256 slaSeconds) returns()
func (_DeliveryTrackerMock *DeliveryTrackerMockTransactor) CreateShipment(opts *bind.TransactOpts, orderId *big.Int, buyer common.Address, slaSeconds *big.Int) (*types.Transaction, error) {
	return _DeliveryTrackerMock.contract.Transact(opts, "createShipment", orderId, buyer, slaSeconds)
}

// CreateShipment is a paid mutator transaction binding the contract method 0x9cb2da19.
//
// Solidity: function createShipment(uint256 orderId, address buyer, uint256 slaSeconds) returns()
func (_DeliveryTrackerMock *DeliveryTrackerMockSession) CreateShipment(orderId *big.Int, buyer common.Address, slaSeconds *big.Int) (*types.Transaction, error) {
	return _DeliveryTrackerMock.Contract.CreateShipment(&_DeliveryTrackerMock.TransactOpts, orderId, buyer, slaSeconds)
}

// CreateShipment is a paid mutator transaction binding the contract method 0x9cb2da19.
//
// Solidity: function createShipment(uint256 orderId, address buyer, uint256 slaSeconds) returns()
func (_DeliveryTrackerMock *DeliveryTrackerMockTransactorSession) CreateShipment(orderId *big.Int, buyer common.Address, slaSeconds *big.Int) (*types.Transaction, error) {
	return _DeliveryTrackerMock.Contract.CreateShipment(&_DeliveryTrackerMock.TransactOpts, orderId, buyer, slaSeconds)
}

// ReportDeliveryFailure is a paid mutator transaction binding the contract method 0x613e9883.
//
// Solidity: function reportDeliveryFailure(uint256 orderId) returns()
func (_DeliveryTrackerMock *DeliveryTrackerMockTransactor) ReportDeliveryFailure(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _DeliveryTrackerMock.contract.Transact(opts, "reportDeliveryFailure", orderId)
}

// ReportDeliveryFailure is a paid mutator transaction binding the contract method 0x613e9883.
//
// Solidity: function reportDeliveryFailure(uint256 orderId) returns()
func (_DeliveryTrackerMock *DeliveryTrackerMockSession) ReportDeliveryFailure(orderId *big.Int) (*types.Transaction, error) {
	return _DeliveryTrackerMock.Contract.ReportDeliveryFailure(&_DeliveryTrackerMock.TransactOpts, orderId)
}

// ReportDeliveryFailure is a paid mutator transaction binding the contract method 0x613e9883.
//
// Solidity: function reportDeliveryFailure(uint256 orderId) returns()
func (_DeliveryTrackerMock *DeliveryTrackerMockTransactorSession) ReportDeliveryFailure(orderId *big.Int) (*types.Transaction, error) {
	return _DeliveryTrackerMock.Contract.ReportDeliveryFailure(&_DeliveryTrackerMock.TransactOpts, orderId)
}

// DeliveryTrackerMockDeliveryConfirmedIterator is returned from FilterDeliveryConfirmed and is used to iterate over the raw logs and unpacked data for DeliveryConfirmed events raised by the DeliveryTrackerMock contract.
type DeliveryTrackerMockDeliveryConfirmedIterator struct {
	Event *DeliveryTrackerMockDeliveryConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeliveryTrackerMockDeliveryConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeliveryTrackerMockDeliveryConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeliveryTrackerMockDeliveryConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeliveryTrackerMockDeliveryConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeliveryTrackerMockDeliveryConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeliveryTrackerMockDeliveryConfirmed represents a DeliveryConfirmed event raised by the DeliveryTrackerMock contract.
type DeliveryTrackerMockDeliveryConfirmed struct {
	OrderId   *big.Int
	Buyer     common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeliveryConfirmed is a free log retrieval operation binding the contract event 0x2f4a952729db32131cb3d658cc231127e0ce7e50f83508ece9a49c985f2384d0.
//
// Solidity: event DeliveryConfirmed(uint256 indexed orderId, address indexed buyer, uint256 timestamp)
func (_DeliveryTrackerMock *DeliveryTrackerMockFilterer) FilterDeliveryConfirmed(opts *bind.FilterOpts, orderId []*big.Int, buyer []common.Address) (*DeliveryTrackerMockDeliveryConfirmedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _DeliveryTrackerMock.contract.FilterLogs(opts, "DeliveryConfirmed", orderIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &DeliveryTrackerMockDeliveryConfirmedIterator{contract: _DeliveryTrackerMock.contract, event: "DeliveryConfirmed", logs: logs, sub: sub}, nil
}

// WatchDeliveryConfirmed is a free log subscription operation binding the contract event 0x2f4a952729db32131cb3d658cc231127e0ce7e50f83508ece9a49c985f2384d0.
//
// Solidity: event DeliveryConfirmed(uint256 indexed orderId, address indexed buyer, uint256 timestamp)
func (_DeliveryTrackerMock *DeliveryTrackerMockFilterer) WatchDeliveryConfirmed(opts *bind.WatchOpts, sink chan<- *DeliveryTrackerMockDeliveryConfirmed, orderId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _DeliveryTrackerMock.contract.WatchLogs(opts, "DeliveryConfirmed", orderIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeliveryTrackerMockDeliveryConfirmed)
				if err := _DeliveryTrackerMock.contract.UnpackLog(event, "DeliveryConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeliveryConfirmed is a log parse operation binding the contract event 0x2f4a952729db32131cb3d658cc231127e0ce7e50f83508ece9a49c985f2384d0.
//
// Solidity: event DeliveryConfirmed(uint256 indexed orderId, address indexed buyer, uint256 timestamp)
func (_DeliveryTrackerMock *DeliveryTrackerMockFilterer) ParseDeliveryConfirmed(log types.Log) (*DeliveryTrackerMockDeliveryConfirmed, error) {
	event := new(DeliveryTrackerMockDeliveryConfirmed)
	if err := _DeliveryTrackerMock.contract.UnpackLog(event, "DeliveryConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeliveryTrackerMockDeliveryFailedIterator is returned from FilterDeliveryFailed and is used to iterate over the raw logs and unpacked data for DeliveryFailed events raised by the DeliveryTrackerMock contract.
type DeliveryTrackerMockDeliveryFailedIterator struct {
	Event *DeliveryTrackerMockDeliveryFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeliveryTrackerMockDeliveryFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeliveryTrackerMockDeliveryFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeliveryTrackerMockDeliveryFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeliveryTrackerMockDeliveryFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeliveryTrackerMockDeliveryFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeliveryTrackerMockDeliveryFailed represents a DeliveryFailed event raised by the DeliveryTrackerMock contract.
type DeliveryTrackerMockDeliveryFailed struct {
	OrderId   *big.Int
	Buyer     common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeliveryFailed is a free log retrieval operation binding the contract event 0xe4ce927580aec56981151f3af94fdeab8c1fb94f1d388862dbaf4c6fcbdec4dc.
//
// Solidity: event DeliveryFailed(uint256 indexed orderId, address indexed buyer, uint256 timestamp)
func (_DeliveryTrackerMock *DeliveryTrackerMockFilterer) FilterDeliveryFailed(opts *bind.FilterOpts, orderId []*big.Int, buyer []common.Address) (*DeliveryTrackerMockDeliveryFailedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _DeliveryTrackerMock.contract.FilterLogs(opts, "DeliveryFailed", orderIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &DeliveryTrackerMockDeliveryFailedIterator{contract: _DeliveryTrackerMock.contract, event: "DeliveryFailed", logs: logs, sub: sub}, nil
}

// WatchDeliveryFailed is a free log subscription operation binding the contract event 0xe4ce927580aec56981151f3af94fdeab8c1fb94f1d388862dbaf4c6fcbdec4dc.
//
// Solidity: event DeliveryFailed(uint256 indexed orderId, address indexed buyer, uint256 timestamp)
func (_DeliveryTrackerMock *DeliveryTrackerMockFilterer) WatchDeliveryFailed(opts *bind.WatchOpts, sink chan<- *DeliveryTrackerMockDeliveryFailed, orderId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _DeliveryTrackerMock.contract.WatchLogs(opts, "DeliveryFailed", orderIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeliveryTrackerMockDeliveryFailed)
				if err := _DeliveryTrackerMock.contract.UnpackLog(event, "DeliveryFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeliveryFailed is a log parse operation binding the contract event 0xe4ce927580aec56981151f3af94fdeab8c1fb94f1d388862dbaf4c6fcbdec4dc.
//
// Solidity: event DeliveryFailed(uint256 indexed orderId, address indexed buyer, uint256 timestamp)
func (_DeliveryTrackerMock *DeliveryTrackerMockFilterer) ParseDeliveryFailed(log types.Log) (*DeliveryTrackerMockDeliveryFailed, error) {
	event := new(DeliveryTrackerMockDeliveryFailed)
	if err := _DeliveryTrackerMock.contract.UnpackLog(event, "DeliveryFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeliveryTrackerMockShipmentCreatedIterator is returned from FilterShipmentCreated and is used to iterate over the raw logs and unpacked data for ShipmentCreated events raised by the DeliveryTrackerMock contract.
type DeliveryTrackerMockShipmentCreatedIterator struct {
	Event *DeliveryTrackerMockShipmentCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeliveryTrackerMockShipmentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeliveryTrackerMockShipmentCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeliveryTrackerMockShipmentCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeliveryTrackerMockShipmentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeliveryTrackerMockShipmentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeliveryTrackerMockShipmentCreated represents a ShipmentCreated event raised by the DeliveryTrackerMock contract.
type DeliveryTrackerMockShipmentCreated struct {
	OrderId     *big.Int
	Buyer       common.Address
	SlaDeadline *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterShipmentCreated is a free log retrieval operation binding the contract event 0xedf892d41ca950980cd2f16b337b699f8ef55ad673f3c2810af49b6a1c20004d.
//
// Solidity: event ShipmentCreated(uint256 indexed orderId, address indexed buyer, uint256 slaDeadline)
func (_DeliveryTrackerMock *DeliveryTrackerMockFilterer) FilterShipmentCreated(opts *bind.FilterOpts, orderId []*big.Int, buyer []common.Address) (*DeliveryTrackerMockShipmentCreatedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _DeliveryTrackerMock.contract.FilterLogs(opts, "ShipmentCreated", orderIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &DeliveryTrackerMockShipmentCreatedIterator{contract: _DeliveryTrackerMock.contract, event: "ShipmentCreated", logs: logs, sub: sub}, nil
}

// WatchShipmentCreated is a free log subscription operation binding the contract event 0xedf892d41ca950980cd2f16b337b699f8ef55ad673f3c2810af49b6a1c20004d.
//
// Solidity: event ShipmentCreated(uint256 indexed orderId, address indexed buyer, uint256 slaDeadline)
func (_DeliveryTrackerMock *DeliveryTrackerMockFilterer) WatchShipmentCreated(opts *bind.WatchOpts, sink chan<- *DeliveryTrackerMockShipmentCreated, orderId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _DeliveryTrackerMock.contract.WatchLogs(opts, "ShipmentCreated", orderIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeliveryTrackerMockShipmentCreated)
				if err := _DeliveryTrackerMock.contract.UnpackLog(event, "ShipmentCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseShipmentCreated is a log parse operation binding the contract event 0xedf892d41ca950980cd2f16b337b699f8ef55ad673f3c2810af49b6a1c20004d.
//
// Solidity: event ShipmentCreated(uint256 indexed orderId, address indexed buyer, uint256 slaDeadline)
func (_DeliveryTrackerMock *DeliveryTrackerMockFilterer) ParseShipmentCreated(log types.Log) (*DeliveryTrackerMockShipmentCreated, error) {
	event := new(DeliveryTrackerMockShipmentCreated)
	if err := _DeliveryTrackerMock.contract.UnpackLog(event, "ShipmentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
