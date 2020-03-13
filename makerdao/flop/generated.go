// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mcdflop

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

// McdflopABI is the input ABI used to generate the binding from.
const McdflopABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gem_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gal\",\"type\":\"address\"}],\"name\":\"Kick\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"beg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"tic\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"end\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"deal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"dent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"contractGemLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gal\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"kick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kicks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pad\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tau\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"tick\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ttl\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// McdflopBin is the compiled bytecode used for deploying new contracts.
var McdflopBin = "0x6080604052670e92596fd62900006004556714d1120d7b160000600555612a30600660006101000a81548165ffffffffffff021916908365ffffffffffff1602179055506202a3006006806101000a81548165ffffffffffff021916908365ffffffffffff160217905550600060075534801561007b57600080fd5b506040516122043803806122048339818101604052604081101561009e57600080fd5b81019080805190602001909291908051906020019092919050505060016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600881905550505061206c806101986000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80637d780d82116100ad578063bf353dbb11610071578063bf353dbb146104e7578063c959c42b1461053f578063cfc4af551461056d578063cfdd33021461059b578063fc7b6aee146105b95761012c565b80637d780d82146103dd5780639361266c146103fb578063957aa58c146104195780639c52a7f114610437578063b7e9cd241461047b5761012c565b80635ff3a382116100f45780635ff3a382146102b9578063626cb3c5146102fb57806365fae35e1461034557806369245009146103895780637bd2bea7146103935761012c565b806326e027f11461013157806329ae81141461015f57806336569e77146101975780634423c5f1146101e15780634e8b1dd51461028b575b600080fd5b61015d6004803603602081101561014757600080fd5b81019080803590602001909291905050506105e7565b005b6101956004803603604081101561017557600080fd5b810190808035906020019092919080359060200190929190505050610950565b005b61019f610bb7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61020d600480360360208110156101f757600080fd5b8101908080359060200190929190505050610bdd565b604051808681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018365ffffffffffff1665ffffffffffff1681526020018265ffffffffffff1665ffffffffffff1681526020019550505050505060405180910390f35b610293610c57565b604051808265ffffffffffff1665ffffffffffff16815260200191505060405180910390f35b6102f9600480360360608110156102cf57600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610c6f565b005b61030361132c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103876004803603602081101561035b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611352565b005b610391611480565b005b61039b6115b2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103e56115d8565b6040518082815260200191505060405180910390f35b6104036115de565b6040518082815260200191505060405180910390f35b6104216115e4565b6040518082815260200191505060405180910390f35b6104796004803603602081101561044d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115ea565b005b6104d16004803603606081101561049157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050611718565b6040518082815260200191505060405180910390f35b610529600480360360208110156104fd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611a33565b6040518082815260200191505060405180910390f35b61056b6004803603602081101561055557600080fd5b8101908080359060200190929190505050611a4b565b005b610575611da1565b604051808265ffffffffffff1665ffffffffffff16815260200191505060405180910390f35b6105a3611db8565b6040518082815260200191505060405180910390f35b6105e5600480360360208110156105cf57600080fd5b8101908080359060200190929190505050611dbe565b005b60006008541461065f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f466c6f707065722f7374696c6c2d6c697665000000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610738576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f466c6f707065722f6775792d6e6f742d7365740000000000000000000000000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f24e23eb600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166001600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660016000868152602001908152602001600020600001546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561088357600080fd5b505af1158015610897573d6000803e3d6000fd5b505050506001600082815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160146101000a81549065ffffffffffff021916905560028201601a6101000a81549065ffffffffffff021916905550505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610a04576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c6f707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b7f6265670000000000000000000000000000000000000000000000000000000000821415610a385780600481905550610b80565b7f7061640000000000000000000000000000000000000000000000000000000000821415610a6c5780600581905550610b7f565b7f74746c0000000000000000000000000000000000000000000000000000000000821415610abe5780600660006101000a81548165ffffffffffff021916908365ffffffffffff160217905550610b7e565b7f7461750000000000000000000000000000000000000000000000000000000000821415610b0f57806006806101000a81548165ffffffffffff021916908365ffffffffffff160217905550610b7d565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f466c6f707065722f66696c652d756e7265636f676e697a65642d706172616d0081525060200191505060405180910390fd5b5b5b5b5961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a4505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160149054906101000a900465ffffffffffff169080600201601a9054906101000a900465ffffffffffff16905085565b600660009054906101000a900465ffffffffffff1681565b600160085414610ce7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f466c6f707065722f6e6f742d6c6976650000000000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166001600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610dc0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f466c6f707065722f6775792d6e6f742d7365740000000000000000000000000081525060200191505060405180910390fd5b426001600085815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff161180610e2e575060006001600085815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff16145b610ea0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f466c6f707065722f616c72656164792d66696e69736865642d7469630000000081525060200191505060405180910390fd5b4260016000858152602001908152602001600020600201601a9054906101000a900465ffffffffffff1665ffffffffffff1611610f45576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f466c6f707065722f616c72656164792d66696e69736865642d656e640000000081525060200191505060405180910390fd5b60016000848152602001908152602001600020600001548114610fd0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f466c6f707065722f6e6f742d6d61746368696e672d626964000000000000000081525060200191505060405180910390fd5b6001600084815260200190815260200160002060010154821061105b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f466c6f707065722f6c6f742d6e6f742d6c6f776572000000000000000000000081525060200191505060405180910390fd5b6110836001600085815260200190815260200160002060010154670de0b6b3a7640000611fe1565b61108f60045484611fe1565b1115611103576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f466c6f707065722f696e73756666696369656e742d646563726561736500000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b336001600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561121657600080fd5b505af115801561122a573d6000803e3d6000fd5b50505050336001600085815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160016000858152602001908152602001600020600101819055506112bc42600660009054906101000a900465ffffffffffff1661200d565b6001600085815260200190815260200160002060020160146101000a81548165ffffffffffff021916908365ffffffffffff1602179055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a450505050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414611406576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c6f707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414611534576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c6f707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b600060088190555033600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a450565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b60055481565b60085481565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541461169e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c6f707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b600060016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146117ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f466c6f707065722f6e6f742d617574686f72697a65640000000000000000000081525060200191505060405180910390fd5b600160085414611846576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f466c6f707065722f6e6f742d6c6976650000000000000000000000000000000081525060200191505060405180910390fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600754106118dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f466c6f707065722f6f766572666c6f770000000000000000000000000000000081525060200191505060405180910390fd5b6007600081546001019190508190559050816001600083815260200190815260200160002060000181905550826001600083815260200190815260200160002060010181905550836001600083815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611996426006809054906101000a900465ffffffffffff1661200d565b60016000838152602001908152602001600020600201601a6101000a81548165ffffffffffff021916908365ffffffffffff1602179055508373ffffffffffffffffffffffffffffffffffffffff167f7e8881001566f9f89aedb9c5dc3d856a2b81e5235a8196413ed484be91cc0df682858560405180848152602001838152602001828152602001935050505060405180910390a29392505050565b60006020528060005260406000206000915090505481565b600160085414611ac3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f466c6f707065722f6e6f742d6c6976650000000000000000000000000000000081525060200191505060405180910390fd5b60006001600083815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff1614158015611b6d5750426001600083815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff161080611b6c57504260016000838152602001908152602001600020600201601a9054906101000a900465ffffffffffff1665ffffffffffff16105b5b611bdf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f466c6f707065722f6e6f742d66696e697368656400000000000000000000000081525060200191505060405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f196001600084815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660016000858152602001908152602001600020600101546040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015611cd457600080fd5b505af1158015611ce8573d6000803e3d6000fd5b505050506001600082815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160146101000a81549065ffffffffffff021916905560028201601a6101000a81549065ffffffffffff021916905550505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b6006809054906101000a900465ffffffffffff1681565b60075481565b4260016000838152602001908152602001600020600201601a9054906101000a900465ffffffffffff1665ffffffffffff1610611e63576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f466c6f707065722f6e6f742d66696e697368656400000000000000000000000081525060200191505060405180910390fd5b60006001600083815260200190815260200160002060020160149054906101000a900465ffffffffffff1665ffffffffffff1614611f09576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f466c6f707065722f6269642d616c72656164792d706c6163656400000000000081525060200191505060405180910390fd5b670de0b6b3a7640000611f346005546001600085815260200190815260200160002060010154611fe1565b81611f3b57fe5b046001600083815260200190815260200160002060010181905550611f73426006809054906101000a900465ffffffffffff1661200d565b60016000838152602001908152602001600020600201601a6101000a81548165ffffffffffff021916908365ffffffffffff1602179055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b600080821480611ffe5750828283850292508281611ffb57fe5b04145b61200757600080fd5b92915050565b60008265ffffffffffff1682840191508165ffffffffffff16101561203157600080fd5b9291505056fea265627a7a7231582027a32232cdfd81f358da7402624811cae6466ebed6f2d37b16e9bee1f44da97764736f6c634300050c003200000000000000000000000035d1b3f3d7966a1dfe207aa4514c12a259a0492b0000000000000000000000009f8f72aa9304c8b593d555f12ef6589cc3a579a2"

const McdflopHexAddress = "0x4d95a049d5b0b7d32058cd3f2163015747522e99"

// DeployMcdflop deploys a new Ethereum contract, binding an instance of Mcdflop to it.
func DeployMcdflop(auth *bind.TransactOpts, backend bind.ContractBackend, vat_ common.Address, gem_ common.Address) (common.Address, *types.Transaction, *Mcdflop, error) {
	parsed, err := abi.JSON(strings.NewReader(McdflopABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(McdflopBin), backend, vat_, gem_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mcdflop{McdflopCaller: McdflopCaller{contract: contract}, McdflopTransactor: McdflopTransactor{contract: contract}, McdflopFilterer: McdflopFilterer{contract: contract}}, nil
}

// Mcdflop is an auto generated Go binding around an Ethereum contract.
type Mcdflop struct {
	McdflopCaller     // Read-only binding to the contract
	McdflopTransactor // Write-only binding to the contract
	McdflopFilterer   // Log filterer for contract events
}

// McdflopCaller is an auto generated read-only Go binding around an Ethereum contract.
type McdflopCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McdflopTransactor is an auto generated write-only Go binding around an Ethereum contract.
type McdflopTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McdflopFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type McdflopFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McdflopSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type McdflopSession struct {
	Contract     *Mcdflop          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// McdflopCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type McdflopCallerSession struct {
	Contract *McdflopCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// McdflopTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type McdflopTransactorSession struct {
	Contract     *McdflopTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// McdflopRaw is an auto generated low-level Go binding around an Ethereum contract.
type McdflopRaw struct {
	Contract *Mcdflop // Generic contract binding to access the raw methods on
}

// McdflopCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type McdflopCallerRaw struct {
	Contract *McdflopCaller // Generic read-only contract binding to access the raw methods on
}

// McdflopTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type McdflopTransactorRaw struct {
	Contract *McdflopTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMcdflop creates a new instance of Mcdflop, bound to a specific deployed contract.
func NewMcdflop(address common.Address, backend bind.ContractBackend) (*Mcdflop, error) {
	contract, err := bindMcdflop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mcdflop{McdflopCaller: McdflopCaller{contract: contract}, McdflopTransactor: McdflopTransactor{contract: contract}, McdflopFilterer: McdflopFilterer{contract: contract}}, nil
}

// NewMcdflopCaller creates a new read-only instance of Mcdflop, bound to a specific deployed contract.
func NewMcdflopCaller(address common.Address, caller bind.ContractCaller) (*McdflopCaller, error) {
	contract, err := bindMcdflop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &McdflopCaller{contract: contract}, nil
}

// NewMcdflopTransactor creates a new write-only instance of Mcdflop, bound to a specific deployed contract.
func NewMcdflopTransactor(address common.Address, transactor bind.ContractTransactor) (*McdflopTransactor, error) {
	contract, err := bindMcdflop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &McdflopTransactor{contract: contract}, nil
}

// NewMcdflopFilterer creates a new log filterer instance of Mcdflop, bound to a specific deployed contract.
func NewMcdflopFilterer(address common.Address, filterer bind.ContractFilterer) (*McdflopFilterer, error) {
	contract, err := bindMcdflop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &McdflopFilterer{contract: contract}, nil
}

// bindMcdflop binds a generic wrapper to an already deployed contract.
func bindMcdflop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(McdflopABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mcdflop *McdflopRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mcdflop.Contract.McdflopCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mcdflop *McdflopRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcdflop.Contract.McdflopTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mcdflop *McdflopRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mcdflop.Contract.McdflopTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mcdflop *McdflopCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mcdflop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mcdflop *McdflopTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcdflop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mcdflop *McdflopTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mcdflop.Contract.contract.Transact(opts, method, params...)
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() constant returns(uint256)
func (_Mcdflop *McdflopCaller) Beg(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflop.contract.Call(opts, out, "beg")
	return *ret0, err
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() constant returns(uint256)
func (_Mcdflop *McdflopSession) Beg() (*big.Int, error) {
	return _Mcdflop.Contract.Beg(&_Mcdflop.CallOpts)
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() constant returns(uint256)
func (_Mcdflop *McdflopCallerSession) Beg() (*big.Int, error) {
	return _Mcdflop.Contract.Beg(&_Mcdflop.CallOpts)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) constant returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end)
func (_Mcdflop *McdflopCaller) Bids(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Mcdflop.contract.Call(opts, out, "bids", arg0)
	return *ret, err
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) constant returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end)
func (_Mcdflop *McdflopSession) Bids(arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
}, error) {
	return _Mcdflop.Contract.Bids(&_Mcdflop.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) constant returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end)
func (_Mcdflop *McdflopCallerSession) Bids(arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
}, error) {
	return _Mcdflop.Contract.Bids(&_Mcdflop.CallOpts, arg0)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() constant returns(address)
func (_Mcdflop *McdflopCaller) Gem(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mcdflop.contract.Call(opts, out, "gem")
	return *ret0, err
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() constant returns(address)
func (_Mcdflop *McdflopSession) Gem() (common.Address, error) {
	return _Mcdflop.Contract.Gem(&_Mcdflop.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() constant returns(address)
func (_Mcdflop *McdflopCallerSession) Gem() (common.Address, error) {
	return _Mcdflop.Contract.Gem(&_Mcdflop.CallOpts)
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() constant returns(uint256)
func (_Mcdflop *McdflopCaller) Kicks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflop.contract.Call(opts, out, "kicks")
	return *ret0, err
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() constant returns(uint256)
func (_Mcdflop *McdflopSession) Kicks() (*big.Int, error) {
	return _Mcdflop.Contract.Kicks(&_Mcdflop.CallOpts)
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() constant returns(uint256)
func (_Mcdflop *McdflopCallerSession) Kicks() (*big.Int, error) {
	return _Mcdflop.Contract.Kicks(&_Mcdflop.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Mcdflop *McdflopCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflop.contract.Call(opts, out, "live")
	return *ret0, err
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Mcdflop *McdflopSession) Live() (*big.Int, error) {
	return _Mcdflop.Contract.Live(&_Mcdflop.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Mcdflop *McdflopCallerSession) Live() (*big.Int, error) {
	return _Mcdflop.Contract.Live(&_Mcdflop.CallOpts)
}

// Pad is a free data retrieval call binding the contract method 0x9361266c.
//
// Solidity: function pad() constant returns(uint256)
func (_Mcdflop *McdflopCaller) Pad(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflop.contract.Call(opts, out, "pad")
	return *ret0, err
}

// Pad is a free data retrieval call binding the contract method 0x9361266c.
//
// Solidity: function pad() constant returns(uint256)
func (_Mcdflop *McdflopSession) Pad() (*big.Int, error) {
	return _Mcdflop.Contract.Pad(&_Mcdflop.CallOpts)
}

// Pad is a free data retrieval call binding the contract method 0x9361266c.
//
// Solidity: function pad() constant returns(uint256)
func (_Mcdflop *McdflopCallerSession) Pad() (*big.Int, error) {
	return _Mcdflop.Contract.Pad(&_Mcdflop.CallOpts)
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() constant returns(uint48)
func (_Mcdflop *McdflopCaller) Tau(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflop.contract.Call(opts, out, "tau")
	return *ret0, err
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() constant returns(uint48)
func (_Mcdflop *McdflopSession) Tau() (*big.Int, error) {
	return _Mcdflop.Contract.Tau(&_Mcdflop.CallOpts)
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() constant returns(uint48)
func (_Mcdflop *McdflopCallerSession) Tau() (*big.Int, error) {
	return _Mcdflop.Contract.Tau(&_Mcdflop.CallOpts)
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() constant returns(uint48)
func (_Mcdflop *McdflopCaller) Ttl(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflop.contract.Call(opts, out, "ttl")
	return *ret0, err
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() constant returns(uint48)
func (_Mcdflop *McdflopSession) Ttl() (*big.Int, error) {
	return _Mcdflop.Contract.Ttl(&_Mcdflop.CallOpts)
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() constant returns(uint48)
func (_Mcdflop *McdflopCallerSession) Ttl() (*big.Int, error) {
	return _Mcdflop.Contract.Ttl(&_Mcdflop.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Mcdflop *McdflopCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mcdflop.contract.Call(opts, out, "vat")
	return *ret0, err
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Mcdflop *McdflopSession) Vat() (common.Address, error) {
	return _Mcdflop.Contract.Vat(&_Mcdflop.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Mcdflop *McdflopCallerSession) Vat() (common.Address, error) {
	return _Mcdflop.Contract.Vat(&_Mcdflop.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() constant returns(address)
func (_Mcdflop *McdflopCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mcdflop.contract.Call(opts, out, "vow")
	return *ret0, err
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() constant returns(address)
func (_Mcdflop *McdflopSession) Vow() (common.Address, error) {
	return _Mcdflop.Contract.Vow(&_Mcdflop.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() constant returns(address)
func (_Mcdflop *McdflopCallerSession) Vow() (common.Address, error) {
	return _Mcdflop.Contract.Vow(&_Mcdflop.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Mcdflop *McdflopCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mcdflop.contract.Call(opts, out, "wards", arg0)
	return *ret0, err
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Mcdflop *McdflopSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Mcdflop.Contract.Wards(&_Mcdflop.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Mcdflop *McdflopCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Mcdflop.Contract.Wards(&_Mcdflop.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Mcdflop *McdflopTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcdflop.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Mcdflop *McdflopSession) Cage() (*types.Transaction, error) {
	return _Mcdflop.Contract.Cage(&_Mcdflop.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Mcdflop *McdflopTransactorSession) Cage() (*types.Transaction, error) {
	return _Mcdflop.Contract.Cage(&_Mcdflop.TransactOpts)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_Mcdflop *McdflopTransactor) Deal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mcdflop.contract.Transact(opts, "deal", id)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_Mcdflop *McdflopSession) Deal(id *big.Int) (*types.Transaction, error) {
	return _Mcdflop.Contract.Deal(&_Mcdflop.TransactOpts, id)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_Mcdflop *McdflopTransactorSession) Deal(id *big.Int) (*types.Transaction, error) {
	return _Mcdflop.Contract.Deal(&_Mcdflop.TransactOpts, id)
}

// Dent is a paid mutator transaction binding the contract method 0x5ff3a382.
//
// Solidity: function dent(uint256 id, uint256 lot, uint256 bid) returns()
func (_Mcdflop *McdflopTransactor) Dent(opts *bind.TransactOpts, id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflop.contract.Transact(opts, "dent", id, lot, bid)
}

// Dent is a paid mutator transaction binding the contract method 0x5ff3a382.
//
// Solidity: function dent(uint256 id, uint256 lot, uint256 bid) returns()
func (_Mcdflop *McdflopSession) Dent(id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflop.Contract.Dent(&_Mcdflop.TransactOpts, id, lot, bid)
}

// Dent is a paid mutator transaction binding the contract method 0x5ff3a382.
//
// Solidity: function dent(uint256 id, uint256 lot, uint256 bid) returns()
func (_Mcdflop *McdflopTransactorSession) Dent(id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflop.Contract.Dent(&_Mcdflop.TransactOpts, id, lot, bid)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Mcdflop *McdflopTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Mcdflop.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Mcdflop *McdflopSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Mcdflop.Contract.Deny(&_Mcdflop.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Mcdflop *McdflopTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Mcdflop.Contract.Deny(&_Mcdflop.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Mcdflop *McdflopTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Mcdflop.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Mcdflop *McdflopSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Mcdflop.Contract.File(&_Mcdflop.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Mcdflop *McdflopTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Mcdflop.Contract.File(&_Mcdflop.TransactOpts, what, data)
}

// Kick is a paid mutator transaction binding the contract method 0xb7e9cd24.
//
// Solidity: function kick(address gal, uint256 lot, uint256 bid) returns(uint256 id)
func (_Mcdflop *McdflopTransactor) Kick(opts *bind.TransactOpts, gal common.Address, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflop.contract.Transact(opts, "kick", gal, lot, bid)
}

// Kick is a paid mutator transaction binding the contract method 0xb7e9cd24.
//
// Solidity: function kick(address gal, uint256 lot, uint256 bid) returns(uint256 id)
func (_Mcdflop *McdflopSession) Kick(gal common.Address, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflop.Contract.Kick(&_Mcdflop.TransactOpts, gal, lot, bid)
}

// Kick is a paid mutator transaction binding the contract method 0xb7e9cd24.
//
// Solidity: function kick(address gal, uint256 lot, uint256 bid) returns(uint256 id)
func (_Mcdflop *McdflopTransactorSession) Kick(gal common.Address, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _Mcdflop.Contract.Kick(&_Mcdflop.TransactOpts, gal, lot, bid)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Mcdflop *McdflopTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Mcdflop.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Mcdflop *McdflopSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Mcdflop.Contract.Rely(&_Mcdflop.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Mcdflop *McdflopTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Mcdflop.Contract.Rely(&_Mcdflop.TransactOpts, usr)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_Mcdflop *McdflopTransactor) Tick(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mcdflop.contract.Transact(opts, "tick", id)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_Mcdflop *McdflopSession) Tick(id *big.Int) (*types.Transaction, error) {
	return _Mcdflop.Contract.Tick(&_Mcdflop.TransactOpts, id)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_Mcdflop *McdflopTransactorSession) Tick(id *big.Int) (*types.Transaction, error) {
	return _Mcdflop.Contract.Tick(&_Mcdflop.TransactOpts, id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_Mcdflop *McdflopTransactor) Yank(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mcdflop.contract.Transact(opts, "yank", id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_Mcdflop *McdflopSession) Yank(id *big.Int) (*types.Transaction, error) {
	return _Mcdflop.Contract.Yank(&_Mcdflop.TransactOpts, id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_Mcdflop *McdflopTransactorSession) Yank(id *big.Int) (*types.Transaction, error) {
	return _Mcdflop.Contract.Yank(&_Mcdflop.TransactOpts, id)
}

// McdflopKickIterator is returned from FilterKick and is used to iterate over the raw logs and unpacked data for Kick events raised by the Mcdflop contract.
type McdflopKickIterator struct {
	Event *McdflopKick // Event containing the contract specifics and raw log

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
func (it *McdflopKickIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(McdflopKick)
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
		it.Event = new(McdflopKick)
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
func (it *McdflopKickIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *McdflopKickIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// McdflopKick represents a Kick event raised by the Mcdflop contract.
type McdflopKick struct {
	Id  *big.Int
	Lot *big.Int
	Bid *big.Int
	Gal common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKick is a free log retrieval operation binding the contract event 0x7e8881001566f9f89aedb9c5dc3d856a2b81e5235a8196413ed484be91cc0df6.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid, address indexed gal)
func (_Mcdflop *McdflopFilterer) FilterKick(opts *bind.FilterOpts, gal []common.Address) (*McdflopKickIterator, error) {

	var galRule []interface{}
	for _, galItem := range gal {
		galRule = append(galRule, galItem)
	}

	logs, sub, err := _Mcdflop.contract.FilterLogs(opts, "Kick", galRule)
	if err != nil {
		return nil, err
	}
	return &McdflopKickIterator{contract: _Mcdflop.contract, event: "Kick", logs: logs, sub: sub}, nil
}

// WatchKick is a free log subscription operation binding the contract event 0x7e8881001566f9f89aedb9c5dc3d856a2b81e5235a8196413ed484be91cc0df6.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid, address indexed gal)
func (_Mcdflop *McdflopFilterer) WatchKick(opts *bind.WatchOpts, sink chan<- *McdflopKick, gal []common.Address) (event.Subscription, error) {

	var galRule []interface{}
	for _, galItem := range gal {
		galRule = append(galRule, galItem)
	}

	logs, sub, err := _Mcdflop.contract.WatchLogs(opts, "Kick", galRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(McdflopKick)
				if err := _Mcdflop.contract.UnpackLog(event, "Kick", log); err != nil {
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

// ParseKick is a log parse operation binding the contract event 0x7e8881001566f9f89aedb9c5dc3d856a2b81e5235a8196413ed484be91cc0df6.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid, address indexed gal)
func (_Mcdflop *McdflopFilterer) ParseKick(log types.Log) (*McdflopKick, error) {
	event := new(McdflopKick)
	if err := _Mcdflop.contract.UnpackLog(event, "Kick", log); err != nil {
		return nil, err
	}
	return event, nil
}
