// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package APIConsumer

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
const APIConsumerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"_callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"_expiration\",\"type\":\"uint256\"}],\"name\":\"cancelRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_path\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_times\",\"type\":\"int256\"}],\"name\":\"createRequestTo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_data\",\"type\":\"uint256\"}],\"name\":\"fulfill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainlinkToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// APIConsumerBin is the compiled bytecode used for deploying new contracts.
var APIConsumerBin = "0x608060405260016004553480156200001657600080fd5b5060405162001f1938038062001f19833981810160405260208110156200003c57600080fd5b810190808051906020019092919050505060006200005f6200016360201b60201c565b905080600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156200014a57620001446200016b60201b60201c565b6200015c565b6200015b816200021560201b60201c565b5b5062000259565b600033905090565b6200021373c89bd4e1632d3a43cb03aaad5262cbe4038bc57173ffffffffffffffffffffffffffffffffffffffff166338cc48316040518163ffffffff1660e01b815260040160206040518083038186803b158015620001ca57600080fd5b505afa158015620001df573d6000803e3d6000fd5b505050506040513d6020811015620001f657600080fd5b81019080805190602001909291905050506200021560201b60201c565b565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611cb080620002696000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80638da5cb5b116100665780638da5cb5b146102f15780638dc654a21461033b578063ea3d508a14610345578063ec65d0f8146103a1578063f2fde38b1461040c5761009e565b8063165d35e1146100a357806316ef7f1a146100ed5780634357855e14610291578063715018a6146102c957806373d4a13a146102d3575b600080fd5b6100ab610450565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61027b600480360360c081101561010357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019064010000000081111561015457600080fd5b82018360208201111561016657600080fd5b8035906020019184600183028401116401000000008311171561018857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156101eb57600080fd5b8201836020820111156101fd57600080fd5b8035906020019184600183028401116401000000008311171561021f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019092919050505061045f565b6040518082815260200191505060405180910390f35b6102c7600480360360408110156102a757600080fd5b810190808035906020019092919080359060200190929190505050610647565b005b6102d161076e565b005b6102db6108de565b6040518082815260200191505060405180910390f35b6102f96108e4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61034361090e565b005b61034d610bb7565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b61040a600480360360808110156103b757600080fd5b81019080803590602001909291908035906020019092919080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916906020019092919080359060200190929190505050610bca565b005b61044e6004803603602081101561042257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c8b565b005b600061045a610e80565b905090565b6000610469610eaa565b73ffffffffffffffffffffffffffffffffffffffff166104876108e4565b73ffffffffffffffffffffffffffffffffffffffff1614610510576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b634357855e60e01b600860006101000a81548163ffffffff021916908360e01c021790555061053d611b68565b61054f8730634357855e60e01b610eb2565b905061059b6040518060400160405280600381526020017f67657400000000000000000000000000000000000000000000000000000000008152508683610ee39092919063ffffffff16565b6105e56040518060400160405280600481526020017f70617468000000000000000000000000000000000000000000000000000000008152508583610ee39092919063ffffffff16565b61062f6040518060400160405280600581526020017f74696d65730000000000000000000000000000000000000000000000000000008152508483610f169092919063ffffffff16565b61063a888288610f49565b9150509695505050505050565b816005600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180611c536028913960400191505060405180910390fd5b6005600082815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055807f7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a60405160405180910390a281600781905550505050565b610776610eaa565b73ffffffffffffffffffffffffffffffffffffffff166107946108e4565b73ffffffffffffffffffffffffffffffffffffffff161461081d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60075481565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610916610eaa565b73ffffffffffffffffffffffffffffffffffffffff166109346108e4565b73ffffffffffffffffffffffffffffffffffffffff16146109bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b60006109c7610e80565b90508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb338373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610a6357600080fd5b505afa158015610a77573d6000803e3d6000fd5b505050506040513d6020811015610a8d57600080fd5b81019080805190602001909291905050506040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610b0757600080fd5b505af1158015610b1b573d6000803e3d6000fd5b505050506040513d6020811015610b3157600080fd5b8101908080519060200190929190505050610bb4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f556e61626c6520746f207472616e73666572000000000000000000000000000081525060200191505060405180910390fd5b50565b600860009054906101000a900460e01b81565b610bd2610eaa565b73ffffffffffffffffffffffffffffffffffffffff16610bf06108e4565b73ffffffffffffffffffffffffffffffffffffffff1614610c79576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b610c8584848484611201565b50505050565b610c93610eaa565b73ffffffffffffffffffffffffffffffffffffffff16610cb16108e4565b73ffffffffffffffffffffffffffffffffffffffff1614610d3a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610dc0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180611c0a6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600033905090565b610eba611b68565b610ec2611b68565b610ed985858584611364909392919063ffffffff16565b9150509392505050565b610efa82846080015161141490919063ffffffff16565b610f1181846080015161141490919063ffffffff16565b505050565b610f2d82846080015161141490919063ffffffff16565b610f4481846080015161143990919063ffffffff16565b505050565b600030600454604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140182815260200192505050604051602081830303815290604052805190602001209050600454836060018181525050836005600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550807fb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af960405160405180910390a2600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634000aea08584611081876114dd565b6040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156111085780820151818401526020810190506110ed565b50505050905090810190601f1680156111355780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561115657600080fd5b505af115801561116a573d6000803e3d6000fd5b505050506040513d602081101561118057600080fd5b81019080805190602001909291905050506111e6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611c306023913960400191505060405180910390fd5b60016004600082825401925050819055508090509392505050565b60006005600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506005600086815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055847fe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c560405160405180910390a28073ffffffffffffffffffffffffffffffffffffffff16636ee4d553868686866040518563ffffffff1660e01b815260040180858152602001848152602001837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001828152602001945050505050600060405180830381600087803b15801561134557600080fd5b505af1158015611359573d6000803e3d6000fd5b505050505050505050565b61136c611b68565b61137c85608001516101006116e4565b508385600001818152505082856020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508185604001907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681525050849050949350505050565b6114218260038351611738565b611434818361187d90919063ffffffff16565b505050565b7fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008112156114705761146b828261189f565b6114d9565b67ffffffffffffffff81131561148f5761148a828261190d565b6114d8565b600081126114a8576114a382600083611738565b6114d7565b6114d6826001837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03611738565b5b5b5b5050565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340429946905060e01b60008084600001518560200151866040015187606001516001896080015160000151604051602401808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018881526020018781526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001857bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561164757808201518184015260208101905061162c565b50505050905090810190601f1680156116745780820380516001836020036101000a031916815260200191505b509950505050505050505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050919050565b6116ec611bd5565b6000602083816116f857fe5b0614611711576020828161170857fe5b06602003820191505b81836020018181525050604051808452600081528281016020016040525082905092915050565b601781116117655761175f8160058460ff16901b60ff16178461195990919063ffffffff16565b50611878565b60ff81116117a75761178a601860058460ff16901b178461195990919063ffffffff16565b506117a1816001856119799092919063ffffffff16565b50611877565b61ffff81116117ea576117cd601960058460ff16901b178461195990919063ffffffff16565b506117e4816002856119799092919063ffffffff16565b50611876565b63ffffffff811161182f57611812601a60058460ff16901b178461195990919063ffffffff16565b50611829816004856119799092919063ffffffff16565b50611875565b67ffffffffffffffff81116118745761185b601b60058460ff16901b178461195990919063ffffffff16565b50611872816008856119799092919063ffffffff16565b505b5b5b5b5b505050565b611885611bd5565b6118978384600001515184855161199b565b905092915050565b6118bd60036005600660ff16901b178361195990919063ffffffff16565b5061190982827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0360405160200180828152602001915050604051602081830303815290604052611a54565b5050565b61192b60026005600660ff16901b178361195990919063ffffffff16565b50611955828260405160200180828152602001915050604051602081830303815290604052611a54565b5050565b611961611bd5565b6119718384600001515184611a79565b905092915050565b611981611bd5565b611992848560000151518585611ac7565b90509392505050565b6119a3611bd5565b82518211156119b157600080fd5b846020015182850111156119dc576119db8560026119d58860200151888701611b28565b02611b44565b5b6000808651805187602083010193508088870111156119fb5787860182525b50506020850190505b60208410611a275780518252602082019150602081019050602084039350611a04565b60006001856020036101000a03905080198251168184511681811785525050879350505050949350505050565b611a618260028351611738565b611a74818361187d90919063ffffffff16565b505050565b611a81611bd5565b83602001518310611a9e57611a9d846002866020015102611b44565b5b835180516020858301018481535080851415611abb576001810182525b50508390509392505050565b611acf611bd5565b84602001518483011115611aed57611aec85600286850102611b44565b5b60006001836101000a039050855183868201018583198251161781525080518487011115611b1b5783860181525b5085915050949350505050565b600081831115611b3a57829050611b3e565b8190505b92915050565b606082600001519050611b5783836116e4565b50611b62838261187d565b50505050565b6040518060a0016040528060008019168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200160008152602001611bcf611bef565b81525090565b604051806040016040528060608152602001600081525090565b60405180604001604052806060815260200160008152509056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373756e61626c6520746f207472616e73666572416e6443616c6c20746f206f7261636c65536f75726365206d75737420626520746865206f7261636c65206f66207468652072657175657374a2646970667358221220d11c684f5bf351544a6757af65619bf2b40cc8be37878a2a732af1131fbe26cd64736f6c63430006000033"

// DeployAPIConsumer deploys a new Ethereum contract, binding an instance of APIConsumer to it.
func DeployAPIConsumer(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address) (common.Address, *types.Transaction, *APIConsumer, error) {
	parsed, err := abi.JSON(strings.NewReader(APIConsumerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(APIConsumerBin), backend, _link)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &APIConsumer{APIConsumerCaller: APIConsumerCaller{contract: contract}, APIConsumerTransactor: APIConsumerTransactor{contract: contract}, APIConsumerFilterer: APIConsumerFilterer{contract: contract}}, nil
}

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

// Selector is a free data retrieval call binding the contract method 0xea3d508a.
//
// Solidity: function selector() view returns(bytes4)
func (_APIConsumer *APIConsumerCaller) Selector(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _APIConsumer.contract.Call(opts, &out, "selector")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// Selector is a free data retrieval call binding the contract method 0xea3d508a.
//
// Solidity: function selector() view returns(bytes4)
func (_APIConsumer *APIConsumerSession) Selector() ([4]byte, error) {
	return _APIConsumer.Contract.Selector(&_APIConsumer.CallOpts)
}

// Selector is a free data retrieval call binding the contract method 0xea3d508a.
//
// Solidity: function selector() view returns(bytes4)
func (_APIConsumer *APIConsumerCallerSession) Selector() ([4]byte, error) {
	return _APIConsumer.Contract.Selector(&_APIConsumer.CallOpts)
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
