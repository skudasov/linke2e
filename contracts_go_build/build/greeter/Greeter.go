// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package greeter

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GreeterABI is the input ABI used to generate the binding from.
const GreeterABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_greeting\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"greet\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_greeting\",\"type\":\"string\"}],\"name\":\"setGreeting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GreeterBin is the compiled bytecode used for deploying new contracts.
var GreeterBin = "0x608060405234801561001057600080fd5b506040516104bf3803806104bf8339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186600182028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019080838360005b838110156100ba57808201518184015260208101905061009f565b50505050905090810190601f1680156100e75780820380516001836020036101000a031916815260200191505b50604052505050806000908051906020019061010492919061010b565b50506101a8565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014c57805160ff191683800117855561017a565b8280016001018555821561017a579182015b8281111561017957825182559160200191906001019061015e565b5b509050610187919061018b565b5090565b5b808211156101a457600081600090555060010161018c565b5090565b610308806101b76000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063a41368621461003b578063cfae3217146100f6575b600080fd5b6100f46004803603602081101561005157600080fd5b810190808035906020019064010000000081111561006e57600080fd5b82018360208201111561008057600080fd5b803590602001918460018302840111640100000000831117156100a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610179565b005b6100fe610193565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561013e578082015181840152602081019050610123565b50505050905090810190601f16801561016b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b806000908051906020019061018f929190610235565b5050565b606060008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561022b5780601f106102005761010080835404028352916020019161022b565b820191906000526020600020905b81548152906001019060200180831161020e57829003601f168201915b5050505050905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061027657805160ff19168380011785556102a4565b828001600101855582156102a4579182015b828111156102a3578251825591602001919060010190610288565b5b5090506102b191906102b5565b5090565b5b808211156102ce5760008160009055506001016102b6565b509056fea264697066735822122001d1cb22f9c63d2f3c17b03dda40178911f3e6dd366e3b24826395d38cd2326c64736f6c63430007000033"

// DeployGreeter deploys a new Ethereum contract, binding an instance of Greeter to it.
func DeployGreeter(auth *bind.TransactOpts, backend bind.ContractBackend, _greeting string) (common.Address, *types.Transaction, *Greeter, error) {
	parsed, err := abi.JSON(strings.NewReader(GreeterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GreeterBin), backend, _greeting)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Greeter{GreeterCaller: GreeterCaller{contract: contract}, GreeterTransactor: GreeterTransactor{contract: contract}, GreeterFilterer: GreeterFilterer{contract: contract}}, nil
}

// Greeter is an auto generated Go binding around an Ethereum contract.
type Greeter struct {
	GreeterCaller     // Read-only binding to the contract
	GreeterTransactor // Write-only binding to the contract
	GreeterFilterer   // Log filterer for contract events
}

// GreeterCaller is an auto generated read-only Go binding around an Ethereum contract.
type GreeterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GreeterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GreeterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GreeterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GreeterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GreeterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GreeterSession struct {
	Contract     *Greeter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GreeterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GreeterCallerSession struct {
	Contract *GreeterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GreeterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GreeterTransactorSession struct {
	Contract     *GreeterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GreeterRaw is an auto generated low-level Go binding around an Ethereum contract.
type GreeterRaw struct {
	Contract *Greeter // Generic contract binding to access the raw methods on
}

// GreeterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GreeterCallerRaw struct {
	Contract *GreeterCaller // Generic read-only contract binding to access the raw methods on
}

// GreeterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GreeterTransactorRaw struct {
	Contract *GreeterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGreeter creates a new instance of Greeter, bound to a specific deployed contract.
func NewGreeter(address common.Address, backend bind.ContractBackend) (*Greeter, error) {
	contract, err := bindGreeter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Greeter{GreeterCaller: GreeterCaller{contract: contract}, GreeterTransactor: GreeterTransactor{contract: contract}, GreeterFilterer: GreeterFilterer{contract: contract}}, nil
}

// NewGreeterCaller creates a new read-only instance of Greeter, bound to a specific deployed contract.
func NewGreeterCaller(address common.Address, caller bind.ContractCaller) (*GreeterCaller, error) {
	contract, err := bindGreeter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GreeterCaller{contract: contract}, nil
}

// NewGreeterTransactor creates a new write-only instance of Greeter, bound to a specific deployed contract.
func NewGreeterTransactor(address common.Address, transactor bind.ContractTransactor) (*GreeterTransactor, error) {
	contract, err := bindGreeter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GreeterTransactor{contract: contract}, nil
}

// NewGreeterFilterer creates a new log filterer instance of Greeter, bound to a specific deployed contract.
func NewGreeterFilterer(address common.Address, filterer bind.ContractFilterer) (*GreeterFilterer, error) {
	contract, err := bindGreeter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GreeterFilterer{contract: contract}, nil
}

// bindGreeter binds a generic wrapper to an already deployed contract.
func bindGreeter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GreeterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Greeter *GreeterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Greeter.Contract.GreeterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Greeter *GreeterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Greeter.Contract.GreeterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Greeter *GreeterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Greeter.Contract.GreeterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Greeter *GreeterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Greeter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Greeter *GreeterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Greeter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Greeter *GreeterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Greeter.Contract.contract.Transact(opts, method, params...)
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() view returns(string)
func (_Greeter *GreeterCaller) Greet(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Greeter.contract.Call(opts, &out, "greet")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() view returns(string)
func (_Greeter *GreeterSession) Greet() (string, error) {
	return _Greeter.Contract.Greet(&_Greeter.CallOpts)
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() view returns(string)
func (_Greeter *GreeterCallerSession) Greet() (string, error) {
	return _Greeter.Contract.Greet(&_Greeter.CallOpts)
}

// SetGreeting is a paid mutator transaction binding the contract method 0xa4136862.
//
// Solidity: function setGreeting(string _greeting) returns()
func (_Greeter *GreeterTransactor) SetGreeting(opts *bind.TransactOpts, _greeting string) (*types.Transaction, error) {
	return _Greeter.contract.Transact(opts, "setGreeting", _greeting)
}

// SetGreeting is a paid mutator transaction binding the contract method 0xa4136862.
//
// Solidity: function setGreeting(string _greeting) returns()
func (_Greeter *GreeterSession) SetGreeting(_greeting string) (*types.Transaction, error) {
	return _Greeter.Contract.SetGreeting(&_Greeter.TransactOpts, _greeting)
}

// SetGreeting is a paid mutator transaction binding the contract method 0xa4136862.
//
// Solidity: function setGreeting(string _greeting) returns()
func (_Greeter *GreeterTransactorSession) SetGreeting(_greeting string) (*types.Transaction, error) {
	return _Greeter.Contract.SetGreeting(&_Greeter.TransactOpts, _greeting)
}
