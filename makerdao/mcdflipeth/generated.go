// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mcdflipeth

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

// McdflipethABI is the input ABI used to generate the binding from.
const McdflipethABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gal\",\"type\":\"address\"}],\"name\":\"Kick\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"beg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"tic\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"end\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gal\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"deal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"dent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ilk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gal\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"kick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kicks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tau\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"tend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"tick\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ttl\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// McdflipethBin is the compiled bytecode used for deploying new contracts.
var McdflipethBin = "0x6080604052670e92596fd6290000600455612a30600560006101000a81548165ffffffffffff021916908365ffffffffffff1602179055506202a300600560066101000a81548165ffffffffffff021916908365ffffffffffff160217905550600060065534801561007057600080fd5b50604051612d5c380380612d5c8339818101604052604081101561009357600080fd5b81019080805190602001909291908051906020019092919050505081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060038190555060016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050612c118061014b6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806365fae35e116100a2578063c5ce281e11610071578063c5ce281e1461051d578063c959c42b1461053b578063cfc4af5514610569578063cfdd330214610597578063fc7b6aee146105b55761010b565b806365fae35e1461041f5780637d780d82146104635780639c52a7f114610481578063bf353dbb146104c55761010b565b80634423c5f1116100de5780634423c5f1146102565780634b43ed121461036d5780634e8b1dd5146103af5780635ff3a382146103dd5761010b565b806326e027f11461011057806329ae81141461013e578063351de6001461017657806336569e771461020c575b600080fd5b61013c6004803603602081101561012657600080fd5b81019080803590602001909291905050506105e3565b005b6101746004803603604081101561015457600080fd5b810190808035906020019092919080359060200190929190505050610b72565b005b6101f6600480360360a081101561018c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190505050610da5565b6040518082815260200191505060405180910390f35b61021461122e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102826004803603602081101561026c57600080fd5b8101908080359060200190929190505050611254565b604051808981526020018881526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018665ffffffffffff1665ffffffffffff1681526020018565ffffffffffff1665ffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019850505050505050505060405180910390f35b6103ad6004803603606081101561038357600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050611320565b005b6103b7611b6a565b604051808265ffffffffffff1665ffffffffffff16815260200191505060405180910390f35b61041d600480360360608110156103f357600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050611b82565b005b6104616004803603602081101561043557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061239f565b005b61046b6124cd565b6040518082815260200191505060405180910390f35b6104c36004803603602081101561049757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506124d3565b005b610507600480360360208110156104db57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612601565b6040518082815260200191505060405180910390f35b610525612619565b6040518082815260200191505060405180910390f35b6105676004803603602081101561055157600080fd5b810190808035906020019092919050505061261f565b005b610571612991565b604051808265ffffffffffff1665ffffffffffff16815260200191505060405180910390f35b61059f6129a9565b6040518082815260200191505060405180910390f35b6105e1600480360360208110156105cb57600080fd5b81019080803590602001909291905050506129af565b005b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610697576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c69707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610770576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f466c69707065722f6775792d6e6f742d7365740000000000000000000000000081525060200191505060405180910390fd5b6001600082815260200190815260200160002060050154600160008381526020019081526020016000206000015410610811576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f466c69707065722f616c72656164792d64656e742d706861736500000000000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636111be2e600354303360016000878152602001908152602001600020600101546040518563ffffffff1660e01b8152600401808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001945050505050600060405180830381600087803b15801561090e57600080fd5b505af1158015610922573d6000803e3d6000fd5b50505050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b336001600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660016000868152602001908152602001600020600001546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015610a4f57600080fd5b505af1158015610a63573d6000803e3d6000fd5b505050506001600082815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160146101000a81549065ffffffffffff021916905560028201601a6101000a81549065ffffffffffff02191690556003820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600582016000905550505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610c26576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c69707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b7f6265670000000000000000000000000000000000000000000000000000000000821415610c5a5780600481905550610d6e565b7f74746c0000000000000000000000000000000000000000000000000000000000821415610cac5780600560006101000a81548165ffffffffffff021916908365ffffffffffff160217905550610d6d565b7f7461750000000000000000000000000000000000000000000000000000000000821415610cfe5780600560066101000a81548165ffffffffffff021916908365ffffffffffff160217905550610d6c565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f466c69707065722f66696c652d756e7265636f676e697a65642d706172616d0081525060200191505060405180910390fd5b5b5b5961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a4505050565b600060016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610e5b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c69707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60065410610ef2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f466c69707065722f6f766572666c6f770000000000000000000000000000000081525060200191505060405180910390fd5b6006600081546001019190508190559050816001600083815260200190815260200160002060000181905550826001600083815260200190815260200160002060010181905550336001600083815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610fac42600560069054906101000a900465ffffffffffff16612b86565b60016000838152602001908152602001600020600201601a6101000a81548165ffffffffffff021916908365ffffffffffff160217905550856001600083815260200190815260200160002060030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550846001600083815260200190815260200160002060040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836001600083815260200190815260200160002060050181905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636111be2e6003543330876040518563ffffffff1660e01b8152600401808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001945050505050600060405180830381600087803b15801561119057600080fd5b505af11580156111a4573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fc84ce3a1172f0dec3173f04caaa6005151a4bfe40d4c9f3ea28dba5f719b2a7a838686896040518085815260200184815260200183815260200182815260200194505050505060405180910390a395945050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160149054906101000a900465ffffffffffff169080600201601a9054906101000a900465ffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905088565b600073ffffffffffffffffffffffffffffffffffffffff166001600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156113f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f466c69707065722f6775792d6e6f742d7365740000000000000000000000000081525060200191505060405180910390fd5b426001600085815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff161180611467575060006001600085815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff16145b6114d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f466c69707065722f616c72656164792d66696e69736865642d7469630000000081525060200191505060405180910390fd5b4260016000858152602001908152602001600020600201601a9054906101000a900465ffffffffffff1665ffffffffffff161161157e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f466c69707065722f616c72656164792d66696e69736865642d656e640000000081525060200191505060405180910390fd5b60016000848152602001908152602001600020600101548214611609576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f466c69707065722f6c6f742d6e6f742d6d61746368696e67000000000000000081525060200191505060405180910390fd5b6001600084815260200190815260200160002060050154811115611695576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f466c69707065722f6869676865722d7468616e2d74616200000000000000000081525060200191505060405180910390fd5b60016000848152602001908152602001600020600001548111611720576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c69707065722f6269642d6e6f742d6869676865720000000000000000000081525060200191505060405180910390fd5b6117426004546001600086815260200190815260200160002060000154612bb0565b61175482670de0b6b3a7640000612bb0565b1015806117765750600160008481526020019081526020016000206005015481145b6117e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f466c69707065722f696e73756666696369656e742d696e63726561736500000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b336001600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660016000888152602001908152602001600020600001546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561191157600080fd5b505af1158015611925573d6000803e3d6000fd5b50505050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b336001600087815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160008881526020019081526020016000206000015485036040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015611a5457600080fd5b505af1158015611a68573d6000803e3d6000fd5b50505050336001600085815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806001600085815260200190815260200160002060000181905550611afa42600560009054906101000a900465ffffffffffff16612b86565b6001600085815260200190815260200160002060020160146101000a81548165ffffffffffff021916908365ffffffffffff1602179055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a450505050565b600560009054906101000a900465ffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff166001600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611c5b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f466c69707065722f6775792d6e6f742d7365740000000000000000000000000081525060200191505060405180910390fd5b426001600085815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff161180611cc9575060006001600085815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff16145b611d3b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f466c69707065722f616c72656164792d66696e69736865642d7469630000000081525060200191505060405180910390fd5b4260016000858152602001908152602001600020600201601a9054906101000a900465ffffffffffff1665ffffffffffff1611611de0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f466c69707065722f616c72656164792d66696e69736865642d656e640000000081525060200191505060405180910390fd5b60016000848152602001908152602001600020600001548114611e6b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f466c69707065722f6e6f742d6d61746368696e672d626964000000000000000081525060200191505060405180910390fd5b60016000848152602001908152602001600020600501548114611ef6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f466c69707065722f74656e642d6e6f742d66696e69736865640000000000000081525060200191505060405180910390fd5b60016000848152602001908152602001600020600101548210611f81576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f466c69707065722f6c6f742d6e6f742d6c6f776572000000000000000000000081525060200191505060405180910390fd5b611fa96001600085815260200190815260200160002060010154670de0b6b3a7640000612bb0565b611fb560045484612bb0565b1115612029576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f466c69707065722f696e73756666696369656e742d646563726561736500000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b336001600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561213c57600080fd5b505af1158015612150573d6000803e3d6000fd5b50505050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636111be2e600354306001600088815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1686600160008a815260200190815260200160002060010154036040518563ffffffff1660e01b8152600401808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001945050505050600060405180830381600087803b15801561228957600080fd5b505af115801561229d573d6000803e3d6000fd5b50505050336001600085815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160008581526020019081526020016000206001018190555061232f42600560009054906101000a900465ffffffffffff16612b86565b6001600085815260200190815260200160002060020160146101000a81548165ffffffffffff021916908365ffffffffffff1602179055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a450505050565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414612453576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c69707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60045481565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414612587576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c69707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60006020528060005260406000206000915090505481565b60035481565b60006001600083815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff16141580156126c95750426001600083815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff1610806126c857504260016000838152602001908152602001600020600201601a9054906101000a900465ffffffffffff1665ffffffffffff16105b5b61273b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f466c69707065722f6e6f742d66696e697368656400000000000000000000000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636111be2e600354306001600086815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660016000878152602001908152602001600020600101546040518563ffffffff1660e01b8152600401808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001945050505050600060405180830381600087803b15801561286e57600080fd5b505af1158015612882573d6000803e3d6000fd5b505050506001600082815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160146101000a81549065ffffffffffff021916905560028201601a6101000a81549065ffffffffffff02191690556003820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600582016000905550505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b600560069054906101000a900465ffffffffffff1681565b60065481565b4260016000838152602001908152602001600020600201601a9054906101000a900465ffffffffffff1665ffffffffffff1610612a54576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f466c69707065722f6e6f742d66696e697368656400000000000000000000000081525060200191505060405180910390fd5b60006001600083815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff1614612afa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f466c69707065722f6269642d616c72656164792d706c6163656400000000000081525060200191505060405180910390fd5b612b1842600560069054906101000a900465ffffffffffff16612b86565b60016000838152602001908152602001600020600201601a6101000a81548165ffffffffffff021916908365ffffffffffff1602179055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60008265ffffffffffff1682840191508165ffffffffffff161015612baa57600080fd5b92915050565b600080821480612bcd5750828283850292508281612bca57fe5b04145b612bd657600080fd5b9291505056fea265627a7a723158204dc61355cd6fa8e35aaa1844cbf949745778e1cd4d9194a4628fcd2253ddbcfc64736f6c634300050c003200000000000000000000000035d1b3f3d7966a1dfe207aa4514c12a259a0492b4554482d41000000000000000000000000000000000000000000000000000000"

const McdflipethHexAddress = "0xd8a04f5412223f513dc55f839574430f5ec15531"

// DeployMcdflipeth deploys a new Ethereum contract, binding an instance of Mcdflipeth to it.
func DeployMcdflipeth(auth *bind.TransactOpts, backend bind.ContractBackend, vat_ common.Address, ilk_ [32]byte) (common.Address, *types.Transaction, *Mcdflipeth, error) {
	parsed, err := abi.JSON(strings.NewReader(McdflipethABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(McdflipethBin), backend, vat_, ilk_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mcdflipeth{McdflipethCaller: McdflipethCaller{contract: contract}, McdflipethTransactor: McdflipethTransactor{contract: contract}, McdflipethFilterer: McdflipethFilterer{contract: contract}}, nil
}

// Mcdflipeth is an auto generated Go binding around an Ethereum contract.
type Mcdflipeth struct {
	McdflipethCaller     // Read-only binding to the contract
	McdflipethTransactor // Write-only binding to the contract
	McdflipethFilterer   // Log filterer for contract events
}

// McdflipethCaller is an auto generated read-only Go binding around an Ethereum contract.
type McdflipethCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McdflipethTransactor is an auto generated write-only Go binding around an Ethereum contract.
type McdflipethTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McdflipethFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type McdflipethFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McdflipethSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type McdflipethSession struct {
	Contract     *Mcdflipeth       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// McdflipethCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type McdflipethCallerSession struct {
	Contract *McdflipethCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// McdflipethTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type McdflipethTransactorSession struct {
	Contract     *McdflipethTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// McdflipethRaw is an auto generated low-level Go binding around an Ethereum contract.
type McdflipethRaw struct {
	Contract *Mcdflipeth // Generic contract binding to access the raw methods on
}

// McdflipethCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type McdflipethCallerRaw struct {
	Contract *McdflipethCaller // Generic read-only contract binding to access the raw methods on
}

// McdflipethTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type McdflipethTransactorRaw struct {
	Contract *McdflipethTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMcdflipeth creates a new instance of Mcdflipeth, bound to a specific deployed contract.
func NewMcdflipeth(address common.Address, backend bind.ContractBackend) (*Mcdflipeth, error) {
	contract, err := bindMcdflipeth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mcdflipeth{McdflipethCaller: McdflipethCaller{contract: contract}, McdflipethTransactor: McdflipethTransactor{contract: contract}, McdflipethFilterer: McdflipethFilterer{contract: contract}}, nil
}

// NewMcdflipethCaller creates a new read-only instance of Mcdflipeth, bound to a specific deployed contract.
func NewMcdflipethCaller(address common.Address, caller bind.ContractCaller) (*McdflipethCaller, error) {
	contract, err := bindMcdflipeth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &McdflipethCaller{contract: contract}, nil
}

// NewMcdflipethTransactor creates a new write-only instance of Mcdflipeth, bound to a specific deployed contract.
func NewMcdflipethTransactor(address common.Address, transactor bind.ContractTransactor) (*McdflipethTransactor, error) {
	contract, err := bindMcdflipeth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &McdflipethTransactor{contract: contract}, nil
}

// NewMcdflipethFilterer creates a new log filterer instance of Mcdflipeth, bound to a specific deployed contract.
func NewMcdflipethFilterer(address common.Address, filterer bind.ContractFilterer) (*McdflipethFilterer, error) {
	contract, err := bindMcdflipeth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &McdflipethFilterer{contract: contract}, nil
}

// bindMcdflipeth binds a generic wrapper to an already deployed contract.
func bindMcdflipeth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(McdflipethABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mcdflipeth *McdflipethRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mcdflipeth.Contract.McdflipethCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mcdflipeth *McdflipethRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.McdflipethTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mcdflipeth *McdflipethRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.McdflipethTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mcdflipeth *McdflipethCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mcdflipeth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mcdflipeth *McdflipethTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mcdflipeth *McdflipethTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.contract.Transact(opts, method, params...)
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() constant returns(uint256)
func (_Mcdflipeth *McdflipethCaller) Beg(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflipeth.contract.Call(opts, out, "beg")
	return *ret0, err
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() constant returns(uint256)
func (_Mcdflipeth *McdflipethSession) Beg() (*big.Int, error) {
	return _Mcdflipeth.Contract.Beg(&_Mcdflipeth.CallOpts)
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() constant returns(uint256)
func (_Mcdflipeth *McdflipethCallerSession) Beg() (*big.Int, error) {
	return _Mcdflipeth.Contract.Beg(&_Mcdflipeth.CallOpts)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) constant returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end, address usr, address gal, uint256 tab)
func (_Mcdflipeth *McdflipethCaller) Bids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
	Usr common.Address
	Gal common.Address
	Tab *big.Int
}, error) {
	ret := new(struct {
		Bid *big.Int
		Lot *big.Int
		Guy common.Address
		Tic *big.Int
		End *big.Int
		Usr common.Address
		Gal common.Address
		Tab *big.Int
	})
	out := ret
	err := _Mcdflipeth.contract.Call(opts, out, "bids", arg0)
	return *ret, err
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) constant returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end, address usr, address gal, uint256 tab)
func (_Mcdflipeth *McdflipethSession) Bids(arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
	Usr common.Address
	Gal common.Address
	Tab *big.Int
}, error) {
	return _Mcdflipeth.Contract.Bids(&_Mcdflipeth.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) constant returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end, address usr, address gal, uint256 tab)
func (_Mcdflipeth *McdflipethCallerSession) Bids(arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
	Usr common.Address
	Gal common.Address
	Tab *big.Int
}, error) {
	return _Mcdflipeth.Contract.Bids(&_Mcdflipeth.CallOpts, arg0)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() constant returns(bytes32)
func (_Mcdflipeth *McdflipethCaller) Ilk(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mcdflipeth.contract.Call(opts, out, "ilk")
	return *ret0, err
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() constant returns(bytes32)
func (_Mcdflipeth *McdflipethSession) Ilk() ([32]byte, error) {
	return _Mcdflipeth.Contract.Ilk(&_Mcdflipeth.CallOpts)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() constant returns(bytes32)
func (_Mcdflipeth *McdflipethCallerSession) Ilk() ([32]byte, error) {
	return _Mcdflipeth.Contract.Ilk(&_Mcdflipeth.CallOpts)
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() constant returns(uint256)
func (_Mcdflipeth *McdflipethCaller) Kicks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflipeth.contract.Call(opts, out, "kicks")
	return *ret0, err
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() constant returns(uint256)
func (_Mcdflipeth *McdflipethSession) Kicks() (*big.Int, error) {
	return _Mcdflipeth.Contract.Kicks(&_Mcdflipeth.CallOpts)
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() constant returns(uint256)
func (_Mcdflipeth *McdflipethCallerSession) Kicks() (*big.Int, error) {
	return _Mcdflipeth.Contract.Kicks(&_Mcdflipeth.CallOpts)
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() constant returns(uint48)
func (_Mcdflipeth *McdflipethCaller) Tau(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflipeth.contract.Call(opts, out, "tau")
	return *ret0, err
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() constant returns(uint48)
func (_Mcdflipeth *McdflipethSession) Tau() (*big.Int, error) {
	return _Mcdflipeth.Contract.Tau(&_Mcdflipeth.CallOpts)
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() constant returns(uint48)
func (_Mcdflipeth *McdflipethCallerSession) Tau() (*big.Int, error) {
	return _Mcdflipeth.Contract.Tau(&_Mcdflipeth.CallOpts)
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() constant returns(uint48)
func (_Mcdflipeth *McdflipethCaller) Ttl(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflipeth.contract.Call(opts, out, "ttl")
	return *ret0, err
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() constant returns(uint48)
func (_Mcdflipeth *McdflipethSession) Ttl() (*big.Int, error) {
	return _Mcdflipeth.Contract.Ttl(&_Mcdflipeth.CallOpts)
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() constant returns(uint48)
func (_Mcdflipeth *McdflipethCallerSession) Ttl() (*big.Int, error) {
	return _Mcdflipeth.Contract.Ttl(&_Mcdflipeth.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Mcdflipeth *McdflipethCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mcdflipeth.contract.Call(opts, out, "vat")
	return *ret0, err
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Mcdflipeth *McdflipethSession) Vat() (common.Address, error) {
	return _Mcdflipeth.Contract.Vat(&_Mcdflipeth.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Mcdflipeth *McdflipethCallerSession) Vat() (common.Address, error) {
	return _Mcdflipeth.Contract.Vat(&_Mcdflipeth.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Mcdflipeth *McdflipethCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflipeth.contract.Call(opts, out, "wards", arg0)
	return *ret0, err
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Mcdflipeth *McdflipethSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Mcdflipeth.Contract.Wards(&_Mcdflipeth.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Mcdflipeth *McdflipethCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Mcdflipeth.Contract.Wards(&_Mcdflipeth.CallOpts, arg0)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_Mcdflipeth *McdflipethTransactor) Deal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.contract.Transact(opts, "deal", id)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_Mcdflipeth *McdflipethSession) Deal(id *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Deal(&_Mcdflipeth.TransactOpts, id)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_Mcdflipeth *McdflipethTransactorSession) Deal(id *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Deal(&_Mcdflipeth.TransactOpts, id)
}

// Dent is a paid mutator transaction binding the contract method 0x5ff3a382.
//
// Solidity: function dent(uint256 id, uint256 lot, uint256 bid) returns()
func (_Mcdflipeth *McdflipethTransactor) Dent(opts *bind.TransactOpts, id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.contract.Transact(opts, "dent", id, lot, bid)
}

// Dent is a paid mutator transaction binding the contract method 0x5ff3a382.
//
// Solidity: function dent(uint256 id, uint256 lot, uint256 bid) returns()
func (_Mcdflipeth *McdflipethSession) Dent(id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Dent(&_Mcdflipeth.TransactOpts, id, lot, bid)
}

// Dent is a paid mutator transaction binding the contract method 0x5ff3a382.
//
// Solidity: function dent(uint256 id, uint256 lot, uint256 bid) returns()
func (_Mcdflipeth *McdflipethTransactorSession) Dent(id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Dent(&_Mcdflipeth.TransactOpts, id, lot, bid)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Mcdflipeth *McdflipethTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Mcdflipeth.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Mcdflipeth *McdflipethSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Deny(&_Mcdflipeth.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Mcdflipeth *McdflipethTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Deny(&_Mcdflipeth.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Mcdflipeth *McdflipethTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Mcdflipeth *McdflipethSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.File(&_Mcdflipeth.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Mcdflipeth *McdflipethTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.File(&_Mcdflipeth.TransactOpts, what, data)
}

// Kick is a paid mutator transaction binding the contract method 0x351de600.
//
// Solidity: function kick(address usr, address gal, uint256 tab, uint256 lot, uint256 bid) returns(uint256 id)
func (_Mcdflipeth *McdflipethTransactor) Kick(opts *bind.TransactOpts, usr common.Address, gal common.Address, tab *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.contract.Transact(opts, "kick", usr, gal, tab, lot, bid)
}

// Kick is a paid mutator transaction binding the contract method 0x351de600.
//
// Solidity: function kick(address usr, address gal, uint256 tab, uint256 lot, uint256 bid) returns(uint256 id)
func (_Mcdflipeth *McdflipethSession) Kick(usr common.Address, gal common.Address, tab *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Kick(&_Mcdflipeth.TransactOpts, usr, gal, tab, lot, bid)
}

// Kick is a paid mutator transaction binding the contract method 0x351de600.
//
// Solidity: function kick(address usr, address gal, uint256 tab, uint256 lot, uint256 bid) returns(uint256 id)
func (_Mcdflipeth *McdflipethTransactorSession) Kick(usr common.Address, gal common.Address, tab *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Kick(&_Mcdflipeth.TransactOpts, usr, gal, tab, lot, bid)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Mcdflipeth *McdflipethTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Mcdflipeth.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Mcdflipeth *McdflipethSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Rely(&_Mcdflipeth.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Mcdflipeth *McdflipethTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Rely(&_Mcdflipeth.TransactOpts, usr)
}

// Tend is a paid mutator transaction binding the contract method 0x4b43ed12.
//
// Solidity: function tend(uint256 id, uint256 lot, uint256 bid) returns()
func (_Mcdflipeth *McdflipethTransactor) Tend(opts *bind.TransactOpts, id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.contract.Transact(opts, "tend", id, lot, bid)
}

// Tend is a paid mutator transaction binding the contract method 0x4b43ed12.
//
// Solidity: function tend(uint256 id, uint256 lot, uint256 bid) returns()
func (_Mcdflipeth *McdflipethSession) Tend(id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Tend(&_Mcdflipeth.TransactOpts, id, lot, bid)
}

// Tend is a paid mutator transaction binding the contract method 0x4b43ed12.
//
// Solidity: function tend(uint256 id, uint256 lot, uint256 bid) returns()
func (_Mcdflipeth *McdflipethTransactorSession) Tend(id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Tend(&_Mcdflipeth.TransactOpts, id, lot, bid)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_Mcdflipeth *McdflipethTransactor) Tick(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.contract.Transact(opts, "tick", id)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_Mcdflipeth *McdflipethSession) Tick(id *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Tick(&_Mcdflipeth.TransactOpts, id)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_Mcdflipeth *McdflipethTransactorSession) Tick(id *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Tick(&_Mcdflipeth.TransactOpts, id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_Mcdflipeth *McdflipethTransactor) Yank(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.contract.Transact(opts, "yank", id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_Mcdflipeth *McdflipethSession) Yank(id *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Yank(&_Mcdflipeth.TransactOpts, id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_Mcdflipeth *McdflipethTransactorSession) Yank(id *big.Int) (*types.Transaction, error) {
	return _Mcdflipeth.Contract.Yank(&_Mcdflipeth.TransactOpts, id)
}

// McdflipethKickIterator is returned from FilterKick and is used to iterate over the raw logs and unpacked data for Kick events raised by the Mcdflipeth contract.
type McdflipethKickIterator struct {
	Event *McdflipethKick // Event containing the contract specifics and raw log

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
func (it *McdflipethKickIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(McdflipethKick)
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
		it.Event = new(McdflipethKick)
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
func (it *McdflipethKickIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *McdflipethKickIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// McdflipethKick represents a Kick event raised by the Mcdflipeth contract.
type McdflipethKick struct {
	Id  *big.Int
	Lot *big.Int
	Bid *big.Int
	Tab *big.Int
	Usr common.Address
	Gal common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKick is a free log retrieval operation binding the contract event 0xc84ce3a1172f0dec3173f04caaa6005151a4bfe40d4c9f3ea28dba5f719b2a7a.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid, uint256 tab, address indexed usr, address indexed gal)
func (_Mcdflipeth *McdflipethFilterer) FilterKick(opts *bind.FilterOpts, usr []common.Address, gal []common.Address) (*McdflipethKickIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var galRule []interface{}
	for _, galItem := range gal {
		galRule = append(galRule, galItem)
	}

	logs, sub, err := _Mcdflipeth.contract.FilterLogs(opts, "Kick", usrRule, galRule)
	if err != nil {
		return nil, err
	}
	return &McdflipethKickIterator{contract: _Mcdflipeth.contract, event: "Kick", logs: logs, sub: sub}, nil
}

// WatchKick is a free log subscription operation binding the contract event 0xc84ce3a1172f0dec3173f04caaa6005151a4bfe40d4c9f3ea28dba5f719b2a7a.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid, uint256 tab, address indexed usr, address indexed gal)
func (_Mcdflipeth *McdflipethFilterer) WatchKick(opts *bind.WatchOpts, sink chan<- *McdflipethKick, usr []common.Address, gal []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var galRule []interface{}
	for _, galItem := range gal {
		galRule = append(galRule, galItem)
	}

	logs, sub, err := _Mcdflipeth.contract.WatchLogs(opts, "Kick", usrRule, galRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(McdflipethKick)
				if err := _Mcdflipeth.contract.UnpackLog(event, "Kick", log); err != nil {
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

// ParseKick is a log parse operation binding the contract event 0xc84ce3a1172f0dec3173f04caaa6005151a4bfe40d4c9f3ea28dba5f719b2a7a.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid, uint256 tab, address indexed usr, address indexed gal)
func (_Mcdflipeth *McdflipethFilterer) ParseKick(log types.Log) (*McdflipethKick, error) {
	event := new(McdflipethKick)
	if err := _Mcdflipeth.contract.UnpackLog(event, "Kick", log); err != nil {
		return nil, err
	}
	return event, nil
}
