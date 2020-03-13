package osm

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

// Osm is a wrapper around BoundContract, enforcing type checking and including
// QoL helper methods
type Osm struct {
	bind.BoundContract
}

// NewOsm creates a new instance of Osm, bound to a specific deployed contract.
func NewOsm(address common.Address, backend bind.ContractBackend) (*Osm, error) {
	a, err := abi.JSON(strings.NewReader(OsmABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, a, backend, backend, backend)
	return &Osm{*contract}, nil
}

//////////////////////////////////////////////////////
//		Structs
////////////////////////////////////////////////////

//////////////////////////////////////////////////////
//		Data Calls
////////////////////////////////////////////////////

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
// - Solidity: function authority() constant returns(address)
func (_Osm *Osm) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Osm.Call(opts, out, "authority")
	return *ret0, err
}

// Osms is a free data retrieval call binding the contract method 0x6c4ba760.
// - Solidity: function osms(bytes32 ) constant returns(address)
func (_Osm *Osm) Osms(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Osm.Call(opts, out, "osms", arg0)
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
// - Solidity: function owner() constant returns(address)
func (_Osm *Osm) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Osm.Call(opts, out, "owner")
	return *ret0, err
}

//////////////////////////////////////////////////////
//		Transactions
////////////////////////////////////////////////////

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
// - Solidity: function setAuthority(address authority_) returns()
func (_Osm *Osm) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Osm.Transact(opts, "setAuthority", authority_)
}

// SetOsm is a paid mutator transaction binding the contract method 0xc98cdf86.
// - Solidity: function setOsm(bytes32 ilk, address osm) returns()
func (_Osm *Osm) SetOsm(opts *bind.TransactOpts, ilk [32]byte, osm common.Address) (*types.Transaction, error) {
	return _Osm.Transact(opts, "setOsm", ilk, osm)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
// - Solidity: function setOwner(address owner_) returns()
func (_Osm *Osm) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Osm.Transact(opts, "setOwner", owner_)
}

// Stop is a paid mutator transaction binding the contract method 0x63c4f031.
// - Solidity: function stop(bytes32 ilk) returns()
func (_Osm *Osm) Stop(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Osm.Transact(opts, "stop", ilk)
}

//////////////////////////////////////////////////////
//		Events
////////////////////////////////////////////////////

/*
Mux can be copied and pasted to save ya a quick minute when distinguishing between log data
func Mux(c *Osm, log types.Log) error {
	switch log.Topics[0].Hex() {
	}
}
*/

//////////////////////////////////////////////////////
//		Bin and ABI
////////////////////////////////////////////////////

// OsmBin is used to deploy the generated contract
const OsmBin = "0x"

// OsmABI is used to communicate with the compiled solidity code of the generated contract
const OsmABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"osms\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"osm\",\"type\":\"address\"}],\"name\":\"setOsm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
