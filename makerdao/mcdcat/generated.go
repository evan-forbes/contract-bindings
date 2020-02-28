// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mcdcat

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

// McdcatABI is the input ABI used to generate the binding from.
const McdcatABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"art\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Bite\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"}],\"name\":\"bite\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chop\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lump\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"contractVowLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"
const McdcatHexAddress = "0x78f2c2af65126834c51822f56be0d7469d7a523e"

// McdcatBin is the compiled bytecode used for deploying new contracts.
var McdcatBin = "0x608060405234801561001057600080fd5b5060405161190b38038061190b8339818101604052602081101561003357600080fd5b810190808051906020019092919050505060016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016002819055505061182a806100e16000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063957aa58c11610071578063957aa58c1461023f5780639c52a7f11461025d578063bf353dbb146102a1578063d4e8be83146102f9578063d9638d3614610347578063ebecb39d146103c3576100b4565b80631a0b287e146100b957806336569e77146100fb57806345cf223014610145578063626cb3c5146101a757806365fae35e146101f15780636924500914610235575b600080fd5b6100f9600480360360608110156100cf57600080fd5b8101908080359060200190929190803590602001909291908035906020019092919050505061041b565b005b610103610606565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101916004803603604081101561015b57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061062c565b6040518082815260200191505060405180910390f35b6101af610e5e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102336004803603602081101561020757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e84565b005b61023d610fb2565b005b6102476110a3565b6040518082815260200191505060405180910390f35b61029f6004803603602081101561027357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110a9565b005b6102e3600480360360208110156102b757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111d7565b6040518082815260200191505060405180910390f35b6103456004803603604081101561030f57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111ef565b005b6103736004803603602081101561035d57600080fd5b81019080803590602001909291905050506113b6565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390f35b610419600480360360608110156103d957600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611400565b005b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146104cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4361742f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b7f63686f7000000000000000000000000000000000000000000000000000000000821415610517578060016000858152602001908152602001600020600101819055506105ce565b7f6c756d700000000000000000000000000000000000000000000000000000000082141561055f578060016000858152602001908152602001600020600201819055506105cd565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4361742f66696c652d756e7265636f676e697a65642d706172616d000000000081525060200191505060405180910390fd5b5b5961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a450505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d9638d36866040518263ffffffff1660e01b81526004018082815260200191505060606040518083038186803b1580156106a457600080fd5b505afa1580156106b8573d6000803e3d6000fd5b505050506040513d60608110156106ce57600080fd5b810190808051906020019092919080519060200190929190805190602001909291905050509250925050600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632424be5c88886040518363ffffffff1660e01b8152600401808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050604080518083038186803b1580156107a157600080fd5b505afa1580156107b5573d6000803e3d6000fd5b505050506040513d60408110156107cb57600080fd5b81019080805190602001909291908051906020019092919050505091509150600160025414610862576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f4361742f6e6f742d6c697665000000000000000000000000000000000000000081525060200191505060405180910390fd5b60008311801561088357506108778185611784565b6108818385611784565b105b6108f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f4361742f6e6f742d756e7361666500000000000000000000000000000000000081525060200191505060405180910390fd5b600061091783600160008b8152602001908152602001600020600201546117b0565b905061093582846109288486611784565b8161092f57fe5b046117b0565b91507f8000000000000000000000000000000000000000000000000000000000000000811115801561098757507f80000000000000000000000000000000000000000000000000000000000000008211155b6109f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f4361742f6f766572666c6f77000000000000000000000000000000000000000081525060200191505060405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637bab3f40898930600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1686600003886000036040518763ffffffff1660e01b8152600401808781526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019650505050505050600060405180830381600087803b158015610b4257600080fd5b505af1158015610b56573d6000803e3d6000fd5b50505050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663697efb78610ba28488611784565b6040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015610bd857600080fd5b505af1158015610bec573d6000803e3d6000fd5b505050506001600089815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663351de60088600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610c8f610c73878b611784565b600160008f8152602001908152602001600020600101546117cc565b8560006040518663ffffffff1660e01b8152600401808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b158015610d3c57600080fd5b505af1158015610d50573d6000803e3d6000fd5b505050506040513d6020811015610d6657600080fd5b810190808051906020019092919050505095508673ffffffffffffffffffffffffffffffffffffffff16887fa716da86bc1fb6d43d1493373f34d7a418b619681cd7b90f7ea667ba1489be288385610dbe878b611784565b600160008f815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168c604051808681526020018581526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019550505050505060405180910390a3505050505092915050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610f38576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4361742f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414611066576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4361742f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b60006002819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a450565b60025481565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541461115d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4361742f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60006020528060005260406000206000915090505481565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146112a3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4361742f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b7f766f7700000000000000000000000000000000000000000000000000000000008214156113115780600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061137f565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4361742f66696c652d756e7265636f676e697a65642d706172616d000000000081525060200191505060405180910390fd5b5961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a4505050565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154905083565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146114b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4361742f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b7f666c6970000000000000000000000000000000000000000000000000000000008214156116de57600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dc4d20fa6001600086815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b1580156115b357600080fd5b505af11580156115c7573d6000803e3d6000fd5b50505050806001600085815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a3b22fc4826040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b1580156116c157600080fd5b505af11580156116d5573d6000803e3d6000fd5b5050505061174c565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4361742f66696c652d756e7265636f676e697a65642d706172616d000000000081525060200191505060405180910390fd5b5961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a450505050565b6000808214806117a1575082828385029250828161179e57fe5b04145b6117aa57600080fd5b92915050565b6000818311156117c2578190506117c6565b8290505b92915050565b60006b033b2e3c9fd0803ce80000006117e58484611784565b816117ec57fe5b0490509291505056fea265627a7a72315820af1c351369240a774be0dc633a48ecb2630342c953a17ff7c11aefe0f6a7760164736f6c634300050c003200000000000000000000000035d1b3f3d7966a1dfe207aa4514c12a259a0492b"

// DeployMcdcat deploys a new Ethereum contract, binding an instance of Mcdcat to it.
func DeployMcdcat(auth *bind.TransactOpts, backend bind.ContractBackend, vat_ common.Address) (common.Address, *types.Transaction, *Mcdcat, error) {
	parsed, err := abi.JSON(strings.NewReader(McdcatABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(McdcatBin), backend, vat_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mcdcat{McdcatCaller: McdcatCaller{contract: contract}, McdcatTransactor: McdcatTransactor{contract: contract}, McdcatFilterer: McdcatFilterer{contract: contract}}, nil
}

// Mcdcat is an auto generated Go binding around an Ethereum contract.
type Mcdcat struct {
	McdcatCaller     // Read-only binding to the contract
	McdcatTransactor // Write-only binding to the contract
	McdcatFilterer   // Log filterer for contract events
}

// McdcatCaller is an auto generated read-only Go binding around an Ethereum contract.
type McdcatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McdcatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type McdcatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McdcatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type McdcatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McdcatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type McdcatSession struct {
	Contract     *Mcdcat           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// McdcatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type McdcatCallerSession struct {
	Contract *McdcatCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// McdcatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type McdcatTransactorSession struct {
	Contract     *McdcatTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// McdcatRaw is an auto generated low-level Go binding around an Ethereum contract.
type McdcatRaw struct {
	Contract *Mcdcat // Generic contract binding to access the raw methods on
}

// McdcatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type McdcatCallerRaw struct {
	Contract *McdcatCaller // Generic read-only contract binding to access the raw methods on
}

// McdcatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type McdcatTransactorRaw struct {
	Contract *McdcatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMcdcat creates a new instance of Mcdcat, bound to a specific deployed contract.
func NewMcdcat(address common.Address, backend bind.ContractBackend) (*Mcdcat, error) {
	contract, err := bindMcdcat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mcdcat{McdcatCaller: McdcatCaller{contract: contract}, McdcatTransactor: McdcatTransactor{contract: contract}, McdcatFilterer: McdcatFilterer{contract: contract}}, nil
}

// NewMcdcatCaller creates a new read-only instance of Mcdcat, bound to a specific deployed contract.
func NewMcdcatCaller(address common.Address, caller bind.ContractCaller) (*McdcatCaller, error) {
	contract, err := bindMcdcat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &McdcatCaller{contract: contract}, nil
}

// NewMcdcatTransactor creates a new write-only instance of Mcdcat, bound to a specific deployed contract.
func NewMcdcatTransactor(address common.Address, transactor bind.ContractTransactor) (*McdcatTransactor, error) {
	contract, err := bindMcdcat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &McdcatTransactor{contract: contract}, nil
}

// NewMcdcatFilterer creates a new log filterer instance of Mcdcat, bound to a specific deployed contract.
func NewMcdcatFilterer(address common.Address, filterer bind.ContractFilterer) (*McdcatFilterer, error) {
	contract, err := bindMcdcat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &McdcatFilterer{contract: contract}, nil
}

// bindMcdcat binds a generic wrapper to an already deployed contract.
func bindMcdcat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(McdcatABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mcdcat *McdcatRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mcdcat.Contract.McdcatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mcdcat *McdcatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcdcat.Contract.McdcatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mcdcat *McdcatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mcdcat.Contract.McdcatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mcdcat *McdcatCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mcdcat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mcdcat *McdcatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcdcat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mcdcat *McdcatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mcdcat.Contract.contract.Transact(opts, method, params...)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) constant returns(address flip, uint256 chop, uint256 lump)
func (_Mcdcat *McdcatCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Flip common.Address
	Chop *big.Int
	Lump *big.Int
}, error) {
	ret := new(struct {
		Flip common.Address
		Chop *big.Int
		Lump *big.Int
	})
	out := ret
	err := _Mcdcat.contract.Call(opts, out, "ilks", arg0)
	return *ret, err
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) constant returns(address flip, uint256 chop, uint256 lump)
func (_Mcdcat *McdcatSession) Ilks(arg0 [32]byte) (struct {
	Flip common.Address
	Chop *big.Int
	Lump *big.Int
}, error) {
	return _Mcdcat.Contract.Ilks(&_Mcdcat.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) constant returns(address flip, uint256 chop, uint256 lump)
func (_Mcdcat *McdcatCallerSession) Ilks(arg0 [32]byte) (struct {
	Flip common.Address
	Chop *big.Int
	Lump *big.Int
}, error) {
	return _Mcdcat.Contract.Ilks(&_Mcdcat.CallOpts, arg0)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Mcdcat *McdcatCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdcat.contract.Call(opts, out, "live")
	return *ret0, err
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Mcdcat *McdcatSession) Live() (*big.Int, error) {
	return _Mcdcat.Contract.Live(&_Mcdcat.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Mcdcat *McdcatCallerSession) Live() (*big.Int, error) {
	return _Mcdcat.Contract.Live(&_Mcdcat.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Mcdcat *McdcatCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mcdcat.contract.Call(opts, out, "vat")
	return *ret0, err
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Mcdcat *McdcatSession) Vat() (common.Address, error) {
	return _Mcdcat.Contract.Vat(&_Mcdcat.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Mcdcat *McdcatCallerSession) Vat() (common.Address, error) {
	return _Mcdcat.Contract.Vat(&_Mcdcat.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() constant returns(address)
func (_Mcdcat *McdcatCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mcdcat.contract.Call(opts, out, "vow")
	return *ret0, err
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() constant returns(address)
func (_Mcdcat *McdcatSession) Vow() (common.Address, error) {
	return _Mcdcat.Contract.Vow(&_Mcdcat.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() constant returns(address)
func (_Mcdcat *McdcatCallerSession) Vow() (common.Address, error) {
	return _Mcdcat.Contract.Vow(&_Mcdcat.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Mcdcat *McdcatCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdcat.contract.Call(opts, out, "wards", arg0)
	return *ret0, err
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Mcdcat *McdcatSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Mcdcat.Contract.Wards(&_Mcdcat.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Mcdcat *McdcatCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Mcdcat.Contract.Wards(&_Mcdcat.CallOpts, arg0)
}

// Bite is a paid mutator transaction binding the contract method 0x45cf2230.
//
// Solidity: function bite(bytes32 ilk, address urn) returns(uint256 id)
func (_Mcdcat *McdcatTransactor) Bite(opts *bind.TransactOpts, ilk [32]byte, urn common.Address) (*types.Transaction, error) {
	return _Mcdcat.contract.Transact(opts, "bite", ilk, urn)
}

// Bite is a paid mutator transaction binding the contract method 0x45cf2230.
//
// Solidity: function bite(bytes32 ilk, address urn) returns(uint256 id)
func (_Mcdcat *McdcatSession) Bite(ilk [32]byte, urn common.Address) (*types.Transaction, error) {
	return _Mcdcat.Contract.Bite(&_Mcdcat.TransactOpts, ilk, urn)
}

// Bite is a paid mutator transaction binding the contract method 0x45cf2230.
//
// Solidity: function bite(bytes32 ilk, address urn) returns(uint256 id)
func (_Mcdcat *McdcatTransactorSession) Bite(ilk [32]byte, urn common.Address) (*types.Transaction, error) {
	return _Mcdcat.Contract.Bite(&_Mcdcat.TransactOpts, ilk, urn)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Mcdcat *McdcatTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcdcat.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Mcdcat *McdcatSession) Cage() (*types.Transaction, error) {
	return _Mcdcat.Contract.Cage(&_Mcdcat.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Mcdcat *McdcatTransactorSession) Cage() (*types.Transaction, error) {
	return _Mcdcat.Contract.Cage(&_Mcdcat.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Mcdcat *McdcatTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Mcdcat.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Mcdcat *McdcatSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Mcdcat.Contract.Deny(&_Mcdcat.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Mcdcat *McdcatTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Mcdcat.Contract.Deny(&_Mcdcat.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Mcdcat *McdcatTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Mcdcat.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Mcdcat *McdcatSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Mcdcat.Contract.File(&_Mcdcat.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Mcdcat *McdcatTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Mcdcat.Contract.File(&_Mcdcat.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Mcdcat *McdcatTransactor) File0(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Mcdcat.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Mcdcat *McdcatSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Mcdcat.Contract.File0(&_Mcdcat.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Mcdcat *McdcatTransactorSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Mcdcat.Contract.File0(&_Mcdcat.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address flip) returns()
func (_Mcdcat *McdcatTransactor) File1(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, flip common.Address) (*types.Transaction, error) {
	return _Mcdcat.contract.Transact(opts, "file1", ilk, what, flip)
}

// File1 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address flip) returns()
func (_Mcdcat *McdcatSession) File1(ilk [32]byte, what [32]byte, flip common.Address) (*types.Transaction, error) {
	return _Mcdcat.Contract.File1(&_Mcdcat.TransactOpts, ilk, what, flip)
}

// File1 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address flip) returns()
func (_Mcdcat *McdcatTransactorSession) File1(ilk [32]byte, what [32]byte, flip common.Address) (*types.Transaction, error) {
	return _Mcdcat.Contract.File1(&_Mcdcat.TransactOpts, ilk, what, flip)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Mcdcat *McdcatTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Mcdcat.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Mcdcat *McdcatSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Mcdcat.Contract.Rely(&_Mcdcat.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Mcdcat *McdcatTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Mcdcat.Contract.Rely(&_Mcdcat.TransactOpts, usr)
}

// McdcatBiteIterator is returned from FilterBite and is used to iterate over the raw logs and unpacked data for Bite events raised by the Mcdcat contract.
type McdcatBiteIterator struct {
	Event *McdcatBite // Event containing the contract specifics and raw log

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
func (it *McdcatBiteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(McdcatBite)
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
		it.Event = new(McdcatBite)
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
func (it *McdcatBiteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *McdcatBiteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// McdcatBite represents a Bite event raised by the Mcdcat contract.
type McdcatBite struct {
	Ilk  [32]byte
	Urn  common.Address
	Ink  *big.Int
	Art  *big.Int
	Tab  *big.Int
	Flip common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBite is a free log retrieval operation binding the contract event 0xa716da86bc1fb6d43d1493373f34d7a418b619681cd7b90f7ea667ba1489be28.
//
// Solidity: event Bite(bytes32 indexed ilk, address indexed urn, uint256 ink, uint256 art, uint256 tab, address flip, uint256 id)
func (_Mcdcat *McdcatFilterer) FilterBite(opts *bind.FilterOpts, ilk [][32]byte, urn []common.Address) (*McdcatBiteIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var urnRule []interface{}
	for _, urnItem := range urn {
		urnRule = append(urnRule, urnItem)
	}

	logs, sub, err := _Mcdcat.contract.FilterLogs(opts, "Bite", ilkRule, urnRule)
	if err != nil {
		return nil, err
	}
	return &McdcatBiteIterator{contract: _Mcdcat.contract, event: "Bite", logs: logs, sub: sub}, nil
}

// WatchBite is a free log subscription operation binding the contract event 0xa716da86bc1fb6d43d1493373f34d7a418b619681cd7b90f7ea667ba1489be28.
//
// Solidity: event Bite(bytes32 indexed ilk, address indexed urn, uint256 ink, uint256 art, uint256 tab, address flip, uint256 id)
func (_Mcdcat *McdcatFilterer) WatchBite(opts *bind.WatchOpts, sink chan<- *McdcatBite, ilk [][32]byte, urn []common.Address) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var urnRule []interface{}
	for _, urnItem := range urn {
		urnRule = append(urnRule, urnItem)
	}

	logs, sub, err := _Mcdcat.contract.WatchLogs(opts, "Bite", ilkRule, urnRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(McdcatBite)
				if err := _Mcdcat.contract.UnpackLog(event, "Bite", log); err != nil {
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

// ParseBite is a log parse operation binding the contract event 0xa716da86bc1fb6d43d1493373f34d7a418b619681cd7b90f7ea667ba1489be28.
//
// Solidity: event Bite(bytes32 indexed ilk, address indexed urn, uint256 ink, uint256 art, uint256 tab, address flip, uint256 id)
func (_Mcdcat *McdcatFilterer) ParseBite(log types.Log) (*McdcatBite, error) {
	event := new(McdcatBite)
	if err := _Mcdcat.contract.UnpackLog(event, "Bite", log); err != nil {
		return nil, err
	}
	return event, nil
}
