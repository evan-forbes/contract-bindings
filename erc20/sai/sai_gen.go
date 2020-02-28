package sai

// Sai is a wrapper around BoundContract, enforcing type checking and including
// QoL helper methods
type Sai bind.BoundContract

// NewSai creates a new instance of Sai, bound to a specific deployed contract.
func NewSai(address common.Address, backend bind.ContractBackend) (*Sai, error) {
	a, err := abi.JSON(strings.NewReader(SaiABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, a, backend, backend, backend)
	return &Sai(*contract), nil
}

//////////////////////////////////////////////////////
//		Deployment
////////////////////////////////////////////////////

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
	return address, tx, &Sai(*contract), nil
}

//////////////////////////////////////////////////////
//		Structs
////////////////////////////////////////////////////

//////////////////////////////////////////////////////
//		Data Calls
////////////////////////////////////////////////////

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
// - Solidity: function allowance(address src, address guy) constant returns(uint256)
func (_Sai *Sai) Allowance(opts *bind.CallOpts, src common.Address, guy common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sai.Call(opts, out, "allowance", src, guy)
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
// - Solidity: function authority() constant returns(address)
func (_Sai *Sai) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sai.Call(opts, out, "authority")
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
// - Solidity: function balanceOf(address src) constant returns(uint256)
func (_Sai *Sai) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sai.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
// - Solidity: function decimals() constant returns(uint256)
func (_Sai *Sai) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sai.Call(opts, out, "decimals")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
// - Solidity: function name() constant returns(bytes32)
func (_Sai *Sai) Name(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Sai.Call(opts, out, "name")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
// - Solidity: function owner() constant returns(address)
func (_Sai *Sai) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sai.Call(opts, out, "owner")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
// - Solidity: function stopped() constant returns(bool)
func (_Sai *Sai) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Sai.Call(opts, out, "stopped")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
// - Solidity: function symbol() constant returns(bytes32)
func (_Sai *Sai) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Sai.Call(opts, out, "symbol")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
// - Solidity: function totalSupply() constant returns(uint256)
func (_Sai *Sai) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sai.Call(opts, out, "totalSupply")
	return *ret0, err
}

//////////////////////////////////////////////////////
//		Transactions
////////////////////////////////////////////////////

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
// - Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Sai *Sai) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Transact(opts, "approve", guy, wad)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdaea85c5.
// - Solidity: function approve(address guy) returns(bool)
func (_Sai *Sai) Approve0(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Sai.Transact(opts, "approve0", guy)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
// - Solidity: function burn(uint256 wad) returns()
func (_Sai *Sai) Burn(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Transact(opts, "burn", wad)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
// - Solidity: function burn(address guy, uint256 wad) returns()
func (_Sai *Sai) Burn0(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Transact(opts, "burn0", guy, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
// - Solidity: function mint(address guy, uint256 wad) returns()
func (_Sai *Sai) Mint(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Transact(opts, "mint", guy, wad)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
// - Solidity: function mint(uint256 wad) returns()
func (_Sai *Sai) Mint0(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Transact(opts, "mint0", wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
// - Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Sai *Sai) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Transact(opts, "move", src, dst, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
// - Solidity: function pull(address src, uint256 wad) returns()
func (_Sai *Sai) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Transact(opts, "pull", src, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
// - Solidity: function push(address dst, uint256 wad) returns()
func (_Sai *Sai) Push(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Transact(opts, "push", dst, wad)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
// - Solidity: function setAuthority(address authority_) returns()
func (_Sai *Sai) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Sai.Transact(opts, "setAuthority", authority_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
// - Solidity: function setName(bytes32 name_) returns()
func (_Sai *Sai) SetName(opts *bind.TransactOpts, name_ [32]byte) (*types.Transaction, error) {
	return _Sai.Transact(opts, "setName", name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
// - Solidity: function setOwner(address owner_) returns()
func (_Sai *Sai) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Sai.Transact(opts, "setOwner", owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
// - Solidity: function start() returns()
func (_Sai *Sai) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sai.Transact(opts, "start")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
// - Solidity: function stop() returns()
func (_Sai *Sai) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sai.Transact(opts, "stop")
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
// - Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Sai *Sai) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Transact(opts, "transfer", dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
// - Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Sai *Sai) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Sai.Transact(opts, "transferFrom", src, dst, wad)
}

//////////////////////////////////////////////////////
//		Events
////////////////////////////////////////////////////

//////// Approval ////////
// ApprovalID is the hex of the Topic Hash
const SaiApprovalID = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"

// ApprovalLog represents a Approval event raised by the Sai contract.
type ApprovalLog struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// UnpackApprovalLog is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Sai *Sai) UnpackApprovalLog(log types.Log) (*ApprovalLog, error) {
	event := new(ApprovalLog)
	if err := _Sai.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// Burn ////////
// BurnID is the hex of the Topic Hash
const SaiBurnID = "0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5"

// BurnLog represents a Burn event raised by the Sai contract.
type BurnLog struct {
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// UnpackBurnLog is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5
// Solidity: event Burn(address indexed guy, uint256 wad)
func (_Sai *Sai) UnpackBurnLog(log types.Log) (*BurnLog, error) {
	event := new(BurnLog)
	if err := _Sai.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogSetAuthority ////////
// LogSetAuthorityID is the hex of the Topic Hash
const SaiLogSetAuthorityID = "0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4"

// LogSetAuthorityLog represents a LogSetAuthority event raised by the Sai contract.
type LogSetAuthorityLog struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// UnpackLogSetAuthorityLog is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4
// Solidity: event LogSetAuthority(address indexed authority)
func (_Sai *Sai) UnpackLogSetAuthorityLog(log types.Log) (*LogSetAuthorityLog, error) {
	event := new(LogSetAuthorityLog)
	if err := _Sai.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// LogSetOwner ////////
// LogSetOwnerID is the hex of the Topic Hash
const SaiLogSetOwnerID = "0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94"

// LogSetOwnerLog represents a LogSetOwner event raised by the Sai contract.
type LogSetOwnerLog struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// UnpackLogSetOwnerLog is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94
// Solidity: event LogSetOwner(address indexed owner)
func (_Sai *Sai) UnpackLogSetOwnerLog(log types.Log) (*LogSetOwnerLog, error) {
	event := new(LogSetOwnerLog)
	if err := _Sai.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// Mint ////////
// MintID is the hex of the Topic Hash
const SaiMintID = "0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885"

// MintLog represents a Mint event raised by the Sai contract.
type MintLog struct {
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// UnpackMintLog is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885
// Solidity: event Mint(address indexed guy, uint256 wad)
func (_Sai *Sai) UnpackMintLog(log types.Log) (*MintLog, error) {
	event := new(MintLog)
	if err := _Sai.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

//////// Transfer ////////
// TransferID is the hex of the Topic Hash
const SaiTransferID = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

// TransferLog represents a Transfer event raised by the Sai contract.
type TransferLog struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// UnpackTransferLog is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Sai *Sai) UnpackTransferLog(log types.Log) (*TransferLog, error) {
	event := new(TransferLog)
	if err := _Sai.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

/*
Mux can be copied and pasted to save ya a quick minute when distinguishing between log data
func Mux(c *Sai, log types.Log) error {
	switch log.Topics[0].Hex() {
	case sai.SaiApprovalID:
		ulog, err := c.UnpackApprovalLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case sai.SaiBurnID:
		ulog, err := c.UnpackBurnLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case sai.SaiLogSetAuthorityID:
		ulog, err := c.UnpackLogSetAuthorityLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case sai.SaiLogSetOwnerID:
		ulog, err := c.UnpackLogSetOwnerLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case sai.SaiMintID:
		ulog, err := c.UnpackMintLog(log)
		if err != nil {
			return err
		}
		// insert additional code here

	case sai.SaiTransferID:
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
// SaiBin is used to deploy the generated contract
const SaiBin = "0x606060405260126006556000600755341561001957600080fd5b604051602080610dd283398101604052808051600160a060020a03331660008181526001602052604080822082905590805560048054600160a060020a0319168317905591935091507fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94905160405180910390a2600555610d338061009f6000396000f3006060604052600436106101485763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461014d57806307da68f514610172578063095ea7b31461018757806313af4035146101bd57806318160ddd146101dc57806323b872dd146101ef578063313ce5671461021757806340c10f191461022a57806342966c681461024c5780635ac801fe1461026257806370a082311461027857806375f12b21146102975780637a9e5e4b146102aa5780638da5cb5b146102c957806395d89b41146102f85780639dc29fac1461030b578063a0712d681461032d578063a9059cbb14610343578063b753a98c14610365578063bb35783b14610387578063be9a6555146103af578063bf7e214f146103c2578063daea85c5146103d5578063dd62ed3e146103f4578063f2d5d56b14610419575b600080fd5b341561015857600080fd5b61016061043b565b60405190815260200160405180910390f35b341561017d57600080fd5b610185610441565b005b341561019257600080fd5b6101a9600160a060020a03600435166024356104e0565b604051901515815260200160405180910390f35b34156101c857600080fd5b610185600160a060020a036004351661050d565b34156101e757600080fd5b61016061058c565b34156101fa57600080fd5b6101a9600160a060020a0360043581169060243516604435610592565b341561022257600080fd5b610160610707565b341561023557600080fd5b610185600160a060020a036004351660243561070d565b341561025757600080fd5b6101856004356107d3565b341561026d57600080fd5b6101856004356107e0565b341561028357600080fd5b610160600160a060020a0360043516610806565b34156102a257600080fd5b6101a9610821565b34156102b557600080fd5b610185600160a060020a0360043516610831565b34156102d457600080fd5b6102dc6108b0565b604051600160a060020a03909116815260200160405180910390f35b341561030357600080fd5b6101606108bf565b341561031657600080fd5b610185600160a060020a03600435166024356108c5565b341561033857600080fd5b610185600435610a33565b341561034e57600080fd5b6101a9600160a060020a0360043516602435610a3d565b341561037057600080fd5b610185600160a060020a0360043516602435610a4a565b341561039257600080fd5b610185600160a060020a0360043581169060243516604435610a5a565b34156103ba57600080fd5b610185610a6b565b34156103cd57600080fd5b6102dc610b04565b34156103e057600080fd5b6101a9600160a060020a0360043516610b13565b34156103ff57600080fd5b610160600160a060020a0360043581169060243516610b39565b341561042457600080fd5b610185600160a060020a0360043516602435610b64565b60075481565b61045733600035600160e060020a031916610b6f565b151561046257600080fd5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a450506004805474ff0000000000000000000000000000000000000000191660a060020a179055565b60045460009060a060020a900460ff16156104fa57600080fd5b6105048383610c7b565b90505b92915050565b61052333600035600160e060020a031916610b6f565b151561052e57600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a250565b60005490565b60045460009060a060020a900460ff16156105ac57600080fd5b33600160a060020a031684600160a060020a0316141580156105f65750600160a060020a038085166000908152600260209081526040808320339094168352929052205460001914155b1561065457600160a060020a038085166000908152600260209081526040808320339094168352929052205461062c9083610ce7565b600160a060020a03808616600090815260026020908152604080832033909416835292905220555b600160a060020a0384166000908152600160205260409020546106779083610ce7565b600160a060020a0380861660009081526001602052604080822093909355908516815220546106a69083610cf7565b600160a060020a03808516600081815260016020526040908190209390935591908616907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60065481565b61072333600035600160e060020a031916610b6f565b151561072e57600080fd5b60045460a060020a900460ff161561074557600080fd5b600160a060020a0382166000908152600160205260409020546107689082610cf7565b600160a060020a0383166000908152600160205260408120919091555461078f9082610cf7565b600055600160a060020a0382167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405190815260200160405180910390a25050565b6107dd33826108c5565b50565b6107f633600035600160e060020a031916610b6f565b151561080157600080fd5b600755565b600160a060020a031660009081526001602052604090205490565b60045460a060020a900460ff1681565b61084733600035600160e060020a031916610b6f565b151561085257600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a250565b600454600160a060020a031681565b60055481565b6108db33600035600160e060020a031916610b6f565b15156108e657600080fd5b60045460a060020a900460ff16156108fd57600080fd5b33600160a060020a031682600160a060020a0316141580156109475750600160a060020a038083166000908152600260209081526040808320339094168352929052205460001914155b156109a557600160a060020a038083166000908152600260209081526040808320339094168352929052205461097d9082610ce7565b600160a060020a03808416600090815260026020908152604080832033909416835292905220555b600160a060020a0382166000908152600160205260409020546109c89082610ce7565b600160a060020a038316600090815260016020526040812091909155546109ef9082610ce7565b600055600160a060020a0382167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405190815260200160405180910390a25050565b6107dd338261070d565b6000610504338484610592565b610a55338383610592565b505050565b610a65838383610592565b50505050565b610a8133600035600160e060020a031916610b6f565b1515610a8c57600080fd5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a450506004805474ff000000000000000000000000000000000000000019169055565b600354600160a060020a031681565b60045460009060a060020a900460ff1615610b2d57600080fd5b61050782600019610c7b565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b610a55823383610592565b600030600160a060020a031683600160a060020a03161415610b9357506001610507565b600454600160a060020a0384811691161415610bb157506001610507565b600354600160a060020a03161515610bcb57506000610507565b600354600160a060020a031663b70096138430856000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b1515610c5957600080fd5b6102c65a03f11515610c6a57600080fd5b505050604051805190509050610507565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b8082038281111561050757600080fd5b8082018281101561050757600080fd00a165627a7a723058200877df264aa5d498c61a45dfe4fc04c68d11820448cf0cc7a275283a271173b400294441490000000000000000000000000000000000000000000000000000000000"

// SaiABI is used to communicate with the compiled solidity code of the generated contract
const SaiABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"bytes32\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"symbol_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"
