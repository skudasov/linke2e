// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockoracle

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

// MockOracleABI is the input ABI used to generate the binding from.
const MockOracleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"CancelOracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"specId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callbackAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cancelExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dataVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EXPIRY_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"_expiration\",\"type\":\"uint256\"}],\"name\":\"cancelOracleRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"fulfillOracleRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainlinkToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_specId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dataVersion\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"oracleRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MockOracleBin is the compiled bytecode used for deploying new contracts.
var MockOracleBin = "0x608060405234801561001057600080fd5b506040516113ba3803806113ba8339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050611326806100946000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063165d35e1146100675780631f8f238c146100b157806340429946146101015780634b6022821461020c5780636ee4d5531461022a578063a4c0ed3614610295575b600080fd5b61006f61037a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100e7600480360360408110156100c757600080fd5b8101908080359060200190929190803590602001909291905050506103a3565b604051808215151515815260200191505060405180910390f35b61020a600480360361010081101561011857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690602001909291908035906020019092919080359060200190929190803590602001906401000000008111156101c657600080fd5b8201836020820111156101d857600080fd5b803590602001918460018302840111640100000000831117156101fa57600080fd5b9091929391929390505050610761565b005b610214610c2b565b6040518082815260200191505060405180910390f35b6102936004803603608081101561024057600080fd5b81019080803590602001909291908035906020019092919080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916906020019092919080359060200190929190505050610c31565b005b610378600480360360608110156102ab57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156102f257600080fd5b82018360208201111561030457600080fd5b8035906020019184600183028401116401000000008311171561032657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610ef0565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600082600073ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561047f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d757374206861766520612076616c696420726571756573744964000000000081525060200191505060405180910390fd5b6104876112a1565b600160008681526020019081526020016000206040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681525050905060016000868152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549063ffffffff0219169055505062061a805a101561061f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4d7573742070726f7669646520636f6e73756d657220656e6f7567682067617381525060200191505060405180910390fd5b6000816000015173ffffffffffffffffffffffffffffffffffffffff16826020015187876040516024018083815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040518082805190602001908083835b602083106106ea57805182526020820191506020810190506020830392506106c7565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461074c576040519150601f19603f3d011682016040523d82523d6000602084013e610751565b606091505b5050905080935050505092915050565b61076961037a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610809576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f4d75737420757365204c494e4b20746f6b656e0000000000000000000000000081525060200191505060405180910390fd5b856000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156108cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f43616e6e6f742063616c6c6261636b20746f204c494e4b00000000000000000081525060200191505060405180910390fd5b60008a86604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140182815260200192505050604051602081830303815290604052805190602001209050600073ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a07576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f4d75737420757365206120756e6971756520494400000000000000000000000081525060200191505060405180910390fd5b6000610a1e6104b04261121990919063ffffffff16565b905060405180604001604052808a73ffffffffffffffffffffffffffffffffffffffff168152602001897bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152506001600084815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548163ffffffff021916908360e01c0217905550905050897fd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c658d848e8d8d878d8d8d604051808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018881526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001867bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001858152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f8201169050808301925050509a505050505050505050505060405180910390a2505050505050505050505050565b6104b081565b600073ffffffffffffffffffffffffffffffffffffffff166001600086815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610d0a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f4d75737420757365206120756e6971756520494400000000000000000000000081525060200191505060405180910390fd5b42811115610d80576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f52657175657374206973206e6f7420657870697265640000000000000000000081525060200191505060405180910390fd5b60016000858152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549063ffffffff02191690555050837fa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e9360405160405180910390a26000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610ea957600080fd5b505af1158015610ebd573d6000803e3d6000fd5b505050506040513d6020811015610ed357600080fd5b8101908080519060200190929190505050610eea57fe5b50505050565b610ef861037a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f98576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f4d75737420757365204c494e4b20746f6b656e0000000000000000000000000081525060200191505060405180910390fd5b80600260200260040181511015611017576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f496e76616c69642072657175657374206c656e6774680000000000000000000081525060200191505060405180910390fd5b81600060208201519050634042994660e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146110db576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f4d757374207573652077686974656c69737465642066756e6374696f6e73000081525060200191505060405180910390fd5b85602485015284604485015260003073ffffffffffffffffffffffffffffffffffffffff16856040518082805190602001908083835b602083106111345780518252602082019150602081019050602083039250611111565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114611194576040519150601f19603f3d011682016040523d82523d6000602084013e611199565b606091505b5050905080611210576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f556e61626c6520746f206372656174652072657175657374000000000000000081525060200191505060405180910390fd5b50505050505050565b600080828401905083811015611297576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152509056fea26469706673582212205971fa613949da60cc06b6f0072240f7e136f60ad1797a9b44c9e66eb017ae8764736f6c63430006060033"

// DeployMockOracle deploys a new Ethereum contract, binding an instance of MockOracle to it.
func DeployMockOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address) (common.Address, *types.Transaction, *MockOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(MockOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MockOracleBin), backend, _link)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockOracle{MockOracleCaller: MockOracleCaller{contract: contract}, MockOracleTransactor: MockOracleTransactor{contract: contract}, MockOracleFilterer: MockOracleFilterer{contract: contract}}, nil
}

// MockOracle is an auto generated Go binding around an Ethereum contract.
type MockOracle struct {
	MockOracleCaller     // Read-only binding to the contract
	MockOracleTransactor // Write-only binding to the contract
	MockOracleFilterer   // Log filterer for contract events
}

// MockOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockOracleSession struct {
	Contract     *MockOracle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockOracleCallerSession struct {
	Contract *MockOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MockOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockOracleTransactorSession struct {
	Contract     *MockOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MockOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockOracleRaw struct {
	Contract *MockOracle // Generic contract binding to access the raw methods on
}

// MockOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockOracleCallerRaw struct {
	Contract *MockOracleCaller // Generic read-only contract binding to access the raw methods on
}

// MockOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockOracleTransactorRaw struct {
	Contract *MockOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockOracle creates a new instance of MockOracle, bound to a specific deployed contract.
func NewMockOracle(address common.Address, backend bind.ContractBackend) (*MockOracle, error) {
	contract, err := bindMockOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockOracle{MockOracleCaller: MockOracleCaller{contract: contract}, MockOracleTransactor: MockOracleTransactor{contract: contract}, MockOracleFilterer: MockOracleFilterer{contract: contract}}, nil
}

// NewMockOracleCaller creates a new read-only instance of MockOracle, bound to a specific deployed contract.
func NewMockOracleCaller(address common.Address, caller bind.ContractCaller) (*MockOracleCaller, error) {
	contract, err := bindMockOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockOracleCaller{contract: contract}, nil
}

// NewMockOracleTransactor creates a new write-only instance of MockOracle, bound to a specific deployed contract.
func NewMockOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*MockOracleTransactor, error) {
	contract, err := bindMockOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockOracleTransactor{contract: contract}, nil
}

// NewMockOracleFilterer creates a new log filterer instance of MockOracle, bound to a specific deployed contract.
func NewMockOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*MockOracleFilterer, error) {
	contract, err := bindMockOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockOracleFilterer{contract: contract}, nil
}

// bindMockOracle binds a generic wrapper to an already deployed contract.
func bindMockOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockOracle *MockOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockOracle.Contract.MockOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockOracle *MockOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOracle.Contract.MockOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockOracle *MockOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockOracle.Contract.MockOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockOracle *MockOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockOracle *MockOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockOracle *MockOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockOracle.Contract.contract.Transact(opts, method, params...)
}

// EXPIRYTIME is a free data retrieval call binding the contract method 0x4b602282.
//
// Solidity: function EXPIRY_TIME() view returns(uint256)
func (_MockOracle *MockOracleCaller) EXPIRYTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockOracle.contract.Call(opts, &out, "EXPIRY_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXPIRYTIME is a free data retrieval call binding the contract method 0x4b602282.
//
// Solidity: function EXPIRY_TIME() view returns(uint256)
func (_MockOracle *MockOracleSession) EXPIRYTIME() (*big.Int, error) {
	return _MockOracle.Contract.EXPIRYTIME(&_MockOracle.CallOpts)
}

// EXPIRYTIME is a free data retrieval call binding the contract method 0x4b602282.
//
// Solidity: function EXPIRY_TIME() view returns(uint256)
func (_MockOracle *MockOracleCallerSession) EXPIRYTIME() (*big.Int, error) {
	return _MockOracle.Contract.EXPIRYTIME(&_MockOracle.CallOpts)
}

// GetChainlinkToken is a free data retrieval call binding the contract method 0x165d35e1.
//
// Solidity: function getChainlinkToken() view returns(address)
func (_MockOracle *MockOracleCaller) GetChainlinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOracle.contract.Call(opts, &out, "getChainlinkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChainlinkToken is a free data retrieval call binding the contract method 0x165d35e1.
//
// Solidity: function getChainlinkToken() view returns(address)
func (_MockOracle *MockOracleSession) GetChainlinkToken() (common.Address, error) {
	return _MockOracle.Contract.GetChainlinkToken(&_MockOracle.CallOpts)
}

// GetChainlinkToken is a free data retrieval call binding the contract method 0x165d35e1.
//
// Solidity: function getChainlinkToken() view returns(address)
func (_MockOracle *MockOracleCallerSession) GetChainlinkToken() (common.Address, error) {
	return _MockOracle.Contract.GetChainlinkToken(&_MockOracle.CallOpts)
}

// CancelOracleRequest is a paid mutator transaction binding the contract method 0x6ee4d553.
//
// Solidity: function cancelOracleRequest(bytes32 _requestId, uint256 _payment, bytes4 , uint256 _expiration) returns()
func (_MockOracle *MockOracleTransactor) CancelOracleRequest(opts *bind.TransactOpts, _requestId [32]byte, _payment *big.Int, arg2 [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _MockOracle.contract.Transact(opts, "cancelOracleRequest", _requestId, _payment, arg2, _expiration)
}

// CancelOracleRequest is a paid mutator transaction binding the contract method 0x6ee4d553.
//
// Solidity: function cancelOracleRequest(bytes32 _requestId, uint256 _payment, bytes4 , uint256 _expiration) returns()
func (_MockOracle *MockOracleSession) CancelOracleRequest(_requestId [32]byte, _payment *big.Int, arg2 [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _MockOracle.Contract.CancelOracleRequest(&_MockOracle.TransactOpts, _requestId, _payment, arg2, _expiration)
}

// CancelOracleRequest is a paid mutator transaction binding the contract method 0x6ee4d553.
//
// Solidity: function cancelOracleRequest(bytes32 _requestId, uint256 _payment, bytes4 , uint256 _expiration) returns()
func (_MockOracle *MockOracleTransactorSession) CancelOracleRequest(_requestId [32]byte, _payment *big.Int, arg2 [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _MockOracle.Contract.CancelOracleRequest(&_MockOracle.TransactOpts, _requestId, _payment, arg2, _expiration)
}

// FulfillOracleRequest is a paid mutator transaction binding the contract method 0x1f8f238c.
//
// Solidity: function fulfillOracleRequest(bytes32 _requestId, bytes32 _data) returns(bool)
func (_MockOracle *MockOracleTransactor) FulfillOracleRequest(opts *bind.TransactOpts, _requestId [32]byte, _data [32]byte) (*types.Transaction, error) {
	return _MockOracle.contract.Transact(opts, "fulfillOracleRequest", _requestId, _data)
}

// FulfillOracleRequest is a paid mutator transaction binding the contract method 0x1f8f238c.
//
// Solidity: function fulfillOracleRequest(bytes32 _requestId, bytes32 _data) returns(bool)
func (_MockOracle *MockOracleSession) FulfillOracleRequest(_requestId [32]byte, _data [32]byte) (*types.Transaction, error) {
	return _MockOracle.Contract.FulfillOracleRequest(&_MockOracle.TransactOpts, _requestId, _data)
}

// FulfillOracleRequest is a paid mutator transaction binding the contract method 0x1f8f238c.
//
// Solidity: function fulfillOracleRequest(bytes32 _requestId, bytes32 _data) returns(bool)
func (_MockOracle *MockOracleTransactorSession) FulfillOracleRequest(_requestId [32]byte, _data [32]byte) (*types.Transaction, error) {
	return _MockOracle.Contract.FulfillOracleRequest(&_MockOracle.TransactOpts, _requestId, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _amount, bytes _data) returns()
func (_MockOracle *MockOracleTransactor) OnTokenTransfer(opts *bind.TransactOpts, _sender common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockOracle.contract.Transact(opts, "onTokenTransfer", _sender, _amount, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _amount, bytes _data) returns()
func (_MockOracle *MockOracleSession) OnTokenTransfer(_sender common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockOracle.Contract.OnTokenTransfer(&_MockOracle.TransactOpts, _sender, _amount, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _amount, bytes _data) returns()
func (_MockOracle *MockOracleTransactorSession) OnTokenTransfer(_sender common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockOracle.Contract.OnTokenTransfer(&_MockOracle.TransactOpts, _sender, _amount, _data)
}

// OracleRequest is a paid mutator transaction binding the contract method 0x40429946.
//
// Solidity: function oracleRequest(address _sender, uint256 _payment, bytes32 _specId, address _callbackAddress, bytes4 _callbackFunctionId, uint256 _nonce, uint256 _dataVersion, bytes _data) returns()
func (_MockOracle *MockOracleTransactor) OracleRequest(opts *bind.TransactOpts, _sender common.Address, _payment *big.Int, _specId [32]byte, _callbackAddress common.Address, _callbackFunctionId [4]byte, _nonce *big.Int, _dataVersion *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockOracle.contract.Transact(opts, "oracleRequest", _sender, _payment, _specId, _callbackAddress, _callbackFunctionId, _nonce, _dataVersion, _data)
}

// OracleRequest is a paid mutator transaction binding the contract method 0x40429946.
//
// Solidity: function oracleRequest(address _sender, uint256 _payment, bytes32 _specId, address _callbackAddress, bytes4 _callbackFunctionId, uint256 _nonce, uint256 _dataVersion, bytes _data) returns()
func (_MockOracle *MockOracleSession) OracleRequest(_sender common.Address, _payment *big.Int, _specId [32]byte, _callbackAddress common.Address, _callbackFunctionId [4]byte, _nonce *big.Int, _dataVersion *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockOracle.Contract.OracleRequest(&_MockOracle.TransactOpts, _sender, _payment, _specId, _callbackAddress, _callbackFunctionId, _nonce, _dataVersion, _data)
}

// OracleRequest is a paid mutator transaction binding the contract method 0x40429946.
//
// Solidity: function oracleRequest(address _sender, uint256 _payment, bytes32 _specId, address _callbackAddress, bytes4 _callbackFunctionId, uint256 _nonce, uint256 _dataVersion, bytes _data) returns()
func (_MockOracle *MockOracleTransactorSession) OracleRequest(_sender common.Address, _payment *big.Int, _specId [32]byte, _callbackAddress common.Address, _callbackFunctionId [4]byte, _nonce *big.Int, _dataVersion *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockOracle.Contract.OracleRequest(&_MockOracle.TransactOpts, _sender, _payment, _specId, _callbackAddress, _callbackFunctionId, _nonce, _dataVersion, _data)
}

// MockOracleCancelOracleRequestIterator is returned from FilterCancelOracleRequest and is used to iterate over the raw logs and unpacked data for CancelOracleRequest events raised by the MockOracle contract.
type MockOracleCancelOracleRequestIterator struct {
	Event *MockOracleCancelOracleRequest // Event containing the contract specifics and raw log

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
func (it *MockOracleCancelOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockOracleCancelOracleRequest)
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
		it.Event = new(MockOracleCancelOracleRequest)
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
func (it *MockOracleCancelOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockOracleCancelOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockOracleCancelOracleRequest represents a CancelOracleRequest event raised by the MockOracle contract.
type MockOracleCancelOracleRequest struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCancelOracleRequest is a free log retrieval operation binding the contract event 0xa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e93.
//
// Solidity: event CancelOracleRequest(bytes32 indexed requestId)
func (_MockOracle *MockOracleFilterer) FilterCancelOracleRequest(opts *bind.FilterOpts, requestId [][32]byte) (*MockOracleCancelOracleRequestIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _MockOracle.contract.FilterLogs(opts, "CancelOracleRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &MockOracleCancelOracleRequestIterator{contract: _MockOracle.contract, event: "CancelOracleRequest", logs: logs, sub: sub}, nil
}

// WatchCancelOracleRequest is a free log subscription operation binding the contract event 0xa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e93.
//
// Solidity: event CancelOracleRequest(bytes32 indexed requestId)
func (_MockOracle *MockOracleFilterer) WatchCancelOracleRequest(opts *bind.WatchOpts, sink chan<- *MockOracleCancelOracleRequest, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _MockOracle.contract.WatchLogs(opts, "CancelOracleRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockOracleCancelOracleRequest)
				if err := _MockOracle.contract.UnpackLog(event, "CancelOracleRequest", log); err != nil {
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

// ParseCancelOracleRequest is a log parse operation binding the contract event 0xa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e93.
//
// Solidity: event CancelOracleRequest(bytes32 indexed requestId)
func (_MockOracle *MockOracleFilterer) ParseCancelOracleRequest(log types.Log) (*MockOracleCancelOracleRequest, error) {
	event := new(MockOracleCancelOracleRequest)
	if err := _MockOracle.contract.UnpackLog(event, "CancelOracleRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockOracleOracleRequestIterator is returned from FilterOracleRequest and is used to iterate over the raw logs and unpacked data for OracleRequest events raised by the MockOracle contract.
type MockOracleOracleRequestIterator struct {
	Event *MockOracleOracleRequest // Event containing the contract specifics and raw log

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
func (it *MockOracleOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockOracleOracleRequest)
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
		it.Event = new(MockOracleOracleRequest)
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
func (it *MockOracleOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockOracleOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockOracleOracleRequest represents a OracleRequest event raised by the MockOracle contract.
type MockOracleOracleRequest struct {
	SpecId             [32]byte
	Requester          common.Address
	RequestId          [32]byte
	Payment            *big.Int
	CallbackAddr       common.Address
	CallbackFunctionId [4]byte
	CancelExpiration   *big.Int
	DataVersion        *big.Int
	Data               []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOracleRequest is a free log retrieval operation binding the contract event 0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65.
//
// Solidity: event OracleRequest(bytes32 indexed specId, address requester, bytes32 requestId, uint256 payment, address callbackAddr, bytes4 callbackFunctionId, uint256 cancelExpiration, uint256 dataVersion, bytes data)
func (_MockOracle *MockOracleFilterer) FilterOracleRequest(opts *bind.FilterOpts, specId [][32]byte) (*MockOracleOracleRequestIterator, error) {

	var specIdRule []interface{}
	for _, specIdItem := range specId {
		specIdRule = append(specIdRule, specIdItem)
	}

	logs, sub, err := _MockOracle.contract.FilterLogs(opts, "OracleRequest", specIdRule)
	if err != nil {
		return nil, err
	}
	return &MockOracleOracleRequestIterator{contract: _MockOracle.contract, event: "OracleRequest", logs: logs, sub: sub}, nil
}

// WatchOracleRequest is a free log subscription operation binding the contract event 0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65.
//
// Solidity: event OracleRequest(bytes32 indexed specId, address requester, bytes32 requestId, uint256 payment, address callbackAddr, bytes4 callbackFunctionId, uint256 cancelExpiration, uint256 dataVersion, bytes data)
func (_MockOracle *MockOracleFilterer) WatchOracleRequest(opts *bind.WatchOpts, sink chan<- *MockOracleOracleRequest, specId [][32]byte) (event.Subscription, error) {

	var specIdRule []interface{}
	for _, specIdItem := range specId {
		specIdRule = append(specIdRule, specIdItem)
	}

	logs, sub, err := _MockOracle.contract.WatchLogs(opts, "OracleRequest", specIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockOracleOracleRequest)
				if err := _MockOracle.contract.UnpackLog(event, "OracleRequest", log); err != nil {
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

// ParseOracleRequest is a log parse operation binding the contract event 0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65.
//
// Solidity: event OracleRequest(bytes32 indexed specId, address requester, bytes32 requestId, uint256 payment, address callbackAddr, bytes4 callbackFunctionId, uint256 cancelExpiration, uint256 dataVersion, bytes data)
func (_MockOracle *MockOracleFilterer) ParseOracleRequest(log types.Log) (*MockOracleOracleRequest, error) {
	event := new(MockOracleOracleRequest)
	if err := _MockOracle.contract.UnpackLog(event, "OracleRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
