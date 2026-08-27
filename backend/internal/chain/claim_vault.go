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

// INativeQueryVerifierMerkleProofEntry is an auto generated low-level Go binding around an user-defined struct.
type INativeQueryVerifierMerkleProofEntry struct {
	Hash   [32]byte
	IsLeft bool
}

// ClaimVaultMetaData contains all meta data concerning the ClaimVault contract.
var ClaimVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"sourceContract_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialWorker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPayoutCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DELIVERY_FAILED_EVENT_SIGNATURE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINativeQueryVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fundPool\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"orders\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"protectionAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payoutCap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processedQueries\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOrder\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"protectionAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPayoutCap\",\"inputs\":[{\"name\":\"newCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWorker\",\"inputs\":[{\"name\":\"newWorker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sourceContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitClaim\",\"inputs\":[{\"name\":\"chainKey\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"encodedTransaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"siblings\",\"type\":\"tuple[]\",\"internalType\":\"structINativeQueryVerifier.MerkleProofEntry[]\",\"components\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isLeft\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"lowerEndpointDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"continuityRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"suggestedPayout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"payoutAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"worker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ClaimPaid\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"queryId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderRegistered\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"protectionAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayoutCapUpdated\",\"inputs\":[{\"name\":\"previousCap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newCap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFunded\",\"inputs\":[{\"name\":\"funder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkerUpdated\",\"inputs\":[{\"name\":\"previousWorker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newWorker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DeliveryFailedEventNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWorker\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderAlreadyClaimed\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OrderAlreadyExists\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OrderBuyerMismatch\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OrderNotFound\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PayoutTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QueryAlreadyProcessed\",\"inputs\":[{\"name\":\"queryId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SourceTransactionNotSuccessful\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60e0346100ef57601f61109e38819003918201601f19168301916001600160401b038311848410176100f3578084926060946040528339810103126100ef5761004781610107565b604061005560208401610107565b920151916001600160a01b038216156100e0576001600160a01b03169081156100e0573360805260a0525f80546001600160a01b031916919091179055600155610fd260c052604051610f82908161011c823960805181818160cc0152818161016c01526102b0015260a0518181816102510152610e40015260c0518181816104430152610cb90152f35b63d92e233d60e01b5f5260045ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100ef5756fe6080806040526004361015610012575f80fd5b5f3560e01c90816308c84e7014610ca757508063282c89ce14610c8a57806333cc99ab14610c5b5780634d547ada14610c3457806357dc1dbf14610b2357806370f37d2714610aea57806371441c7b14610ab057806381c284fb146102df5780638da5cb5b1461029b57806396365d4414610280578063a444ae411461023c578063a85c38ef146101ee578063c26f6d44146101445763e3643600146100b6575f80fd5b34610140576020366003190112610140576004357f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610131577f16b91e3df5f11d1a549caeac074e2d268332ae28ab46cbefa4909847be96270860406001548151908152836020820152a1600155005b6330cd747160e01b5f5260045ffd5b5f80fd5b34610140576020366003190112610140576004356001600160a01b03811690819003610140577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036101315780156101df575f548160018060a01b0382167f98b88aa89cb5f247008e613dc8529d633ab05a62f7120c07ebcfcdd852fc2a8d5f80a36001600160a01b031916175f55005b63d92e233d60e01b5f5260045ffd5b34610140576020366003190112610140576004355f526002602052606060405f2060018060a01b038154169060ff600260018301549201541690604051928352602083015215156040820152f35b34610140575f366003190112610140576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610140575f36600319011261014057602047604051908152f35b34610140575f366003190112610140576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461014057610100366003190112610140576004356001600160401b0381169081810361014057602435916001600160401b0383169283810361014057604435916001600160401b0383116101405736602384011215610140578260040135916001600160401b038311610140576024840193602484369201011161014057608435946001600160401b0386116101405736602387011215610140578560040135966001600160401b0388116101405760248860061b8801013681116101405760c435916001600160401b0383116101405736602384011215610140578260040135936001600160401b0385116101405760248560051b8501019a368c116101405760e4359a604051946103f286610d03565b606435865261040083610d3f565b9261040e6040519485610d1e565b83526024602084019201915b818310610a715750505060208481019190915260405163113e17c760e21b8152600481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169791818061047c6024820189610d76565b03818b5afa9081156109a6575f91610a37575b506040516001600160c01b031960c094851b81166020830190815293851b811660288301529190931b16603083015260188252906104ce603882610d1e565b519020998a5f52600360205260ff60405f205416610a2457604096949392919651926104f984610d03565b60a435845261050785610d3f565b946105156040519687610d1e565b8552602401602085015b828210610a145750505060606105786020928385019586526040519889976302f4d16760e01b89526004890152602488015260a0604488015261056660a488018b8d610d56565b87810360031901606489015290610d76565b93600319868603016084870152604085019351855251936040838201528451809452019201905f5b8181106109fb5750505091815f8160209503925af19081156109a6575f916109c0575b50156109b157610609915f91858352600360205260408320600160ff1982541617905560405193849283926306eae3bf60e11b8452602060048501526024840191610d56565b038173e671304e9d91b7df3a5d9019ea58e6fc33eb4bb35af49081156109a6575f91610799575b50600160ff8251160361078a57604061064a910151610e3e565b5f8281526002602052604090208054929390926001600160a01b0316908115610777576001600160a01b0383169182900361076457600284019081549460ff861661075157945f94888695869586956001809c9b0154808210881461074a57505b8b5490818111881461074057509a8b955b60ff1916179055887f1452c7c49daaf3c16d43eec536038b6301740b5b1d895e7ea5f0f03de02281776020604051878152a45af13d1561073b573d61070081610ddd565b9061070e6040519283610d1e565b81525f60203d92013e5b1561072c5760409182519182526020820152f35b631486dc3f60e21b5f5260045ffd5b610718565b90509a8b956106bc565b90506106ab565b8663616d044f60e01b5f5260045260245ffd5b84630aca528160e31b5f5260045260245ffd5b846313a42eb760e21b5f5260045260245ffd5b633c6dedfb60e21b5f5260045ffd5b90503d805f833e6107aa8183610d1e565b810190602081830312610140578051906001600160401b03821161014057016080818303126101405760405191608083018381106001600160401b0382111761099257604052815160ff8116810361014057835261080a60208301610dc9565b602084015260408201516001600160401b03811161014057820181601f820112156101405780519061083b82610d3f565b916108496040519384610d1e565b80835260208084019160051b830101918483116101405760208101915b83831061089e5750505050604084015260608201516001600160401b038111610140576108939201610df8565b606082015283610630565b82516001600160401b038111610140578201906060828803601f19011261014057604051906108cc82610ce8565b60208301516001600160a01b038116810361014057825260408301516001600160401b0381116101405760209084010188601f820112156101405780519061091382610d3f565b916109216040519384610d1e565b80835260208084019160051b830101918b831161014057602001905b8282106109825750505060208301526060830151916001600160401b0383116101405761097289602080969581960101610df8565b6040820152815201920191610866565b815181526020918201910161093d565b634e487b7160e01b5f52604160045260245ffd5b6040513d5f823e3d90fd5b6309bde33960e01b5f5260045ffd5b90506020813d6020116109f3575b816109db60209383610d1e565b810103126101405751801515810361014057856105c3565b3d91506109ce565b82518452869450602093840193909201916001016105a0565b813581526020918201910161051f565b8a6362e48a6560e01b5f5260045260245ffd5b90506020813d602011610a69575b81610a5260209383610d1e565b8101031261014057610a6390610dc9565b8d61048f565b3d9150610a45565b6040833603126101405760405190610a8882610d03565b833582526020840135908115158203610140578260209283604095015281520192019161041a565b34610140575f3660031901126101405760206040517fe4ce927580aec56981151f3af94fdeab8c1fb94f1d388862dbaf4c6fcbdec4dc8152f35b5f366003190112610140576040513481527f32173d8e51cec3a6fe484b7a1c3febe760cdf96e03d4cca36a43563a4333e83860203392a2005b34610140576060366003190112610140576024356001600160a01b0381169060043590829003610140575f54604435906001600160a01b03163303610c25575f828152600260205260409020546001600160a01b0316610c125782156101df5760207f2d4441e560168d0759d4f3ef8e50e50f95b5d4ead1742716247b50cc31b4bbd691604051610bb381610ce8565b85815260028382019183835260408101925f8452875f5282865260405f209160018060a01b039051166bffffffffffffffffffffffff60a01b835416178255516001820155019051151560ff80198354169116179055604051908152a3005b506314d0bb6760e31b5f5260045260245ffd5b63fb55adaf60e01b5f5260045ffd5b34610140575f366003190112610140575f546040516001600160a01b039091168152602090f35b34610140576020366003190112610140576004355f526003602052602060ff60405f2054166040519015158152f35b34610140575f366003190112610140576020600154604051908152f35b34610140575f366003190112610140577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b606081019081106001600160401b0382111761099257604052565b604081019081106001600160401b0382111761099257604052565b90601f801991011681019081106001600160401b0382111761099257604052565b6001600160401b0381116109925760051b60200190565b908060209392818452848401375f828201840152601f01601f1916010190565b60206060816040850193805186520151936040838201528451809452019201905f5b818110610da55750505090565b82518051855260209081015115158186015260409094019390920191600101610d98565b51906001600160401b038216820361014057565b6001600160401b03811161099257601f01601f191660200190565b81601f8201121561014057805190610e0f82610ddd565b92610e1d6040519485610d1e565b8284526020838301011161014057815f9260208093018386015e8301015290565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165f5b8251811015610f3d578251811015610ee257600581901b83016020015180516001600160a01b0316831480610f2e575b80610ef6575b610ead5750600101610e6a565b9150506020915001908151805160011015610ee257604001519151805160021015610ee257606001516001600160a01b031690565b634e487b7160e01b5f52603260045260245ffd5b506020810151805115610ee257602001517fe4ce927580aec56981151f3af94fdeab8c1fb94f1d388862dbaf4c6fcbdec4dc14610ea0565b50600360208201515114610e9a565b63f9a4632960e01b5f5260045ffdfea2646970667358221220941385cf49e5a608c1e1ee97f6c169efc7a0eda8ff0329c78cc03020255adb4664736f6c634300081a0033",
}

// ClaimVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimVaultMetaData.ABI instead.
var ClaimVaultABI = ClaimVaultMetaData.ABI

// ClaimVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClaimVaultMetaData.Bin instead.
var ClaimVaultBin = ClaimVaultMetaData.Bin

// DeployClaimVault deploys a new Ethereum contract, binding an instance of ClaimVault to it.
func DeployClaimVault(auth *bind.TransactOpts, backend bind.ContractBackend, sourceContract_ common.Address, initialWorker common.Address, initialPayoutCap *big.Int) (common.Address, *types.Transaction, *ClaimVault, error) {
	parsed, err := ClaimVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClaimVaultBin), backend, sourceContract_, initialWorker, initialPayoutCap)
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

// DELIVERYFAILEDEVENTSIGNATURE is a free data retrieval call binding the contract method 0x71441c7b.
//
// Solidity: function DELIVERY_FAILED_EVENT_SIGNATURE() view returns(bytes32)
func (_ClaimVault *ClaimVaultCaller) DELIVERYFAILEDEVENTSIGNATURE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ClaimVault.contract.Call(opts, &out, "DELIVERY_FAILED_EVENT_SIGNATURE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELIVERYFAILEDEVENTSIGNATURE is a free data retrieval call binding the contract method 0x71441c7b.
//
// Solidity: function DELIVERY_FAILED_EVENT_SIGNATURE() view returns(bytes32)
func (_ClaimVault *ClaimVaultSession) DELIVERYFAILEDEVENTSIGNATURE() ([32]byte, error) {
	return _ClaimVault.Contract.DELIVERYFAILEDEVENTSIGNATURE(&_ClaimVault.CallOpts)
}

// DELIVERYFAILEDEVENTSIGNATURE is a free data retrieval call binding the contract method 0x71441c7b.
//
// Solidity: function DELIVERY_FAILED_EVENT_SIGNATURE() view returns(bytes32)
func (_ClaimVault *ClaimVaultCallerSession) DELIVERYFAILEDEVENTSIGNATURE() ([32]byte, error) {
	return _ClaimVault.Contract.DELIVERYFAILEDEVENTSIGNATURE(&_ClaimVault.CallOpts)
}

// VERIFIER is a free data retrieval call binding the contract method 0x08c84e70.
//
// Solidity: function VERIFIER() view returns(address)
func (_ClaimVault *ClaimVaultCaller) VERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClaimVault.contract.Call(opts, &out, "VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VERIFIER is a free data retrieval call binding the contract method 0x08c84e70.
//
// Solidity: function VERIFIER() view returns(address)
func (_ClaimVault *ClaimVaultSession) VERIFIER() (common.Address, error) {
	return _ClaimVault.Contract.VERIFIER(&_ClaimVault.CallOpts)
}

// VERIFIER is a free data retrieval call binding the contract method 0x08c84e70.
//
// Solidity: function VERIFIER() view returns(address)
func (_ClaimVault *ClaimVaultCallerSession) VERIFIER() (common.Address, error) {
	return _ClaimVault.Contract.VERIFIER(&_ClaimVault.CallOpts)
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

// ProcessedQueries is a free data retrieval call binding the contract method 0x33cc99ab.
//
// Solidity: function processedQueries(bytes32 ) view returns(bool)
func (_ClaimVault *ClaimVaultCaller) ProcessedQueries(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ClaimVault.contract.Call(opts, &out, "processedQueries", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedQueries is a free data retrieval call binding the contract method 0x33cc99ab.
//
// Solidity: function processedQueries(bytes32 ) view returns(bool)
func (_ClaimVault *ClaimVaultSession) ProcessedQueries(arg0 [32]byte) (bool, error) {
	return _ClaimVault.Contract.ProcessedQueries(&_ClaimVault.CallOpts, arg0)
}

// ProcessedQueries is a free data retrieval call binding the contract method 0x33cc99ab.
//
// Solidity: function processedQueries(bytes32 ) view returns(bool)
func (_ClaimVault *ClaimVaultCallerSession) ProcessedQueries(arg0 [32]byte) (bool, error) {
	return _ClaimVault.Contract.ProcessedQueries(&_ClaimVault.CallOpts, arg0)
}

// SourceContract is a free data retrieval call binding the contract method 0xa444ae41.
//
// Solidity: function sourceContract() view returns(address)
func (_ClaimVault *ClaimVaultCaller) SourceContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClaimVault.contract.Call(opts, &out, "sourceContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SourceContract is a free data retrieval call binding the contract method 0xa444ae41.
//
// Solidity: function sourceContract() view returns(address)
func (_ClaimVault *ClaimVaultSession) SourceContract() (common.Address, error) {
	return _ClaimVault.Contract.SourceContract(&_ClaimVault.CallOpts)
}

// SourceContract is a free data retrieval call binding the contract method 0xa444ae41.
//
// Solidity: function sourceContract() view returns(address)
func (_ClaimVault *ClaimVaultCallerSession) SourceContract() (common.Address, error) {
	return _ClaimVault.Contract.SourceContract(&_ClaimVault.CallOpts)
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

// SubmitClaim is a paid mutator transaction binding the contract method 0x81c284fb.
//
// Solidity: function submitClaim(uint64 chainKey, uint64 blockHeight, bytes encodedTransaction, bytes32 merkleRoot, (bytes32,bool)[] siblings, bytes32 lowerEndpointDigest, bytes32[] continuityRoots, uint256 suggestedPayout) returns(uint256 orderId, uint256 payoutAmount)
func (_ClaimVault *ClaimVaultTransactor) SubmitClaim(opts *bind.TransactOpts, chainKey uint64, blockHeight uint64, encodedTransaction []byte, merkleRoot [32]byte, siblings []INativeQueryVerifierMerkleProofEntry, lowerEndpointDigest [32]byte, continuityRoots [][32]byte, suggestedPayout *big.Int) (*types.Transaction, error) {
	return _ClaimVault.contract.Transact(opts, "submitClaim", chainKey, blockHeight, encodedTransaction, merkleRoot, siblings, lowerEndpointDigest, continuityRoots, suggestedPayout)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0x81c284fb.
//
// Solidity: function submitClaim(uint64 chainKey, uint64 blockHeight, bytes encodedTransaction, bytes32 merkleRoot, (bytes32,bool)[] siblings, bytes32 lowerEndpointDigest, bytes32[] continuityRoots, uint256 suggestedPayout) returns(uint256 orderId, uint256 payoutAmount)
func (_ClaimVault *ClaimVaultSession) SubmitClaim(chainKey uint64, blockHeight uint64, encodedTransaction []byte, merkleRoot [32]byte, siblings []INativeQueryVerifierMerkleProofEntry, lowerEndpointDigest [32]byte, continuityRoots [][32]byte, suggestedPayout *big.Int) (*types.Transaction, error) {
	return _ClaimVault.Contract.SubmitClaim(&_ClaimVault.TransactOpts, chainKey, blockHeight, encodedTransaction, merkleRoot, siblings, lowerEndpointDigest, continuityRoots, suggestedPayout)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0x81c284fb.
//
// Solidity: function submitClaim(uint64 chainKey, uint64 blockHeight, bytes encodedTransaction, bytes32 merkleRoot, (bytes32,bool)[] siblings, bytes32 lowerEndpointDigest, bytes32[] continuityRoots, uint256 suggestedPayout) returns(uint256 orderId, uint256 payoutAmount)
func (_ClaimVault *ClaimVaultTransactorSession) SubmitClaim(chainKey uint64, blockHeight uint64, encodedTransaction []byte, merkleRoot [32]byte, siblings []INativeQueryVerifierMerkleProofEntry, lowerEndpointDigest [32]byte, continuityRoots [][32]byte, suggestedPayout *big.Int) (*types.Transaction, error) {
	return _ClaimVault.Contract.SubmitClaim(&_ClaimVault.TransactOpts, chainKey, blockHeight, encodedTransaction, merkleRoot, siblings, lowerEndpointDigest, continuityRoots, suggestedPayout)
}

// ClaimVaultClaimPaidIterator is returned from FilterClaimPaid and is used to iterate over the raw logs and unpacked data for ClaimPaid events raised by the ClaimVault contract.
type ClaimVaultClaimPaidIterator struct {
	Event *ClaimVaultClaimPaid // Event containing the contract specifics and raw log

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
func (it *ClaimVaultClaimPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimVaultClaimPaid)
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
		it.Event = new(ClaimVaultClaimPaid)
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
func (it *ClaimVaultClaimPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimVaultClaimPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimVaultClaimPaid represents a ClaimPaid event raised by the ClaimVault contract.
type ClaimVaultClaimPaid struct {
	OrderId *big.Int
	Buyer   common.Address
	Amount  *big.Int
	QueryId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimPaid is a free log retrieval operation binding the contract event 0x1452c7c49daaf3c16d43eec536038b6301740b5b1d895e7ea5f0f03de0228177.
//
// Solidity: event ClaimPaid(uint256 indexed orderId, address indexed buyer, uint256 amount, bytes32 indexed queryId)
func (_ClaimVault *ClaimVaultFilterer) FilterClaimPaid(opts *bind.FilterOpts, orderId []*big.Int, buyer []common.Address, queryId [][32]byte) (*ClaimVaultClaimPaidIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	var queryIdRule []interface{}
	for _, queryIdItem := range queryId {
		queryIdRule = append(queryIdRule, queryIdItem)
	}

	logs, sub, err := _ClaimVault.contract.FilterLogs(opts, "ClaimPaid", orderIdRule, buyerRule, queryIdRule)
	if err != nil {
		return nil, err
	}
	return &ClaimVaultClaimPaidIterator{contract: _ClaimVault.contract, event: "ClaimPaid", logs: logs, sub: sub}, nil
}

// WatchClaimPaid is a free log subscription operation binding the contract event 0x1452c7c49daaf3c16d43eec536038b6301740b5b1d895e7ea5f0f03de0228177.
//
// Solidity: event ClaimPaid(uint256 indexed orderId, address indexed buyer, uint256 amount, bytes32 indexed queryId)
func (_ClaimVault *ClaimVaultFilterer) WatchClaimPaid(opts *bind.WatchOpts, sink chan<- *ClaimVaultClaimPaid, orderId []*big.Int, buyer []common.Address, queryId [][32]byte) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	var queryIdRule []interface{}
	for _, queryIdItem := range queryId {
		queryIdRule = append(queryIdRule, queryIdItem)
	}

	logs, sub, err := _ClaimVault.contract.WatchLogs(opts, "ClaimPaid", orderIdRule, buyerRule, queryIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimVaultClaimPaid)
				if err := _ClaimVault.contract.UnpackLog(event, "ClaimPaid", log); err != nil {
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

// ParseClaimPaid is a log parse operation binding the contract event 0x1452c7c49daaf3c16d43eec536038b6301740b5b1d895e7ea5f0f03de0228177.
//
// Solidity: event ClaimPaid(uint256 indexed orderId, address indexed buyer, uint256 amount, bytes32 indexed queryId)
func (_ClaimVault *ClaimVaultFilterer) ParseClaimPaid(log types.Log) (*ClaimVaultClaimPaid, error) {
	event := new(ClaimVaultClaimPaid)
	if err := _ClaimVault.contract.UnpackLog(event, "ClaimPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
