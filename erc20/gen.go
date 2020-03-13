package erc20

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

// Erc20 is a wrapper around BoundContract, enforcing type checking and including
// QoL helper methods
type Erc20 struct {
	bind.BoundContract
}

// NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.
func NewErc20(address common.Address, backend bind.ContractBackend) (*Erc20, error) {
	a, err := abi.JSON(strings.NewReader(Erc20ABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, a, backend, backend, backend)
	return &Erc20{*contract}, nil
}

//////////////////////////////////////////////////////
//		Structs
////////////////////////////////////////////////////

//////////////////////////////////////////////////////
//		Data Calls
////////////////////////////////////////////////////

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
// - Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_Erc20 *Erc20) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Erc20.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
// - Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_Erc20 *Erc20) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Erc20.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
// - Solidity: function decimals() constant returns(uint8)
func (_Erc20 *Erc20) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Erc20.Call(opts, out, "decimals")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
// - Solidity: function name() constant returns(string)
func (_Erc20 *Erc20) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Erc20.Call(opts, out, "name")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
// - Solidity: function symbol() constant returns(string)
func (_Erc20 *Erc20) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Erc20.Call(opts, out, "symbol")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
// - Solidity: function totalSupply() constant returns(uint256)
func (_Erc20 *Erc20) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Erc20.Call(opts, out, "totalSupply")
	return *ret0, err
}

//////////////////////////////////////////////////////
//		Transactions
////////////////////////////////////////////////////

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
// - Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Erc20 *Erc20) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.Transact(opts, "approve", _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
// - Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Erc20 *Erc20) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.Transact(opts, "transfer", _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
// - Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Erc20 *Erc20) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.Transact(opts, "transferFrom", _from, _to, _value)
}

//////////////////////////////////////////////////////
//		Events
////////////////////////////////////////////////////

//////// Approval ////////

// ApprovalID is the hex of the Topic Hash
const ApprovalID = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"

// ApprovalLog represents a Approval event raised by the Erc20 contract.
type ApprovalLog struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// UnpackApprovalLog is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20) UnpackApprovalLog(log types.Log) (*ApprovalLog, error) {
	event := new(ApprovalLog)
	if err := _Erc20.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// Transfer ////////

// TransferID is the hex of the Topic Hash
const TransferID = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

// TransferLog represents a Transfer event raised by the Erc20 contract.
type TransferLog struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// UnpackTransferLog is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20) UnpackTransferLog(log types.Log) (*TransferLog, error) {
	event := new(TransferLog)
	if err := _Erc20.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

/*
Mux can be copied and pasted to save ya a quick minute when distinguishing between log data
func Mux(c *Erc20, log types.Log) error {
	switch log.Topics[0].Hex() {
	case erc20.Erc20ApprovalID:
		ulog, err := c.UnpackApprovalLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case erc20.Erc20TransferID:
		ulog, err := c.UnpackTransferLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	}
}
*/

//////////////////////////////////////////////////////
//		Bin and ABI
////////////////////////////////////////////////////

// Erc20Bin is used to deploy the generated contract
const Erc20Bin = "0x"

// Erc20ABI is used to communicate with the compiled solidity code of the generated contract
const Erc20ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"
