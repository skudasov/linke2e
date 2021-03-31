// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package consumer

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

// APIConsumerABI is the input ABI used to generate the binding from.
const APIConsumerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"_callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"_expiration\",\"type\":\"uint256\"}],\"name\":\"cancelRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_path\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_times\",\"type\":\"int256\"}],\"name\":\"createRequestTo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_data\",\"type\":\"uint256\"}],\"name\":\"fulfill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainlinkToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// APIConsumer is an auto generated Go binding around an Ethereum contract.
type APIConsumer struct {
	APIConsumerCaller     // Read-only binding to the contract
	APIConsumerTransactor // Write-only binding to the contract
	APIConsumerFilterer   // Log filterer for contract events
}

// APIConsumerCaller is an auto generated read-only Go binding around an Ethereum contract.
type APIConsumerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APIConsumerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type APIConsumerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APIConsumerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type APIConsumerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APIConsumerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type APIConsumerSession struct {
	Contract     *APIConsumer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// APIConsumerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type APIConsumerCallerSession struct {
	Contract *APIConsumerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// APIConsumerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type APIConsumerTransactorSession struct {
	Contract     *APIConsumerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// APIConsumerRaw is an auto generated low-level Go binding around an Ethereum contract.
type APIConsumerRaw struct {
	Contract *APIConsumer // Generic contract binding to access the raw methods on
}

// APIConsumerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type APIConsumerCallerRaw struct {
	Contract *APIConsumerCaller // Generic read-only contract binding to access the raw methods on
}

// APIConsumerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type APIConsumerTransactorRaw struct {
	Contract *APIConsumerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAPIConsumer creates a new instance of APIConsumer, bound to a specific deployed contract.
func NewAPIConsumer(address common.Address, backend bind.ContractBackend) (*APIConsumer, error) {
	contract, err := bindAPIConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &APIConsumer{APIConsumerCaller: APIConsumerCaller{contract: contract}, APIConsumerTransactor: APIConsumerTransactor{contract: contract}, APIConsumerFilterer: APIConsumerFilterer{contract: contract}}, nil
}

// NewAPIConsumerCaller creates a new read-only instance of APIConsumer, bound to a specific deployed contract.
func NewAPIConsumerCaller(address common.Address, caller bind.ContractCaller) (*APIConsumerCaller, error) {
	contract, err := bindAPIConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &APIConsumerCaller{contract: contract}, nil
}

// NewAPIConsumerTransactor creates a new write-only instance of APIConsumer, bound to a specific deployed contract.
func NewAPIConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*APIConsumerTransactor, error) {
	contract, err := bindAPIConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &APIConsumerTransactor{contract: contract}, nil
}

// NewAPIConsumerFilterer creates a new log filterer instance of APIConsumer, bound to a specific deployed contract.
func NewAPIConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*APIConsumerFilterer, error) {
	contract, err := bindAPIConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &APIConsumerFilterer{contract: contract}, nil
}

// bindAPIConsumer binds a generic wrapper to an already deployed contract.
func bindAPIConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(APIConsumerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APIConsumer *APIConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APIConsumer.Contract.APIConsumerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APIConsumer *APIConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APIConsumer.Contract.APIConsumerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APIConsumer *APIConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APIConsumer.Contract.APIConsumerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APIConsumer *APIConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APIConsumer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APIConsumer *APIConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APIConsumer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APIConsumer *APIConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APIConsumer.Contract.contract.Transact(opts, method, params...)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(uint256)
func (_APIConsumer *APIConsumerCaller) Data(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APIConsumer.contract.Call(opts, &out, "data")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(uint256)
func (_APIConsumer *APIConsumerSession) Data() (*big.Int, error) {
	return _APIConsumer.Contract.Data(&_APIConsumer.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(uint256)
func (_APIConsumer *APIConsumerCallerSession) Data() (*big.Int, error) {
	return _APIConsumer.Contract.Data(&_APIConsumer.CallOpts)
}

// GetChainlinkToken is a free data retrieval call binding the contract method 0x165d35e1.
//
// Solidity: function getChainlinkToken() view returns(address)
func (_APIConsumer *APIConsumerCaller) GetChainlinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _APIConsumer.contract.Call(opts, &out, "getChainlinkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChainlinkToken is a free data retrieval call binding the contract method 0x165d35e1.
//
// Solidity: function getChainlinkToken() view returns(address)
func (_APIConsumer *APIConsumerSession) GetChainlinkToken() (common.Address, error) {
	return _APIConsumer.Contract.GetChainlinkToken(&_APIConsumer.CallOpts)
}

// GetChainlinkToken is a free data retrieval call binding the contract method 0x165d35e1.
//
// Solidity: function getChainlinkToken() view returns(address)
func (_APIConsumer *APIConsumerCallerSession) GetChainlinkToken() (common.Address, error) {
	return _APIConsumer.Contract.GetChainlinkToken(&_APIConsumer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_APIConsumer *APIConsumerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _APIConsumer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_APIConsumer *APIConsumerSession) Owner() (common.Address, error) {
	return _APIConsumer.Contract.Owner(&_APIConsumer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_APIConsumer *APIConsumerCallerSession) Owner() (common.Address, error) {
	return _APIConsumer.Contract.Owner(&_APIConsumer.CallOpts)
}

// CancelRequest is a paid mutator transaction binding the contract method 0xec65d0f8.
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, bytes4 _callbackFunctionId, uint256 _expiration) returns()
func (_APIConsumer *APIConsumerTransactor) CancelRequest(opts *bind.TransactOpts, _requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _APIConsumer.contract.Transact(opts, "cancelRequest", _requestId, _payment, _callbackFunctionId, _expiration)
}

// CancelRequest is a paid mutator transaction binding the contract method 0xec65d0f8.
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, bytes4 _callbackFunctionId, uint256 _expiration) returns()
func (_APIConsumer *APIConsumerSession) CancelRequest(_requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _APIConsumer.Contract.CancelRequest(&_APIConsumer.TransactOpts, _requestId, _payment, _callbackFunctionId, _expiration)
}

// CancelRequest is a paid mutator transaction binding the contract method 0xec65d0f8.
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, bytes4 _callbackFunctionId, uint256 _expiration) returns()
func (_APIConsumer *APIConsumerTransactorSession) CancelRequest(_requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _APIConsumer.Contract.CancelRequest(&_APIConsumer.TransactOpts, _requestId, _payment, _callbackFunctionId, _expiration)
}

// CreateRequestTo is a paid mutator transaction binding the contract method 0x16ef7f1a.
//
// Solidity: function createRequestTo(address _oracle, bytes32 _jobId, uint256 _payment, string _url, string _path, int256 _times) returns(bytes32 requestId)
func (_APIConsumer *APIConsumerTransactor) CreateRequestTo(opts *bind.TransactOpts, _oracle common.Address, _jobId [32]byte, _payment *big.Int, _url string, _path string, _times *big.Int) (*types.Transaction, error) {
	return _APIConsumer.contract.Transact(opts, "createRequestTo", _oracle, _jobId, _payment, _url, _path, _times)
}

// CreateRequestTo is a paid mutator transaction binding the contract method 0x16ef7f1a.
//
// Solidity: function createRequestTo(address _oracle, bytes32 _jobId, uint256 _payment, string _url, string _path, int256 _times) returns(bytes32 requestId)
func (_APIConsumer *APIConsumerSession) CreateRequestTo(_oracle common.Address, _jobId [32]byte, _payment *big.Int, _url string, _path string, _times *big.Int) (*types.Transaction, error) {
	return _APIConsumer.Contract.CreateRequestTo(&_APIConsumer.TransactOpts, _oracle, _jobId, _payment, _url, _path, _times)
}

// CreateRequestTo is a paid mutator transaction binding the contract method 0x16ef7f1a.
//
// Solidity: function createRequestTo(address _oracle, bytes32 _jobId, uint256 _payment, string _url, string _path, int256 _times) returns(bytes32 requestId)
func (_APIConsumer *APIConsumerTransactorSession) CreateRequestTo(_oracle common.Address, _jobId [32]byte, _payment *big.Int, _url string, _path string, _times *big.Int) (*types.Transaction, error) {
	return _APIConsumer.Contract.CreateRequestTo(&_APIConsumer.TransactOpts, _oracle, _jobId, _payment, _url, _path, _times)
}

// Fulfill is a paid mutator transaction binding the contract method 0x4357855e.
//
// Solidity: function fulfill(bytes32 _requestId, uint256 _data) returns()
func (_APIConsumer *APIConsumerTransactor) Fulfill(opts *bind.TransactOpts, _requestId [32]byte, _data *big.Int) (*types.Transaction, error) {
	return _APIConsumer.contract.Transact(opts, "fulfill", _requestId, _data)
}

// Fulfill is a paid mutator transaction binding the contract method 0x4357855e.
//
// Solidity: function fulfill(bytes32 _requestId, uint256 _data) returns()
func (_APIConsumer *APIConsumerSession) Fulfill(_requestId [32]byte, _data *big.Int) (*types.Transaction, error) {
	return _APIConsumer.Contract.Fulfill(&_APIConsumer.TransactOpts, _requestId, _data)
}

// Fulfill is a paid mutator transaction binding the contract method 0x4357855e.
//
// Solidity: function fulfill(bytes32 _requestId, uint256 _data) returns()
func (_APIConsumer *APIConsumerTransactorSession) Fulfill(_requestId [32]byte, _data *big.Int) (*types.Transaction, error) {
	return _APIConsumer.Contract.Fulfill(&_APIConsumer.TransactOpts, _requestId, _data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_APIConsumer *APIConsumerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APIConsumer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_APIConsumer *APIConsumerSession) RenounceOwnership() (*types.Transaction, error) {
	return _APIConsumer.Contract.RenounceOwnership(&_APIConsumer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_APIConsumer *APIConsumerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _APIConsumer.Contract.RenounceOwnership(&_APIConsumer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_APIConsumer *APIConsumerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _APIConsumer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_APIConsumer *APIConsumerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _APIConsumer.Contract.TransferOwnership(&_APIConsumer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_APIConsumer *APIConsumerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _APIConsumer.Contract.TransferOwnership(&_APIConsumer.TransactOpts, newOwner)
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_APIConsumer *APIConsumerTransactor) WithdrawLink(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APIConsumer.contract.Transact(opts, "withdrawLink")
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_APIConsumer *APIConsumerSession) WithdrawLink() (*types.Transaction, error) {
	return _APIConsumer.Contract.WithdrawLink(&_APIConsumer.TransactOpts)
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_APIConsumer *APIConsumerTransactorSession) WithdrawLink() (*types.Transaction, error) {
	return _APIConsumer.Contract.WithdrawLink(&_APIConsumer.TransactOpts)
}

// APIConsumerChainlinkCancelledIterator is returned from FilterChainlinkCancelled and is used to iterate over the raw logs and unpacked data for ChainlinkCancelled events raised by the APIConsumer contract.
type APIConsumerChainlinkCancelledIterator struct {
	Event *APIConsumerChainlinkCancelled // Event containing the contract specifics and raw log

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
func (it *APIConsumerChainlinkCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APIConsumerChainlinkCancelled)
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
		it.Event = new(APIConsumerChainlinkCancelled)
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
func (it *APIConsumerChainlinkCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APIConsumerChainlinkCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APIConsumerChainlinkCancelled represents a ChainlinkCancelled event raised by the APIConsumer contract.
type APIConsumerChainlinkCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkCancelled is a free log retrieval operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_APIConsumer *APIConsumerFilterer) FilterChainlinkCancelled(opts *bind.FilterOpts, id [][32]byte) (*APIConsumerChainlinkCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _APIConsumer.contract.FilterLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &APIConsumerChainlinkCancelledIterator{contract: _APIConsumer.contract, event: "ChainlinkCancelled", logs: logs, sub: sub}, nil
}

// WatchChainlinkCancelled is a free log subscription operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_APIConsumer *APIConsumerFilterer) WatchChainlinkCancelled(opts *bind.WatchOpts, sink chan<- *APIConsumerChainlinkCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _APIConsumer.contract.WatchLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APIConsumerChainlinkCancelled)
				if err := _APIConsumer.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
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

// ParseChainlinkCancelled is a log parse operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_APIConsumer *APIConsumerFilterer) ParseChainlinkCancelled(log types.Log) (*APIConsumerChainlinkCancelled, error) {
	event := new(APIConsumerChainlinkCancelled)
	if err := _APIConsumer.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APIConsumerChainlinkFulfilledIterator is returned from FilterChainlinkFulfilled and is used to iterate over the raw logs and unpacked data for ChainlinkFulfilled events raised by the APIConsumer contract.
type APIConsumerChainlinkFulfilledIterator struct {
	Event *APIConsumerChainlinkFulfilled // Event containing the contract specifics and raw log

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
func (it *APIConsumerChainlinkFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APIConsumerChainlinkFulfilled)
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
		it.Event = new(APIConsumerChainlinkFulfilled)
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
func (it *APIConsumerChainlinkFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APIConsumerChainlinkFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APIConsumerChainlinkFulfilled represents a ChainlinkFulfilled event raised by the APIConsumer contract.
type APIConsumerChainlinkFulfilled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkFulfilled is a free log retrieval operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_APIConsumer *APIConsumerFilterer) FilterChainlinkFulfilled(opts *bind.FilterOpts, id [][32]byte) (*APIConsumerChainlinkFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _APIConsumer.contract.FilterLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &APIConsumerChainlinkFulfilledIterator{contract: _APIConsumer.contract, event: "ChainlinkFulfilled", logs: logs, sub: sub}, nil
}

// WatchChainlinkFulfilled is a free log subscription operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_APIConsumer *APIConsumerFilterer) WatchChainlinkFulfilled(opts *bind.WatchOpts, sink chan<- *APIConsumerChainlinkFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _APIConsumer.contract.WatchLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APIConsumerChainlinkFulfilled)
				if err := _APIConsumer.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
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

// ParseChainlinkFulfilled is a log parse operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_APIConsumer *APIConsumerFilterer) ParseChainlinkFulfilled(log types.Log) (*APIConsumerChainlinkFulfilled, error) {
	event := new(APIConsumerChainlinkFulfilled)
	if err := _APIConsumer.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APIConsumerChainlinkRequestedIterator is returned from FilterChainlinkRequested and is used to iterate over the raw logs and unpacked data for ChainlinkRequested events raised by the APIConsumer contract.
type APIConsumerChainlinkRequestedIterator struct {
	Event *APIConsumerChainlinkRequested // Event containing the contract specifics and raw log

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
func (it *APIConsumerChainlinkRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APIConsumerChainlinkRequested)
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
		it.Event = new(APIConsumerChainlinkRequested)
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
func (it *APIConsumerChainlinkRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APIConsumerChainlinkRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APIConsumerChainlinkRequested represents a ChainlinkRequested event raised by the APIConsumer contract.
type APIConsumerChainlinkRequested struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkRequested is a free log retrieval operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_APIConsumer *APIConsumerFilterer) FilterChainlinkRequested(opts *bind.FilterOpts, id [][32]byte) (*APIConsumerChainlinkRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _APIConsumer.contract.FilterLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return &APIConsumerChainlinkRequestedIterator{contract: _APIConsumer.contract, event: "ChainlinkRequested", logs: logs, sub: sub}, nil
}

// WatchChainlinkRequested is a free log subscription operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_APIConsumer *APIConsumerFilterer) WatchChainlinkRequested(opts *bind.WatchOpts, sink chan<- *APIConsumerChainlinkRequested, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _APIConsumer.contract.WatchLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APIConsumerChainlinkRequested)
				if err := _APIConsumer.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
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

// ParseChainlinkRequested is a log parse operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_APIConsumer *APIConsumerFilterer) ParseChainlinkRequested(log types.Log) (*APIConsumerChainlinkRequested, error) {
	event := new(APIConsumerChainlinkRequested)
	if err := _APIConsumer.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APIConsumerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the APIConsumer contract.
type APIConsumerOwnershipTransferredIterator struct {
	Event *APIConsumerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *APIConsumerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APIConsumerOwnershipTransferred)
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
		it.Event = new(APIConsumerOwnershipTransferred)
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
func (it *APIConsumerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APIConsumerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APIConsumerOwnershipTransferred represents a OwnershipTransferred event raised by the APIConsumer contract.
type APIConsumerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_APIConsumer *APIConsumerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*APIConsumerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _APIConsumer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &APIConsumerOwnershipTransferredIterator{contract: _APIConsumer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_APIConsumer *APIConsumerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *APIConsumerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _APIConsumer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APIConsumerOwnershipTransferred)
				if err := _APIConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_APIConsumer *APIConsumerFilterer) ParseOwnershipTransferred(log types.Log) (*APIConsumerOwnershipTransferred, error) {
	event := new(APIConsumerOwnershipTransferred)
	if err := _APIConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
