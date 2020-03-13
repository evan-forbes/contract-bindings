package median

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

// Median is a wrapper around BoundContract, enforcing type checking and including
// QoL helper methods
type Median struct {
	bind.BoundContract
}

// NewMedian creates a new instance of Median, bound to a specific deployed contract.
func NewMedian(address common.Address, backend bind.ContractBackend) (*Median, error) {
	a, err := abi.JSON(strings.NewReader(MedianABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, a, backend, backend, backend)
	return &Median{*contract}, nil
}

//////////////////////////////////////////////////////
//		Structs
////////////////////////////////////////////////////

//////////////////////////////////////////////////////
//		Data Calls
////////////////////////////////////////////////////

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
// - Solidity: function authority() constant returns(address)
func (_Median *Median) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Median.Call(opts, out, "authority")
	return *ret0, err
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
// - Solidity: function compute() constant returns(bytes32, bool)
func (_Median *Median) Compute(opts *bind.CallOpts) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Median.Call(opts, out, "compute")
	return *ret0, *ret1, err
}

// Indexes is a free data retrieval call binding the contract method 0x2db78d93.
// - Solidity: function indexes(address ) constant returns(bytes12)
func (_Median *Median) Indexes(opts *bind.CallOpts, arg0 common.Address) ([12]byte, error) {
	var (
		ret0 = new([12]byte)
	)
	out := ret0
	err := _Median.Call(opts, out, "indexes", arg0)
	return *ret0, err
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
// - Solidity: function min() constant returns(uint96)
func (_Median *Median) Min(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Median.Call(opts, out, "min")
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
// - Solidity: function next() constant returns(bytes12)
func (_Median *Median) Next(opts *bind.CallOpts) ([12]byte, error) {
	var (
		ret0 = new([12]byte)
	)
	out := ret0
	err := _Median.Call(opts, out, "next")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
// - Solidity: function owner() constant returns(address)
func (_Median *Median) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Median.Call(opts, out, "owner")
	return *ret0, err
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
// - Solidity: function peek() constant returns(bytes32, bool)
func (_Median *Median) Peek(opts *bind.CallOpts) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Median.Call(opts, out, "peek")
	return *ret0, *ret1, err
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
// - Solidity: function read() constant returns(bytes32)
func (_Median *Median) Read(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Median.Call(opts, out, "read")
	return *ret0, err
}

// Values is a free data retrieval call binding the contract method 0x651dd0de.
// - Solidity: function values(bytes12 ) constant returns(address)
func (_Median *Median) Values(opts *bind.CallOpts, arg0 [12]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Median.Call(opts, out, "values", arg0)
	return *ret0, err
}

//////////////////////////////////////////////////////
//		Transactions
////////////////////////////////////////////////////

// Poke is a paid mutator transaction binding the contract method 0x1504460f.
// - Solidity: function poke(bytes32 ) returns()
func (_Median *Median) Poke(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _Median.Transact(opts, "poke", arg0)
}

// Poke0 is a paid mutator transaction binding the contract method 0x18178358.
// - Solidity: function poke() returns()
func (_Median *Median) Poke0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Median.Transact(opts, "poke0")
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
// - Solidity: function set(address wat) returns()
func (_Median *Median) Set(opts *bind.TransactOpts, wat common.Address) (*types.Transaction, error) {
	return _Median.Transact(opts, "set", wat)
}

// Set0 is a paid mutator transaction binding the contract method 0xbeb38b43.
// - Solidity: function set(bytes12 pos, address wat) returns()
func (_Median *Median) Set0(opts *bind.TransactOpts, pos [12]byte, wat common.Address) (*types.Transaction, error) {
	return _Median.Transact(opts, "set0", pos, wat)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
// - Solidity: function setAuthority(address authority_) returns()
func (_Median *Median) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Median.Transact(opts, "setAuthority", authority_)
}

// SetMin is a paid mutator transaction binding the contract method 0x6ba5ef0d.
// - Solidity: function setMin(uint96 min_) returns()
func (_Median *Median) SetMin(opts *bind.TransactOpts, min_ *big.Int) (*types.Transaction, error) {
	return _Median.Transact(opts, "setMin", min_)
}

// SetNext is a paid mutator transaction binding the contract method 0xf2c5925d.
// - Solidity: function setNext(bytes12 next_) returns()
func (_Median *Median) SetNext(opts *bind.TransactOpts, next_ [12]byte) (*types.Transaction, error) {
	return _Median.Transact(opts, "setNext", next_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
// - Solidity: function setOwner(address owner_) returns()
func (_Median *Median) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Median.Transact(opts, "setOwner", owner_)
}

// Unset is a paid mutator transaction binding the contract method 0x2966d1b9.
// - Solidity: function unset(address wat) returns()
func (_Median *Median) Unset(opts *bind.TransactOpts, wat common.Address) (*types.Transaction, error) {
	return _Median.Transact(opts, "unset", wat)
}

// Unset0 is a paid mutator transaction binding the contract method 0xe0a1fdad.
// - Solidity: function unset(bytes12 pos) returns()
func (_Median *Median) Unset0(opts *bind.TransactOpts, pos [12]byte) (*types.Transaction, error) {
	return _Median.Transact(opts, "unset0", pos)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
// - Solidity: function void() returns()
func (_Median *Median) Void(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Median.Transact(opts, "void")
}

//////////////////////////////////////////////////////
//		Events
////////////////////////////////////////////////////

//////// LogSetAuthority ////////

// LogSetAuthorityID is the hex of the Topic Hash
const LogSetAuthorityID = "0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4"

// LogSetAuthorityLog represents a LogSetAuthority event raised by the Median contract.
type LogSetAuthorityLog struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// UnpackLogSetAuthorityLog is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4
// Solidity: event LogSetAuthority(address indexed authority)
func (_Median *Median) UnpackLogSetAuthorityLog(log types.Log) (*LogSetAuthorityLog, error) {
	event := new(LogSetAuthorityLog)
	if err := _Median.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogSetOwner ////////

// LogSetOwnerID is the hex of the Topic Hash
const LogSetOwnerID = "0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94"

// LogSetOwnerLog represents a LogSetOwner event raised by the Median contract.
type LogSetOwnerLog struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// UnpackLogSetOwnerLog is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94
// Solidity: event LogSetOwner(address indexed owner)
func (_Median *Median) UnpackLogSetOwnerLog(log types.Log) (*LogSetOwnerLog, error) {
	event := new(LogSetOwnerLog)
	if err := _Median.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}

/*
Mux can be copied and pasted to save ya a quick minute when distinguishing between log data
func Mux(c *Median, log types.Log) error {
	switch log.Topics[0].Hex() {
	case median.MedianLogSetAuthorityID:
		ulog, err := c.UnpackLogSetAuthorityLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case median.MedianLogSetOwnerID:
		ulog, err := c.UnpackLogSetOwnerLog(log)
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

// MedianBin is used to deploy the generated contract
const MedianBin = "0x"

// MedianABI is used to communicate with the compiled solidity code of the generated contract
const MedianABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"poke\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"poke\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compute\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wat\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wat\",\"type\":\"address\"}],\"name\":\"unset\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"indexes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"next\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"read\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peek\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes12\"}],\"name\":\"values\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"min_\",\"type\":\"uint96\"}],\"name\":\"setMin\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"void\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pos\",\"type\":\"bytes12\"},{\"name\":\"wat\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pos\",\"type\":\"bytes12\"}],\"name\":\"unset\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next_\",\"type\":\"bytes12\"}],\"name\":\"setNext\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"min\",\"outputs\":[{\"name\":\"\",\"type\":\"uint96\"}],\"payable\":false,\"type\":\"function\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"
