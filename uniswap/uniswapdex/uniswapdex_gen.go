// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapdex

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

// UniswapdexABI is the input ABI used to generate the binding from.
const UniswapdexABI = "[{\"name\":\"TokenPurchase\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_sold\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"EthPurchase\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_bought\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_amount\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"token_amount\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_amount\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"token_amount\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"setup\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":175875},{\"name\":\"addLiquidity\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"min_liquidity\"},{\"type\":\"uint256\",\"name\":\"max_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":82605},{\"name\":\"removeLiquidity\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\"},{\"type\":\"uint256\",\"name\":\"min_eth\"},{\"type\":\"uint256\",\"name\":\"min_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":116814},{\"name\":\"__default__\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":true,\"type\":\"function\"},{\"name\":\"ethToTokenSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"min_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":12757},{\"name\":\"ethToTokenTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"min_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":12965},{\"name\":\"ethToTokenSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":50455},{\"name\":\"ethToTokenTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":50663},{\"name\":\"tokenToEthSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_eth\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":47503},{\"name\":\"tokenToEthTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_eth\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":47712},{\"name\":\"tokenToEthSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":50175},{\"name\":\"tokenToEthTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":50384},{\"name\":\"tokenToTokenSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":51007},{\"name\":\"tokenToTokenTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":51098},{\"name\":\"tokenToTokenSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":54928},{\"name\":\"tokenToTokenTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":55019},{\"name\":\"tokenToExchangeSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":49342},{\"name\":\"tokenToExchangeTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":49532},{\"name\":\"tokenToExchangeSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":53233},{\"name\":\"tokenToExchangeTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":53423},{\"name\":\"getEthToTokenInputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_sold\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":5542},{\"name\":\"getEthToTokenOutputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6872},{\"name\":\"getTokenToEthInputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":5637},{\"name\":\"getTokenToEthOutputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_bought\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6897},{\"name\":\"tokenAddress\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1413},{\"name\":\"factoryAddress\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1443},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1645},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":75034},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":110907},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":38769},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"address\",\"name\":\"_spender\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1925},{\"name\":\"name\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1623},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1653},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1683},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1713}]"

// Uniswapdex is an auto generated Go binding around an Ethereum contract.
type Uniswapdex struct {
	UniswapdexCaller     // Read-only binding to the contract
	UniswapdexTransactor // Write-only binding to the contract
	UniswapdexFilterer   // Log filterer for contract events
}

// UniswapdexCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapdexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapdexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapdexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapdexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapdexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapdexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapdexSession struct {
	Contract     *Uniswapdex       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapdexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapdexCallerSession struct {
	Contract *UniswapdexCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UniswapdexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapdexTransactorSession struct {
	Contract     *UniswapdexTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UniswapdexRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapdexRaw struct {
	Contract *Uniswapdex // Generic contract binding to access the raw methods on
}

// UniswapdexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapdexCallerRaw struct {
	Contract *UniswapdexCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapdexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapdexTransactorRaw struct {
	Contract *UniswapdexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapdex creates a new instance of Uniswapdex, bound to a specific deployed contract.
func NewUniswapdex(address common.Address, backend bind.ContractBackend) (*Uniswapdex, error) {
	contract, err := bindUniswapdex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapdex{UniswapdexCaller: UniswapdexCaller{contract: contract}, UniswapdexTransactor: UniswapdexTransactor{contract: contract}, UniswapdexFilterer: UniswapdexFilterer{contract: contract}}, nil
}

// NewUniswapdexCaller creates a new read-only instance of Uniswapdex, bound to a specific deployed contract.
func NewUniswapdexCaller(address common.Address, caller bind.ContractCaller) (*UniswapdexCaller, error) {
	contract, err := bindUniswapdex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapdexCaller{contract: contract}, nil
}

// NewUniswapdexTransactor creates a new write-only instance of Uniswapdex, bound to a specific deployed contract.
func NewUniswapdexTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapdexTransactor, error) {
	contract, err := bindUniswapdex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapdexTransactor{contract: contract}, nil
}

// NewUniswapdexFilterer creates a new log filterer instance of Uniswapdex, bound to a specific deployed contract.
func NewUniswapdexFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapdexFilterer, error) {
	contract, err := bindUniswapdex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapdexFilterer{contract: contract}, nil
}

// bindUniswapdex binds a generic wrapper to an already deployed contract.
func bindUniswapdex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapdexABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapdex *UniswapdexRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Uniswapdex.Contract.UniswapdexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapdex *UniswapdexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapdex.Contract.UniswapdexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapdex *UniswapdexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapdex.Contract.UniswapdexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapdex *UniswapdexCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Uniswapdex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapdex *UniswapdexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapdex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapdex *UniswapdexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapdex.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswapdex.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Uniswapdex.Contract.Allowance(&_Uniswapdex.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Uniswapdex.Contract.Allowance(&_Uniswapdex.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswapdex.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Uniswapdex.Contract.BalanceOf(&_Uniswapdex.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Uniswapdex.Contract.BalanceOf(&_Uniswapdex.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswapdex.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) Decimals() (*big.Int, error) {
	return _Uniswapdex.Contract.Decimals(&_Uniswapdex.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCallerSession) Decimals() (*big.Int, error) {
	return _Uniswapdex.Contract.Decimals(&_Uniswapdex.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() constant returns(address out)
func (_Uniswapdex *UniswapdexCaller) FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Uniswapdex.contract.Call(opts, out, "factoryAddress")
	return *ret0, err
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() constant returns(address out)
func (_Uniswapdex *UniswapdexSession) FactoryAddress() (common.Address, error) {
	return _Uniswapdex.Contract.FactoryAddress(&_Uniswapdex.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() constant returns(address out)
func (_Uniswapdex *UniswapdexCallerSession) FactoryAddress() (common.Address, error) {
	return _Uniswapdex.Contract.FactoryAddress(&_Uniswapdex.CallOpts)
}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 eth_sold) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCaller) GetEthToTokenInputPrice(opts *bind.CallOpts, eth_sold *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswapdex.contract.Call(opts, out, "getEthToTokenInputPrice", eth_sold)
	return *ret0, err
}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 eth_sold) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) GetEthToTokenInputPrice(eth_sold *big.Int) (*big.Int, error) {
	return _Uniswapdex.Contract.GetEthToTokenInputPrice(&_Uniswapdex.CallOpts, eth_sold)
}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 eth_sold) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCallerSession) GetEthToTokenInputPrice(eth_sold *big.Int) (*big.Int, error) {
	return _Uniswapdex.Contract.GetEthToTokenInputPrice(&_Uniswapdex.CallOpts, eth_sold)
}

// GetEthToTokenOutputPrice is a free data retrieval call binding the contract method 0x59e94862.
//
// Solidity: function getEthToTokenOutputPrice(uint256 tokens_bought) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCaller) GetEthToTokenOutputPrice(opts *bind.CallOpts, tokens_bought *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswapdex.contract.Call(opts, out, "getEthToTokenOutputPrice", tokens_bought)
	return *ret0, err
}

// GetEthToTokenOutputPrice is a free data retrieval call binding the contract method 0x59e94862.
//
// Solidity: function getEthToTokenOutputPrice(uint256 tokens_bought) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) GetEthToTokenOutputPrice(tokens_bought *big.Int) (*big.Int, error) {
	return _Uniswapdex.Contract.GetEthToTokenOutputPrice(&_Uniswapdex.CallOpts, tokens_bought)
}

// GetEthToTokenOutputPrice is a free data retrieval call binding the contract method 0x59e94862.
//
// Solidity: function getEthToTokenOutputPrice(uint256 tokens_bought) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCallerSession) GetEthToTokenOutputPrice(tokens_bought *big.Int) (*big.Int, error) {
	return _Uniswapdex.Contract.GetEthToTokenOutputPrice(&_Uniswapdex.CallOpts, tokens_bought)
}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokens_sold) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCaller) GetTokenToEthInputPrice(opts *bind.CallOpts, tokens_sold *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswapdex.contract.Call(opts, out, "getTokenToEthInputPrice", tokens_sold)
	return *ret0, err
}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokens_sold) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) GetTokenToEthInputPrice(tokens_sold *big.Int) (*big.Int, error) {
	return _Uniswapdex.Contract.GetTokenToEthInputPrice(&_Uniswapdex.CallOpts, tokens_sold)
}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokens_sold) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCallerSession) GetTokenToEthInputPrice(tokens_sold *big.Int) (*big.Int, error) {
	return _Uniswapdex.Contract.GetTokenToEthInputPrice(&_Uniswapdex.CallOpts, tokens_sold)
}

// GetTokenToEthOutputPrice is a free data retrieval call binding the contract method 0x2640f62c.
//
// Solidity: function getTokenToEthOutputPrice(uint256 eth_bought) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCaller) GetTokenToEthOutputPrice(opts *bind.CallOpts, eth_bought *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswapdex.contract.Call(opts, out, "getTokenToEthOutputPrice", eth_bought)
	return *ret0, err
}

// GetTokenToEthOutputPrice is a free data retrieval call binding the contract method 0x2640f62c.
//
// Solidity: function getTokenToEthOutputPrice(uint256 eth_bought) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) GetTokenToEthOutputPrice(eth_bought *big.Int) (*big.Int, error) {
	return _Uniswapdex.Contract.GetTokenToEthOutputPrice(&_Uniswapdex.CallOpts, eth_bought)
}

// GetTokenToEthOutputPrice is a free data retrieval call binding the contract method 0x2640f62c.
//
// Solidity: function getTokenToEthOutputPrice(uint256 eth_bought) constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCallerSession) GetTokenToEthOutputPrice(eth_bought *big.Int) (*big.Int, error) {
	return _Uniswapdex.Contract.GetTokenToEthOutputPrice(&_Uniswapdex.CallOpts, eth_bought)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32 out)
func (_Uniswapdex *UniswapdexCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Uniswapdex.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32 out)
func (_Uniswapdex *UniswapdexSession) Name() ([32]byte, error) {
	return _Uniswapdex.Contract.Name(&_Uniswapdex.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32 out)
func (_Uniswapdex *UniswapdexCallerSession) Name() ([32]byte, error) {
	return _Uniswapdex.Contract.Name(&_Uniswapdex.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32 out)
func (_Uniswapdex *UniswapdexCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Uniswapdex.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32 out)
func (_Uniswapdex *UniswapdexSession) Symbol() ([32]byte, error) {
	return _Uniswapdex.Contract.Symbol(&_Uniswapdex.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32 out)
func (_Uniswapdex *UniswapdexCallerSession) Symbol() ([32]byte, error) {
	return _Uniswapdex.Contract.Symbol(&_Uniswapdex.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address out)
func (_Uniswapdex *UniswapdexCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Uniswapdex.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address out)
func (_Uniswapdex *UniswapdexSession) TokenAddress() (common.Address, error) {
	return _Uniswapdex.Contract.TokenAddress(&_Uniswapdex.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address out)
func (_Uniswapdex *UniswapdexCallerSession) TokenAddress() (common.Address, error) {
	return _Uniswapdex.Contract.TokenAddress(&_Uniswapdex.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswapdex.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TotalSupply() (*big.Int, error) {
	return _Uniswapdex.Contract.TotalSupply(&_Uniswapdex.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 out)
func (_Uniswapdex *UniswapdexCallerSession) TotalSupply() (*big.Int, error) {
	return _Uniswapdex.Contract.TotalSupply(&_Uniswapdex.CallOpts)
}

// Default is a paid mutator transaction binding the contract method 0x89402a72.
//
// Solidity: function __default__() returns()
func (_Uniswapdex *UniswapdexTransactor) Default(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "__default__")
}

// Default is a paid mutator transaction binding the contract method 0x89402a72.
//
// Solidity: function __default__() returns()
func (_Uniswapdex *UniswapdexSession) Default() (*types.Transaction, error) {
	return _Uniswapdex.Contract.Default(&_Uniswapdex.TransactOpts)
}

// Default is a paid mutator transaction binding the contract method 0x89402a72.
//
// Solidity: function __default__() returns()
func (_Uniswapdex *UniswapdexTransactorSession) Default() (*types.Transaction, error) {
	return _Uniswapdex.Contract.Default(&_Uniswapdex.TransactOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 min_liquidity, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) AddLiquidity(opts *bind.TransactOpts, min_liquidity *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "addLiquidity", min_liquidity, max_tokens, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 min_liquidity, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) AddLiquidity(min_liquidity *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.AddLiquidity(&_Uniswapdex.TransactOpts, min_liquidity, max_tokens, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 min_liquidity, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) AddLiquidity(min_liquidity *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.AddLiquidity(&_Uniswapdex.TransactOpts, min_liquidity, max_tokens, deadline)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_Uniswapdex *UniswapdexTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_Uniswapdex *UniswapdexSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.Approve(&_Uniswapdex.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_Uniswapdex *UniswapdexTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.Approve(&_Uniswapdex.TransactOpts, _spender, _value)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 min_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) EthToTokenSwapInput(opts *bind.TransactOpts, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "ethToTokenSwapInput", min_tokens, deadline)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 min_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) EthToTokenSwapInput(min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.EthToTokenSwapInput(&_Uniswapdex.TransactOpts, min_tokens, deadline)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 min_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) EthToTokenSwapInput(min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.EthToTokenSwapInput(&_Uniswapdex.TransactOpts, min_tokens, deadline)
}

// EthToTokenSwapOutput is a paid mutator transaction binding the contract method 0x6b1d4db7.
//
// Solidity: function ethToTokenSwapOutput(uint256 tokens_bought, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) EthToTokenSwapOutput(opts *bind.TransactOpts, tokens_bought *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "ethToTokenSwapOutput", tokens_bought, deadline)
}

// EthToTokenSwapOutput is a paid mutator transaction binding the contract method 0x6b1d4db7.
//
// Solidity: function ethToTokenSwapOutput(uint256 tokens_bought, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) EthToTokenSwapOutput(tokens_bought *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.EthToTokenSwapOutput(&_Uniswapdex.TransactOpts, tokens_bought, deadline)
}

// EthToTokenSwapOutput is a paid mutator transaction binding the contract method 0x6b1d4db7.
//
// Solidity: function ethToTokenSwapOutput(uint256 tokens_bought, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) EthToTokenSwapOutput(tokens_bought *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.EthToTokenSwapOutput(&_Uniswapdex.TransactOpts, tokens_bought, deadline)
}

// EthToTokenTransferInput is a paid mutator transaction binding the contract method 0xad65d76d.
//
// Solidity: function ethToTokenTransferInput(uint256 min_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) EthToTokenTransferInput(opts *bind.TransactOpts, min_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "ethToTokenTransferInput", min_tokens, deadline, recipient)
}

// EthToTokenTransferInput is a paid mutator transaction binding the contract method 0xad65d76d.
//
// Solidity: function ethToTokenTransferInput(uint256 min_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) EthToTokenTransferInput(min_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.EthToTokenTransferInput(&_Uniswapdex.TransactOpts, min_tokens, deadline, recipient)
}

// EthToTokenTransferInput is a paid mutator transaction binding the contract method 0xad65d76d.
//
// Solidity: function ethToTokenTransferInput(uint256 min_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) EthToTokenTransferInput(min_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.EthToTokenTransferInput(&_Uniswapdex.TransactOpts, min_tokens, deadline, recipient)
}

// EthToTokenTransferOutput is a paid mutator transaction binding the contract method 0x0b573638.
//
// Solidity: function ethToTokenTransferOutput(uint256 tokens_bought, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) EthToTokenTransferOutput(opts *bind.TransactOpts, tokens_bought *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "ethToTokenTransferOutput", tokens_bought, deadline, recipient)
}

// EthToTokenTransferOutput is a paid mutator transaction binding the contract method 0x0b573638.
//
// Solidity: function ethToTokenTransferOutput(uint256 tokens_bought, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) EthToTokenTransferOutput(tokens_bought *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.EthToTokenTransferOutput(&_Uniswapdex.TransactOpts, tokens_bought, deadline, recipient)
}

// EthToTokenTransferOutput is a paid mutator transaction binding the contract method 0x0b573638.
//
// Solidity: function ethToTokenTransferOutput(uint256 tokens_bought, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) EthToTokenTransferOutput(tokens_bought *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.EthToTokenTransferOutput(&_Uniswapdex.TransactOpts, tokens_bought, deadline, recipient)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xf88bf15a.
//
// Solidity: function removeLiquidity(uint256 amount, uint256 min_eth, uint256 min_tokens, uint256 deadline) returns(uint256 out, uint256 out)
func (_Uniswapdex *UniswapdexTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, min_eth *big.Int, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "removeLiquidity", amount, min_eth, min_tokens, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xf88bf15a.
//
// Solidity: function removeLiquidity(uint256 amount, uint256 min_eth, uint256 min_tokens, uint256 deadline) returns(uint256 out, uint256 out)
func (_Uniswapdex *UniswapdexSession) RemoveLiquidity(amount *big.Int, min_eth *big.Int, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.RemoveLiquidity(&_Uniswapdex.TransactOpts, amount, min_eth, min_tokens, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xf88bf15a.
//
// Solidity: function removeLiquidity(uint256 amount, uint256 min_eth, uint256 min_tokens, uint256 deadline) returns(uint256 out, uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) RemoveLiquidity(amount *big.Int, min_eth *big.Int, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.RemoveLiquidity(&_Uniswapdex.TransactOpts, amount, min_eth, min_tokens, deadline)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address token_addr) returns()
func (_Uniswapdex *UniswapdexTransactor) Setup(opts *bind.TransactOpts, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "setup", token_addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address token_addr) returns()
func (_Uniswapdex *UniswapdexSession) Setup(token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.Setup(&_Uniswapdex.TransactOpts, token_addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address token_addr) returns()
func (_Uniswapdex *UniswapdexTransactorSession) Setup(token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.Setup(&_Uniswapdex.TransactOpts, token_addr)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) TokenToEthSwapInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "tokenToEthSwapInput", tokens_sold, min_eth, deadline)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TokenToEthSwapInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToEthSwapInput(&_Uniswapdex.TransactOpts, tokens_sold, min_eth, deadline)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) TokenToEthSwapInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToEthSwapInput(&_Uniswapdex.TransactOpts, tokens_sold, min_eth, deadline)
}

// TokenToEthSwapOutput is a paid mutator transaction binding the contract method 0x013efd8b.
//
// Solidity: function tokenToEthSwapOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) TokenToEthSwapOutput(opts *bind.TransactOpts, eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "tokenToEthSwapOutput", eth_bought, max_tokens, deadline)
}

// TokenToEthSwapOutput is a paid mutator transaction binding the contract method 0x013efd8b.
//
// Solidity: function tokenToEthSwapOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TokenToEthSwapOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToEthSwapOutput(&_Uniswapdex.TransactOpts, eth_bought, max_tokens, deadline)
}

// TokenToEthSwapOutput is a paid mutator transaction binding the contract method 0x013efd8b.
//
// Solidity: function tokenToEthSwapOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) TokenToEthSwapOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToEthSwapOutput(&_Uniswapdex.TransactOpts, eth_bought, max_tokens, deadline)
}

// TokenToEthTransferInput is a paid mutator transaction binding the contract method 0x7237e031.
//
// Solidity: function tokenToEthTransferInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) TokenToEthTransferInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "tokenToEthTransferInput", tokens_sold, min_eth, deadline, recipient)
}

// TokenToEthTransferInput is a paid mutator transaction binding the contract method 0x7237e031.
//
// Solidity: function tokenToEthTransferInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TokenToEthTransferInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToEthTransferInput(&_Uniswapdex.TransactOpts, tokens_sold, min_eth, deadline, recipient)
}

// TokenToEthTransferInput is a paid mutator transaction binding the contract method 0x7237e031.
//
// Solidity: function tokenToEthTransferInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) TokenToEthTransferInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToEthTransferInput(&_Uniswapdex.TransactOpts, tokens_sold, min_eth, deadline, recipient)
}

// TokenToEthTransferOutput is a paid mutator transaction binding the contract method 0xd4e4841d.
//
// Solidity: function tokenToEthTransferOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) TokenToEthTransferOutput(opts *bind.TransactOpts, eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "tokenToEthTransferOutput", eth_bought, max_tokens, deadline, recipient)
}

// TokenToEthTransferOutput is a paid mutator transaction binding the contract method 0xd4e4841d.
//
// Solidity: function tokenToEthTransferOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TokenToEthTransferOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToEthTransferOutput(&_Uniswapdex.TransactOpts, eth_bought, max_tokens, deadline, recipient)
}

// TokenToEthTransferOutput is a paid mutator transaction binding the contract method 0xd4e4841d.
//
// Solidity: function tokenToEthTransferOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) TokenToEthTransferOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToEthTransferOutput(&_Uniswapdex.TransactOpts, eth_bought, max_tokens, deadline, recipient)
}

// TokenToExchangeSwapInput is a paid mutator transaction binding the contract method 0xb1cb43bf.
//
// Solidity: function tokenToExchangeSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) TokenToExchangeSwapInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "tokenToExchangeSwapInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, exchange_addr)
}

// TokenToExchangeSwapInput is a paid mutator transaction binding the contract method 0xb1cb43bf.
//
// Solidity: function tokenToExchangeSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TokenToExchangeSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToExchangeSwapInput(&_Uniswapdex.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, exchange_addr)
}

// TokenToExchangeSwapInput is a paid mutator transaction binding the contract method 0xb1cb43bf.
//
// Solidity: function tokenToExchangeSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) TokenToExchangeSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToExchangeSwapInput(&_Uniswapdex.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, exchange_addr)
}

// TokenToExchangeSwapOutput is a paid mutator transaction binding the contract method 0xea650c7d.
//
// Solidity: function tokenToExchangeSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) TokenToExchangeSwapOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "tokenToExchangeSwapOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, exchange_addr)
}

// TokenToExchangeSwapOutput is a paid mutator transaction binding the contract method 0xea650c7d.
//
// Solidity: function tokenToExchangeSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TokenToExchangeSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToExchangeSwapOutput(&_Uniswapdex.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, exchange_addr)
}

// TokenToExchangeSwapOutput is a paid mutator transaction binding the contract method 0xea650c7d.
//
// Solidity: function tokenToExchangeSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) TokenToExchangeSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToExchangeSwapOutput(&_Uniswapdex.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, exchange_addr)
}

// TokenToExchangeTransferInput is a paid mutator transaction binding the contract method 0xec384a3e.
//
// Solidity: function tokenToExchangeTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) TokenToExchangeTransferInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "tokenToExchangeTransferInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferInput is a paid mutator transaction binding the contract method 0xec384a3e.
//
// Solidity: function tokenToExchangeTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TokenToExchangeTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToExchangeTransferInput(&_Uniswapdex.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferInput is a paid mutator transaction binding the contract method 0xec384a3e.
//
// Solidity: function tokenToExchangeTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) TokenToExchangeTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToExchangeTransferInput(&_Uniswapdex.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferOutput is a paid mutator transaction binding the contract method 0x981a1327.
//
// Solidity: function tokenToExchangeTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) TokenToExchangeTransferOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "tokenToExchangeTransferOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferOutput is a paid mutator transaction binding the contract method 0x981a1327.
//
// Solidity: function tokenToExchangeTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TokenToExchangeTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToExchangeTransferOutput(&_Uniswapdex.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferOutput is a paid mutator transaction binding the contract method 0x981a1327.
//
// Solidity: function tokenToExchangeTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) TokenToExchangeTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToExchangeTransferOutput(&_Uniswapdex.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, exchange_addr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address token_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) TokenToTokenSwapInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "tokenToTokenSwapInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, token_addr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address token_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TokenToTokenSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToTokenSwapInput(&_Uniswapdex.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, token_addr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address token_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) TokenToTokenSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToTokenSwapInput(&_Uniswapdex.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, token_addr)
}

// TokenToTokenSwapOutput is a paid mutator transaction binding the contract method 0xb040d545.
//
// Solidity: function tokenToTokenSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address token_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) TokenToTokenSwapOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "tokenToTokenSwapOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, token_addr)
}

// TokenToTokenSwapOutput is a paid mutator transaction binding the contract method 0xb040d545.
//
// Solidity: function tokenToTokenSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address token_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TokenToTokenSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToTokenSwapOutput(&_Uniswapdex.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, token_addr)
}

// TokenToTokenSwapOutput is a paid mutator transaction binding the contract method 0xb040d545.
//
// Solidity: function tokenToTokenSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address token_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) TokenToTokenSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToTokenSwapOutput(&_Uniswapdex.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, token_addr)
}

// TokenToTokenTransferInput is a paid mutator transaction binding the contract method 0xf552d91b.
//
// Solidity: function tokenToTokenTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) TokenToTokenTransferInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "tokenToTokenTransferInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, token_addr)
}

// TokenToTokenTransferInput is a paid mutator transaction binding the contract method 0xf552d91b.
//
// Solidity: function tokenToTokenTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TokenToTokenTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToTokenTransferInput(&_Uniswapdex.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, token_addr)
}

// TokenToTokenTransferInput is a paid mutator transaction binding the contract method 0xf552d91b.
//
// Solidity: function tokenToTokenTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) TokenToTokenTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToTokenTransferInput(&_Uniswapdex.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, token_addr)
}

// TokenToTokenTransferOutput is a paid mutator transaction binding the contract method 0xf3c0efe9.
//
// Solidity: function tokenToTokenTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactor) TokenToTokenTransferOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "tokenToTokenTransferOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, token_addr)
}

// TokenToTokenTransferOutput is a paid mutator transaction binding the contract method 0xf3c0efe9.
//
// Solidity: function tokenToTokenTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexSession) TokenToTokenTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToTokenTransferOutput(&_Uniswapdex.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, token_addr)
}

// TokenToTokenTransferOutput is a paid mutator transaction binding the contract method 0xf3c0efe9.
//
// Solidity: function tokenToTokenTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Uniswapdex *UniswapdexTransactorSession) TokenToTokenTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TokenToTokenTransferOutput(&_Uniswapdex.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, token_addr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_Uniswapdex *UniswapdexTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_Uniswapdex *UniswapdexSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.Transfer(&_Uniswapdex.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_Uniswapdex *UniswapdexTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.Transfer(&_Uniswapdex.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_Uniswapdex *UniswapdexTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_Uniswapdex *UniswapdexSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TransferFrom(&_Uniswapdex.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_Uniswapdex *UniswapdexTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswapdex.Contract.TransferFrom(&_Uniswapdex.TransactOpts, _from, _to, _value)
}

// UniswapdexAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Uniswapdex contract.
type UniswapdexAddLiquidityIterator struct {
	Event *UniswapdexAddLiquidity // Event containing the contract specifics and raw log

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
func (it *UniswapdexAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapdexAddLiquidity)
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
		it.Event = new(UniswapdexAddLiquidity)
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
func (it *UniswapdexAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapdexAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapdexAddLiquidity represents a AddLiquidity event raised by the Uniswapdex contract.
type UniswapdexAddLiquidity struct {
	Provider    common.Address
	EthAmount   *big.Int
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Uniswapdex *UniswapdexFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (*UniswapdexAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _Uniswapdex.contract.FilterLogs(opts, "AddLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return &UniswapdexAddLiquidityIterator{contract: _Uniswapdex.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Uniswapdex *UniswapdexFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *UniswapdexAddLiquidity, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _Uniswapdex.contract.WatchLogs(opts, "AddLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapdexAddLiquidity)
				if err := _Uniswapdex.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// UniswapdexApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Uniswapdex contract.
type UniswapdexApprovalIterator struct {
	Event *UniswapdexApproval // Event containing the contract specifics and raw log

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
func (it *UniswapdexApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapdexApproval)
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
		it.Event = new(UniswapdexApproval)
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
func (it *UniswapdexApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapdexApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapdexApproval represents a Approval event raised by the Uniswapdex contract.
type UniswapdexApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Uniswapdex *UniswapdexFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*UniswapdexApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Uniswapdex.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &UniswapdexApprovalIterator{contract: _Uniswapdex.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Uniswapdex *UniswapdexFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UniswapdexApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Uniswapdex.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapdexApproval)
				if err := _Uniswapdex.contract.UnpackLog(event, "Approval", log); err != nil {
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

// UniswapdexEthPurchaseIterator is returned from FilterEthPurchase and is used to iterate over the raw logs and unpacked data for EthPurchase events raised by the Uniswapdex contract.
type UniswapdexEthPurchaseIterator struct {
	Event *UniswapdexEthPurchase // Event containing the contract specifics and raw log

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
func (it *UniswapdexEthPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapdexEthPurchase)
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
		it.Event = new(UniswapdexEthPurchase)
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
func (it *UniswapdexEthPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapdexEthPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapdexEthPurchase represents a EthPurchase event raised by the Uniswapdex contract.
type UniswapdexEthPurchase struct {
	Buyer      common.Address
	TokensSold *big.Int
	EthBought  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEthPurchase is a free log retrieval operation binding the contract event 0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705.
//
// Solidity: event EthPurchase(address indexed buyer, uint256 indexed tokens_sold, uint256 indexed eth_bought)
func (_Uniswapdex *UniswapdexFilterer) FilterEthPurchase(opts *bind.FilterOpts, buyer []common.Address, tokens_sold []*big.Int, eth_bought []*big.Int) (*UniswapdexEthPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokens_soldRule []interface{}
	for _, tokens_soldItem := range tokens_sold {
		tokens_soldRule = append(tokens_soldRule, tokens_soldItem)
	}
	var eth_boughtRule []interface{}
	for _, eth_boughtItem := range eth_bought {
		eth_boughtRule = append(eth_boughtRule, eth_boughtItem)
	}

	logs, sub, err := _Uniswapdex.contract.FilterLogs(opts, "EthPurchase", buyerRule, tokens_soldRule, eth_boughtRule)
	if err != nil {
		return nil, err
	}
	return &UniswapdexEthPurchaseIterator{contract: _Uniswapdex.contract, event: "EthPurchase", logs: logs, sub: sub}, nil
}

// WatchEthPurchase is a free log subscription operation binding the contract event 0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705.
//
// Solidity: event EthPurchase(address indexed buyer, uint256 indexed tokens_sold, uint256 indexed eth_bought)
func (_Uniswapdex *UniswapdexFilterer) WatchEthPurchase(opts *bind.WatchOpts, sink chan<- *UniswapdexEthPurchase, buyer []common.Address, tokens_sold []*big.Int, eth_bought []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokens_soldRule []interface{}
	for _, tokens_soldItem := range tokens_sold {
		tokens_soldRule = append(tokens_soldRule, tokens_soldItem)
	}
	var eth_boughtRule []interface{}
	for _, eth_boughtItem := range eth_bought {
		eth_boughtRule = append(eth_boughtRule, eth_boughtItem)
	}

	logs, sub, err := _Uniswapdex.contract.WatchLogs(opts, "EthPurchase", buyerRule, tokens_soldRule, eth_boughtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapdexEthPurchase)
				if err := _Uniswapdex.contract.UnpackLog(event, "EthPurchase", log); err != nil {
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

// UniswapdexRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Uniswapdex contract.
type UniswapdexRemoveLiquidityIterator struct {
	Event *UniswapdexRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *UniswapdexRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapdexRemoveLiquidity)
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
		it.Event = new(UniswapdexRemoveLiquidity)
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
func (it *UniswapdexRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapdexRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapdexRemoveLiquidity represents a RemoveLiquidity event raised by the Uniswapdex contract.
type UniswapdexRemoveLiquidity struct {
	Provider    common.Address
	EthAmount   *big.Int
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Uniswapdex *UniswapdexFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (*UniswapdexRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _Uniswapdex.contract.FilterLogs(opts, "RemoveLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return &UniswapdexRemoveLiquidityIterator{contract: _Uniswapdex.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Uniswapdex *UniswapdexFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *UniswapdexRemoveLiquidity, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _Uniswapdex.contract.WatchLogs(opts, "RemoveLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapdexRemoveLiquidity)
				if err := _Uniswapdex.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// UniswapdexTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the Uniswapdex contract.
type UniswapdexTokenPurchaseIterator struct {
	Event *UniswapdexTokenPurchase // Event containing the contract specifics and raw log

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
func (it *UniswapdexTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapdexTokenPurchase)
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
		it.Event = new(UniswapdexTokenPurchase)
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
func (it *UniswapdexTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapdexTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapdexTokenPurchase represents a TokenPurchase event raised by the Uniswapdex contract.
type UniswapdexTokenPurchase struct {
	Buyer        common.Address
	EthSold      *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f.
//
// Solidity: event TokenPurchase(address indexed buyer, uint256 indexed eth_sold, uint256 indexed tokens_bought)
func (_Uniswapdex *UniswapdexFilterer) FilterTokenPurchase(opts *bind.FilterOpts, buyer []common.Address, eth_sold []*big.Int, tokens_bought []*big.Int) (*UniswapdexTokenPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var eth_soldRule []interface{}
	for _, eth_soldItem := range eth_sold {
		eth_soldRule = append(eth_soldRule, eth_soldItem)
	}
	var tokens_boughtRule []interface{}
	for _, tokens_boughtItem := range tokens_bought {
		tokens_boughtRule = append(tokens_boughtRule, tokens_boughtItem)
	}

	logs, sub, err := _Uniswapdex.contract.FilterLogs(opts, "TokenPurchase", buyerRule, eth_soldRule, tokens_boughtRule)
	if err != nil {
		return nil, err
	}
	return &UniswapdexTokenPurchaseIterator{contract: _Uniswapdex.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f.
//
// Solidity: event TokenPurchase(address indexed buyer, uint256 indexed eth_sold, uint256 indexed tokens_bought)
func (_Uniswapdex *UniswapdexFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *UniswapdexTokenPurchase, buyer []common.Address, eth_sold []*big.Int, tokens_bought []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var eth_soldRule []interface{}
	for _, eth_soldItem := range eth_sold {
		eth_soldRule = append(eth_soldRule, eth_soldItem)
	}
	var tokens_boughtRule []interface{}
	for _, tokens_boughtItem := range tokens_bought {
		tokens_boughtRule = append(tokens_boughtRule, tokens_boughtItem)
	}

	logs, sub, err := _Uniswapdex.contract.WatchLogs(opts, "TokenPurchase", buyerRule, eth_soldRule, tokens_boughtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapdexTokenPurchase)
				if err := _Uniswapdex.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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

// UniswapdexTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Uniswapdex contract.
type UniswapdexTransferIterator struct {
	Event *UniswapdexTransfer // Event containing the contract specifics and raw log

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
func (it *UniswapdexTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapdexTransfer)
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
		it.Event = new(UniswapdexTransfer)
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
func (it *UniswapdexTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapdexTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapdexTransfer represents a Transfer event raised by the Uniswapdex contract.
type UniswapdexTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Uniswapdex *UniswapdexFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*UniswapdexTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Uniswapdex.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &UniswapdexTransferIterator{contract: _Uniswapdex.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Uniswapdex *UniswapdexFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UniswapdexTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Uniswapdex.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapdexTransfer)
				if err := _Uniswapdex.contract.UnpackLog(event, "Transfer", log); err != nil {
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
