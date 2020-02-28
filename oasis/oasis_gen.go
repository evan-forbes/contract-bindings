package oasis

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

// Oasis is a wrapper around BoundContract, enforcing type checking and including
// QoL helper methods
type Oasis struct {
	bind.BoundContract
}

// NewOasis creates a new instance of Oasis, bound to a specific deployed contract.
func NewOasis(address common.Address, backend bind.ContractBackend) (*Oasis, error) {
	a, err := abi.JSON(strings.NewReader(OasisABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, a, backend, backend, backend)
	return &Oasis{*contract}, nil
}

//////////////////////////////////////////////////////
//		Deployment
////////////////////////////////////////////////////

// DeployOasis deploys a new Ethereum contract, binding an instance of Oasis to it.
func DeployOasis(auth *bind.TransactOpts, backend bind.ContractBackend, close_time uint64) (common.Address, *types.Transaction, *Oasis, error) {
	parsed, err := abi.JSON(strings.NewReader(OasisABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OasisBin), backend, close_time)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oasis{*contract}, nil
}

//////////////////////////////////////////////////////
//		Structs
////////////////////////////////////////////////////

//////////////////////////////////////////////////////
//		Data Calls
////////////////////////////////////////////////////

// Best is a free data retrieval call binding the contract method 0x74c1d7d3.
// - Solidity: function _best(address , address ) constant returns(uint256)
func (_Oasis *Oasis) Best(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "_best", arg0, arg1)
	return *ret0, err
}

// Dust is a free data retrieval call binding the contract method 0x91be90c8.
// - Solidity: function _dust(address ) constant returns(uint256)
func (_Oasis *Oasis) Dust(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "_dust", arg0)
	return *ret0, err
}

// Near is a free data retrieval call binding the contract method 0xa78d4316.
// - Solidity: function _near(uint256 ) constant returns(uint256)
func (_Oasis *Oasis) Near(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "_near", arg0)
	return *ret0, err
}

// Rank is a free data retrieval call binding the contract method 0xc2d526aa.
// - Solidity: function _rank(uint256 ) constant returns(uint256 next, uint256 prev, uint256 delb)
func (_Oasis *Oasis) Rank(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Next *big.Int
	Prev *big.Int
	Delb *big.Int
}, error) {
	ret := new(struct {
		Next *big.Int
		Prev *big.Int
		Delb *big.Int
	})
	out := ret
	err := _Oasis.Call(opts, out, "_rank", arg0)
	return *ret, err
}

// Span is a free data retrieval call binding the contract method 0x677170e1.
// - Solidity: function _span(address , address ) constant returns(uint256)
func (_Oasis *Oasis) Span(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "_span", arg0, arg1)
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
// - Solidity: function authority() constant returns(address)
func (_Oasis *Oasis) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "authority")
	return *ret0, err
}

// BuyEnabled is a free data retrieval call binding the contract method 0xf582d293.
// - Solidity: function buyEnabled() constant returns(bool)
func (_Oasis *Oasis) BuyEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "buyEnabled")
	return *ret0, err
}

// CloseTime is a free data retrieval call binding the contract method 0x6377ebca.
// - Solidity: function close_time() constant returns(uint64)
func (_Oasis *Oasis) CloseTime(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "close_time")
	return *ret0, err
}

// DustId is a free data retrieval call binding the contract method 0x56ad8764.
// - Solidity: function dustId() constant returns(uint256)
func (_Oasis *Oasis) DustId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "dustId")
	return *ret0, err
}

// GetBestOffer is a free data retrieval call binding the contract method 0x0374fc6f.
// - Solidity: function getBestOffer(address sell_gem, address buy_gem) constant returns(uint256)
func (_Oasis *Oasis) GetBestOffer(opts *bind.CallOpts, sell_gem common.Address, buy_gem common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "getBestOffer", sell_gem, buy_gem)
	return *ret0, err
}

// GetBetterOffer is a free data retrieval call binding the contract method 0x911550f4.
// - Solidity: function getBetterOffer(uint256 id) constant returns(uint256)
func (_Oasis *Oasis) GetBetterOffer(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "getBetterOffer", id)
	return *ret0, err
}

// GetBuyAmount is a free data retrieval call binding the contract method 0x144a2752.
// - Solidity: function getBuyAmount(address buy_gem, address pay_gem, uint256 pay_amt) constant returns(uint256 fill_amt)
func (_Oasis *Oasis) GetBuyAmount(opts *bind.CallOpts, buy_gem common.Address, pay_gem common.Address, pay_amt *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "getBuyAmount", buy_gem, pay_gem, pay_amt)
	return *ret0, err
}

// GetFirstUnsortedOffer is a free data retrieval call binding the contract method 0x8af82a2e.
// - Solidity: function getFirstUnsortedOffer() constant returns(uint256)
func (_Oasis *Oasis) GetFirstUnsortedOffer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "getFirstUnsortedOffer")
	return *ret0, err
}

// GetMinSell is a free data retrieval call binding the contract method 0x511fa487.
// - Solidity: function getMinSell(address pay_gem) constant returns(uint256)
func (_Oasis *Oasis) GetMinSell(opts *bind.CallOpts, pay_gem common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "getMinSell", pay_gem)
	return *ret0, err
}

// GetNextUnsortedOffer is a free data retrieval call binding the contract method 0x61f54a79.
// - Solidity: function getNextUnsortedOffer(uint256 id) constant returns(uint256)
func (_Oasis *Oasis) GetNextUnsortedOffer(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "getNextUnsortedOffer", id)
	return *ret0, err
}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
// - Solidity: function getOffer(uint256 id) constant returns(uint256, address, uint256, address)
func (_Oasis *Oasis) GetOffer(opts *bind.CallOpts, id *big.Int) (*big.Int, common.Address, *big.Int, common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Oasis.Call(opts, out, "getOffer", id)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetOfferCount is a free data retrieval call binding the contract method 0x7ca9429a.
// - Solidity: function getOfferCount(address sell_gem, address buy_gem) constant returns(uint256)
func (_Oasis *Oasis) GetOfferCount(opts *bind.CallOpts, sell_gem common.Address, buy_gem common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "getOfferCount", sell_gem, buy_gem)
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
// - Solidity: function getOwner(uint256 id) constant returns(address owner)
func (_Oasis *Oasis) GetOwner(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "getOwner", id)
	return *ret0, err
}

// GetPayAmount is a free data retrieval call binding the contract method 0xff1fd974.
// - Solidity: function getPayAmount(address pay_gem, address buy_gem, uint256 buy_amt) constant returns(uint256 fill_amt)
func (_Oasis *Oasis) GetPayAmount(opts *bind.CallOpts, pay_gem common.Address, buy_gem common.Address, buy_amt *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "getPayAmount", pay_gem, buy_gem, buy_amt)
	return *ret0, err
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
// - Solidity: function getTime() constant returns(uint64)
func (_Oasis *Oasis) GetTime(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "getTime")
	return *ret0, err
}

// GetWorseOffer is a free data retrieval call binding the contract method 0x943911bc.
// - Solidity: function getWorseOffer(uint256 id) constant returns(uint256)
func (_Oasis *Oasis) GetWorseOffer(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "getWorseOffer", id)
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
// - Solidity: function isActive(uint256 id) constant returns(bool active)
func (_Oasis *Oasis) IsActive(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "isActive", id)
	return *ret0, err
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
// - Solidity: function isClosed() constant returns(bool closed)
func (_Oasis *Oasis) IsClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "isClosed")
	return *ret0, err
}

// IsOfferSorted is a free data retrieval call binding the contract method 0xd2b420ce.
// - Solidity: function isOfferSorted(uint256 id) constant returns(bool)
func (_Oasis *Oasis) IsOfferSorted(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "isOfferSorted", id)
	return *ret0, err
}

// LastOfferId is a free data retrieval call binding the contract method 0x232cae0b.
// - Solidity: function last_offer_id() constant returns(uint256)
func (_Oasis *Oasis) LastOfferId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "last_offer_id")
	return *ret0, err
}

// MatchingEnabled is a free data retrieval call binding the contract method 0x01492a0b.
// - Solidity: function matchingEnabled() constant returns(bool)
func (_Oasis *Oasis) MatchingEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "matchingEnabled")
	return *ret0, err
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
// - Solidity: function offers(uint256 ) constant returns(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, address owner, uint64 timestamp)
func (_Oasis *Oasis) Offers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PayAmt    *big.Int
	PayGem    common.Address
	BuyAmt    *big.Int
	BuyGem    common.Address
	Owner     common.Address
	Timestamp uint64
}, error) {
	ret := new(struct {
		PayAmt    *big.Int
		PayGem    common.Address
		BuyAmt    *big.Int
		BuyGem    common.Address
		Owner     common.Address
		Timestamp uint64
	})
	out := ret
	err := _Oasis.Call(opts, out, "offers", arg0)
	return *ret, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
// - Solidity: function owner() constant returns(address)
func (_Oasis *Oasis) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "owner")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
// - Solidity: function stopped() constant returns(bool)
func (_Oasis *Oasis) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oasis.Call(opts, out, "stopped")
	return *ret0, err
}

//////////////////////////////////////////////////////
//		Transactions
////////////////////////////////////////////////////

// Bump is a paid mutator transaction binding the contract method 0x779997c3.
// - Solidity: function bump(bytes32 id_) returns()
func (_Oasis *Oasis) Bump(opts *bind.TransactOpts, id_ [32]byte) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "bump", id_)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
// - Solidity: function buy(uint256 id, uint256 amount) returns(bool)
func (_Oasis *Oasis) Buy(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "buy", id, amount)
}

// BuyAllAmount is a paid mutator transaction binding the contract method 0x8185402b.
// - Solidity: function buyAllAmount(address buy_gem, uint256 buy_amt, address pay_gem, uint256 max_fill_amount) returns(uint256 fill_amt)
func (_Oasis *Oasis) BuyAllAmount(opts *bind.TransactOpts, buy_gem common.Address, buy_amt *big.Int, pay_gem common.Address, max_fill_amount *big.Int) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "buyAllAmount", buy_gem, buy_amt, pay_gem, max_fill_amount)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
// - Solidity: function cancel(uint256 id) returns(bool success)
func (_Oasis *Oasis) Cancel(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "cancel", id)
}

// DelRank is a paid mutator transaction binding the contract method 0x467f0b7b.
// - Solidity: function del_rank(uint256 id) returns(bool)
func (_Oasis *Oasis) DelRank(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "del_rank", id)
}

// Insert is a paid mutator transaction binding the contract method 0x1d834a1b.
// - Solidity: function insert(uint256 id, uint256 pos) returns(bool)
func (_Oasis *Oasis) Insert(opts *bind.TransactOpts, id *big.Int, pos *big.Int) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "insert", id, pos)
}

// Kill is a paid mutator transaction binding the contract method 0xb4f9b6c8.
// - Solidity: function kill(bytes32 id) returns()
func (_Oasis *Oasis) Kill(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "kill", id)
}

// Make is a paid mutator transaction binding the contract method 0x093f5198.
// - Solidity: function make(address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt) returns(bytes32)
func (_Oasis *Oasis) Make(opts *bind.TransactOpts, pay_gem common.Address, buy_gem common.Address, pay_amt *big.Int, buy_amt *big.Int) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "make", pay_gem, buy_gem, pay_amt, buy_amt)
}

// Offer is a paid mutator transaction binding the contract method 0x1b33d412.
// - Solidity: function offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, uint256 pos) returns(uint256)
func (_Oasis *Oasis) Offer(opts *bind.TransactOpts, pay_amt *big.Int, pay_gem common.Address, buy_amt *big.Int, buy_gem common.Address, pos *big.Int) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "offer", pay_amt, pay_gem, buy_amt, buy_gem, pos)
}

// Offer0 is a paid mutator transaction binding the contract method 0xe1a6f014.
// - Solidity: function offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, uint256 pos, bool rounding) returns(uint256)
func (_Oasis *Oasis) Offer0(opts *bind.TransactOpts, pay_amt *big.Int, pay_gem common.Address, buy_amt *big.Int, buy_gem common.Address, pos *big.Int, rounding bool) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "offer0", pay_amt, pay_gem, buy_amt, buy_gem, pos, rounding)
}

// Offer1 is a paid mutator transaction binding the contract method 0xf09ea2a6.
// - Solidity: function offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem) returns(uint256)
func (_Oasis *Oasis) Offer1(opts *bind.TransactOpts, pay_amt *big.Int, pay_gem common.Address, buy_amt *big.Int, buy_gem common.Address) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "offer1", pay_amt, pay_gem, buy_amt, buy_gem)
}

// SellAllAmount is a paid mutator transaction binding the contract method 0x0621b4f6.
// - Solidity: function sellAllAmount(address pay_gem, uint256 pay_amt, address buy_gem, uint256 min_fill_amount) returns(uint256 fill_amt)
func (_Oasis *Oasis) SellAllAmount(opts *bind.TransactOpts, pay_gem common.Address, pay_amt *big.Int, buy_gem common.Address, min_fill_amount *big.Int) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "sellAllAmount", pay_gem, pay_amt, buy_gem, min_fill_amount)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
// - Solidity: function setAuthority(address authority_) returns()
func (_Oasis *Oasis) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "setAuthority", authority_)
}

// SetBuyEnabled is a paid mutator transaction binding the contract method 0xd6f15469.
// - Solidity: function setBuyEnabled(bool buyEnabled_) returns(bool)
func (_Oasis *Oasis) SetBuyEnabled(opts *bind.TransactOpts, buyEnabled_ bool) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "setBuyEnabled", buyEnabled_)
}

// SetMatchingEnabled is a paid mutator transaction binding the contract method 0x2aed1905.
// - Solidity: function setMatchingEnabled(bool matchingEnabled_) returns(bool)
func (_Oasis *Oasis) SetMatchingEnabled(opts *bind.TransactOpts, matchingEnabled_ bool) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "setMatchingEnabled", matchingEnabled_)
}

// SetMinSell is a paid mutator transaction binding the contract method 0xbf7c734e.
// - Solidity: function setMinSell(address pay_gem, uint256 dust) returns(bool)
func (_Oasis *Oasis) SetMinSell(opts *bind.TransactOpts, pay_gem common.Address, dust *big.Int) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "setMinSell", pay_gem, dust)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
// - Solidity: function setOwner(address owner_) returns()
func (_Oasis *Oasis) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "setOwner", owner_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
// - Solidity: function stop() returns()
func (_Oasis *Oasis) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "stop")
}

// Take is a paid mutator transaction binding the contract method 0x49606455.
// - Solidity: function take(bytes32 id, uint128 maxTakeAmount) returns()
func (_Oasis *Oasis) Take(opts *bind.TransactOpts, id [32]byte, maxTakeAmount *big.Int) (*types.Transaction, error) {
	return _Oasis.Transact(opts, "take", id, maxTakeAmount)
}

//////////////////////////////////////////////////////
//		Events
////////////////////////////////////////////////////

//////// LogBump ////////
// LogBumpID is the hex of the Topic Hash
const OasisLogBumpID = "0x70a14c213064359ede031fd2a1645a11ce2ec825ffe6ab5cfb5b160c3ef4d0a2"

// LogBumpLog represents a LogBump event raised by the Oasis contract.
type LogBumpLog struct {
	Id        [32]byte
	Pair      [32]byte
	Maker     common.Address
	PayGem    common.Address
	BuyGem    common.Address
	PayAmt    *big.Int
	BuyAmt    *big.Int
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// UnpackLogBumpLog is a log parse operation binding the contract event 0x70a14c213064359ede031fd2a1645a11ce2ec825ffe6ab5cfb5b160c3ef4d0a2
// Solidity: event LogBump(bytes32 indexed id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt, uint64 timestamp)
func (_Oasis *Oasis) UnpackLogBumpLog(log types.Log) (*LogBumpLog, error) {
	event := new(LogBumpLog)
	if err := _Oasis.UnpackLog(event, "LogBump", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogBuyEnabled ////////
// LogBuyEnabledID is the hex of the Topic Hash
const OasisLogBuyEnabledID = "0x7089e4f0bcc948f9f723a361590c32d9c2284da7ab1981b1249ad2edb9f953c1"

// LogBuyEnabledLog represents a LogBuyEnabled event raised by the Oasis contract.
type LogBuyEnabledLog struct {
	IsEnabled bool
	Raw       types.Log // Blockchain specific contextual infos
}

// UnpackLogBuyEnabledLog is a log parse operation binding the contract event 0x7089e4f0bcc948f9f723a361590c32d9c2284da7ab1981b1249ad2edb9f953c1
// Solidity: event LogBuyEnabled(bool isEnabled)
func (_Oasis *Oasis) UnpackLogBuyEnabledLog(log types.Log) (*LogBuyEnabledLog, error) {
	event := new(LogBuyEnabledLog)
	if err := _Oasis.UnpackLog(event, "LogBuyEnabled", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogDelete ////////
// LogDeleteID is the hex of the Topic Hash
const OasisLogDeleteID = "0xcb9d6176c6aac6478ebb9a2754cdce22a944de29ed1f2642f8613884eba4b40c"

// LogDeleteLog represents a LogDelete event raised by the Oasis contract.
type LogDeleteLog struct {
	Keeper common.Address
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// UnpackLogDeleteLog is a log parse operation binding the contract event 0xcb9d6176c6aac6478ebb9a2754cdce22a944de29ed1f2642f8613884eba4b40c
// Solidity: event LogDelete(address keeper, uint256 id)
func (_Oasis *Oasis) UnpackLogDeleteLog(log types.Log) (*LogDeleteLog, error) {
	event := new(LogDeleteLog)
	if err := _Oasis.UnpackLog(event, "LogDelete", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogInsert ////////
// LogInsertID is the hex of the Topic Hash
const OasisLogInsertID = "0x6d5c16212bdea16850dce4d9fa2314c446bd30ce84700d9c36c7677c6d283940"

// LogInsertLog represents a LogInsert event raised by the Oasis contract.
type LogInsertLog struct {
	Keeper common.Address
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// UnpackLogInsertLog is a log parse operation binding the contract event 0x6d5c16212bdea16850dce4d9fa2314c446bd30ce84700d9c36c7677c6d283940
// Solidity: event LogInsert(address keeper, uint256 id)
func (_Oasis *Oasis) UnpackLogInsertLog(log types.Log) (*LogInsertLog, error) {
	event := new(LogInsertLog)
	if err := _Oasis.UnpackLog(event, "LogInsert", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogItemUpdate ////////
// LogItemUpdateID is the hex of the Topic Hash
const OasisLogItemUpdateID = "0xa2c251311b1a7a475913900a2a73dc9789a21b04bc737e050bbc506dd4eb3488"

// LogItemUpdateLog represents a LogItemUpdate event raised by the Oasis contract.
type LogItemUpdateLog struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// UnpackLogItemUpdateLog is a log parse operation binding the contract event 0xa2c251311b1a7a475913900a2a73dc9789a21b04bc737e050bbc506dd4eb3488
// Solidity: event LogItemUpdate(uint256 id)
func (_Oasis *Oasis) UnpackLogItemUpdateLog(log types.Log) (*LogItemUpdateLog, error) {
	event := new(LogItemUpdateLog)
	if err := _Oasis.UnpackLog(event, "LogItemUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogKill ////////
// LogKillID is the hex of the Topic Hash
const OasisLogKillID = "0x9577941d28fff863bfbee4694a6a4a56fb09e169619189d2eaa750b5b4819995"

// LogKillLog represents a LogKill event raised by the Oasis contract.
type LogKillLog struct {
	Id        [32]byte
	Pair      [32]byte
	Maker     common.Address
	PayGem    common.Address
	BuyGem    common.Address
	PayAmt    *big.Int
	BuyAmt    *big.Int
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// UnpackLogKillLog is a log parse operation binding the contract event 0x9577941d28fff863bfbee4694a6a4a56fb09e169619189d2eaa750b5b4819995
// Solidity: event LogKill(bytes32 indexed id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt, uint64 timestamp)
func (_Oasis *Oasis) UnpackLogKillLog(log types.Log) (*LogKillLog, error) {
	event := new(LogKillLog)
	if err := _Oasis.UnpackLog(event, "LogKill", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogMake ////////
// LogMakeID is the hex of the Topic Hash
const OasisLogMakeID = "0x773ff502687307abfa024ac9f62f9752a0d210dac2ffd9a29e38e12e2ea82c82"

// LogMakeLog represents a LogMake event raised by the Oasis contract.
type LogMakeLog struct {
	Id        [32]byte
	Pair      [32]byte
	Maker     common.Address
	PayGem    common.Address
	BuyGem    common.Address
	PayAmt    *big.Int
	BuyAmt    *big.Int
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// UnpackLogMakeLog is a log parse operation binding the contract event 0x773ff502687307abfa024ac9f62f9752a0d210dac2ffd9a29e38e12e2ea82c82
// Solidity: event LogMake(bytes32 indexed id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt, uint64 timestamp)
func (_Oasis *Oasis) UnpackLogMakeLog(log types.Log) (*LogMakeLog, error) {
	event := new(LogMakeLog)
	if err := _Oasis.UnpackLog(event, "LogMake", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogMatchingEnabled ////////
// LogMatchingEnabledID is the hex of the Topic Hash
const OasisLogMatchingEnabledID = "0xea11e00ec1642be9b494019b756440e2c57dbe9e59242c4f9c64ce33fb4f41d9"

// LogMatchingEnabledLog represents a LogMatchingEnabled event raised by the Oasis contract.
type LogMatchingEnabledLog struct {
	IsEnabled bool
	Raw       types.Log // Blockchain specific contextual infos
}

// UnpackLogMatchingEnabledLog is a log parse operation binding the contract event 0xea11e00ec1642be9b494019b756440e2c57dbe9e59242c4f9c64ce33fb4f41d9
// Solidity: event LogMatchingEnabled(bool isEnabled)
func (_Oasis *Oasis) UnpackLogMatchingEnabledLog(log types.Log) (*LogMatchingEnabledLog, error) {
	event := new(LogMatchingEnabledLog)
	if err := _Oasis.UnpackLog(event, "LogMatchingEnabled", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogMinSell ////////
// LogMinSellID is the hex of the Topic Hash
const OasisLogMinSellID = "0xc28d56449b0bb31e64ee7487e061f57a2e72aea8019d810832f26dda099823d0"

// LogMinSellLog represents a LogMinSell event raised by the Oasis contract.
type LogMinSellLog struct {
	PayGem    common.Address
	MinAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// UnpackLogMinSellLog is a log parse operation binding the contract event 0xc28d56449b0bb31e64ee7487e061f57a2e72aea8019d810832f26dda099823d0
// Solidity: event LogMinSell(address pay_gem, uint256 min_amount)
func (_Oasis *Oasis) UnpackLogMinSellLog(log types.Log) (*LogMinSellLog, error) {
	event := new(LogMinSellLog)
	if err := _Oasis.UnpackLog(event, "LogMinSell", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogSetAuthority ////////
// LogSetAuthorityID is the hex of the Topic Hash
const OasisLogSetAuthorityID = "0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4"

// LogSetAuthorityLog represents a LogSetAuthority event raised by the Oasis contract.
type LogSetAuthorityLog struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// UnpackLogSetAuthorityLog is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4
// Solidity: event LogSetAuthority(address indexed authority)
func (_Oasis *Oasis) UnpackLogSetAuthorityLog(log types.Log) (*LogSetAuthorityLog, error) {
	event := new(LogSetAuthorityLog)
	if err := _Oasis.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogSetOwner ////////
// LogSetOwnerID is the hex of the Topic Hash
const OasisLogSetOwnerID = "0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94"

// LogSetOwnerLog represents a LogSetOwner event raised by the Oasis contract.
type LogSetOwnerLog struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// UnpackLogSetOwnerLog is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94
// Solidity: event LogSetOwner(address indexed owner)
func (_Oasis *Oasis) UnpackLogSetOwnerLog(log types.Log) (*LogSetOwnerLog, error) {
	event := new(LogSetOwnerLog)
	if err := _Oasis.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogSortedOffer ////////
// LogSortedOfferID is the hex of the Topic Hash
const OasisLogSortedOfferID = "0x20fb9bad86c18f7e22e8065258790d9416a7d2df8ff05f80f82c46d38b925acd"

// LogSortedOfferLog represents a LogSortedOffer event raised by the Oasis contract.
type LogSortedOfferLog struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// UnpackLogSortedOfferLog is a log parse operation binding the contract event 0x20fb9bad86c18f7e22e8065258790d9416a7d2df8ff05f80f82c46d38b925acd
// Solidity: event LogSortedOffer(uint256 id)
func (_Oasis *Oasis) UnpackLogSortedOfferLog(log types.Log) (*LogSortedOfferLog, error) {
	event := new(LogSortedOfferLog)
	if err := _Oasis.UnpackLog(event, "LogSortedOffer", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogTake ////////
// LogTakeID is the hex of the Topic Hash
const OasisLogTakeID = "0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f"

// LogTakeLog represents a LogTake event raised by the Oasis contract.
type LogTakeLog struct {
	Id        [32]byte
	Pair      [32]byte
	Maker     common.Address
	PayGem    common.Address
	BuyGem    common.Address
	Taker     common.Address
	TakeAmt   *big.Int
	GiveAmt   *big.Int
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// UnpackLogTakeLog is a log parse operation binding the contract event 0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f
// Solidity: event LogTake(bytes32 id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, address indexed taker, uint128 take_amt, uint128 give_amt, uint64 timestamp)
func (_Oasis *Oasis) UnpackLogTakeLog(log types.Log) (*LogTakeLog, error) {
	event := new(LogTakeLog)
	if err := _Oasis.UnpackLog(event, "LogTake", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogTrade ////////
// LogTradeID is the hex of the Topic Hash
const OasisLogTradeID = "0x819e390338feffe95e2de57172d6faf337853dfd15c7a09a32d76f7fd2443875"

// LogTradeLog represents a LogTrade event raised by the Oasis contract.
type LogTradeLog struct {
	PayAmt *big.Int
	PayGem common.Address
	BuyAmt *big.Int
	BuyGem common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// UnpackLogTradeLog is a log parse operation binding the contract event 0x819e390338feffe95e2de57172d6faf337853dfd15c7a09a32d76f7fd2443875
// Solidity: event LogTrade(uint256 pay_amt, address indexed pay_gem, uint256 buy_amt, address indexed buy_gem)
func (_Oasis *Oasis) UnpackLogTradeLog(log types.Log) (*LogTradeLog, error) {
	event := new(LogTradeLog)
	if err := _Oasis.UnpackLog(event, "LogTrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogUnsortedOffer ////////
// LogUnsortedOfferID is the hex of the Topic Hash
const OasisLogUnsortedOfferID = "0x8173832a493e0a3989e521458e55bfe9feac9f9b675a94e100b9d5a85f814862"

// LogUnsortedOfferLog represents a LogUnsortedOffer event raised by the Oasis contract.
type LogUnsortedOfferLog struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// UnpackLogUnsortedOfferLog is a log parse operation binding the contract event 0x8173832a493e0a3989e521458e55bfe9feac9f9b675a94e100b9d5a85f814862
// Solidity: event LogUnsortedOffer(uint256 id)
func (_Oasis *Oasis) UnpackLogUnsortedOfferLog(log types.Log) (*LogUnsortedOfferLog, error) {
	event := new(LogUnsortedOfferLog)
	if err := _Oasis.UnpackLog(event, "LogUnsortedOffer", log); err != nil {
		return nil, err
	}
	return event, nil
}

/*
Mux can be copied and pasted to save ya a quick minute when distinguishing between log data
func Mux(c *Oasis, log types.Log) error {
	switch log.Topics[0].Hex() {
	case oasis.OasisLogBumpID:
		ulog, err := c.UnpackLogBumpLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogBuyEnabledID:
		ulog, err := c.UnpackLogBuyEnabledLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogDeleteID:
		ulog, err := c.UnpackLogDeleteLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogInsertID:
		ulog, err := c.UnpackLogInsertLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogItemUpdateID:
		ulog, err := c.UnpackLogItemUpdateLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogKillID:
		ulog, err := c.UnpackLogKillLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogMakeID:
		ulog, err := c.UnpackLogMakeLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogMatchingEnabledID:
		ulog, err := c.UnpackLogMatchingEnabledLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogMinSellID:
		ulog, err := c.UnpackLogMinSellLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogSetAuthorityID:
		ulog, err := c.UnpackLogSetAuthorityLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogSetOwnerID:
		ulog, err := c.UnpackLogSetOwnerLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogSortedOfferID:
		ulog, err := c.UnpackLogSortedOfferLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogTakeID:
		ulog, err := c.UnpackLogTakeLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogTradeID:
		ulog, err := c.UnpackLogTradeLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case oasis.OasisLogUnsortedOfferID:
		ulog, err := c.UnpackLogUnsortedOfferLog(log)
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
// OasisBin is used to deploy the generated contract
const OasisBin = "0x60806040526004805460ff60581b1960ff60501b199091166a010000000000000000000017166b01000000000000000000000017905534801561004157600080fd5b506040516135063803806135068339818101604052602081101561006457600080fd5b5051600180546001600160a01b031916339081179091556040518291907fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9490600090a2600480546001600160401b0390921661010002610100600160481b031990921691909117905550613429806100dd6000396000f3fe608060405234801561001057600080fd5b50600436106102a05760003560e01c80637a9e5e4b11610167578063bf7c734e116100ce578063d6f1546911610087578063d6f15469146108a6578063d6febde8146108c5578063e1a6f014146108e8578063f09ea2a614610930578063f582d2931461096c578063ff1fd97414610974576102a0565b8063bf7c734e146107f5578063bf7e214f14610821578063c2b6b58c14610829578063c2d526aa14610831578063c41a360a1461086c578063d2b420ce14610889576102a0565b80638da5cb5b116101205780638da5cb5b14610737578063911550f41461075b57806391be90c814610778578063943911bc1461079e578063a78d4316146107bb578063b4f9b6c8146107d8576102a0565b80637a9e5e4b1461061e5780637ca9429a146106445780638185402b1461067257806382afd23b146106ac5780638a72ea6a146106c95780638af82a2e1461072f576102a0565b80634579268a1161020b57806361f54a79116101c457806361f54a79146105785780636377ebca14610595578063677170e11461059d57806374c1d7d3146105cb57806375f12b21146105f9578063779997c314610601576102a0565b80634579268a1461048f578063467f0b7b146104dd57806349606455146104fa578063511fa48714610526578063557ed1ba1461054c57806356ad876414610570576102a0565b8063144a27521161025d578063144a2752146103b25780631b33d412146103e85780631d834a1b14610428578063232cae0b1461044b5780632aed19051461045357806340e58ee514610472576102a0565b806301492a0b146102a55780630374fc6f146102c15780630621b4f61461030157806307da68f51461033b578063093f51981461034557806313af40351461038c575b600080fd5b6102ad6109aa565b604080519115158252519081900360200190f35b6102ef600480360360408110156102d757600080fd5b506001600160a01b03813581169160200135166109ba565b60408051918252519081900360200190f35b6102ef6004803603608081101561031757600080fd5b506001600160a01b03813581169160208101359160408201351690606001356109e7565b610343610b72565b005b6102ef6004803603608081101561035b57600080fd5b506001600160a01b0381358116916020810135909116906001600160801b0360408201358116916060013516610beb565b610343600480360360208110156103a257600080fd5b50356001600160a01b0316610c14565b6102ef600480360360608110156103c857600080fd5b506001600160a01b03813581169160208101359091169060400135610cc2565b6102ef600480360360a08110156103fe57600080fd5b508035906001600160a01b0360208201358116916040810135916060820135169060800135610d87565b6102ad6004803603604081101561043e57600080fd5b5080359060200135610db4565b6102ef610e81565b6102ad6004803603602081101561046957600080fd5b50351515610e87565b6102ad6004803603602081101561048857600080fd5b5035610f4b565b6104ac600480360360208110156104a557600080fd5b50356110ab565b604080519485526001600160a01b039384166020860152848101929092529091166060830152519081900360800190f35b6102ad600480360360208110156104f357600080fd5b5035611138565b6103436004803603604081101561051057600080fd5b50803590602001356001600160801b0316611239565b6102ef6004803603602081101561053c57600080fd5b50356001600160a01b0316611259565b610554611274565b604080516001600160401b039092168252519081900360200190f35b6102ef611278565b6102ef6004803603602081101561058e57600080fd5b503561127e565b610554611290565b6102ef600480360360408110156105b357600080fd5b506001600160a01b03813581169160200135166112a4565b6102ef600480360360408110156105e157600080fd5b506001600160a01b03813581169160200135166112c1565b6102ad6112de565b6103436004803603602081101561061757600080fd5b50356112ee565b6103436004803603602081101561063457600080fd5b50356001600160a01b03166113fa565b6102ef6004803603604081101561065a57600080fd5b506001600160a01b03813581169160200135166114a4565b6102ef6004803603608081101561068857600080fd5b506001600160a01b03813581169160208101359160408201351690606001356114cf565b6102ad600480360360208110156106c257600080fd5b503561163a565b6106e6600480360360208110156106df57600080fd5b5035611661565b604080519687526001600160a01b03958616602088015286810194909452918416606086015290921660808401526001600160401b0390911660a0830152519081900360c00190f35b6102ef6116b2565b61073f6116b8565b604080516001600160a01b039092168252519081900360200190f35b6102ef6004803603602081101561077157600080fd5b50356116c7565b6102ef6004803603602081101561078e57600080fd5b50356001600160a01b03166116d9565b6102ef600480360360208110156107b457600080fd5b50356116eb565b6102ef600480360360208110156107d157600080fd5b5035611700565b610343600480360360208110156107ee57600080fd5b5035611712565b6102ad6004803603604081101561080b57600080fd5b506001600160a01b038135169060200135611727565b61073f61184c565b6102ad61185b565b61084e6004803603602081101561084757600080fd5b503561189a565b60408051938452602084019290925282820152519081900360600190f35b61073f6004803603602081101561088257600080fd5b50356118bb565b6102ad6004803603602081101561089f57600080fd5b50356118d9565b6102ad600480360360208110156108bc57600080fd5b5035151561194a565b6102ad600480360360408110156108db57600080fd5b5080359060200135611a0d565b6102ef600480360360c08110156108fe57600080fd5b508035906001600160a01b0360208201358116916040810135916060820135169060808101359060a001351515611ab2565b6102ef6004803603608081101561094657600080fd5b508035906001600160a01b03602082013581169160408101359160609091013516611b76565b6102ad611bf9565b6102ef6004803603606081101561098a57600080fd5b506001600160a01b03813581169160208101359091169060400135611c09565b600454600160581b900460ff1681565b6001600160a01b038083166000908152600660209081526040808320938516835292905220545b92915050565b60045460009060ff1615610a37576040805162461bcd60e51b81526020600482015260126024820152711499595b9d1c985b98de48185d1d195b5c1d60721b604482015290519081900360640190fd5b60005b8415610b5c57610a4a84876109ba565b905080610a5657600080fd5b600081815260036020526040902060028101549054610a759190611cbd565b85670de0b6b3a7640000021015610a8b57610b5c565b6000818152600360205260409020600201548510610aff57600081815260036020526040902054610abd908390611cec565b600082815260036020526040902060020154909250610add908690611d3b565b600082815260036020526040902054909550610afa908290611239565b610b57565b60008181526003602052604081208054600290910154633b9aca0091610b319189840291610b2c91611d8b565b611da7565b81610b3857fe5b049050610b458382611cec565b9250610b518282611239565b60009550505b610a3a565b82821015610b6957600080fd5b50949350505050565b610b88336000356001600160e01b031916611dd7565b610bd0576040805162461bcd60e51b8152602060048201526014602482015273191ccb585d5d1a0b5d5b985d5d1a1bdc9a5e995960621b604482015290519081900360640190fd5b6004805469ff0000000000000000001916600160481b179055565b6000610c0b836001600160801b031686846001600160801b031687611b76565b95945050505050565b610c2a336000356001600160e01b031916611dd7565b610c72576040805162461bcd60e51b8152602060048201526014602482015273191ccb585d5d1a0b5d5b985d5d1a1bdc9a5e995960621b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0383811691909117918290556040519116907fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9490600090a250565b600080610ccf85856109ba565b90505b600081815260036020526040902060020154831115610d4757600081815260036020526040902054610d05908390611cec565b600082815260036020526040902060020154909250610d25908490611d3b565b92508215610d4257610d36816116eb565b905080610d4257600080fd5b610cd2565b60008181526003602052604090208054600290910154610c0b918491633b9aca0091610d7a9188840291610b2c91611d8b565b81610d8157fe5b04611cec565b6000610d9161185b565b15610d9b57600080fd5b610daa86868686866001611ab2565b9695505050505050565b60045460009060ff1615610e04576040805162461bcd60e51b81526020600482015260126024820152711499595b9d1c985b98de48185d1d195b5c1d60721b604482015290519081900360640190fd5b610e0d836118d9565b15610e1757600080fd5b610e208361163a565b610e2957600080fd5b610e3283611ebe565b50610e3d8383611f68565b604080513381526020810185905281517f6d5c16212bdea16850dce4d9fa2314c446bd30ce84700d9c36c7677c6d283940929181900390910190a150600192915050565b60025481565b6000610e9f336000356001600160e01b031916611dd7565b610ee7576040805162461bcd60e51b8152602060048201526014602482015273191ccb585d5d1a0b5d5b985d5d1a1bdc9a5e995960621b604482015290519081900360640190fd5b60048054831515600160581b90810260ff60581b199092169190911791829055604080519190920460ff161515815290517fea11e00ec1642be9b494019b756440e2c57dbe9e59242c4f9c64ce33fb4f41d99181900360200190a15060015b919050565b600081610f578161163a565b610f925760405162461bcd60e51b815260040180806020018281038252602d815260200180613354602d913960400191505060405180910390fd5b610f9a61185b565b80610fbe5750610fa9816118bb565b6001600160a01b0316336001600160a01b0316145b80610fca5750600b5481145b6110055760405162461bcd60e51b81526004018080602001828103825260748152602001806133816074913960800191505060405180910390fd5b60045460ff1615611052576040805162461bcd60e51b81526020600482015260126024820152711499595b9d1c985b98de48185d1d195b5c1d60721b604482015290519081900360640190fd5b600454600160581b900460ff161561109b5761106d836118d9565b156110895761107b83612109565b61108457600080fd5b61109b565b61109283611ebe565b61109b57600080fd5b6110a4836122c5565b9392505050565b6000806000806110b961331c565b5050506000928352505060036020818152604092839020835160c081018552815480825260018301546001600160a01b039081169483018590526002840154968301879052948301548516606083018190526004909301549485166080830152600160a01b9094046001600160401b031660a090910152919390929190565b60045460009060ff1615611188576040805162461bcd60e51b81526020600482015260126024820152711499595b9d1c985b98de48185d1d195b5c1d60721b604482015290519081900360640190fd5b6111918261163a565b1580156111ae575060008281526005602052604090206002015415155b80156111ce57506000828152600560205260409020600201546009194301115b6111d757600080fd5b60008281526005602090815260408083208381556001810184905560020192909255815133815290810184905281517fcb9d6176c6aac6478ebb9a2754cdce22a944de29ed1f2642f8613884eba4b40c929181900390910190a1506001919050565b61124c826001600160801b038316611a0d565b61125557600080fd5b5050565b6001600160a01b031660009081526008602052604090205490565b4290565b600b5481565b60009081526009602052604090205490565b60045461010090046001600160401b031681565b600760209081526000928352604080842090915290825290205481565b600660209081526000928352604080842090915290825290205481565b600454600160481b900460ff1681565b806112f88161163a565b61130157600080fd5b61130961185b565b1561131357600080fd5b6000828152600360208181526040808420600481015460018201548286015484516001600160601b0319606084811b8216838a015283901b1660348201528551602881830301815260488201808852815191890191909120998c90529790965283546002909401546001600160a01b03928316909752811660688601526001600160801b0392831660888601529190941660a88401526001600160401b03600160a01b85041660c8840152905186949190931692909184917f70a14c213064359ede031fd2a1645a11ce2ec825ffe6ab5cfb5b160c3ef4d0a29181900360e80190a4505050565b611410336000356001600160e01b031916611dd7565b611458576040805162461bcd60e51b8152602060048201526014602482015273191ccb585d5d1a0b5d5b985d5d1a1bdc9a5e995960621b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b03838116919091178083556040519116917f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada491a250565b6001600160a01b03918216600090815260076020908152604080832093909416825291909152205490565b60045460009060ff161561151f576040805162461bcd60e51b81526020600482015260126024820152711499595b9d1c985b98de48185d1d195b5c1d60721b604482015290519081900360640190fd5b60005b841561162d5761153286856109ba565b90508061153e57600080fd5b6000818152600360205260409020805460029091015461155e9190611cbd565b85670de0b6b3a76400000210156115745761162d565b60008181526003602052604090205485106115e5576000818152600360205260409020600201546115a6908390611cec565b6000828152600360205260409020549092506115c3908690611d3b565b6000828152600360205260409020549095506115e0908290611239565b611628565b600081815260036020526040902060028101549054611617918491633b9aca0091610d7a918a840291610b2c91611d8b565b91506116238186611239565b600094505b611522565b82821115610b6957600080fd5b600090815260036020526040902060040154600160a01b90046001600160401b0316151590565b60036020819052600091825260409091208054600182015460028301549383015460049093015491936001600160a01b0391821693909290821691811690600160a01b90046001600160401b031686565b600a5490565b6001546001600160a01b031681565b60009081526005602052604090205490565b60086020526000908152604090205481565b60009081526005602052604090206001015490565b60096020526000908152604090205481565b61171b81610f4b565b61172457600080fd5b50565b600061173f336000356001600160e01b031916611dd7565b611787576040805162461bcd60e51b8152602060048201526014602482015273191ccb585d5d1a0b5d5b985d5d1a1bdc9a5e995960621b604482015290519081900360640190fd5b604080513480825260208201838152369383018490526004359360243593849286923392600080356001600160e01b03191693889391929060608201848480828437600083820152604051601f909101601f1916909201829003965090945050505050a46001600160a01b0386166000818152600860209081526040918290208890558151928352820187905280517fc28d56449b0bb31e64ee7487e061f57a2e72aea8019d810832f26dda099823d09281900390910190a150600195945050505050565b6000546001600160a01b031681565b600454600090600160481b900460ff1680611895575060045461010090046001600160401b031661188a611274565b6001600160401b0316115b905090565b60056020526000908152604090208054600182015460029092015490919083565b6000908152600360205260409020600401546001600160a01b031690565b600081815260056020526040812054151580611905575060008281526005602052604090206001015415155b806109e1575050600081815260036020818152604080842060018101546001600160a01b03908116865260068452828620919094015490931684529190529020541490565b6000611962336000356001600160e01b031916611dd7565b6119aa576040805162461bcd60e51b8152602060048201526014602482015273191ccb585d5d1a0b5d5b985d5d1a1bdc9a5e995960621b604482015290519081900360640190fd5b60048054831515600160501b90810260ff60501b199092169190911791829055604080519190920460ff161515815290517f7089e4f0bcc948f9f723a361590c32d9c2284da7ab1981b1249ad2edb9f953c19181900360200190a1506001919050565b600082611a198161163a565b611a2257600080fd5b611a2a61185b565b15611a3457600080fd5b60045460ff1615611a81576040805162461bcd60e51b81526020600482015260126024820152711499595b9d1c985b98de48185d1d195b5c1d60721b604482015290519081900360640190fd5b60045461335190600160581b900460ff16611a9e576126d3611aa2565b612b2d5b9050610c0b85858363ffffffff16565b6000611abc61185b565b15611ac657600080fd5b60045460ff1615611b13576040805162461bcd60e51b81526020600482015260126024820152711499595b9d1c985b98de48185d1d195b5c1d60721b604482015290519081900360640190fd5b6001600160a01b038616600090815260086020526040902054871015611b3857600080fd5b600454600160581b900460ff1615611b5f57611b58878787878787612bf8565b9050610daa565b611b6b87878787612d5a565b979650505050505050565b60045460009060ff1615611bc6576040805162461bcd60e51b81526020600482015260126024820152711499595b9d1c985b98de48185d1d195b5c1d60721b604482015290519081900360640190fd5b60045461335190600160581b900460ff16611be357612d5a611be7565b6130695b9050610daa868686868563ffffffff16565b600454600160501b900460ff1681565b600080611c1684866109ba565b90505b600081815260036020526040902054831115611c8b57600081815260036020526040902060020154611c4c908390611cec565b600082815260036020526040902054909250611c69908490611d3b565b92508215611c8657611c7a816116eb565b905080611c8657600080fd5b611c19565b600081815260036020526040902060028101549054610c0b918491633b9aca0091610d7a9188840291610b2c91611d8b565b600081611cdd611cd585670de0b6b3a76400006130f3565b600285610d81565b81611ce457fe5b049392505050565b808201828110156109e1576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b604482015290519081900360640190fd5b808203828111156109e1576040805162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b604482015290519081900360640190fd5b600081611cdd611cd5856b033b2e3c9fd0803ce80000006130f3565b60006b033b2e3c9fd0803ce8000000611cdd611dc385856130f3565b60026b033b2e3c9fd0803ce8000000610d81565b60006001600160a01b038316301415611df2575060016109e1565b6001546001600160a01b0384811691161415611e10575060016109e1565b6000546001600160a01b0316611e28575060006109e1565b6000546040805163b700961360e01b81526001600160a01b0386811660048301523060248301526001600160e01b0319861660448301529151919092169163b7009613916064808301926020929190829003018186803b158015611e8b57600080fd5b505afa158015611e9f573d6000803e3d6000fd5b505050506040513d6020811015611eb557600080fd5b505190506109e1565b600a5460009080611ece846118d9565b15611ed857600080fd5b83600a541415611f005750505060008181526009602052604081208054600a55556001610f46565b5b600082118015611f115750838214155b15611f2c575060008181526009602052604090205490611f01565b838214611f3e57600092505050610f46565b60008481526009602052604080822080549383529082209290925584815290555060019050919050565b611f718261163a565b611f7a57600080fd5b60008281526003602081905260408220908101546001909101546001600160a01b0391821692911690831580611fcd57506000848152600360205260409020600101546001600160a01b03838116911614155b80611ff65750600084815260036020819052604090912001546001600160a01b03848116911614155b806120075750612005846118d9565b155b61201a576120158585613156565b612023565b61202385613228565b935083156120515750600083815260056020526040808220600101805490879055868352912084905561207e565b506001600160a01b0381811660009081526006602090815260408083209386168352929052208054908590555b80156120a25760008181526005602052604080822087905586825290206001018190555b6001600160a01b03808316600090815260076020908152604080832093871683529281529082902080546001019055815187815291517f20fb9bad86c18f7e22e8065258790d9416a7d2df8ff05f80f82c46d38b925acd9281900390910190a15050505050565b6000818152600360208181526040808420928301546001909301546001600160a01b039081168086526007845282862091909416808652925283205490919061215157600080fd5b6000848152600560205260409020600201541580156121745750612174846118d9565b61217d57600080fd5b6001600160a01b0380821660009081526006602090815260408083209386168352929052205484146121f05760008481526005602052604080822054825290206001015484146121cc57600080fd5b60008481526005602052604080822060018082015491548452919092200155612226565b6000848152600560209081526040808320600101546001600160a01b038086168552600684528285209087168552909252909120555b6000848152600560205260409020600101541561227d57600084815260056020526040808220600101548252902054841461226057600080fd5b600084815260056020526040808220805460019091015483529120555b6001600160a01b039081166000908152600760209081526040808320949093168252928352818120805460001901905593845260059091529091204360029091015550600190565b6000816122d18161163a565b61230c5760405162461bcd60e51b815260040180806020018281038252602d815260200180613354602d913960400191505060405180910390fd5b61231461185b565b806123385750612323816118bb565b6001600160a01b0316336001600160a01b0316145b806123445750600b5481145b61237f5760405162461bcd60e51b81526004018080602001828103825260748152602001806133816074913960800191505060405180910390fd5b60045460ff161561238f57600080fd5b6004805460ff191660011790556123a461331c565b600360008581526020019081526020016000206040518060c0016040529081600082015481526020016001820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001600282015481526020016003820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016004820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016004820160149054906101000a90046001600160401b03166001600160401b03166001600160401b0316815250509050600360008581526020019081526020016000206000808201600090556001820160006101000a8154906001600160a01b03021916905560028201600090556003820160006101000a8154906001600160a01b0302191690556004820160006101000a8154906001600160a01b0302191690556004820160146101000a8154906001600160401b030219169055505080602001516001600160a01b031663a9059cbb826080015183600001516040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561258e57600080fd5b505af11580156125a2573d6000803e3d6000fd5b505050506040513d60208110156125b857600080fd5b50516125c357600080fd5b6040805185815290517fa2c251311b1a7a475913900a2a73dc9789a21b04bc737e050bbc506dd4eb34889181900360200190a160808101516020808301805160608086018051604080516001600160601b031995851b8616818901529190931b9093166034840152815160288185030181526048840180845281519190960120935190518751838901516001600160a01b0393841690975290821660688501526001600160801b03908116608885015290941660a88301526001600160401b03421660c8830152519290931692909187917f9577941d28fff863bfbee4694a6a4a56fb09e169619189d2eaa750b5b4819995919081900360e80190a450506004805460ff19169055506001919050565b6000826126df8161163a565b6126e857600080fd5b6126f061185b565b156126fa57600080fd5b60045460ff161561270a57600080fd5b6004805460ff1916600117905561271f61331c565b506000848152600360208181526040808420815160c081018352815480825260018301546001600160a01b0390811695830195909552600283015493820184905294820154841660608201526004909101549283166080820152600160a01b9092046001600160401b031660a08301529092919061279e9087906130f3565b816127a557fe5b04905080816001600160801b0316146127bd57600080fd5b84856001600160801b0316146127d257600080fd5b8415806127dd575080155b806127e85750815185115b806127f65750816040015181115b1561280657600093505050612b1c565b81516128129086611d3b565b60008781526003602052604090819020919091558201516128339082611d3b565b6000878152600360209081526040808320600201939093556060850151608086015184516323b872dd60e01b81523360048201526001600160a01b0391821660248201526044810187905294519116936323b872dd9360648083019493928390030190829087803b1580156128a757600080fd5b505af11580156128bb573d6000803e3d6000fd5b505050506040513d60208110156128d157600080fd5b50516128dc57600080fd5b6020808301516040805163a9059cbb60e01b81523360048201526024810189905290516001600160a01b039092169263a9059cbb926044808401938290030181600087803b15801561292d57600080fd5b505af1158015612941573d6000803e3d6000fd5b505050506040513d602081101561295757600080fd5b505161296257600080fd5b6040805187815290517fa2c251311b1a7a475913900a2a73dc9789a21b04bc737e050bbc506dd4eb34889181900360200190a160808201516020808401805160608087018051604080516001600160601b031995851b8616818901529190931b909316603484015281516028818503018152604884018084528151919096012093519051948c90526001600160a01b03908116606884015293841660888301526001600160801b03808b1660a8840152861660c88301526001600160401b03421660e8830152513394909316927f3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f918190036101080190a481606001516001600160a01b031682602001516001600160a01b03167f819e390338feffe95e2de57172d6faf337853dfd15c7a09a32d76f7fd24438758784604051808381526020018281526020019250505060405180910390a3600086815260036020526040902054612b1557600086815260036020819052604082208281556001810180546001600160a01b03199081169091556002820193909355908101805490921690915560040180546001600160e01b03191690555b6001935050505b506004805460ff1916905592915050565b600454600090600160501b900460ff16612b4657600080fd5b600083815260036020526040902054821415612b8457612b65836118d9565b15612b7957612b7383612109565b50612b84565b612b8283611ebe565b505b612b8e83836126d3565b612b9757600080fd5b612ba08361163a565b8015612bda5750600083815260036020818152604080842060018101546001600160a01b0316855260088352908420549387905291905254105b15612bef57600b839055612bed83610f4b565b505b50600192915050565b60008060008060005b6001600160a01b038089166000908152600660209081526040808320938e168352929052205415612cfb5750506001600160a01b038087166000908152600660209081526040808320938c168352928152828220548083526003909152919020600281015490549193509085612c78576000612c80565b808b8a840101015b612c8a8c836130f3565b01612c95838b6130f3565b1115612ca057612cfb565b612cb384612cae838c6132a7565b611a0d565b50889250612cca89612cc5838c6132a7565b611d3b565b985082612cd78a8d6130f3565b81612cde57fe5b049a508a1580612cec575088155b15612cf657612cfb565b612c01565b600089118015612d0b575060008b115b8015612d2f57506001600160a01b038a166000908152600860205260409020548b10155b15612d4c57612d408b8b8b8b612d5a565b9450612d4c8588611f68565b505050509695505050505050565b6000612d6461185b565b15612d6e57600080fd5b60045460ff1615612d7e57600080fd5b6004805460ff191660011790556001600160801b0385168514612da057600080fd5b82836001600160801b031614612db557600080fd5b60008511612dc257600080fd5b6001600160a01b038416612dd557600080fd5b60008311612de257600080fd5b6001600160a01b038216612df557600080fd5b816001600160a01b0316846001600160a01b03161415612e1457600080fd5b612e1c61331c565b8581526001600160a01b03808616602083015260408201859052831660608201523360808201526001600160401b03421660a0820152612e5a6132be565b600081815260036020818152604080842086518155828701516001820180546001600160a01b039283166001600160a01b031991821617909155838901516002840155606089015195830180549683169682169690961790955560808801516004928301805460a08b01516001600160401b0316600160a01b0267ffffffffffffffff60a01b199385169190981617919091169590951790945581516323b872dd60e01b81523391810191909152306024820152604481018c90529051949650918916936323b872dd936064808501948390030190829087803b158015612f4057600080fd5b505af1158015612f54573d6000803e3d6000fd5b505050506040513d6020811015612f6a57600080fd5b5051612f7557600080fd5b6040805183815290517fa2c251311b1a7a475913900a2a73dc9789a21b04bc737e050bbc506dd4eb34889181900360200190a1604080516001600160601b0319606088811b82166020808501919091529087901b90911660348301528251602881840301815260488301808552815191909201206001600160a01b0389811690925290861660688301526001600160801b03808a166088840152871660a88301526001600160401b03421660c8830152915133929185917f773ff502687307abfa024ac9f62f9752a0d210dac2ffd9a29e38e12e2ea82c829181900360e80190a4506004805460ff19169055949350505050565b6001600160a01b03831660009081526008602052604081205485101561308e57600080fd5b61309a85858585612d5a565b600a80546000838152600960209081526040918290209290925591839055815183815291519293507f8173832a493e0a3989e521458e55bfe9feac9f9b675a94e100b9d5a85f81486292918290030190a1949350505050565b600081158061310e5750508082028282828161310b57fe5b04145b6109e1576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6d756c2d6f766572666c6f7760601b604482015290519081900360640190fd5b600080831161316457600080fd5b811580159061317957506131778261163a565b155b15613197576000918252600560205260409091206001015490613164565b816131ac576131a583613228565b90506109e1565b6131b683836132cc565b156131f05760005b82158015906131d257506131d284846132cc565b156131a55750600082815260056020526040902060010154916131be565b8115801590613206575061320483836132cc565b155b156132215760009182526005602052604090912054906131f0565b50806109e1565b600080821161323657600080fd5b6000828152600360208181526040808420928301546001909301546001600160a01b039081168086526006845282862091909416808652925283205490925b8115801590613289575061328986836132cc565b15610c0b575060008181526005602052604090206001015490613275565b6000818311156132b757816110a4565b5090919050565b600280546001019081905590565b60008181526003602052604080822060020154848352908220546132f091906130f3565b6000848152600360205260408082206002015485835291205461331391906130f3565b10159392505050565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565bfefe4f66666572207761732064656c65746564206f722074616b656e2c206f72206e6576657220657869737465642e4f666665722063616e206e6f742062652063616e63656c6c656420626563617573652075736572206973206e6f74206f776e65722c20616e64206d61726b6574206973206f70656e2c20616e64206f666665722073656c6c7320726571756972656420616d6f756e74206f6620746f6b656e732ea265627a7a723158205b0c151f0af00c06b5e7e444003a6a00e95e2e416fe2b80a78015489ce8f2d8064736f6c634300050c003200000000000000000000000000000000000000000000000000000000601be1c0"

// OasisABI is used to communicate with the compiled solidity code of the generated contract
const OasisABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"close_time\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pair\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"pay_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"buy_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"LogBump\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"LogBuyEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LogDelete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LogInsert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LogItemUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pair\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"pay_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"buy_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"LogKill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pair\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"pay_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"buy_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"LogMake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"LogMatchingEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"min_amount\",\"type\":\"uint256\"}],\"name\":\"LogMinSell\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LogSortedOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pair\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"take_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"give_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"LogTake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buy_gem\",\"type\":\"address\"}],\"name\":\"LogTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LogUnsortedOffer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_best\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_dust\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_near\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_rank\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"next\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delb\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_span\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id_\",\"type\":\"bytes32\"}],\"name\":\"bump\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"max_fill_amount\",\"type\":\"uint256\"}],\"name\":\"buyAllAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fill_amt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"buyEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"close_time\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"del_rank\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dustId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"sell_gem\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"}],\"name\":\"getBestOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getBetterOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"}],\"name\":\"getBuyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fill_amt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFirstUnsortedOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"}],\"name\":\"getMinSell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getNextUnsortedOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"sell_gem\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"}],\"name\":\"getOfferCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"}],\"name\":\"getPayAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fill_amt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getWorseOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"insert\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"closed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isOfferSorted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"last_offer_id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"pay_amt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"buy_amt\",\"type\":\"uint128\"}],\"name\":\"make\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"matchingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"offer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"rounding\",\"type\":\"bool\"}],\"name\":\"offer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"}],\"name\":\"offer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"min_fill_amount\",\"type\":\"uint256\"}],\"name\":\"sellAllAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fill_amt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"buyEnabled_\",\"type\":\"bool\"}],\"name\":\"setBuyEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"matchingEnabled_\",\"type\":\"bool\"}],\"name\":\"setMatchingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dust\",\"type\":\"uint256\"}],\"name\":\"setMinSell\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"maxTakeAmount\",\"type\":\"uint128\"}],\"name\":\"take\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
