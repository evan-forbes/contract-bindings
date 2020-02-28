// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mcdflap

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

// McdflapABI is the input ABI used to generate the binding from.
const McdflapABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gem_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"Kick\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"beg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"tic\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"end\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"deal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"contractGemLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"kick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kicks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tau\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"tend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"tick\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ttl\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// McdflapBin is the compiled bytecode used for deploying new contracts.
var McdflapBin = "0x6080604052670e92596fd6290000600455612a30600560006101000a81548165ffffffffffff021916908365ffffffffffff1602179055506202a300600560066101000a81548165ffffffffffff021916908365ffffffffffff160217905550600060065534801561007057600080fd5b506040516124733803806124738339818101604052604081101561009357600080fd5b81019080805190602001909291908051906020019092919050505060016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160078190555050506122e68061018d6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063957aa58c116100a2578063c959c42b11610071578063c959c42b14610479578063ca40c419146104a7578063cfc4af55146104f3578063cfdd330214610521578063fc7b6aee1461053f57610116565b8063957aa58c146103915780639c52a7f1146103af578063a2f91af2146103f3578063bf353dbb1461042157610116565b80634b43ed12116100e95780634b43ed12146102755780634e8b1dd5146102b757806365fae35e146102e55780637bd2bea7146103295780637d780d821461037357610116565b806326e027f11461011b57806329ae81141461014957806336569e77146101815780634423c5f1146101cb575b600080fd5b6101476004803603602081101561013157600080fd5b810190808035906020019092919050505061056d565b005b61017f6004803603604081101561015f57600080fd5b8101908080359060200190929190803590602001909291905050506108b4565b005b610189610ae7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101f7600480360360208110156101e157600080fd5b8101908080359060200190929190505050610b0d565b604051808681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018365ffffffffffff1665ffffffffffff1681526020018265ffffffffffff1665ffffffffffff1681526020019550505050505060405180910390f35b6102b56004803603606081101561028b57600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610b87565b005b6102bf611367565b604051808265ffffffffffff1665ffffffffffff16815260200191505060405180910390f35b610327600480360360208110156102fb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061137f565b005b6103316114ad565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61037b6114d3565b6040518082815260200191505060405180910390f35b6103996114d9565b6040518082815260200191505060405180910390f35b6103f1600480360360208110156103c557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506114df565b005b61041f6004803603602081101561040957600080fd5b810190808035906020019092919050505061160d565b005b6104636004803603602081101561043757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506117f4565b6040518082815260200191505060405180910390f35b6104a56004803603602081101561048f57600080fd5b810190808035906020019092919050505061180c565b005b6104dd600480360360408110156104bd57600080fd5b810190808035906020019092919080359060200190929190505050611c6d565b6040518082815260200191505060405180910390f35b6104fb612066565b604051808265ffffffffffff1665ffffffffffff16815260200191505060405180910390f35b61052961207e565b6040518082815260200191505060405180910390f35b61056b6004803603602081101561055557600080fd5b8101908080359060200190929190505050612084565b005b6000600754146105e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f466c61707065722f7374696c6c2d6c697665000000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156106be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f466c61707065722f6775792d6e6f742d7365740000000000000000000000000081525060200191505060405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b306001600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660016000868152602001908152602001600020600001546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b1580156107e757600080fd5b505af11580156107fb573d6000803e3d6000fd5b505050506001600082815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160146101000a81549065ffffffffffff021916905560028201601a6101000a81549065ffffffffffff021916905550505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610968576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c61707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b7f626567000000000000000000000000000000000000000000000000000000000082141561099c5780600481905550610ab0565b7f74746c00000000000000000000000000000000000000000000000000000000008214156109ee5780600560006101000a81548165ffffffffffff021916908365ffffffffffff160217905550610aaf565b7f7461750000000000000000000000000000000000000000000000000000000000821415610a405780600560066101000a81548165ffffffffffff021916908365ffffffffffff160217905550610aae565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f466c61707065722f66696c652d756e7265636f676e697a65642d706172616d0081525060200191505060405180910390fd5b5b5b5961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a4505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160149054906101000a900465ffffffffffff169080600201601a9054906101000a900465ffffffffffff16905085565b600160075414610bff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f466c61707065722f6e6f742d6c6976650000000000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166001600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610cd8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f466c61707065722f6775792d6e6f742d7365740000000000000000000000000081525060200191505060405180910390fd5b426001600085815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff161180610d46575060006001600085815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff16145b610db8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f466c61707065722f616c72656164792d66696e69736865642d7469630000000081525060200191505060405180910390fd5b4260016000858152602001908152602001600020600201601a9054906101000a900465ffffffffffff1665ffffffffffff1611610e5d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f466c61707065722f616c72656164792d66696e69736865642d656e640000000081525060200191505060405180910390fd5b60016000848152602001908152602001600020600101548214610ee8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f466c61707065722f6c6f742d6e6f742d6d61746368696e67000000000000000081525060200191505060405180910390fd5b60016000848152602001908152602001600020600001548111610f73576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c61707065722f6269642d6e6f742d6869676865720000000000000000000081525060200191505060405180910390fd5b610f95600454600160008681526020019081526020016000206000015461225b565b610fa782670de0b6b3a764000061225b565b101561101b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f466c61707065722f696e73756666696369656e742d696e63726561736500000081525060200191505060405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b336001600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660016000888152602001908152602001600020600001546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561114457600080fd5b505af1158015611158573d6000803e3d6000fd5b50505050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b3330600160008881526020019081526020016000206000015485036040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561125157600080fd5b505af1158015611265573d6000803e3d6000fd5b50505050336001600085815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060016000858152602001908152602001600020600001819055506112f742600560009054906101000a900465ffffffffffff16612287565b6001600085815260200190815260200160002060020160146101000a81548165ffffffffffff021916908365ffffffffffff1602179055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a450505050565b600560009054906101000a900465ffffffffffff1681565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414611433576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c61707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b60075481565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414611593576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c61707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146116c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c61707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b6000600781905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b3033846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b1580156117a657600080fd5b505af11580156117ba573d6000803e3d6000fd5b505050505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60006020528060005260406000206000915090505481565b600160075414611884576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f466c61707065722f6e6f742d6c6976650000000000000000000000000000000081525060200191505060405180910390fd5b60006001600083815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff161415801561192e5750426001600083815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff16108061192d57504260016000838152602001908152602001600020600201601a9054906101000a900465ffffffffffff1665ffffffffffff16105b5b6119a0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f466c61707065722f6e6f742d66696e697368656400000000000000000000000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b306001600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660016000868152602001908152602001600020600101546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015611ac957600080fd5b505af1158015611add573d6000803e3d6000fd5b50505050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac3060016000858152602001908152602001600020600001546040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015611ba057600080fd5b505af1158015611bb4573d6000803e3d6000fd5b505050506001600082815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160146101000a81549065ffffffffffff021916905560028201601a6101000a81549065ffffffffffff021916905550505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b600060016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414611d23576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c61707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b600160075414611d9b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f466c61707065722f6e6f742d6c6976650000000000000000000000000000000081525060200191505060405180910390fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60065410611e32576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f466c61707065722f6f766572666c6f770000000000000000000000000000000081525060200191505060405180910390fd5b6006600081546001019190508190559050816001600083815260200190815260200160002060000181905550826001600083815260200190815260200160002060010181905550336001600083815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611eec42600560069054906101000a900465ffffffffffff16612287565b60016000838152602001908152602001600020600201601a6101000a81548165ffffffffffff021916908365ffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b3330866040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561200157600080fd5b505af1158015612015573d6000803e3d6000fd5b505050507fe6dde59cbc017becba89714a037778d234a84ce7f0a137487142a007e580d60981848460405180848152602001838152602001828152602001935050505060405180910390a192915050565b600560069054906101000a900465ffffffffffff1681565b60065481565b4260016000838152602001908152602001600020600201601a9054906101000a900465ffffffffffff1665ffffffffffff1610612129576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f466c61707065722f6e6f742d66696e697368656400000000000000000000000081525060200191505060405180910390fd5b60006001600083815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff16146121cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f466c61707065722f6269642d616c72656164792d706c6163656400000000000081525060200191505060405180910390fd5b6121ed42600560069054906101000a900465ffffffffffff16612287565b60016000838152602001908152602001600020600201601a6101000a81548165ffffffffffff021916908365ffffffffffff1602179055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b600080821480612278575082828385029250828161227557fe5b04145b61228157600080fd5b92915050565b60008265ffffffffffff1682840191508165ffffffffffff1610156122ab57600080fd5b9291505056fea265627a7a7231582092bfb1420ca4c8f363a93d62c6c31bbdcb469b676d8368d6806b8b488049be9864736f6c634300050c003200000000000000000000000035d1b3f3d7966a1dfe207aa4514c12a259a0492b0000000000000000000000009f8f72aa9304c8b593d555f12ef6589cc3a579a2"

const McdflapHexAddress = "0xdfe0fb1be2a52cdbf8fb962d5701d7fd0902db9f"

// DeployMcdflap deploys a new Ethereum contract, binding an instance of Mcdflap to it.
func DeployMcdflap(auth *bind.TransactOpts, backend bind.ContractBackend, vat_ common.Address, gem_ common.Address) (common.Address, *types.Transaction, *Mcdflap, error) {
	parsed, err := abi.JSON(strings.NewReader(McdflapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(McdflapBin), backend, vat_, gem_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mcdflap{McdflapCaller: McdflapCaller{contract: contract}, McdflapTransactor: McdflapTransactor{contract: contract}, McdflapFilterer: McdflapFilterer{contract: contract}}, nil
}

// Mcdflap is an auto generated Go binding around an Ethereum contract.
type Mcdflap struct {
	McdflapCaller     // Read-only binding to the contract
	McdflapTransactor // Write-only binding to the contract
	McdflapFilterer   // Log filterer for contract events
}

// McdflapCaller is an auto generated read-only Go binding around an Ethereum contract.
type McdflapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McdflapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type McdflapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McdflapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type McdflapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McdflapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type McdflapSession struct {
	Contract     *Mcdflap          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// McdflapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type McdflapCallerSession struct {
	Contract *McdflapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// McdflapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type McdflapTransactorSession struct {
	Contract     *McdflapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// McdflapRaw is an auto generated low-level Go binding around an Ethereum contract.
type McdflapRaw struct {
	Contract *Mcdflap // Generic contract binding to access the raw methods on
}

// McdflapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type McdflapCallerRaw struct {
	Contract *McdflapCaller // Generic read-only contract binding to access the raw methods on
}

// McdflapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type McdflapTransactorRaw struct {
	Contract *McdflapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMcdflap creates a new instance of Mcdflap, bound to a specific deployed contract.
func NewMcdflap(address common.Address, backend bind.ContractBackend) (*Mcdflap, error) {
	contract, err := bindMcdflap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mcdflap{McdflapCaller: McdflapCaller{contract: contract}, McdflapTransactor: McdflapTransactor{contract: contract}, McdflapFilterer: McdflapFilterer{contract: contract}}, nil
}

// NewMcdflapCaller creates a new read-only instance of Mcdflap, bound to a specific deployed contract.
func NewMcdflapCaller(address common.Address, caller bind.ContractCaller) (*McdflapCaller, error) {
	contract, err := bindMcdflap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &McdflapCaller{contract: contract}, nil
}

// NewMcdflapTransactor creates a new write-only instance of Mcdflap, bound to a specific deployed contract.
func NewMcdflapTransactor(address common.Address, transactor bind.ContractTransactor) (*McdflapTransactor, error) {
	contract, err := bindMcdflap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &McdflapTransactor{contract: contract}, nil
}

// NewMcdflapFilterer creates a new log filterer instance of Mcdflap, bound to a specific deployed contract.
func NewMcdflapFilterer(address common.Address, filterer bind.ContractFilterer) (*McdflapFilterer, error) {
	contract, err := bindMcdflap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &McdflapFilterer{contract: contract}, nil
}

// bindMcdflap binds a generic wrapper to an already deployed contract.
func bindMcdflap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(McdflapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mcdflap *McdflapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mcdflap.Contract.McdflapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mcdflap *McdflapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcdflap.Contract.McdflapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mcdflap *McdflapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mcdflap.Contract.McdflapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mcdflap *McdflapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mcdflap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mcdflap *McdflapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcdflap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mcdflap *McdflapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mcdflap.Contract.contract.Transact(opts, method, params...)
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() constant returns(uint256)
func (_Mcdflap *McdflapCaller) Beg(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflap.contract.Call(opts, out, "beg")
	return *ret0, err
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() constant returns(uint256)
func (_Mcdflap *McdflapSession) Beg() (*big.Int, error) {
	return _Mcdflap.Contract.Beg(&_Mcdflap.CallOpts)
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() constant returns(uint256)
func (_Mcdflap *McdflapCallerSession) Beg() (*big.Int, error) {
	return _Mcdflap.Contract.Beg(&_Mcdflap.CallOpts)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) constant returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end)
func (_Mcdflap *McdflapCaller) Bids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
}, error) {
	ret := new(struct {
		Bid *big.Int
		Lot *big.Int
		Guy common.Address
		Tic *big.Int
		End *big.Int
	})
	out := ret
	err := _Mcdflap.contract.Call(opts, out, "bids", arg0)
	return *ret, err
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) constant returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end)
func (_Mcdflap *McdflapSession) Bids(arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
}, error) {
	return _Mcdflap.Contract.Bids(&_Mcdflap.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) constant returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end)
func (_Mcdflap *McdflapCallerSession) Bids(arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
}, error) {
	return _Mcdflap.Contract.Bids(&_Mcdflap.CallOpts, arg0)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() constant returns(address)
func (_Mcdflap *McdflapCaller) Gem(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mcdflap.contract.Call(opts, out, "gem")
	return *ret0, err
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() constant returns(address)
func (_Mcdflap *McdflapSession) Gem() (common.Address, error) {
	return _Mcdflap.Contract.Gem(&_Mcdflap.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() constant returns(address)
func (_Mcdflap *McdflapCallerSession) Gem() (common.Address, error) {
	return _Mcdflap.Contract.Gem(&_Mcdflap.CallOpts)
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() constant returns(uint256)
func (_Mcdflap *McdflapCaller) Kicks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflap.contract.Call(opts, out, "kicks")
	return *ret0, err
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() constant returns(uint256)
func (_Mcdflap *McdflapSession) Kicks() (*big.Int, error) {
	return _Mcdflap.Contract.Kicks(&_Mcdflap.CallOpts)
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() constant returns(uint256)
func (_Mcdflap *McdflapCallerSession) Kicks() (*big.Int, error) {
	return _Mcdflap.Contract.Kicks(&_Mcdflap.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Mcdflap *McdflapCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflap.contract.Call(opts, out, "live")
	return *ret0, err
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Mcdflap *McdflapSession) Live() (*big.Int, error) {
	return _Mcdflap.Contract.Live(&_Mcdflap.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Mcdflap *McdflapCallerSession) Live() (*big.Int, error) {
	return _Mcdflap.Contract.Live(&_Mcdflap.CallOpts)
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() constant returns(uint48)
func (_Mcdflap *McdflapCaller) Tau(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflap.contract.Call(opts, out, "tau")
	return *ret0, err
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() constant returns(uint48)
func (_Mcdflap *McdflapSession) Tau() (*big.Int, error) {
	return _Mcdflap.Contract.Tau(&_Mcdflap.CallOpts)
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() constant returns(uint48)
func (_Mcdflap *McdflapCallerSession) Tau() (*big.Int, error) {
	return _Mcdflap.Contract.Tau(&_Mcdflap.CallOpts)
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() constant returns(uint48)
func (_Mcdflap *McdflapCaller) Ttl(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflap.contract.Call(opts, out, "ttl")
	return *ret0, err
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() constant returns(uint48)
func (_Mcdflap *McdflapSession) Ttl() (*big.Int, error) {
	return _Mcdflap.Contract.Ttl(&_Mcdflap.CallOpts)
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() constant returns(uint48)
func (_Mcdflap *McdflapCallerSession) Ttl() (*big.Int, error) {
	return _Mcdflap.Contract.Ttl(&_Mcdflap.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Mcdflap *McdflapCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mcdflap.contract.Call(opts, out, "vat")
	return *ret0, err
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Mcdflap *McdflapSession) Vat() (common.Address, error) {
	return _Mcdflap.Contract.Vat(&_Mcdflap.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Mcdflap *McdflapCallerSession) Vat() (common.Address, error) {
	return _Mcdflap.Contract.Vat(&_Mcdflap.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Mcdflap *McdflapCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflap.contract.Call(opts, out, "wards", arg0)
	return *ret0, err
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Mcdflap *McdflapSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Mcdflap.Contract.Wards(&_Mcdflap.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Mcdflap *McdflapCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Mcdflap.Contract.Wards(&_Mcdflap.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0xa2f91af2.
//
// Solidity: function cage(uint256 rad) returns()
func (_Mcdflap *McdflapTransactor) Cage(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _Mcdflap.contract.Transact(opts, "cage", rad)
}

// Cage is a paid mutator transaction binding the contract method 0xa2f91af2.
//
// Solidity: function cage(uint256 rad) returns()
func (_Mcdflap *McdflapSession) Cage(rad *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.Cage(&_Mcdflap.TransactOpts, rad)
}

// Cage is a paid mutator transaction binding the contract method 0xa2f91af2.
//
// Solidity: function cage(uint256 rad) returns()
func (_Mcdflap *McdflapTransactorSession) Cage(rad *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.Cage(&_Mcdflap.TransactOpts, rad)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_Mcdflap *McdflapTransactor) Deal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mcdflap.contract.Transact(opts, "deal", id)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_Mcdflap *McdflapSession) Deal(id *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.Deal(&_Mcdflap.TransactOpts, id)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_Mcdflap *McdflapTransactorSession) Deal(id *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.Deal(&_Mcdflap.TransactOpts, id)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Mcdflap *McdflapTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Mcdflap.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Mcdflap *McdflapSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Mcdflap.Contract.Deny(&_Mcdflap.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Mcdflap *McdflapTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Mcdflap.Contract.Deny(&_Mcdflap.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Mcdflap *McdflapTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Mcdflap.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Mcdflap *McdflapSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.File(&_Mcdflap.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Mcdflap *McdflapTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.File(&_Mcdflap.TransactOpts, what, data)
}

// Kick is a paid mutator transaction binding the contract method 0xca40c419.
//
// Solidity: function kick(uint256 lot, uint256 bid) returns(uint256 id)
func (_Mcdflap *McdflapTransactor) Kick(opts *bind.TransactOpts, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflap.contract.Transact(opts, "kick", lot, bid)
}

// Kick is a paid mutator transaction binding the contract method 0xca40c419.
//
// Solidity: function kick(uint256 lot, uint256 bid) returns(uint256 id)
func (_Mcdflap *McdflapSession) Kick(lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.Kick(&_Mcdflap.TransactOpts, lot, bid)
}

// Kick is a paid mutator transaction binding the contract method 0xca40c419.
//
// Solidity: function kick(uint256 lot, uint256 bid) returns(uint256 id)
func (_Mcdflap *McdflapTransactorSession) Kick(lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.Kick(&_Mcdflap.TransactOpts, lot, bid)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Mcdflap *McdflapTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Mcdflap.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Mcdflap *McdflapSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Mcdflap.Contract.Rely(&_Mcdflap.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Mcdflap *McdflapTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Mcdflap.Contract.Rely(&_Mcdflap.TransactOpts, usr)
}

// Tend is a paid mutator transaction binding the contract method 0x4b43ed12.
//
// Solidity: function tend(uint256 id, uint256 lot, uint256 bid) returns()
func (_Mcdflap *McdflapTransactor) Tend(opts *bind.TransactOpts, id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflap.contract.Transact(opts, "tend", id, lot, bid)
}

// Tend is a paid mutator transaction binding the contract method 0x4b43ed12.
//
// Solidity: function tend(uint256 id, uint256 lot, uint256 bid) returns()
func (_Mcdflap *McdflapSession) Tend(id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.Tend(&_Mcdflap.TransactOpts, id, lot, bid)
}

// Tend is a paid mutator transaction binding the contract method 0x4b43ed12.
//
// Solidity: function tend(uint256 id, uint256 lot, uint256 bid) returns()
func (_Mcdflap *McdflapTransactorSession) Tend(id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.Tend(&_Mcdflap.TransactOpts, id, lot, bid)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_Mcdflap *McdflapTransactor) Tick(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mcdflap.contract.Transact(opts, "tick", id)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_Mcdflap *McdflapSession) Tick(id *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.Tick(&_Mcdflap.TransactOpts, id)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_Mcdflap *McdflapTransactorSession) Tick(id *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.Tick(&_Mcdflap.TransactOpts, id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_Mcdflap *McdflapTransactor) Yank(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mcdflap.contract.Transact(opts, "yank", id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_Mcdflap *McdflapSession) Yank(id *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.Yank(&_Mcdflap.TransactOpts, id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_Mcdflap *McdflapTransactorSession) Yank(id *big.Int) (*types.Transaction, error) {
	return _Mcdflap.Contract.Yank(&_Mcdflap.TransactOpts, id)
}

// McdflapKickIterator is returned from FilterKick and is used to iterate over the raw logs and unpacked data for Kick events raised by the Mcdflap contract.
type McdflapKickIterator struct {
	Event *McdflapKick // Event containing the contract specifics and raw log

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
func (it *McdflapKickIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(McdflapKick)
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
		it.Event = new(McdflapKick)
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
func (it *McdflapKickIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *McdflapKickIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// McdflapKick represents a Kick event raised by the Mcdflap contract.
type McdflapKick struct {
	Id  *big.Int
	Lot *big.Int
	Bid *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKick is a free log retrieval operation binding the contract event 0xe6dde59cbc017becba89714a037778d234a84ce7f0a137487142a007e580d609.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid)
func (_Mcdflap *McdflapFilterer) FilterKick(opts *bind.FilterOpts) (*McdflapKickIterator, error) {

	logs, sub, err := _Mcdflap.contract.FilterLogs(opts, "Kick")
	if err != nil {
		return nil, err
	}
	return &McdflapKickIterator{contract: _Mcdflap.contract, event: "Kick", logs: logs, sub: sub}, nil
}

// WatchKick is a free log subscription operation binding the contract event 0xe6dde59cbc017becba89714a037778d234a84ce7f0a137487142a007e580d609.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid)
func (_Mcdflap *McdflapFilterer) WatchKick(opts *bind.WatchOpts, sink chan<- *McdflapKick) (event.Subscription, error) {

	logs, sub, err := _Mcdflap.contract.WatchLogs(opts, "Kick")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(McdflapKick)
				if err := _Mcdflap.contract.UnpackLog(event, "Kick", log); err != nil {
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

// ParseKick is a log parse operation binding the contract event 0xe6dde59cbc017becba89714a037778d234a84ce7f0a137487142a007e580d609.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid)
func (_Mcdflap *McdflapFilterer) ParseKick(log types.Log) (*McdflapKick, error) {
	event := new(McdflapKick)
	if err := _Mcdflap.contract.UnpackLog(event, "Kick", log); err != nil {
		return nil, err
	}
	return event, nil
}
