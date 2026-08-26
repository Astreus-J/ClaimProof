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

// ClaimVaultMetaData contains all meta data concerning the ClaimVault contract.
var ClaimVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialWorker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPayoutCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fundPool\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"orders\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"protectionAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payoutCap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOrder\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"protectionAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPayoutCap\",\"inputs\":[{\"name\":\"newCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWorker\",\"inputs\":[{\"name\":\"newWorker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"worker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OrderRegistered\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"protectionAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayoutCapUpdated\",\"inputs\":[{\"name\":\"previousCap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newCap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFunded\",\"inputs\":[{\"name\":\"funder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkerUpdated\",\"inputs\":[{\"name\":\"previousWorker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newWorker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderAlreadyExists\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a0604052348015600e575f80fd5b5060405161061f38038061061f833981016040819052602b91607c565b6001600160a01b03821660515760405163d92e233d60e01b815260040160405180910390fd5b336080525f80546001600160a01b0319166001600160a01b03939093169290921790915560015560b1565b5f8060408385031215608c575f80fd5b82516001600160a01b038116811460a1575f80fd5b6020939093015192949293505050565b6080516105486100d75f395f818161012001528181610345015261040e01526105485ff3fe608060405260043610610084575f3560e01c80638da5cb5b116100575780638da5cb5b1461010f57806396365d4414610142578063a85c38ef14610154578063c26f6d44146101c3578063e3643600146101e2575f80fd5b8063282c89ce146100885780634d547ada146100b057806357dc1dbf146100e657806370f37d2714610107575b5f80fd5b348015610093575f80fd5b5061009d60015481565b6040519081526020015b60405180910390f35b3480156100bb575f80fd5b505f546100ce906001600160a01b031681565b6040516001600160a01b0390911681526020016100a7565b3480156100f1575f80fd5b506101056101003660046104a8565b610201565b005b610105610303565b34801561011a575f80fd5b506100ce7f000000000000000000000000000000000000000000000000000000000000000081565b34801561014d575f80fd5b504761009d565b34801561015f575f80fd5b5061019c61016e3660046104db565b600260208190525f91825260409091208054600182015491909201546001600160a01b039092169160ff1683565b604080516001600160a01b03909416845260208401929092521515908201526060016100a7565b3480156101ce575f80fd5b506101056101dd3660046104f2565b61033a565b3480156101ed575f80fd5b506101056101fc3660046104db565b610403565b5f838152600260205260409020546001600160a01b03161561023d576040516314d0bb6760e31b81526004810184905260240160405180910390fd5b6001600160a01b0382166102645760405163d92e233d60e01b815260040160405180910390fd5b604080516060810182526001600160a01b0384811680835260208084018681525f8587018181528a8252600280855291889020965187546001600160a01b031916961695909517865590516001860155925193909201805460ff1916931515939093179092559151838152909185917f2d4441e560168d0759d4f3ef8e50e50f95b5d4ead1742716247b50cc31b4bbd6910160405180910390a3505050565b60405134815233907f32173d8e51cec3a6fe484b7a1c3febe760cdf96e03d4cca36a43563a4333e8389060200160405180910390a2565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610383576040516330cd747160e01b815260040160405180910390fd5b6001600160a01b0381166103aa5760405163d92e233d60e01b815260040160405180910390fd5b5f80546040516001600160a01b03808516939216917f98b88aa89cb5f247008e613dc8529d633ab05a62f7120c07ebcfcdd852fc2a8d91a35f80546001600160a01b0319166001600160a01b0392909216919091179055565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461044c576040516330cd747160e01b815260040160405180910390fd5b60015460408051918252602082018390527f16b91e3df5f11d1a549caeac074e2d268332ae28ab46cbefa4909847be962708910160405180910390a1600155565b80356001600160a01b03811681146104a3575f80fd5b919050565b5f805f606084860312156104ba575f80fd5b833592506104ca6020850161048d565b929592945050506040919091013590565b5f602082840312156104eb575f80fd5b5035919050565b5f60208284031215610502575f80fd5b61050b8261048d565b939250505056fea2646970667358221220c9809f085a1455b4a45348c751525434a82de9292b279814655f1dcbb6ca002964736f6c634300081a0033",
}

// ClaimVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimVaultMetaData.ABI instead.
var ClaimVaultABI = ClaimVaultMetaData.ABI

// ClaimVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClaimVaultMetaData.Bin instead.
var ClaimVaultBin = ClaimVaultMetaData.Bin

// DeployClaimVault deploys a new Ethereum contract, binding an instance of ClaimVault to it.
func DeployClaimVault(auth *bind.TransactOpts, backend bind.ContractBackend, initialWorker common.Address, initialPayoutCap *big.Int) (common.Address, *types.Transaction, *ClaimVault, error) {
	parsed, err := ClaimVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClaimVaultBin), backend, initialWorker, initialPayoutCap)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClaimVault{ClaimVaultCaller: ClaimVaultCaller{contract: contract}, ClaimVaultTransactor: ClaimVaultTransactor{contract: contract}, ClaimVaultFilterer: ClaimVaultFilterer{contract: contract}}, nil
}

// ClaimVault is an auto generated Go binding around an Ethereum contract.
type ClaimVault struct {
	ClaimVaultCaller     // Read-only binding to the contract
	ClaimVaultTransactor // Write-only binding to the contract
	ClaimVaultFilterer   // Log filterer for contract events
}

// ClaimVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimVaultSession struct {
	Contract     *ClaimVault       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimVaultCallerSession struct {
	Contract *ClaimVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ClaimVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimVaultTransactorSession struct {
	Contract     *ClaimVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ClaimVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimVaultRaw struct {
	Contract *ClaimVault // Generic contract binding to access the raw methods on
}

// ClaimVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimVaultCallerRaw struct {
	Contract *ClaimVaultCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimVaultTransactorRaw struct {
	Contract *ClaimVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimVault creates a new instance of ClaimVault, bound to a specific deployed contract.
func NewClaimVault(address common.Address, backend bind.ContractBackend) (*ClaimVault, error) {
	contract, err := bindClaimVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimVault{ClaimVaultCaller: ClaimVaultCaller{contract: contract}, ClaimVaultTransactor: ClaimVaultTransactor{contract: contract}, ClaimVaultFilterer: ClaimVaultFilterer{contract: contract}}, nil
}

// NewClaimVaultCaller creates a new read-only instance of ClaimVault, bound to a specific deployed contract.
func NewClaimVaultCaller(address common.Address, caller bind.ContractCaller) (*ClaimVaultCaller, error) {
	contract, err := bindClaimVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimVaultCaller{contract: contract}, nil
}

// NewClaimVaultTransactor creates a new write-only instance of ClaimVault, bound to a specific deployed contract.
func NewClaimVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimVaultTransactor, error) {
	contract, err := bindClaimVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimVaultTransactor{contract: contract}, nil
}

// NewClaimVaultFilterer creates a new log filterer instance of ClaimVault, bound to a specific deployed contract.
func NewClaimVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimVaultFilterer, error) {
	contract, err := bindClaimVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimVaultFilterer{contract: contract}, nil
}

// bindClaimVault binds a generic wrapper to an already deployed contract.
func bindClaimVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClaimVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimVault *ClaimVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimVault.Contract.ClaimVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimVault *ClaimVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimVault.Contract.ClaimVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimVault *ClaimVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimVault.Contract.ClaimVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimVault *ClaimVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimVault *ClaimVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimVault *ClaimVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimVault.Contract.contract.Transact(opts, method, params...)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address buyer, uint256 protectionAmount, bool claimed)
func (_ClaimVault *ClaimVaultCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Buyer            common.Address
	ProtectionAmount *big.Int
	Claimed          bool
}, error) {
	var out []interface{}
	err := _ClaimVault.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Buyer            common.Address
		ProtectionAmount *big.Int
		Claimed          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Buyer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ProtectionAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address buyer, uint256 protectionAmount, bool claimed)
func (_ClaimVault *ClaimVaultSession) Orders(arg0 *big.Int) (struct {
	Buyer            common.Address
	ProtectionAmount *big.Int
	Claimed          bool
}, error) {
	return _ClaimVault.Contract.Orders(&_ClaimVault.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address buyer, uint256 protectionAmount, bool claimed)
func (_ClaimVault *ClaimVaultCallerSession) Orders(arg0 *big.Int) (struct {
	Buyer            common.Address
	ProtectionAmount *big.Int
	Claimed          bool
}, error) {
	return _ClaimVault.Contract.Orders(&_ClaimVault.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClaimVault *ClaimVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClaimVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClaimVault *ClaimVaultSession) Owner() (common.Address, error) {
	return _ClaimVault.Contract.Owner(&_ClaimVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClaimVault *ClaimVaultCallerSession) Owner() (common.Address, error) {
	return _ClaimVault.Contract.Owner(&_ClaimVault.CallOpts)
}

// PayoutCap is a free data retrieval call binding the contract method 0x282c89ce.
//
// Solidity: function payoutCap() view returns(uint256)
func (_ClaimVault *ClaimVaultCaller) PayoutCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClaimVault.contract.Call(opts, &out, "payoutCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayoutCap is a free data retrieval call binding the contract method 0x282c89ce.
//
// Solidity: function payoutCap() view returns(uint256)
func (_ClaimVault *ClaimVaultSession) PayoutCap() (*big.Int, error) {
	return _ClaimVault.Contract.PayoutCap(&_ClaimVault.CallOpts)
}

// PayoutCap is a free data retrieval call binding the contract method 0x282c89ce.
//
// Solidity: function payoutCap() view returns(uint256)
func (_ClaimVault *ClaimVaultCallerSession) PayoutCap() (*big.Int, error) {
	return _ClaimVault.Contract.PayoutCap(&_ClaimVault.CallOpts)
}

// PoolBalance is a free data retrieval call binding the contract method 0x96365d44.
//
// Solidity: function poolBalance() view returns(uint256)
func (_ClaimVault *ClaimVaultCaller) PoolBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClaimVault.contract.Call(opts, &out, "poolBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolBalance is a free data retrieval call binding the contract method 0x96365d44.
//
// Solidity: function poolBalance() view returns(uint256)
func (_ClaimVault *ClaimVaultSession) PoolBalance() (*big.Int, error) {
	return _ClaimVault.Contract.PoolBalance(&_ClaimVault.CallOpts)
}

// PoolBalance is a free data retrieval call binding the contract method 0x96365d44.
//
// Solidity: function poolBalance() view returns(uint256)
func (_ClaimVault *ClaimVaultCallerSession) PoolBalance() (*big.Int, error) {
	return _ClaimVault.Contract.PoolBalance(&_ClaimVault.CallOpts)
}

// Worker is a free data retrieval call binding the contract method 0x4d547ada.
//
// Solidity: function worker() view returns(address)
func (_ClaimVault *ClaimVaultCaller) Worker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClaimVault.contract.Call(opts, &out, "worker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Worker is a free data retrieval call binding the contract method 0x4d547ada.
//
// Solidity: function worker() view returns(address)
func (_ClaimVault *ClaimVaultSession) Worker() (common.Address, error) {
	return _ClaimVault.Contract.Worker(&_ClaimVault.CallOpts)
}

// Worker is a free data retrieval call binding the contract method 0x4d547ada.
//
// Solidity: function worker() view returns(address)
func (_ClaimVault *ClaimVaultCallerSession) Worker() (common.Address, error) {
	return _ClaimVault.Contract.Worker(&_ClaimVault.CallOpts)
}

// FundPool is a paid mutator transaction binding the contract method 0x70f37d27.
//
// Solidity: function fundPool() payable returns()
func (_ClaimVault *ClaimVaultTransactor) FundPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimVault.contract.Transact(opts, "fundPool")
}

// FundPool is a paid mutator transaction binding the contract method 0x70f37d27.
//
// Solidity: function fundPool() payable returns()
func (_ClaimVault *ClaimVaultSession) FundPool() (*types.Transaction, error) {
	return _ClaimVault.Contract.FundPool(&_ClaimVault.TransactOpts)
}

// FundPool is a paid mutator transaction binding the contract method 0x70f37d27.
//
// Solidity: function fundPool() payable returns()
func (_ClaimVault *ClaimVaultTransactorSession) FundPool() (*types.Transaction, error) {
	return _ClaimVault.Contract.FundPool(&_ClaimVault.TransactOpts)
}

// RegisterOrder is a paid mutator transaction binding the contract method 0x57dc1dbf.
//
// Solidity: function registerOrder(uint256 orderId, address buyer, uint256 protectionAmount) returns()
func (_ClaimVault *ClaimVaultTransactor) RegisterOrder(opts *bind.TransactOpts, orderId *big.Int, buyer common.Address, protectionAmount *big.Int) (*types.Transaction, error) {
	return _ClaimVault.contract.Transact(opts, "registerOrder", orderId, buyer, protectionAmount)
}

// RegisterOrder is a paid mutator transaction binding the contract method 0x57dc1dbf.
//
// Solidity: function registerOrder(uint256 orderId, address buyer, uint256 protectionAmount) returns()
func (_ClaimVault *ClaimVaultSession) RegisterOrder(orderId *big.Int, buyer common.Address, protectionAmount *big.Int) (*types.Transaction, error) {
	return _ClaimVault.Contract.RegisterOrder(&_ClaimVault.TransactOpts, orderId, buyer, protectionAmount)
}

// RegisterOrder is a paid mutator transaction binding the contract method 0x57dc1dbf.
//
// Solidity: function registerOrder(uint256 orderId, address buyer, uint256 protectionAmount) returns()
func (_ClaimVault *ClaimVaultTransactorSession) RegisterOrder(orderId *big.Int, buyer common.Address, protectionAmount *big.Int) (*types.Transaction, error) {
	return _ClaimVault.Contract.RegisterOrder(&_ClaimVault.TransactOpts, orderId, buyer, protectionAmount)
}

// SetPayoutCap is a paid mutator transaction binding the contract method 0xe3643600.
//
// Solidity: function setPayoutCap(uint256 newCap) returns()
func (_ClaimVault *ClaimVaultTransactor) SetPayoutCap(opts *bind.TransactOpts, newCap *big.Int) (*types.Transaction, error) {
	return _ClaimVault.contract.Transact(opts, "setPayoutCap", newCap)
}

// SetPayoutCap is a paid mutator transaction binding the contract method 0xe3643600.
//
// Solidity: function setPayoutCap(uint256 newCap) returns()
func (_ClaimVault *ClaimVaultSession) SetPayoutCap(newCap *big.Int) (*types.Transaction, error) {
	return _ClaimVault.Contract.SetPayoutCap(&_ClaimVault.TransactOpts, newCap)
}

// SetPayoutCap is a paid mutator transaction binding the contract method 0xe3643600.
//
// Solidity: function setPayoutCap(uint256 newCap) returns()
func (_ClaimVault *ClaimVaultTransactorSession) SetPayoutCap(newCap *big.Int) (*types.Transaction, error) {
	return _ClaimVault.Contract.SetPayoutCap(&_ClaimVault.TransactOpts, newCap)
}

// SetWorker is a paid mutator transaction binding the contract method 0xc26f6d44.
//
// Solidity: function setWorker(address newWorker) returns()
func (_ClaimVault *ClaimVaultTransactor) SetWorker(opts *bind.TransactOpts, newWorker common.Address) (*types.Transaction, error) {
	return _ClaimVault.contract.Transact(opts, "setWorker", newWorker)
}

// SetWorker is a paid mutator transaction binding the contract method 0xc26f6d44.
//
// Solidity: function setWorker(address newWorker) returns()
func (_ClaimVault *ClaimVaultSession) SetWorker(newWorker common.Address) (*types.Transaction, error) {
	return _ClaimVault.Contract.SetWorker(&_ClaimVault.TransactOpts, newWorker)
}

// SetWorker is a paid mutator transaction binding the contract method 0xc26f6d44.
//
// Solidity: function setWorker(address newWorker) returns()
func (_ClaimVault *ClaimVaultTransactorSession) SetWorker(newWorker common.Address) (*types.Transaction, error) {
	return _ClaimVault.Contract.SetWorker(&_ClaimVault.TransactOpts, newWorker)
}

// ClaimVaultOrderRegisteredIterator is returned from FilterOrderRegistered and is used to iterate over the raw logs and unpacked data for OrderRegistered events raised by the ClaimVault contract.
type ClaimVaultOrderRegisteredIterator struct {
	Event *ClaimVaultOrderRegistered // Event containing the contract specifics and raw log

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
func (it *ClaimVaultOrderRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimVaultOrderRegistered)
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
		it.Event = new(ClaimVaultOrderRegistered)
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
func (it *ClaimVaultOrderRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimVaultOrderRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimVaultOrderRegistered represents a OrderRegistered event raised by the ClaimVault contract.
type ClaimVaultOrderRegistered struct {
	OrderId          *big.Int
	Buyer            common.Address
	ProtectionAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOrderRegistered is a free log retrieval operation binding the contract event 0x2d4441e560168d0759d4f3ef8e50e50f95b5d4ead1742716247b50cc31b4bbd6.
//
// Solidity: event OrderRegistered(uint256 indexed orderId, address indexed buyer, uint256 protectionAmount)
func (_ClaimVault *ClaimVaultFilterer) FilterOrderRegistered(opts *bind.FilterOpts, orderId []*big.Int, buyer []common.Address) (*ClaimVaultOrderRegisteredIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _ClaimVault.contract.FilterLogs(opts, "OrderRegistered", orderIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimVaultOrderRegisteredIterator{contract: _ClaimVault.contract, event: "OrderRegistered", logs: logs, sub: sub}, nil
}

// WatchOrderRegistered is a free log subscription operation binding the contract event 0x2d4441e560168d0759d4f3ef8e50e50f95b5d4ead1742716247b50cc31b4bbd6.
//
// Solidity: event OrderRegistered(uint256 indexed orderId, address indexed buyer, uint256 protectionAmount)
func (_ClaimVault *ClaimVaultFilterer) WatchOrderRegistered(opts *bind.WatchOpts, sink chan<- *ClaimVaultOrderRegistered, orderId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _ClaimVault.contract.WatchLogs(opts, "OrderRegistered", orderIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimVaultOrderRegistered)
				if err := _ClaimVault.contract.UnpackLog(event, "OrderRegistered", log); err != nil {
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

// ParseOrderRegistered is a log parse operation binding the contract event 0x2d4441e560168d0759d4f3ef8e50e50f95b5d4ead1742716247b50cc31b4bbd6.
//
// Solidity: event OrderRegistered(uint256 indexed orderId, address indexed buyer, uint256 protectionAmount)
func (_ClaimVault *ClaimVaultFilterer) ParseOrderRegistered(log types.Log) (*ClaimVaultOrderRegistered, error) {
	event := new(ClaimVaultOrderRegistered)
	if err := _ClaimVault.contract.UnpackLog(event, "OrderRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimVaultPayoutCapUpdatedIterator is returned from FilterPayoutCapUpdated and is used to iterate over the raw logs and unpacked data for PayoutCapUpdated events raised by the ClaimVault contract.
type ClaimVaultPayoutCapUpdatedIterator struct {
	Event *ClaimVaultPayoutCapUpdated // Event containing the contract specifics and raw log

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
func (it *ClaimVaultPayoutCapUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimVaultPayoutCapUpdated)
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
		it.Event = new(ClaimVaultPayoutCapUpdated)
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
func (it *ClaimVaultPayoutCapUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimVaultPayoutCapUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimVaultPayoutCapUpdated represents a PayoutCapUpdated event raised by the ClaimVault contract.
type ClaimVaultPayoutCapUpdated struct {
	PreviousCap *big.Int
	NewCap      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayoutCapUpdated is a free log retrieval operation binding the contract event 0x16b91e3df5f11d1a549caeac074e2d268332ae28ab46cbefa4909847be962708.
//
// Solidity: event PayoutCapUpdated(uint256 previousCap, uint256 newCap)
func (_ClaimVault *ClaimVaultFilterer) FilterPayoutCapUpdated(opts *bind.FilterOpts) (*ClaimVaultPayoutCapUpdatedIterator, error) {

	logs, sub, err := _ClaimVault.contract.FilterLogs(opts, "PayoutCapUpdated")
	if err != nil {
		return nil, err
	}
	return &ClaimVaultPayoutCapUpdatedIterator{contract: _ClaimVault.contract, event: "PayoutCapUpdated", logs: logs, sub: sub}, nil
}

// WatchPayoutCapUpdated is a free log subscription operation binding the contract event 0x16b91e3df5f11d1a549caeac074e2d268332ae28ab46cbefa4909847be962708.
//
// Solidity: event PayoutCapUpdated(uint256 previousCap, uint256 newCap)
func (_ClaimVault *ClaimVaultFilterer) WatchPayoutCapUpdated(opts *bind.WatchOpts, sink chan<- *ClaimVaultPayoutCapUpdated) (event.Subscription, error) {

	logs, sub, err := _ClaimVault.contract.WatchLogs(opts, "PayoutCapUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimVaultPayoutCapUpdated)
				if err := _ClaimVault.contract.UnpackLog(event, "PayoutCapUpdated", log); err != nil {
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

// ParsePayoutCapUpdated is a log parse operation binding the contract event 0x16b91e3df5f11d1a549caeac074e2d268332ae28ab46cbefa4909847be962708.
//
// Solidity: event PayoutCapUpdated(uint256 previousCap, uint256 newCap)
func (_ClaimVault *ClaimVaultFilterer) ParsePayoutCapUpdated(log types.Log) (*ClaimVaultPayoutCapUpdated, error) {
	event := new(ClaimVaultPayoutCapUpdated)
	if err := _ClaimVault.contract.UnpackLog(event, "PayoutCapUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimVaultPoolFundedIterator is returned from FilterPoolFunded and is used to iterate over the raw logs and unpacked data for PoolFunded events raised by the ClaimVault contract.
type ClaimVaultPoolFundedIterator struct {
	Event *ClaimVaultPoolFunded // Event containing the contract specifics and raw log

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
func (it *ClaimVaultPoolFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimVaultPoolFunded)
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
		it.Event = new(ClaimVaultPoolFunded)
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
func (it *ClaimVaultPoolFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimVaultPoolFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimVaultPoolFunded represents a PoolFunded event raised by the ClaimVault contract.
type ClaimVaultPoolFunded struct {
	Funder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolFunded is a free log retrieval operation binding the contract event 0x32173d8e51cec3a6fe484b7a1c3febe760cdf96e03d4cca36a43563a4333e838.
//
// Solidity: event PoolFunded(address indexed funder, uint256 amount)
func (_ClaimVault *ClaimVaultFilterer) FilterPoolFunded(opts *bind.FilterOpts, funder []common.Address) (*ClaimVaultPoolFundedIterator, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}

	logs, sub, err := _ClaimVault.contract.FilterLogs(opts, "PoolFunded", funderRule)
	if err != nil {
		return nil, err
	}
	return &ClaimVaultPoolFundedIterator{contract: _ClaimVault.contract, event: "PoolFunded", logs: logs, sub: sub}, nil
}

// WatchPoolFunded is a free log subscription operation binding the contract event 0x32173d8e51cec3a6fe484b7a1c3febe760cdf96e03d4cca36a43563a4333e838.
//
// Solidity: event PoolFunded(address indexed funder, uint256 amount)
func (_ClaimVault *ClaimVaultFilterer) WatchPoolFunded(opts *bind.WatchOpts, sink chan<- *ClaimVaultPoolFunded, funder []common.Address) (event.Subscription, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}

	logs, sub, err := _ClaimVault.contract.WatchLogs(opts, "PoolFunded", funderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimVaultPoolFunded)
				if err := _ClaimVault.contract.UnpackLog(event, "PoolFunded", log); err != nil {
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

// ParsePoolFunded is a log parse operation binding the contract event 0x32173d8e51cec3a6fe484b7a1c3febe760cdf96e03d4cca36a43563a4333e838.
//
// Solidity: event PoolFunded(address indexed funder, uint256 amount)
func (_ClaimVault *ClaimVaultFilterer) ParsePoolFunded(log types.Log) (*ClaimVaultPoolFunded, error) {
	event := new(ClaimVaultPoolFunded)
	if err := _ClaimVault.contract.UnpackLog(event, "PoolFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimVaultWorkerUpdatedIterator is returned from FilterWorkerUpdated and is used to iterate over the raw logs and unpacked data for WorkerUpdated events raised by the ClaimVault contract.
type ClaimVaultWorkerUpdatedIterator struct {
	Event *ClaimVaultWorkerUpdated // Event containing the contract specifics and raw log

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
func (it *ClaimVaultWorkerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimVaultWorkerUpdated)
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
		it.Event = new(ClaimVaultWorkerUpdated)
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
func (it *ClaimVaultWorkerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimVaultWorkerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimVaultWorkerUpdated represents a WorkerUpdated event raised by the ClaimVault contract.
type ClaimVaultWorkerUpdated struct {
	PreviousWorker common.Address
	NewWorker      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWorkerUpdated is a free log retrieval operation binding the contract event 0x98b88aa89cb5f247008e613dc8529d633ab05a62f7120c07ebcfcdd852fc2a8d.
//
// Solidity: event WorkerUpdated(address indexed previousWorker, address indexed newWorker)
func (_ClaimVault *ClaimVaultFilterer) FilterWorkerUpdated(opts *bind.FilterOpts, previousWorker []common.Address, newWorker []common.Address) (*ClaimVaultWorkerUpdatedIterator, error) {

	var previousWorkerRule []interface{}
	for _, previousWorkerItem := range previousWorker {
		previousWorkerRule = append(previousWorkerRule, previousWorkerItem)
	}
	var newWorkerRule []interface{}
	for _, newWorkerItem := range newWorker {
		newWorkerRule = append(newWorkerRule, newWorkerItem)
	}

	logs, sub, err := _ClaimVault.contract.FilterLogs(opts, "WorkerUpdated", previousWorkerRule, newWorkerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimVaultWorkerUpdatedIterator{contract: _ClaimVault.contract, event: "WorkerUpdated", logs: logs, sub: sub}, nil
}

// WatchWorkerUpdated is a free log subscription operation binding the contract event 0x98b88aa89cb5f247008e613dc8529d633ab05a62f7120c07ebcfcdd852fc2a8d.
//
// Solidity: event WorkerUpdated(address indexed previousWorker, address indexed newWorker)
func (_ClaimVault *ClaimVaultFilterer) WatchWorkerUpdated(opts *bind.WatchOpts, sink chan<- *ClaimVaultWorkerUpdated, previousWorker []common.Address, newWorker []common.Address) (event.Subscription, error) {

	var previousWorkerRule []interface{}
	for _, previousWorkerItem := range previousWorker {
		previousWorkerRule = append(previousWorkerRule, previousWorkerItem)
	}
	var newWorkerRule []interface{}
	for _, newWorkerItem := range newWorker {
		newWorkerRule = append(newWorkerRule, newWorkerItem)
	}

	logs, sub, err := _ClaimVault.contract.WatchLogs(opts, "WorkerUpdated", previousWorkerRule, newWorkerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimVaultWorkerUpdated)
				if err := _ClaimVault.contract.UnpackLog(event, "WorkerUpdated", log); err != nil {
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

// ParseWorkerUpdated is a log parse operation binding the contract event 0x98b88aa89cb5f247008e613dc8529d633ab05a62f7120c07ebcfcdd852fc2a8d.
//
// Solidity: event WorkerUpdated(address indexed previousWorker, address indexed newWorker)
func (_ClaimVault *ClaimVaultFilterer) ParseWorkerUpdated(log types.Log) (*ClaimVaultWorkerUpdated, error) {
	event := new(ClaimVaultWorkerUpdated)
	if err := _ClaimVault.contract.UnpackLog(event, "WorkerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
