// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dsproxy

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DsproxyABI is the input ABI used to generate the binding from.
const DsproxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cacheAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cache\",\"outputs\":[{\"internalType\":\"contractDSProxyCache\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_code\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cacheAddr\",\"type\":\"address\"}],\"name\":\"setCache\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DsproxyBin is the compiled bytecode used for deploying new contracts.
var DsproxyBin = "0x60806040523480156200001157600080fd5b506040516200183238038062001832833981810160405260208110156200003757600080fd5b810190808051906020019092919050505033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a2620000dd81620000e560201b60201c565b50506200059d565b60006200011d336000357fffffffff00000000000000000000000000000000000000000000000000000000166200033b60201b60201c565b62000190576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f64732d617574682d756e617574686f72697a656400000000000000000000000081525060200191505060405180910390fd5b60008060006004359250602435915034905081833373ffffffffffffffffffffffffffffffffffffffff166000357fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168460003660405180848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505094505050505060405180910390a4600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415620002ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f64732d70726f78792d63616368652d616464726573732d72657175697265640081525060200191505060405180910390fd5b84600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019350505050919050565b60003073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156200037c576001905062000597565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415620003dd576001905062000597565b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156200043e576000905062000597565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b70096138430856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001935050505060206040518083038186803b1580156200055757600080fd5b505afa1580156200056c573d6000803e3d6000fd5b505050506040513d60208110156200058357600080fd5b810190808051906020019092919050505090505b92915050565b61128580620005ad6000396000f3fe60806040526004361061007b5760003560e01c80637a9e5e4b1161004e5780637a9e5e4b146104775780638da5cb5b146104c8578063948f50761461051f578063bf7e214f146105885761007b565b806313af40351461007d5780631cff79cd146100ce5780631f6a1eb91461022257806360c7d29514610420575b005b34801561008957600080fd5b506100cc600480360360208110156100a057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105df565b005b6101a7600480360360408110156100e457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561012157600080fd5b82018360208201111561013357600080fd5b8035906020019184600183028401116401000000008311171561015557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610728565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101e75780820151818401526020810190506101cc565b50505050905090810190601f1680156102145780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103726004803603604081101561023857600080fd5b810190808035906020019064010000000081111561025557600080fd5b82018360208201111561026757600080fd5b8035906020019184600183028401116401000000008311171561028957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156102ec57600080fd5b8201836020820111156102fe57600080fd5b8035906020019184600183028401116401000000008311171561032057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061097a565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103e45780820151818401526020810190506103c9565b50505050905090810190601f1680156104115780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34801561042c57600080fd5b50610435610bf3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561048357600080fd5b506104c66004803603602081101561049a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c19565b005b3480156104d457600080fd5b506104dd610d60565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561052b57600080fd5b5061056e6004803603602081101561054257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d86565b604051808215151515815260200191505060405180910390f35b34801561059457600080fd5b5061059d610fd2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61060d336000357fffffffff0000000000000000000000000000000000000000000000000000000016610ff7565b61067f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f64732d617574682d756e617574686f72697a656400000000000000000000000081525060200191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a250565b6060610758336000357fffffffff0000000000000000000000000000000000000000000000000000000016610ff7565b6107ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f64732d617574682d756e617574686f72697a656400000000000000000000000081525060200191505060405180910390fd5b60008060006004359250602435915034905081833373ffffffffffffffffffffffffffffffffffffffff166000357fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168460003660405180848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505094505050505060405180910390a4600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415610927576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f64732d70726f78792d7461726765742d616464726573732d726571756972656481525060200191505060405180910390fd5b600080865160208801896113885a03f43d6040519550601f19601f6020830101168601604052808652806000602088013e8115600181146109675761096e565b8160208801fd5b50505050505092915050565b60006060600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638bf4515c856040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a0c5780820151818401526020810190506109f1565b50505050905090810190601f168015610a395780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015610a5657600080fd5b505afa158015610a6a573d6000803e3d6000fd5b505050506040513d6020811015610a8057600080fd5b81019080805190602001909291905050509150600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610be057600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637ed0c3b2856040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610b56578082015181840152602081019050610b3b565b50505050905090810190601f168015610b835780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015610ba257600080fd5b505af1158015610bb6573d6000803e3d6000fd5b505050506040513d6020811015610bcc57600080fd5b810190808051906020019092919050505091505b610bea8284610728565b90509250929050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610c47336000357fffffffff0000000000000000000000000000000000000000000000000000000016610ff7565b610cb9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f64732d617574682d756e617574686f72697a656400000000000000000000000081525060200191505060405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a250565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000610db6336000357fffffffff0000000000000000000000000000000000000000000000000000000016610ff7565b610e28576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f64732d617574682d756e617574686f72697a656400000000000000000000000081525060200191505060405180910390fd5b60008060006004359250602435915034905081833373ffffffffffffffffffffffffffffffffffffffff166000357fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168460003660405180848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505094505050505060405180910390a4600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415610f85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f64732d70726f78792d63616368652d616464726573732d72657175697265640081525060200191505060405180910390fd5b84600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019350505050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60003073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611036576001905061124a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611095576001905061124a565b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156110f4576000905061124a565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b70096138430856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001935050505060206040518083038186803b15801561120c57600080fd5b505afa158015611220573d6000803e3d6000fd5b505050506040513d602081101561123657600080fd5b810190808051906020019092919050505090505b9291505056fea265627a7a7231582005e2b4eb1dbbd9138937ee9bfb11c4d1a89a89b5bdbe0250d6d172eb3220411364736f6c634300050f0032"

// DeployDsproxy deploys a new Ethereum contract, binding an instance of Dsproxy to it.
func DeployDsproxy(auth *bind.TransactOpts, backend bind.ContractBackend, _cacheAddr common.Address) (common.Address, *types.Transaction, *Dsproxy, error) {
	parsed, err := abi.JSON(strings.NewReader(DsproxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DsproxyBin), backend, _cacheAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dsproxy{DsproxyCaller: DsproxyCaller{contract: contract}, DsproxyTransactor: DsproxyTransactor{contract: contract}, DsproxyFilterer: DsproxyFilterer{contract: contract}}, nil
}

// Dsproxy is an auto generated Go binding around an Ethereum contract.
type Dsproxy struct {
	DsproxyCaller     // Read-only binding to the contract
	DsproxyTransactor // Write-only binding to the contract
	DsproxyFilterer   // Log filterer for contract events
}

// DsproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type DsproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DsproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DsproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DsproxySession struct {
	Contract     *Dsproxy          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DsproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DsproxyCallerSession struct {
	Contract *DsproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DsproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DsproxyTransactorSession struct {
	Contract     *DsproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DsproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type DsproxyRaw struct {
	Contract *Dsproxy // Generic contract binding to access the raw methods on
}

// DsproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DsproxyCallerRaw struct {
	Contract *DsproxyCaller // Generic read-only contract binding to access the raw methods on
}

// DsproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DsproxyTransactorRaw struct {
	Contract *DsproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDsproxy creates a new instance of Dsproxy, bound to a specific deployed contract.
func NewDsproxy(address common.Address, backend bind.ContractBackend) (*Dsproxy, error) {
	contract, err := bindDsproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dsproxy{DsproxyCaller: DsproxyCaller{contract: contract}, DsproxyTransactor: DsproxyTransactor{contract: contract}, DsproxyFilterer: DsproxyFilterer{contract: contract}}, nil
}

// NewDsproxyCaller creates a new read-only instance of Dsproxy, bound to a specific deployed contract.
func NewDsproxyCaller(address common.Address, caller bind.ContractCaller) (*DsproxyCaller, error) {
	contract, err := bindDsproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DsproxyCaller{contract: contract}, nil
}

// NewDsproxyTransactor creates a new write-only instance of Dsproxy, bound to a specific deployed contract.
func NewDsproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*DsproxyTransactor, error) {
	contract, err := bindDsproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DsproxyTransactor{contract: contract}, nil
}

// NewDsproxyFilterer creates a new log filterer instance of Dsproxy, bound to a specific deployed contract.
func NewDsproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*DsproxyFilterer, error) {
	contract, err := bindDsproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DsproxyFilterer{contract: contract}, nil
}

// bindDsproxy binds a generic wrapper to an already deployed contract.
func bindDsproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DsproxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dsproxy *DsproxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dsproxy.Contract.DsproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dsproxy *DsproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dsproxy.Contract.DsproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dsproxy *DsproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dsproxy.Contract.DsproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dsproxy *DsproxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dsproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dsproxy *DsproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dsproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dsproxy *DsproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dsproxy.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Dsproxy *DsproxyCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsproxy.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Dsproxy *DsproxySession) Authority() (common.Address, error) {
	return _Dsproxy.Contract.Authority(&_Dsproxy.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Dsproxy *DsproxyCallerSession) Authority() (common.Address, error) {
	return _Dsproxy.Contract.Authority(&_Dsproxy.CallOpts)
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() constant returns(address)
func (_Dsproxy *DsproxyCaller) Cache(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsproxy.contract.Call(opts, out, "cache")
	return *ret0, err
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() constant returns(address)
func (_Dsproxy *DsproxySession) Cache() (common.Address, error) {
	return _Dsproxy.Contract.Cache(&_Dsproxy.CallOpts)
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() constant returns(address)
func (_Dsproxy *DsproxyCallerSession) Cache() (common.Address, error) {
	return _Dsproxy.Contract.Cache(&_Dsproxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dsproxy *DsproxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsproxy.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dsproxy *DsproxySession) Owner() (common.Address, error) {
	return _Dsproxy.Contract.Owner(&_Dsproxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dsproxy *DsproxyCallerSession) Owner() (common.Address, error) {
	return _Dsproxy.Contract.Owner(&_Dsproxy.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address _target, bytes _data) returns(bytes response)
func (_Dsproxy *DsproxyTransactor) Execute(opts *bind.TransactOpts, _target common.Address, _data []byte) (*types.Transaction, error) {
	return _Dsproxy.contract.Transact(opts, "execute", _target, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address _target, bytes _data) returns(bytes response)
func (_Dsproxy *DsproxySession) Execute(_target common.Address, _data []byte) (*types.Transaction, error) {
	return _Dsproxy.Contract.Execute(&_Dsproxy.TransactOpts, _target, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address _target, bytes _data) returns(bytes response)
func (_Dsproxy *DsproxyTransactorSession) Execute(_target common.Address, _data []byte) (*types.Transaction, error) {
	return _Dsproxy.Contract.Execute(&_Dsproxy.TransactOpts, _target, _data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _code, bytes _data) returns(address target, bytes response)
func (_Dsproxy *DsproxyTransactor) Execute0(opts *bind.TransactOpts, _code []byte, _data []byte) (*types.Transaction, error) {
	return _Dsproxy.contract.Transact(opts, "execute0", _code, _data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _code, bytes _data) returns(address target, bytes response)
func (_Dsproxy *DsproxySession) Execute0(_code []byte, _data []byte) (*types.Transaction, error) {
	return _Dsproxy.Contract.Execute0(&_Dsproxy.TransactOpts, _code, _data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _code, bytes _data) returns(address target, bytes response)
func (_Dsproxy *DsproxyTransactorSession) Execute0(_code []byte, _data []byte) (*types.Transaction, error) {
	return _Dsproxy.Contract.Execute0(&_Dsproxy.TransactOpts, _code, _data)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Dsproxy *DsproxyTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Dsproxy.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Dsproxy *DsproxySession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Dsproxy.Contract.SetAuthority(&_Dsproxy.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Dsproxy *DsproxyTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Dsproxy.Contract.SetAuthority(&_Dsproxy.TransactOpts, authority_)
}

// SetCache is a paid mutator transaction binding the contract method 0x948f5076.
//
// Solidity: function setCache(address _cacheAddr) returns(bool)
func (_Dsproxy *DsproxyTransactor) SetCache(opts *bind.TransactOpts, _cacheAddr common.Address) (*types.Transaction, error) {
	return _Dsproxy.contract.Transact(opts, "setCache", _cacheAddr)
}

// SetCache is a paid mutator transaction binding the contract method 0x948f5076.
//
// Solidity: function setCache(address _cacheAddr) returns(bool)
func (_Dsproxy *DsproxySession) SetCache(_cacheAddr common.Address) (*types.Transaction, error) {
	return _Dsproxy.Contract.SetCache(&_Dsproxy.TransactOpts, _cacheAddr)
}

// SetCache is a paid mutator transaction binding the contract method 0x948f5076.
//
// Solidity: function setCache(address _cacheAddr) returns(bool)
func (_Dsproxy *DsproxyTransactorSession) SetCache(_cacheAddr common.Address) (*types.Transaction, error) {
	return _Dsproxy.Contract.SetCache(&_Dsproxy.TransactOpts, _cacheAddr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Dsproxy *DsproxyTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Dsproxy.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Dsproxy *DsproxySession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Dsproxy.Contract.SetOwner(&_Dsproxy.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Dsproxy *DsproxyTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Dsproxy.Contract.SetOwner(&_Dsproxy.TransactOpts, owner_)
}

// DsproxyLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Dsproxy contract.
type DsproxyLogSetAuthorityIterator struct {
	Event *DsproxyLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *DsproxyLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DsproxyLogSetAuthority)
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
		it.Event = new(DsproxyLogSetAuthority)
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
func (it *DsproxyLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DsproxyLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DsproxyLogSetAuthority represents a LogSetAuthority event raised by the Dsproxy contract.
type DsproxyLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Dsproxy *DsproxyFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*DsproxyLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Dsproxy.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &DsproxyLogSetAuthorityIterator{contract: _Dsproxy.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Dsproxy *DsproxyFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *DsproxyLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Dsproxy.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DsproxyLogSetAuthority)
				if err := _Dsproxy.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// ParseLogSetAuthority is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Dsproxy *DsproxyFilterer) ParseLogSetAuthority(log types.Log) (*DsproxyLogSetAuthority, error) {
	event := new(DsproxyLogSetAuthority)
	if err := _Dsproxy.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DsproxyLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Dsproxy contract.
type DsproxyLogSetOwnerIterator struct {
	Event *DsproxyLogSetOwner // Event containing the contract specifics and raw log

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
func (it *DsproxyLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DsproxyLogSetOwner)
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
		it.Event = new(DsproxyLogSetOwner)
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
func (it *DsproxyLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DsproxyLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DsproxyLogSetOwner represents a LogSetOwner event raised by the Dsproxy contract.
type DsproxyLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Dsproxy *DsproxyFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*DsproxyLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Dsproxy.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DsproxyLogSetOwnerIterator{contract: _Dsproxy.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Dsproxy *DsproxyFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *DsproxyLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Dsproxy.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DsproxyLogSetOwner)
				if err := _Dsproxy.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// ParseLogSetOwner is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Dsproxy *DsproxyFilterer) ParseLogSetOwner(log types.Log) (*DsproxyLogSetOwner, error) {
	event := new(DsproxyLogSetOwner)
	if err := _Dsproxy.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}
