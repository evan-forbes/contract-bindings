package pipeth

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

// Pipeth is a wrapper around BoundContract, enforcing type checking and including
// QoL helper methods
type Pipeth struct {
	bind.BoundContract
}

// NewPipeth creates a new instance of Pipeth, bound to a specific deployed contract.
func NewPipeth(address common.Address, backend bind.ContractBackend) (*Pipeth, error) {
	a, err := abi.JSON(strings.NewReader(PipethABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, a, backend, backend, backend)
	return &Pipeth{*contract}, nil
}

//////////////////////////////////////////////////////
//		Structs
////////////////////////////////////////////////////

//////////////////////////////////////////////////////
//		Data Calls
////////////////////////////////////////////////////

// Bud is a free data retrieval call binding the contract method 0x4fce7a2a.
// - Solidity: function bud(address ) constant returns(uint256)
func (_Pipeth *Pipeth) Bud(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pipeth.Call(opts, out, "bud", arg0)
	return *ret0, err
}

// Hop is a free data retrieval call binding the contract method 0xb0b8579b.
// - Solidity: function hop() constant returns(uint16)
func (_Pipeth *Pipeth) Hop(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Pipeth.Call(opts, out, "hop")
	return *ret0, err
}

// Pass is a free data retrieval call binding the contract method 0xa7a1ed72.
// - Solidity: function pass() constant returns(bool ok)
func (_Pipeth *Pipeth) Pass(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pipeth.Call(opts, out, "pass")
	return *ret0, err
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
// - Solidity: function peek() constant returns(bytes32, bool)
func (_Pipeth *Pipeth) Peek(opts *bind.CallOpts) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Pipeth.Call(opts, out, "peek")
	return *ret0, *ret1, err
}

// Peep is a free data retrieval call binding the contract method 0x0e5a6c70.
// - Solidity: function peep() constant returns(bytes32, bool)
func (_Pipeth *Pipeth) Peep(opts *bind.CallOpts) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Pipeth.Call(opts, out, "peep")
	return *ret0, *ret1, err
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
// - Solidity: function read() constant returns(bytes32)
func (_Pipeth *Pipeth) Read(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Pipeth.Call(opts, out, "read")
	return *ret0, err
}

// Src is a free data retrieval call binding the contract method 0x2e7dc6af.
// - Solidity: function src() constant returns(address)
func (_Pipeth *Pipeth) Src(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pipeth.Call(opts, out, "src")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
// - Solidity: function stopped() constant returns(uint256)
func (_Pipeth *Pipeth) Stopped(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pipeth.Call(opts, out, "stopped")
	return *ret0, err
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
// - Solidity: function wards(address ) constant returns(uint256)
func (_Pipeth *Pipeth) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pipeth.Call(opts, out, "wards", arg0)
	return *ret0, err
}

// Zzz is a free data retrieval call binding the contract method 0xa4dff0a2.
// - Solidity: function zzz() constant returns(uint64)
func (_Pipeth *Pipeth) Zzz(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Pipeth.Call(opts, out, "zzz")
	return *ret0, err
}

//////////////////////////////////////////////////////
//		Transactions
////////////////////////////////////////////////////

// Change is a paid mutator transaction binding the contract method 0x1e77933e.
// - Solidity: function change(address src_) returns()
func (_Pipeth *Pipeth) Change(opts *bind.TransactOpts, src_ common.Address) (*types.Transaction, error) {
	return _Pipeth.Transact(opts, "change", src_)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
// - Solidity: function deny(address usr) returns()
func (_Pipeth *Pipeth) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Pipeth.Transact(opts, "deny", usr)
}

// Diss is a paid mutator transaction binding the contract method 0x46d4577d.
// - Solidity: function diss(address[] a) returns()
func (_Pipeth *Pipeth) Diss(opts *bind.TransactOpts, a []common.Address) (*types.Transaction, error) {
	return _Pipeth.Transact(opts, "diss", a)
}

// Diss0 is a paid mutator transaction binding the contract method 0x65c4ce7a.
// - Solidity: function diss(address a) returns()
func (_Pipeth *Pipeth) Diss0(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Pipeth.Transact(opts, "diss0", a)
}

// Kiss is a paid mutator transaction binding the contract method 0x1b25b65f.
// - Solidity: function kiss(address[] a) returns()
func (_Pipeth *Pipeth) Kiss(opts *bind.TransactOpts, a []common.Address) (*types.Transaction, error) {
	return _Pipeth.Transact(opts, "kiss", a)
}

// Kiss0 is a paid mutator transaction binding the contract method 0xf29c29c4.
// - Solidity: function kiss(address a) returns()
func (_Pipeth *Pipeth) Kiss0(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Pipeth.Transact(opts, "kiss0", a)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
// - Solidity: function poke() returns()
func (_Pipeth *Pipeth) Poke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pipeth.Transact(opts, "poke")
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
// - Solidity: function rely(address usr) returns()
func (_Pipeth *Pipeth) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Pipeth.Transact(opts, "rely", usr)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
// - Solidity: function start() returns()
func (_Pipeth *Pipeth) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pipeth.Transact(opts, "start")
}

// Step is a paid mutator transaction binding the contract method 0xe38e2cfb.
// - Solidity: function step(uint16 ts) returns()
func (_Pipeth *Pipeth) Step(opts *bind.TransactOpts, ts uint16) (*types.Transaction, error) {
	return _Pipeth.Transact(opts, "step", ts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
// - Solidity: function stop() returns()
func (_Pipeth *Pipeth) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pipeth.Transact(opts, "stop")
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
// - Solidity: function void() returns()
func (_Pipeth *Pipeth) Void(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pipeth.Transact(opts, "void")
}

//////////////////////////////////////////////////////
//		Events
////////////////////////////////////////////////////

//////// LogValue ////////

// LogValueID is the hex of the Topic Hash
const LogValueID = "0x296ba4ca62c6c21c95e828080cb8aec7481b71390585605300a8a76f9e95b527"

// LogValueLog represents a LogValue event raised by the Pipeth contract.
type LogValueLog struct {
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// UnpackLogValueLog is a log parse operation binding the contract event 0x296ba4ca62c6c21c95e828080cb8aec7481b71390585605300a8a76f9e95b527
// Solidity: event LogValue(bytes32 val)
func (_Pipeth *Pipeth) UnpackLogValueLog(log types.Log) (*LogValueLog, error) {
	event := new(LogValueLog)
	if err := _Pipeth.UnpackLog(event, "LogValue", log); err != nil {
		return nil, err
	}
	return event, nil
}

/*
Mux can be copied and pasted to save ya a quick minute when distinguishing between log data
func Mux(c *Pipeth, log types.Log) error {
	switch log.Topics[0].Hex() {
	case pipeth.PipethLogValueID:
		ulog, err := c.UnpackLogValueLog(log)
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

// PipethBin is used to deploy the generated contract
const PipethBin = "0x"

// PipethABI is used to communicate with the compiled solidity code of the generated contract
const PipethABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"LogValue\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bud\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src_\",\"type\":\"address\"}],\"name\":\"change\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"a\",\"type\":\"address[]\"}],\"name\":\"diss\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"diss\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hop\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"a\",\"type\":\"address[]\"}],\"name\":\"kiss\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"kiss\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pass\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peek\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"poke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"read\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"src\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"ts\",\"type\":\"uint16\"}],\"name\":\"step\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"void\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zzz\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"
