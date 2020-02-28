// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package peth

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

// PethABI is the input ABI used to generate the binding from.
const PethABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"bytes32\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"symbol_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// PethBin is the compiled bytecode used for deploying new contracts.
var PethBin = "0x606060405260126006556000600755341561001957600080fd5b604051602080610dd283398101604052808051600160a060020a03331660008181526001602052604080822082905590805560048054600160a060020a0319168317905591935091507fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94905160405180910390a2600555610d338061009f6000396000f3006060604052600436106101485763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461014d57806307da68f514610172578063095ea7b31461018757806313af4035146101bd57806318160ddd146101dc57806323b872dd146101ef578063313ce5671461021757806340c10f191461022a57806342966c681461024c5780635ac801fe1461026257806370a082311461027857806375f12b21146102975780637a9e5e4b146102aa5780638da5cb5b146102c957806395d89b41146102f85780639dc29fac1461030b578063a0712d681461032d578063a9059cbb14610343578063b753a98c14610365578063bb35783b14610387578063be9a6555146103af578063bf7e214f146103c2578063daea85c5146103d5578063dd62ed3e146103f4578063f2d5d56b14610419575b600080fd5b341561015857600080fd5b61016061043b565b60405190815260200160405180910390f35b341561017d57600080fd5b610185610441565b005b341561019257600080fd5b6101a9600160a060020a03600435166024356104e0565b604051901515815260200160405180910390f35b34156101c857600080fd5b610185600160a060020a036004351661050d565b34156101e757600080fd5b61016061058c565b34156101fa57600080fd5b6101a9600160a060020a0360043581169060243516604435610592565b341561022257600080fd5b610160610707565b341561023557600080fd5b610185600160a060020a036004351660243561070d565b341561025757600080fd5b6101856004356107d3565b341561026d57600080fd5b6101856004356107e0565b341561028357600080fd5b610160600160a060020a0360043516610806565b34156102a257600080fd5b6101a9610821565b34156102b557600080fd5b610185600160a060020a0360043516610831565b34156102d457600080fd5b6102dc6108b0565b604051600160a060020a03909116815260200160405180910390f35b341561030357600080fd5b6101606108bf565b341561031657600080fd5b610185600160a060020a03600435166024356108c5565b341561033857600080fd5b610185600435610a33565b341561034e57600080fd5b6101a9600160a060020a0360043516602435610a3d565b341561037057600080fd5b610185600160a060020a0360043516602435610a4a565b341561039257600080fd5b610185600160a060020a0360043581169060243516604435610a5a565b34156103ba57600080fd5b610185610a6b565b34156103cd57600080fd5b6102dc610b04565b34156103e057600080fd5b6101a9600160a060020a0360043516610b13565b34156103ff57600080fd5b610160600160a060020a0360043581169060243516610b39565b341561042457600080fd5b610185600160a060020a0360043516602435610b64565b60075481565b61045733600035600160e060020a031916610b6f565b151561046257600080fd5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a450506004805474ff0000000000000000000000000000000000000000191660a060020a179055565b60045460009060a060020a900460ff16156104fa57600080fd5b6105048383610c7b565b90505b92915050565b61052333600035600160e060020a031916610b6f565b151561052e57600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a250565b60005490565b60045460009060a060020a900460ff16156105ac57600080fd5b33600160a060020a031684600160a060020a0316141580156105f65750600160a060020a038085166000908152600260209081526040808320339094168352929052205460001914155b1561065457600160a060020a038085166000908152600260209081526040808320339094168352929052205461062c9083610ce7565b600160a060020a03808616600090815260026020908152604080832033909416835292905220555b600160a060020a0384166000908152600160205260409020546106779083610ce7565b600160a060020a0380861660009081526001602052604080822093909355908516815220546106a69083610cf7565b600160a060020a03808516600081815260016020526040908190209390935591908616907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60065481565b61072333600035600160e060020a031916610b6f565b151561072e57600080fd5b60045460a060020a900460ff161561074557600080fd5b600160a060020a0382166000908152600160205260409020546107689082610cf7565b600160a060020a0383166000908152600160205260408120919091555461078f9082610cf7565b600055600160a060020a0382167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405190815260200160405180910390a25050565b6107dd33826108c5565b50565b6107f633600035600160e060020a031916610b6f565b151561080157600080fd5b600755565b600160a060020a031660009081526001602052604090205490565b60045460a060020a900460ff1681565b61084733600035600160e060020a031916610b6f565b151561085257600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a250565b600454600160a060020a031681565b60055481565b6108db33600035600160e060020a031916610b6f565b15156108e657600080fd5b60045460a060020a900460ff16156108fd57600080fd5b33600160a060020a031682600160a060020a0316141580156109475750600160a060020a038083166000908152600260209081526040808320339094168352929052205460001914155b156109a557600160a060020a038083166000908152600260209081526040808320339094168352929052205461097d9082610ce7565b600160a060020a03808416600090815260026020908152604080832033909416835292905220555b600160a060020a0382166000908152600160205260409020546109c89082610ce7565b600160a060020a038316600090815260016020526040812091909155546109ef9082610ce7565b600055600160a060020a0382167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405190815260200160405180910390a25050565b6107dd338261070d565b6000610504338484610592565b610a55338383610592565b505050565b610a65838383610592565b50505050565b610a8133600035600160e060020a031916610b6f565b1515610a8c57600080fd5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a450506004805474ff000000000000000000000000000000000000000019169055565b600354600160a060020a031681565b60045460009060a060020a900460ff1615610b2d57600080fd5b61050782600019610c7b565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b610a55823383610592565b600030600160a060020a031683600160a060020a03161415610b9357506001610507565b600454600160a060020a0384811691161415610bb157506001610507565b600354600160a060020a03161515610bcb57506000610507565b600354600160a060020a031663b70096138430856000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b1515610c5957600080fd5b6102c65a03f11515610c6a57600080fd5b505050604051805190509050610507565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b8082038281111561050757600080fd5b8082018281101561050757600080fd00a165627a7a723058200877df264aa5d498c61a45dfe4fc04c68d11820448cf0cc7a275283a271173b40029494f550000000000000000000000000000000000000000000000000000000000"

// DeployPeth deploys a new Ethereum contract, binding an instance of Peth to it.
func DeployPeth(auth *bind.TransactOpts, backend bind.ContractBackend, symbol_ [32]byte) (common.Address, *types.Transaction, *Peth, error) {
	parsed, err := abi.JSON(strings.NewReader(PethABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PethBin), backend, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Peth{PethCaller: PethCaller{contract: contract}, PethTransactor: PethTransactor{contract: contract}, PethFilterer: PethFilterer{contract: contract}}, nil
}

// Peth is an auto generated Go binding around an Ethereum contract.
type Peth struct {
	PethCaller     // Read-only binding to the contract
	PethTransactor // Write-only binding to the contract
	PethFilterer   // Log filterer for contract events
}

// PethCaller is an auto generated read-only Go binding around an Ethereum contract.
type PethCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PethTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PethTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PethFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PethFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PethSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PethSession struct {
	Contract     *Peth             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PethCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PethCallerSession struct {
	Contract *PethCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PethTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PethTransactorSession struct {
	Contract     *PethTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PethRaw is an auto generated low-level Go binding around an Ethereum contract.
type PethRaw struct {
	Contract *Peth // Generic contract binding to access the raw methods on
}

// PethCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PethCallerRaw struct {
	Contract *PethCaller // Generic read-only contract binding to access the raw methods on
}

// PethTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PethTransactorRaw struct {
	Contract *PethTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeth creates a new instance of Peth, bound to a specific deployed contract.
func NewPeth(address common.Address, backend bind.ContractBackend) (*Peth, error) {
	contract, err := bindPeth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Peth{PethCaller: PethCaller{contract: contract}, PethTransactor: PethTransactor{contract: contract}, PethFilterer: PethFilterer{contract: contract}}, nil
}

// NewPethCaller creates a new read-only instance of Peth, bound to a specific deployed contract.
func NewPethCaller(address common.Address, caller bind.ContractCaller) (*PethCaller, error) {
	contract, err := bindPeth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PethCaller{contract: contract}, nil
}

// NewPethTransactor creates a new write-only instance of Peth, bound to a specific deployed contract.
func NewPethTransactor(address common.Address, transactor bind.ContractTransactor) (*PethTransactor, error) {
	contract, err := bindPeth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PethTransactor{contract: contract}, nil
}

// NewPethFilterer creates a new log filterer instance of Peth, bound to a specific deployed contract.
func NewPethFilterer(address common.Address, filterer bind.ContractFilterer) (*PethFilterer, error) {
	contract, err := bindPeth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PethFilterer{contract: contract}, nil
}

// bindPeth binds a generic wrapper to an already deployed contract.
func bindPeth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PethABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Peth *PethRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Peth.Contract.PethCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Peth *PethRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peth.Contract.PethTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Peth *PethRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Peth.Contract.PethTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Peth *PethCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Peth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Peth *PethTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Peth *PethTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Peth.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address guy) constant returns(uint256)
func (_Peth *PethCaller) Allowance(opts *bind.CallOpts, src common.Address, guy common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Peth.contract.Call(opts, out, "allowance", src, guy)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address guy) constant returns(uint256)
func (_Peth *PethSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _Peth.Contract.Allowance(&_Peth.CallOpts, src, guy)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address guy) constant returns(uint256)
func (_Peth *PethCallerSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _Peth.Contract.Allowance(&_Peth.CallOpts, src, guy)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Peth *PethCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Peth.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Peth *PethSession) Authority() (common.Address, error) {
	return _Peth.Contract.Authority(&_Peth.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Peth *PethCallerSession) Authority() (common.Address, error) {
	return _Peth.Contract.Authority(&_Peth.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address src) constant returns(uint256)
func (_Peth *PethCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Peth.contract.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address src) constant returns(uint256)
func (_Peth *PethSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _Peth.Contract.BalanceOf(&_Peth.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address src) constant returns(uint256)
func (_Peth *PethCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _Peth.Contract.BalanceOf(&_Peth.CallOpts, src)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Peth *PethCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Peth.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Peth *PethSession) Decimals() (*big.Int, error) {
	return _Peth.Contract.Decimals(&_Peth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Peth *PethCallerSession) Decimals() (*big.Int, error) {
	return _Peth.Contract.Decimals(&_Peth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_Peth *PethCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Peth.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_Peth *PethSession) Name() ([32]byte, error) {
	return _Peth.Contract.Name(&_Peth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_Peth *PethCallerSession) Name() ([32]byte, error) {
	return _Peth.Contract.Name(&_Peth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Peth *PethCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Peth.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Peth *PethSession) Owner() (common.Address, error) {
	return _Peth.Contract.Owner(&_Peth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Peth *PethCallerSession) Owner() (common.Address, error) {
	return _Peth.Contract.Owner(&_Peth.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Peth *PethCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Peth.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Peth *PethSession) Stopped() (bool, error) {
	return _Peth.Contract.Stopped(&_Peth.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Peth *PethCallerSession) Stopped() (bool, error) {
	return _Peth.Contract.Stopped(&_Peth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_Peth *PethCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Peth.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_Peth *PethSession) Symbol() ([32]byte, error) {
	return _Peth.Contract.Symbol(&_Peth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_Peth *PethCallerSession) Symbol() ([32]byte, error) {
	return _Peth.Contract.Symbol(&_Peth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Peth *PethCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Peth.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Peth *PethSession) TotalSupply() (*big.Int, error) {
	return _Peth.Contract.TotalSupply(&_Peth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Peth *PethCallerSession) TotalSupply() (*big.Int, error) {
	return _Peth.Contract.TotalSupply(&_Peth.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Peth *PethTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Peth *PethSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Approve(&_Peth.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Peth *PethTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Approve(&_Peth.TransactOpts, guy, wad)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address guy) returns(bool)
func (_Peth *PethTransactor) Approve0(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "approve0", guy)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address guy) returns(bool)
func (_Peth *PethSession) Approve0(guy common.Address) (*types.Transaction, error) {
	return _Peth.Contract.Approve0(&_Peth.TransactOpts, guy)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address guy) returns(bool)
func (_Peth *PethTransactorSession) Approve0(guy common.Address) (*types.Transaction, error) {
	return _Peth.Contract.Approve0(&_Peth.TransactOpts, guy)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 wad) returns()
func (_Peth *PethTransactor) Burn(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "burn", wad)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 wad) returns()
func (_Peth *PethSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Burn(&_Peth.TransactOpts, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 wad) returns()
func (_Peth *PethTransactorSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Burn(&_Peth.TransactOpts, wad)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address guy, uint256 wad) returns()
func (_Peth *PethTransactor) Burn0(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "burn0", guy, wad)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address guy, uint256 wad) returns()
func (_Peth *PethSession) Burn0(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Burn0(&_Peth.TransactOpts, guy, wad)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address guy, uint256 wad) returns()
func (_Peth *PethTransactorSession) Burn0(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Burn0(&_Peth.TransactOpts, guy, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address guy, uint256 wad) returns()
func (_Peth *PethTransactor) Mint(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "mint", guy, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address guy, uint256 wad) returns()
func (_Peth *PethSession) Mint(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Mint(&_Peth.TransactOpts, guy, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address guy, uint256 wad) returns()
func (_Peth *PethTransactorSession) Mint(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Mint(&_Peth.TransactOpts, guy, wad)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 wad) returns()
func (_Peth *PethTransactor) Mint0(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "mint0", wad)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 wad) returns()
func (_Peth *PethSession) Mint0(wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Mint0(&_Peth.TransactOpts, wad)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 wad) returns()
func (_Peth *PethTransactorSession) Mint0(wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Mint0(&_Peth.TransactOpts, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Peth *PethTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "move", src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Peth *PethSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Move(&_Peth.TransactOpts, src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Peth *PethTransactorSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Move(&_Peth.TransactOpts, src, dst, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address src, uint256 wad) returns()
func (_Peth *PethTransactor) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "pull", src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address src, uint256 wad) returns()
func (_Peth *PethSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Pull(&_Peth.TransactOpts, src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address src, uint256 wad) returns()
func (_Peth *PethTransactorSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Pull(&_Peth.TransactOpts, src, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address dst, uint256 wad) returns()
func (_Peth *PethTransactor) Push(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "push", dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address dst, uint256 wad) returns()
func (_Peth *PethSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Push(&_Peth.TransactOpts, dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address dst, uint256 wad) returns()
func (_Peth *PethTransactorSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Push(&_Peth.TransactOpts, dst, wad)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Peth *PethTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Peth *PethSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Peth.Contract.SetAuthority(&_Peth.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Peth *PethTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Peth.Contract.SetAuthority(&_Peth.TransactOpts, authority_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(bytes32 name_) returns()
func (_Peth *PethTransactor) SetName(opts *bind.TransactOpts, name_ [32]byte) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(bytes32 name_) returns()
func (_Peth *PethSession) SetName(name_ [32]byte) (*types.Transaction, error) {
	return _Peth.Contract.SetName(&_Peth.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(bytes32 name_) returns()
func (_Peth *PethTransactorSession) SetName(name_ [32]byte) (*types.Transaction, error) {
	return _Peth.Contract.SetName(&_Peth.TransactOpts, name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Peth *PethTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Peth *PethSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Peth.Contract.SetOwner(&_Peth.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Peth *PethTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Peth.Contract.SetOwner(&_Peth.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Peth *PethTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Peth *PethSession) Start() (*types.Transaction, error) {
	return _Peth.Contract.Start(&_Peth.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Peth *PethTransactorSession) Start() (*types.Transaction, error) {
	return _Peth.Contract.Start(&_Peth.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Peth *PethTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Peth *PethSession) Stop() (*types.Transaction, error) {
	return _Peth.Contract.Stop(&_Peth.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Peth *PethTransactorSession) Stop() (*types.Transaction, error) {
	return _Peth.Contract.Stop(&_Peth.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Peth *PethTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Peth *PethSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Transfer(&_Peth.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Peth *PethTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.Transfer(&_Peth.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Peth *PethTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Peth *PethSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.TransferFrom(&_Peth.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Peth *PethTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Peth.Contract.TransferFrom(&_Peth.TransactOpts, src, dst, wad)
}

// PethApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Peth contract.
type PethApprovalIterator struct {
	Event *PethApproval // Event containing the contract specifics and raw log

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
func (it *PethApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PethApproval)
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
		it.Event = new(PethApproval)
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
func (it *PethApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PethApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PethApproval represents a Approval event raised by the Peth contract.
type PethApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Peth *PethFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*PethApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Peth.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &PethApprovalIterator{contract: _Peth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Peth *PethFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PethApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Peth.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PethApproval)
				if err := _Peth.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Peth *PethFilterer) ParseApproval(log types.Log) (*PethApproval, error) {
	event := new(PethApproval)
	if err := _Peth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PethBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Peth contract.
type PethBurnIterator struct {
	Event *PethBurn // Event containing the contract specifics and raw log

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
func (it *PethBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PethBurn)
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
		it.Event = new(PethBurn)
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
func (it *PethBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PethBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PethBurn represents a Burn event raised by the Peth contract.
type PethBurn struct {
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed guy, uint256 wad)
func (_Peth *PethFilterer) FilterBurn(opts *bind.FilterOpts, guy []common.Address) (*PethBurnIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Peth.contract.FilterLogs(opts, "Burn", guyRule)
	if err != nil {
		return nil, err
	}
	return &PethBurnIterator{contract: _Peth.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed guy, uint256 wad)
func (_Peth *PethFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PethBurn, guy []common.Address) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Peth.contract.WatchLogs(opts, "Burn", guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PethBurn)
				if err := _Peth.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed guy, uint256 wad)
func (_Peth *PethFilterer) ParseBurn(log types.Log) (*PethBurn, error) {
	event := new(PethBurn)
	if err := _Peth.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PethLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Peth contract.
type PethLogSetAuthorityIterator struct {
	Event *PethLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *PethLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PethLogSetAuthority)
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
		it.Event = new(PethLogSetAuthority)
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
func (it *PethLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PethLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PethLogSetAuthority represents a LogSetAuthority event raised by the Peth contract.
type PethLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Peth *PethFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*PethLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Peth.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &PethLogSetAuthorityIterator{contract: _Peth.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Peth *PethFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *PethLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Peth.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PethLogSetAuthority)
				if err := _Peth.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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
func (_Peth *PethFilterer) ParseLogSetAuthority(log types.Log) (*PethLogSetAuthority, error) {
	event := new(PethLogSetAuthority)
	if err := _Peth.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PethLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Peth contract.
type PethLogSetOwnerIterator struct {
	Event *PethLogSetOwner // Event containing the contract specifics and raw log

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
func (it *PethLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PethLogSetOwner)
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
		it.Event = new(PethLogSetOwner)
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
func (it *PethLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PethLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PethLogSetOwner represents a LogSetOwner event raised by the Peth contract.
type PethLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Peth *PethFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*PethLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Peth.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &PethLogSetOwnerIterator{contract: _Peth.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Peth *PethFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *PethLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Peth.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PethLogSetOwner)
				if err := _Peth.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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
func (_Peth *PethFilterer) ParseLogSetOwner(log types.Log) (*PethLogSetOwner, error) {
	event := new(PethLogSetOwner)
	if err := _Peth.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PethMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Peth contract.
type PethMintIterator struct {
	Event *PethMint // Event containing the contract specifics and raw log

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
func (it *PethMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PethMint)
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
		it.Event = new(PethMint)
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
func (it *PethMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PethMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PethMint represents a Mint event raised by the Peth contract.
type PethMint struct {
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed guy, uint256 wad)
func (_Peth *PethFilterer) FilterMint(opts *bind.FilterOpts, guy []common.Address) (*PethMintIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Peth.contract.FilterLogs(opts, "Mint", guyRule)
	if err != nil {
		return nil, err
	}
	return &PethMintIterator{contract: _Peth.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed guy, uint256 wad)
func (_Peth *PethFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PethMint, guy []common.Address) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Peth.contract.WatchLogs(opts, "Mint", guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PethMint)
				if err := _Peth.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed guy, uint256 wad)
func (_Peth *PethFilterer) ParseMint(log types.Log) (*PethMint, error) {
	event := new(PethMint)
	if err := _Peth.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PethTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Peth contract.
type PethTransferIterator struct {
	Event *PethTransfer // Event containing the contract specifics and raw log

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
func (it *PethTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PethTransfer)
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
		it.Event = new(PethTransfer)
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
func (it *PethTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PethTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PethTransfer represents a Transfer event raised by the Peth contract.
type PethTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Peth *PethFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*PethTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Peth.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &PethTransferIterator{contract: _Peth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Peth *PethFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PethTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Peth.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PethTransfer)
				if err := _Peth.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Peth *PethFilterer) ParseTransfer(log types.Log) (*PethTransfer, error) {
	event := new(PethTransfer)
	if err := _Peth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
