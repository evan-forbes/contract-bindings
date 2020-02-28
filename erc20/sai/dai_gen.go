// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sai

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

// SaiABI is the input ABI used to generate the binding from.
const SaiABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"bytes32\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"symbol_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

const HexAddr = "0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359"

// SaiBin is the compiled bytecode used for deploying new contracts.
const SaiBin = `606060405260126006556000600755341561001957600080fd5b604051602080610dd283398101604052808051600160a060020a03331660008181526001602052604080822082905590805560048054600160a060020a0319168317905591935091507fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94905160405180910390a2600555610d338061009f6000396000f3006060604052600436106101485763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461014d57806307da68f514610172578063095ea7b31461018757806313af4035146101bd57806318160ddd146101dc57806323b872dd146101ef578063313ce5671461021757806340c10f191461022a57806342966c681461024c5780635ac801fe1461026257806370a082311461027857806375f12b21146102975780637a9e5e4b146102aa5780638da5cb5b146102c957806395d89b41146102f85780639dc29fac1461030b578063a0712d681461032d578063a9059cbb14610343578063b753a98c14610365578063bb35783b14610387578063be9a6555146103af578063bf7e214f146103c2578063daea85c5146103d5578063dd62ed3e146103f4578063f2d5d56b14610419575b600080fd5b341561015857600080fd5b61016061043b565b60405190815260200160405180910390f35b341561017d57600080fd5b610185610441565b005b341561019257600080fd5b6101a9600160a060020a03600435166024356104e0565b604051901515815260200160405180910390f35b34156101c857600080fd5b610185600160a060020a036004351661050d565b34156101e757600080fd5b61016061058c565b34156101fa57600080fd5b6101a9600160a060020a0360043581169060243516604435610592565b341561022257600080fd5b610160610707565b341561023557600080fd5b610185600160a060020a036004351660243561070d565b341561025757600080fd5b6101856004356107d3565b341561026d57600080fd5b6101856004356107e0565b341561028357600080fd5b610160600160a060020a0360043516610806565b34156102a257600080fd5b6101a9610821565b34156102b557600080fd5b610185600160a060020a0360043516610831565b34156102d457600080fd5b6102dc6108b0565b604051600160a060020a03909116815260200160405180910390f35b341561030357600080fd5b6101606108bf565b341561031657600080fd5b610185600160a060020a03600435166024356108c5565b341561033857600080fd5b610185600435610a33565b341561034e57600080fd5b6101a9600160a060020a0360043516602435610a3d565b341561037057600080fd5b610185600160a060020a0360043516602435610a4a565b341561039257600080fd5b610185600160a060020a0360043581169060243516604435610a5a565b34156103ba57600080fd5b610185610a6b565b34156103cd57600080fd5b6102dc610b04565b34156103e057600080fd5b6101a9600160a060020a0360043516610b13565b34156103ff57600080fd5b610160600160a060020a0360043581169060243516610b39565b341561042457600080fd5b610185600160a060020a0360043516602435610b64565b60075481565b61045733600035600160e060020a031916610b6f565b151561046257600080fd5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a450506004805474ff0000000000000000000000000000000000000000191660a060020a179055565b60045460009060a060020a900460ff16156104fa57600080fd5b6105048383610c7b565b90505b92915050565b61052333600035600160e060020a031916610b6f565b151561052e57600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a250565b60005490565b60045460009060a060020a900460ff16156105ac57600080fd5b33600160a060020a031684600160a060020a0316141580156105f65750600160a060020a038085166000908152600260209081526040808320339094168352929052205460001914155b1561065457600160a060020a038085166000908152600260209081526040808320339094168352929052205461062c9083610ce7565b600160a060020a03808616600090815260026020908152604080832033909416835292905220555b600160a060020a0384166000908152600160205260409020546106779083610ce7565b600160a060020a0380861660009081526001602052604080822093909355908516815220546106a69083610cf7565b600160a060020a03808516600081815260016020526040908190209390935591908616907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60065481565b61072333600035600160e060020a031916610b6f565b151561072e57600080fd5b60045460a060020a900460ff161561074557600080fd5b600160a060020a0382166000908152600160205260409020546107689082610cf7565b600160a060020a0383166000908152600160205260408120919091555461078f9082610cf7565b600055600160a060020a0382167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405190815260200160405180910390a25050565b6107dd33826108c5565b50565b6107f633600035600160e060020a031916610b6f565b151561080157600080fd5b600755565b600160a060020a031660009081526001602052604090205490565b60045460a060020a900460ff1681565b61084733600035600160e060020a031916610b6f565b151561085257600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a250565b600454600160a060020a031681565b60055481565b6108db33600035600160e060020a031916610b6f565b15156108e657600080fd5b60045460a060020a900460ff16156108fd57600080fd5b33600160a060020a031682600160a060020a0316141580156109475750600160a060020a038083166000908152600260209081526040808320339094168352929052205460001914155b156109a557600160a060020a038083166000908152600260209081526040808320339094168352929052205461097d9082610ce7565b600160a060020a03808416600090815260026020908152604080832033909416835292905220555b600160a060020a0382166000908152600160205260409020546109c89082610ce7565b600160a060020a038316600090815260016020526040812091909155546109ef9082610ce7565b600055600160a060020a0382167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405190815260200160405180910390a25050565b6107dd338261070d565b6000610504338484610592565b610a55338383610592565b505050565b610a65838383610592565b50505050565b610a8133600035600160e060020a031916610b6f565b1515610a8c57600080fd5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a450506004805474ff000000000000000000000000000000000000000019169055565b600354600160a060020a031681565b60045460009060a060020a900460ff1615610b2d57600080fd5b61050782600019610c7b565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b610a55823383610592565b600030600160a060020a031683600160a060020a03161415610b9357506001610507565b600454600160a060020a0384811691161415610bb157506001610507565b600354600160a060020a03161515610bcb57506000610507565b600354600160a060020a031663b70096138430856000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b1515610c5957600080fd5b6102c65a03f11515610c6a57600080fd5b505050604051805190509050610507565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b8082038281111561050757600080fd5b8082018281101561050757600080fd00a165627a7a723058200877df264aa5d498c61a45dfe4fc04c68d11820448cf0cc7a275283a271173b400294441490000000000000000000000000000000000000000000000000000000000`

// DeploySai deploys a new Ethereum contract, binding an instance of Sai to it.
func DeploySai(auth *bind.TransactOpts, backend bind.ContractBackend, symbol_ [32]byte) (common.Address, *types.Transaction, *Sai, error) {
	parsed, err := abi.JSON(strings.NewReader(SaiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SaiBin), backend, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sai{SaiCaller: SaiCaller{contract: contract}, SaiTransactor: SaiTransactor{contract: contract}, SaiFilterer: SaiFilterer{contract: contract}}, nil
}

// Sai is an auto generated Go binding around an Ethereum contract.
type Sai struct {
	SaiCaller     // Read-only binding to the contract
	SaiTransactor // Write-only binding to the contract
	SaiFilterer   // Log filterer for contract events
}

// SaiCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaiSession struct {
	Contract     *Sai              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaiCallerSession struct {
	Contract *SaiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SaiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaiTransactorSession struct {
	Contract     *SaiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaiRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaiRaw struct {
	Contract *Sai // Generic contract binding to access the raw methods on
}

// SaiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaiCallerRaw struct {
	Contract *SaiCaller // Generic read-only contract binding to access the raw methods on
}

// SaiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaiTransactorRaw struct {
	Contract *SaiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSai creates a new instance of Sai, bound to a specific deployed contract.
func NewSai(address common.Address, backend bind.ContractBackend) (*Sai, error) {
	contract, err := bindSai(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sai{SaiCaller: SaiCaller{contract: contract}, SaiTransactor: SaiTransactor{contract: contract}, SaiFilterer: SaiFilterer{contract: contract}}, nil
}

// NewSaiCaller creates a new read-only instance of Sai, bound to a specific deployed contract.
func NewSaiCaller(address common.Address, caller bind.ContractCaller) (*SaiCaller, error) {
	contract, err := bindSai(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaiCaller{contract: contract}, nil
}

// NewSaiTransactor creates a new write-only instance of Sai, bound to a specific deployed contract.
func NewSaiTransactor(address common.Address, transactor bind.ContractTransactor) (*SaiTransactor, error) {
	contract, err := bindSai(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaiTransactor{contract: contract}, nil
}

// NewSaiFilterer creates a new log filterer instance of Sai, bound to a specific deployed contract.
func NewSaiFilterer(address common.Address, filterer bind.ContractFilterer) (*SaiFilterer, error) {
	contract, err := bindSai(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaiFilterer{contract: contract}, nil
}

// bindSai binds a generic wrapper to an already deployed contract.
func bindSai(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SaiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sai *SaiRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sai.Contract.SaiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sai *SaiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sai.Contract.SaiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sai *SaiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sai.Contract.SaiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sai *SaiCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sai.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sai *SaiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sai.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sai *SaiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sai.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address guy) constant returns(uint256)
func (_Sai *SaiCaller) Allowance(opts *bind.CallOpts, src common.Address, guy common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sai.contract.Call(opts, out, "allowance", src, guy)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address guy) constant returns(uint256)
func (_Sai *SaiSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _Sai.Contract.Allowance(&_Sai.CallOpts, src, guy)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address guy) constant returns(uint256)
func (_Sai *SaiCallerSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _Sai.Contract.Allowance(&_Sai.CallOpts, src, guy)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Sai *SaiCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sai.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Sai *SaiSession) Authority() (common.Address, error) {
	return _Sai.Contract.Authority(&_Sai.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Sai *SaiCallerSession) Authority() (common.Address, error) {
	return _Sai.Contract.Authority(&_Sai.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address src) constant returns(uint256)
func (_Sai *SaiCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sai.contract.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address src) constant returns(uint256)
func (_Sai *SaiSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _Sai.Contract.BalanceOf(&_Sai.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address src) constant returns(uint256)
func (_Sai *SaiCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _Sai.Contract.BalanceOf(&_Sai.CallOpts, src)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Sai *SaiCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sai.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Sai *SaiSession) Decimals() (*big.Int, error) {
	return _Sai.Contract.Decimals(&_Sai.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Sai *SaiCallerSession) Decimals() (*big.Int, error) {
	return _Sai.Contract.Decimals(&_Sai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_Sai *SaiCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Sai.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_Sai *SaiSession) Name() ([32]byte, error) {
	return _Sai.Contract.Name(&_Sai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_Sai *SaiCallerSession) Name() ([32]byte, error) {
	return _Sai.Contract.Name(&_Sai.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Sai *SaiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sai.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Sai *SaiSession) Owner() (common.Address, error) {
	return _Sai.Contract.Owner(&_Sai.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Sai *SaiCallerSession) Owner() (common.Address, error) {
	return _Sai.Contract.Owner(&_Sai.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Sai *SaiCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Sai.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Sai *SaiSession) Stopped() (bool, error) {
	return _Sai.Contract.Stopped(&_Sai.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Sai *SaiCallerSession) Stopped() (bool, error) {
	return _Sai.Contract.Stopped(&_Sai.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_Sai *SaiCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Sai.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_Sai *SaiSession) Symbol() ([32]byte, error) {
	return _Sai.Contract.Symbol(&_Sai.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_Sai *SaiCallerSession) Symbol() ([32]byte, error) {
	return _Sai.Contract.Symbol(&_Sai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Sai *SaiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sai.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Sai *SaiSession) TotalSupply() (*big.Int, error) {
	return _Sai.Contract.TotalSupply(&_Sai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Sai *SaiCallerSession) TotalSupply() (*big.Int, error) {
	return _Sai.Contract.TotalSupply(&_Sai.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address guy) returns(bool)
func (_Sai *SaiTransactor) Approve(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "approve", guy)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address guy) returns(bool)
func (_Sai *SaiSession) Approve(guy common.Address) (*types.Transaction, error) {
	return _Sai.Contract.Approve(&_Sai.TransactOpts, guy)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address guy) returns(bool)
func (_Sai *SaiTransactorSession) Approve(guy common.Address) (*types.Transaction, error) {
	return _Sai.Contract.Approve(&_Sai.TransactOpts, guy)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address guy, uint256 wad) returns()
func (_Sai *SaiTransactor) Burn(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "burn", guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address guy, uint256 wad) returns()
func (_Sai *SaiSession) Burn(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.Burn(&_Sai.TransactOpts, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address guy, uint256 wad) returns()
func (_Sai *SaiTransactorSession) Burn(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.Burn(&_Sai.TransactOpts, guy, wad)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 wad) returns()
func (_Sai *SaiTransactor) Mint(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "mint", wad)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 wad) returns()
func (_Sai *SaiSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.Mint(&_Sai.TransactOpts, wad)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 wad) returns()
func (_Sai *SaiTransactorSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.Mint(&_Sai.TransactOpts, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Sai *SaiTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "move", src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Sai *SaiSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.Move(&_Sai.TransactOpts, src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Sai *SaiTransactorSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.Move(&_Sai.TransactOpts, src, dst, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address src, uint256 wad) returns()
func (_Sai *SaiTransactor) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "pull", src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address src, uint256 wad) returns()
func (_Sai *SaiSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.Pull(&_Sai.TransactOpts, src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address src, uint256 wad) returns()
func (_Sai *SaiTransactorSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.Pull(&_Sai.TransactOpts, src, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address dst, uint256 wad) returns()
func (_Sai *SaiTransactor) Push(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "push", dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address dst, uint256 wad) returns()
func (_Sai *SaiSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.Push(&_Sai.TransactOpts, dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address dst, uint256 wad) returns()
func (_Sai *SaiTransactorSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.Push(&_Sai.TransactOpts, dst, wad)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Sai *SaiTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Sai *SaiSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Sai.Contract.SetAuthority(&_Sai.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Sai *SaiTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Sai.Contract.SetAuthority(&_Sai.TransactOpts, authority_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(bytes32 name_) returns()
func (_Sai *SaiTransactor) SetName(opts *bind.TransactOpts, name_ [32]byte) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(bytes32 name_) returns()
func (_Sai *SaiSession) SetName(name_ [32]byte) (*types.Transaction, error) {
	return _Sai.Contract.SetName(&_Sai.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(bytes32 name_) returns()
func (_Sai *SaiTransactorSession) SetName(name_ [32]byte) (*types.Transaction, error) {
	return _Sai.Contract.SetName(&_Sai.TransactOpts, name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Sai *SaiTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Sai *SaiSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Sai.Contract.SetOwner(&_Sai.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Sai *SaiTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Sai.Contract.SetOwner(&_Sai.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Sai *SaiTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Sai *SaiSession) Start() (*types.Transaction, error) {
	return _Sai.Contract.Start(&_Sai.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Sai *SaiTransactorSession) Start() (*types.Transaction, error) {
	return _Sai.Contract.Start(&_Sai.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Sai *SaiTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Sai *SaiSession) Stop() (*types.Transaction, error) {
	return _Sai.Contract.Stop(&_Sai.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Sai *SaiTransactorSession) Stop() (*types.Transaction, error) {
	return _Sai.Contract.Stop(&_Sai.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Sai *SaiTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Sai *SaiSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.Transfer(&_Sai.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Sai *SaiTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.Transfer(&_Sai.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Sai *SaiTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Sai *SaiSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.TransferFrom(&_Sai.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Sai *SaiTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Contract.TransferFrom(&_Sai.TransactOpts, src, dst, wad)
}

// SaiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Sai contract.
type SaiApprovalIterator struct {
	Event *SaiApproval // Event containing the contract specifics and raw log

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
func (it *SaiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaiApproval)
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
		it.Event = new(SaiApproval)
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
func (it *SaiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaiApproval represents a Approval event raised by the Sai contract.
type SaiApproval struct {
	Src common.Address `json:"_from"`
	Guy common.Address `json:"_to"`
	Wad *big.Int       `json:"amount"`
	Raw types.Log      `json:"logData"` // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Sai *SaiFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*SaiApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Sai.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &SaiApprovalIterator{contract: _Sai.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Sai *SaiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SaiApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Sai.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaiApproval)
				if err := _Sai.contract.UnpackLog(event, "Approval", log); err != nil {
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

// SaiBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Sai contract.
type SaiBurnIterator struct {
	Event *SaiBurn // Event containing the contract specifics and raw log

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
func (it *SaiBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaiBurn)
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
		it.Event = new(SaiBurn)
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
func (it *SaiBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaiBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaiBurn represents a Burn event raised by the Sai contract.
type SaiBurn struct {
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed guy, uint256 wad)
func (_Sai *SaiFilterer) FilterBurn(opts *bind.FilterOpts, guy []common.Address) (*SaiBurnIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Sai.contract.FilterLogs(opts, "Burn", guyRule)
	if err != nil {
		return nil, err
	}
	return &SaiBurnIterator{contract: _Sai.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed guy, uint256 wad)
func (_Sai *SaiFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *SaiBurn, guy []common.Address) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Sai.contract.WatchLogs(opts, "Burn", guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaiBurn)
				if err := _Sai.contract.UnpackLog(event, "Burn", log); err != nil {
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

// SaiLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Sai contract.
type SaiLogSetAuthorityIterator struct {
	Event *SaiLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *SaiLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaiLogSetAuthority)
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
		it.Event = new(SaiLogSetAuthority)
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
func (it *SaiLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaiLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaiLogSetAuthority represents a LogSetAuthority event raised by the Sai contract.
type SaiLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Sai *SaiFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*SaiLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Sai.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &SaiLogSetAuthorityIterator{contract: _Sai.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Sai *SaiFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *SaiLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Sai.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaiLogSetAuthority)
				if err := _Sai.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// SaiLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Sai contract.
type SaiLogSetOwnerIterator struct {
	Event *SaiLogSetOwner // Event containing the contract specifics and raw log

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
func (it *SaiLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaiLogSetOwner)
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
		it.Event = new(SaiLogSetOwner)
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
func (it *SaiLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaiLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaiLogSetOwner represents a LogSetOwner event raised by the Sai contract.
type SaiLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Sai *SaiFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*SaiLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Sai.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SaiLogSetOwnerIterator{contract: _Sai.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Sai *SaiFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *SaiLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Sai.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaiLogSetOwner)
				if err := _Sai.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// SaiMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Sai contract.
type SaiMintIterator struct {
	Event *SaiMint // Event containing the contract specifics and raw log

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
func (it *SaiMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaiMint)
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
		it.Event = new(SaiMint)
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
func (it *SaiMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaiMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaiMint represents a Mint event raised by the Sai contract.
type SaiMint struct {
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed guy, uint256 wad)
func (_Sai *SaiFilterer) FilterMint(opts *bind.FilterOpts, guy []common.Address) (*SaiMintIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Sai.contract.FilterLogs(opts, "Mint", guyRule)
	if err != nil {
		return nil, err
	}
	return &SaiMintIterator{contract: _Sai.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed guy, uint256 wad)
func (_Sai *SaiFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *SaiMint, guy []common.Address) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Sai.contract.WatchLogs(opts, "Mint", guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaiMint)
				if err := _Sai.contract.UnpackLog(event, "Mint", log); err != nil {
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

// SaiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Sai contract.
type SaiTransferIterator struct {
	Event *SaiTransfer // Event containing the contract specifics and raw log

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
func (it *SaiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaiTransfer)
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
		it.Event = new(SaiTransfer)
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
func (it *SaiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaiTransfer represents a Transfer event raised by the Sai contract.
type SaiTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Sai *SaiFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*SaiTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Sai.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &SaiTransferIterator{contract: _Sai.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Sai *SaiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SaiTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Sai.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaiTransfer)
				if err := _Sai.contract.UnpackLog(event, "Transfer", log); err != nil {
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
