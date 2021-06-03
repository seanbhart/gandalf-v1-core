// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package single

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

// SingleABI is the input ABI used to generate the binding from.
const SingleABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"title\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SingleBin is the compiled bytecode used for deploying new contracts.
var SingleBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506106f3806100606000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063481c6a75146100515780634a79d50c1461006f578063c45a01551461008d578063f399e22e146100ab575b600080fd5b6100596100c7565b6040516100669190610468565b60405180910390f35b6100776100ed565b6040516100849190610483565b60405180910390f35b61009561017b565b6040516100a29190610468565b60405180910390f35b6100c560048036038101906100c091906103a9565b61019f565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600280546100fa906105ab565b80601f0160208091040260200160405190810160405280929190818152602001828054610126906105ab565b80156101735780601f1061014857610100808354040283529160200191610173565b820191906000526020600020905b81548152906001019060200180831161015657829003601f168201915b505050505081565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461022d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610224906104a5565b60405180910390fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060029080519060200190610284929190610289565b505050565b828054610295906105ab565b90600052602060002090601f0160209004810192826102b757600085556102fe565b82601f106102d057805160ff19168380011785556102fe565b828001600101855582156102fe579182015b828111156102fd5782518255916020019190600101906102e2565b5b50905061030b919061030f565b5090565b5b80821115610328576000816000905550600101610310565b5090565b600061033f61033a846104ea565b6104c5565b90508281526020810184848401111561035757600080fd5b610362848285610569565b509392505050565b600081359050610379816106a6565b92915050565b600082601f83011261039057600080fd5b81356103a084826020860161032c565b91505092915050565b600080604083850312156103bc57600080fd5b60006103ca8582860161036a565b925050602083013567ffffffffffffffff8111156103e757600080fd5b6103f38582860161037f565b9150509250929050565b61040681610537565b82525050565b60006104178261051b565b6104218185610526565b9350610431818560208601610578565b61043a8161066c565b840191505092915050565b6000610452601483610526565b915061045d8261067d565b602082019050919050565b600060208201905061047d60008301846103fd565b92915050565b6000602082019050818103600083015261049d818461040c565b905092915050565b600060208201905081810360008301526104be81610445565b9050919050565b60006104cf6104e0565b90506104db82826105dd565b919050565b6000604051905090565b600067ffffffffffffffff8211156105055761050461063d565b5b61050e8261066c565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600061054282610549565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b82818337600083830152505050565b60005b8381101561059657808201518184015260208101905061057b565b838111156105a5576000848401525b50505050565b600060028204905060018216806105c357607f821691505b602082108114156105d7576105d661060e565b5b50919050565b6105e68261066c565b810181811067ffffffffffffffff821117156106055761060461063d565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f47616e64616c6656313a20464f5242494444454e000000000000000000000000600082015250565b6106af81610537565b81146106ba57600080fd5b5056fea2646970667358221220be315ca1d7d9051cbb9670d04320346113e59dc27ac4bb68bfab381edeb4d33764736f6c63430008040033"

// DeploySingle deploys a new Ethereum contract, binding an instance of Single to it.
func DeploySingle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Single, error) {
	parsed, err := abi.JSON(strings.NewReader(SingleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SingleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Single{SingleCaller: SingleCaller{contract: contract}, SingleTransactor: SingleTransactor{contract: contract}, SingleFilterer: SingleFilterer{contract: contract}}, nil
}

// Single is an auto generated Go binding around an Ethereum contract.
type Single struct {
	SingleCaller     // Read-only binding to the contract
	SingleTransactor // Write-only binding to the contract
	SingleFilterer   // Log filterer for contract events
}

// SingleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SingleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SingleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SingleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SingleSession struct {
	Contract     *Single           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SingleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SingleCallerSession struct {
	Contract *SingleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SingleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SingleTransactorSession struct {
	Contract     *SingleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SingleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SingleRaw struct {
	Contract *Single // Generic contract binding to access the raw methods on
}

// SingleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SingleCallerRaw struct {
	Contract *SingleCaller // Generic read-only contract binding to access the raw methods on
}

// SingleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SingleTransactorRaw struct {
	Contract *SingleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSingle creates a new instance of Single, bound to a specific deployed contract.
func NewSingle(address common.Address, backend bind.ContractBackend) (*Single, error) {
	contract, err := bindSingle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Single{SingleCaller: SingleCaller{contract: contract}, SingleTransactor: SingleTransactor{contract: contract}, SingleFilterer: SingleFilterer{contract: contract}}, nil
}

// NewSingleCaller creates a new read-only instance of Single, bound to a specific deployed contract.
func NewSingleCaller(address common.Address, caller bind.ContractCaller) (*SingleCaller, error) {
	contract, err := bindSingle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SingleCaller{contract: contract}, nil
}

// NewSingleTransactor creates a new write-only instance of Single, bound to a specific deployed contract.
func NewSingleTransactor(address common.Address, transactor bind.ContractTransactor) (*SingleTransactor, error) {
	contract, err := bindSingle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SingleTransactor{contract: contract}, nil
}

// NewSingleFilterer creates a new log filterer instance of Single, bound to a specific deployed contract.
func NewSingleFilterer(address common.Address, filterer bind.ContractFilterer) (*SingleFilterer, error) {
	contract, err := bindSingle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SingleFilterer{contract: contract}, nil
}

// bindSingle binds a generic wrapper to an already deployed contract.
func bindSingle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SingleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Single *SingleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Single.Contract.SingleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Single *SingleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Single.Contract.SingleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Single *SingleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Single.Contract.SingleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Single *SingleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Single.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Single *SingleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Single.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Single *SingleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Single.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Single *SingleCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Single.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Single *SingleSession) Factory() (common.Address, error) {
	return _Single.Contract.Factory(&_Single.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Single *SingleCallerSession) Factory() (common.Address, error) {
	return _Single.Contract.Factory(&_Single.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Single *SingleCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Single.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Single *SingleSession) Manager() (common.Address, error) {
	return _Single.Contract.Manager(&_Single.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Single *SingleCallerSession) Manager() (common.Address, error) {
	return _Single.Contract.Manager(&_Single.CallOpts)
}

// Title is a free data retrieval call binding the contract method 0x4a79d50c.
//
// Solidity: function title() view returns(string)
func (_Single *SingleCaller) Title(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Single.contract.Call(opts, &out, "title")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Title is a free data retrieval call binding the contract method 0x4a79d50c.
//
// Solidity: function title() view returns(string)
func (_Single *SingleSession) Title() (string, error) {
	return _Single.Contract.Title(&_Single.CallOpts)
}

// Title is a free data retrieval call binding the contract method 0x4a79d50c.
//
// Solidity: function title() view returns(string)
func (_Single *SingleCallerSession) Title() (string, error) {
	return _Single.Contract.Title(&_Single.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _manager, string _title) returns()
func (_Single *SingleTransactor) Initialize(opts *bind.TransactOpts, _manager common.Address, _title string) (*types.Transaction, error) {
	return _Single.contract.Transact(opts, "initialize", _manager, _title)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _manager, string _title) returns()
func (_Single *SingleSession) Initialize(_manager common.Address, _title string) (*types.Transaction, error) {
	return _Single.Contract.Initialize(&_Single.TransactOpts, _manager, _title)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _manager, string _title) returns()
func (_Single *SingleTransactorSession) Initialize(_manager common.Address, _title string) (*types.Transaction, error) {
	return _Single.Contract.Initialize(&_Single.TransactOpts, _manager, _title)
}
