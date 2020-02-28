// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapfactory

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

// UniswapfactoryABI is the input ABI used to generate the binding from.
const UniswapfactoryABI = "[{\"name\":\"NewExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"exchange\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"initializeFactory\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"template\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35725},{\"name\":\"createExchange\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":187911},{\"name\":\"getExchange\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":715},{\"name\":\"getToken\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"exchange\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":745},{\"name\":\"getTokenWithId\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"token_id\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":736},{\"name\":\"exchangeTemplate\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":633},{\"name\":\"tokenCount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":663}]"

const HexAddr = "0xc0a47dFe034B400B47bDaD5FecDa2621de6c4d95"

// Uniswapfactory is an auto generated Go binding around an Ethereum contract.
type Uniswapfactory struct {
	UniswapfactoryCaller     // Read-only binding to the contract
	UniswapfactoryTransactor // Write-only binding to the contract
	UniswapfactoryFilterer   // Log filterer for contract events
}

// UniswapfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapfactorySession struct {
	Contract     *Uniswapfactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapfactoryCallerSession struct {
	Contract *UniswapfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UniswapfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapfactoryTransactorSession struct {
	Contract     *UniswapfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UniswapfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapfactoryRaw struct {
	Contract *Uniswapfactory // Generic contract binding to access the raw methods on
}

// UniswapfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapfactoryCallerRaw struct {
	Contract *UniswapfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapfactoryTransactorRaw struct {
	Contract *UniswapfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapfactory creates a new instance of Uniswapfactory, bound to a specific deployed contract.
func NewUniswapfactory(address common.Address, backend bind.ContractBackend) (*Uniswapfactory, error) {
	contract, err := bindUniswapfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapfactory{UniswapfactoryCaller: UniswapfactoryCaller{contract: contract}, UniswapfactoryTransactor: UniswapfactoryTransactor{contract: contract}, UniswapfactoryFilterer: UniswapfactoryFilterer{contract: contract}}, nil
}

// NewUniswapfactoryCaller creates a new read-only instance of Uniswapfactory, bound to a specific deployed contract.
func NewUniswapfactoryCaller(address common.Address, caller bind.ContractCaller) (*UniswapfactoryCaller, error) {
	contract, err := bindUniswapfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapfactoryCaller{contract: contract}, nil
}

// NewUniswapfactoryTransactor creates a new write-only instance of Uniswapfactory, bound to a specific deployed contract.
func NewUniswapfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapfactoryTransactor, error) {
	contract, err := bindUniswapfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapfactoryTransactor{contract: contract}, nil
}

// NewUniswapfactoryFilterer creates a new log filterer instance of Uniswapfactory, bound to a specific deployed contract.
func NewUniswapfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapfactoryFilterer, error) {
	contract, err := bindUniswapfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapfactoryFilterer{contract: contract}, nil
}

// bindUniswapfactory binds a generic wrapper to an already deployed contract.
func bindUniswapfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapfactory *UniswapfactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Uniswapfactory.Contract.UniswapfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapfactory *UniswapfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.UniswapfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapfactory *UniswapfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.UniswapfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapfactory *UniswapfactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Uniswapfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapfactory *UniswapfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapfactory *UniswapfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.contract.Transact(opts, method, params...)
}

// ExchangeTemplate is a free data retrieval call binding the contract method 0x1c2bbd18.
//
// Solidity: function exchangeTemplate() constant returns(address out)
func (_Uniswapfactory *UniswapfactoryCaller) ExchangeTemplate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Uniswapfactory.contract.Call(opts, out, "exchangeTemplate")
	return *ret0, err
}

// ExchangeTemplate is a free data retrieval call binding the contract method 0x1c2bbd18.
//
// Solidity: function exchangeTemplate() constant returns(address out)
func (_Uniswapfactory *UniswapfactorySession) ExchangeTemplate() (common.Address, error) {
	return _Uniswapfactory.Contract.ExchangeTemplate(&_Uniswapfactory.CallOpts)
}

// ExchangeTemplate is a free data retrieval call binding the contract method 0x1c2bbd18.
//
// Solidity: function exchangeTemplate() constant returns(address out)
func (_Uniswapfactory *UniswapfactoryCallerSession) ExchangeTemplate() (common.Address, error) {
	return _Uniswapfactory.Contract.ExchangeTemplate(&_Uniswapfactory.CallOpts)
}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) constant returns(address out)
func (_Uniswapfactory *UniswapfactoryCaller) GetExchange(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Uniswapfactory.contract.Call(opts, out, "getExchange", token)
	return *ret0, err
}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) constant returns(address out)
func (_Uniswapfactory *UniswapfactorySession) GetExchange(token common.Address) (common.Address, error) {
	return _Uniswapfactory.Contract.GetExchange(&_Uniswapfactory.CallOpts, token)
}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) constant returns(address out)
func (_Uniswapfactory *UniswapfactoryCallerSession) GetExchange(token common.Address) (common.Address, error) {
	return _Uniswapfactory.Contract.GetExchange(&_Uniswapfactory.CallOpts, token)
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address exchange) constant returns(address out)
func (_Uniswapfactory *UniswapfactoryCaller) GetToken(opts *bind.CallOpts, exchange common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Uniswapfactory.contract.Call(opts, out, "getToken", exchange)
	return *ret0, err
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address exchange) constant returns(address out)
func (_Uniswapfactory *UniswapfactorySession) GetToken(exchange common.Address) (common.Address, error) {
	return _Uniswapfactory.Contract.GetToken(&_Uniswapfactory.CallOpts, exchange)
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address exchange) constant returns(address out)
func (_Uniswapfactory *UniswapfactoryCallerSession) GetToken(exchange common.Address) (common.Address, error) {
	return _Uniswapfactory.Contract.GetToken(&_Uniswapfactory.CallOpts, exchange)
}

// GetTokenWithId is a free data retrieval call binding the contract method 0xaa65a6c0.
//
// Solidity: function getTokenWithId(uint256 token_id) constant returns(address out)
func (_Uniswapfactory *UniswapfactoryCaller) GetTokenWithId(opts *bind.CallOpts, token_id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Uniswapfactory.contract.Call(opts, out, "getTokenWithId", token_id)
	return *ret0, err
}

// GetTokenWithId is a free data retrieval call binding the contract method 0xaa65a6c0.
//
// Solidity: function getTokenWithId(uint256 token_id) constant returns(address out)
func (_Uniswapfactory *UniswapfactorySession) GetTokenWithId(token_id *big.Int) (common.Address, error) {
	return _Uniswapfactory.Contract.GetTokenWithId(&_Uniswapfactory.CallOpts, token_id)
}

// GetTokenWithId is a free data retrieval call binding the contract method 0xaa65a6c0.
//
// Solidity: function getTokenWithId(uint256 token_id) constant returns(address out)
func (_Uniswapfactory *UniswapfactoryCallerSession) GetTokenWithId(token_id *big.Int) (common.Address, error) {
	return _Uniswapfactory.Contract.GetTokenWithId(&_Uniswapfactory.CallOpts, token_id)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() constant returns(uint256 out)
func (_Uniswapfactory *UniswapfactoryCaller) TokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswapfactory.contract.Call(opts, out, "tokenCount")
	return *ret0, err
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() constant returns(uint256 out)
func (_Uniswapfactory *UniswapfactorySession) TokenCount() (*big.Int, error) {
	return _Uniswapfactory.Contract.TokenCount(&_Uniswapfactory.CallOpts)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() constant returns(uint256 out)
func (_Uniswapfactory *UniswapfactoryCallerSession) TokenCount() (*big.Int, error) {
	return _Uniswapfactory.Contract.TokenCount(&_Uniswapfactory.CallOpts)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x1648f38e.
//
// Solidity: function createExchange(address token) returns(address out)
func (_Uniswapfactory *UniswapfactoryTransactor) CreateExchange(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.contract.Transact(opts, "createExchange", token)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x1648f38e.
//
// Solidity: function createExchange(address token) returns(address out)
func (_Uniswapfactory *UniswapfactorySession) CreateExchange(token common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.CreateExchange(&_Uniswapfactory.TransactOpts, token)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x1648f38e.
//
// Solidity: function createExchange(address token) returns(address out)
func (_Uniswapfactory *UniswapfactoryTransactorSession) CreateExchange(token common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.CreateExchange(&_Uniswapfactory.TransactOpts, token)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0x538a3f0e.
//
// Solidity: function initializeFactory(address template) returns()
func (_Uniswapfactory *UniswapfactoryTransactor) InitializeFactory(opts *bind.TransactOpts, template common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.contract.Transact(opts, "initializeFactory", template)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0x538a3f0e.
//
// Solidity: function initializeFactory(address template) returns()
func (_Uniswapfactory *UniswapfactorySession) InitializeFactory(template common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.InitializeFactory(&_Uniswapfactory.TransactOpts, template)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0x538a3f0e.
//
// Solidity: function initializeFactory(address template) returns()
func (_Uniswapfactory *UniswapfactoryTransactorSession) InitializeFactory(template common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.InitializeFactory(&_Uniswapfactory.TransactOpts, template)
}

// UniswapfactoryNewExchangeIterator is returned from FilterNewExchange and is used to iterate over the raw logs and unpacked data for NewExchange events raised by the Uniswapfactory contract.
type UniswapfactoryNewExchangeIterator struct {
	Event *UniswapfactoryNewExchange // Event containing the contract specifics and raw log

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
func (it *UniswapfactoryNewExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapfactoryNewExchange)
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
		it.Event = new(UniswapfactoryNewExchange)
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
func (it *UniswapfactoryNewExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapfactoryNewExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapfactoryNewExchange represents a NewExchange event raised by the Uniswapfactory contract.
type UniswapfactoryNewExchange struct {
	Token    common.Address
	Exchange common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewExchange is a free log retrieval operation binding the contract event 0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9.
//
// Solidity: event NewExchange(address indexed token, address indexed exchange)
func (_Uniswapfactory *UniswapfactoryFilterer) FilterNewExchange(opts *bind.FilterOpts, token []common.Address, exchange []common.Address) (*UniswapfactoryNewExchangeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var exchangeRule []interface{}
	for _, exchangeItem := range exchange {
		exchangeRule = append(exchangeRule, exchangeItem)
	}

	logs, sub, err := _Uniswapfactory.contract.FilterLogs(opts, "NewExchange", tokenRule, exchangeRule)
	if err != nil {
		return nil, err
	}
	return &UniswapfactoryNewExchangeIterator{contract: _Uniswapfactory.contract, event: "NewExchange", logs: logs, sub: sub}, nil
}

// WatchNewExchange is a free log subscription operation binding the contract event 0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9.
//
// Solidity: event NewExchange(address indexed token, address indexed exchange)
func (_Uniswapfactory *UniswapfactoryFilterer) WatchNewExchange(opts *bind.WatchOpts, sink chan<- *UniswapfactoryNewExchange, token []common.Address, exchange []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var exchangeRule []interface{}
	for _, exchangeItem := range exchange {
		exchangeRule = append(exchangeRule, exchangeItem)
	}

	logs, sub, err := _Uniswapfactory.contract.WatchLogs(opts, "NewExchange", tokenRule, exchangeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapfactoryNewExchange)
				if err := _Uniswapfactory.contract.UnpackLog(event, "NewExchange", log); err != nil {
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

// ParseNewExchange is a log parse operation binding the contract event 0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9.
//
// Solidity: event NewExchange(address indexed token, address indexed exchange)
func (_Uniswapfactory *UniswapfactoryFilterer) ParseNewExchange(log types.Log) (*UniswapfactoryNewExchange, error) {
	event := new(UniswapfactoryNewExchange)
	if err := _Uniswapfactory.contract.UnpackLog(event, "NewExchange", log); err != nil {
		return nil, err
	}
	return event, nil
}
