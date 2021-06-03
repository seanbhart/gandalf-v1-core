// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

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

// FactoryABI is the input ABI used to generate the binding from.
const FactoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fund\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"FundCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fund\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"FundCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"}],\"name\":\"createFund\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"fund\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fundList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundListLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"getFund\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FactoryBin is the compiled bytecode used for deploying new contracts.
var FactoryBin = "0x608060405234801561001057600080fd5b50611247806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630b262d021461005157806350a94e55146100815780636bcc4f09146100b1578063e4b96982146100cf575b600080fd5b61006b60048036038101906100669190610619565b6100ff565b6040516100789190610794565b60405180910390f35b61009b600480360381019061009691906105c5565b61013e565b6040516100a89190610794565b60405180910390f35b6100b9610196565b6040516100c6919061085d565b60405180910390f35b6100e960048036038101906100e491906105c5565b6101a3565b6040516100f69190610794565b60405180910390f35b6001818154811061010f57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000602052816000526040600020818051602081018201805184825260208301602085012081835280955050505050506000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600180549050905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610214576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020b9061083d565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083604051610278919061074a565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146102fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f49061081d565b60405180910390fd5b60006040518060200161030f90610526565b6020820181038252601f19601f8201166040525090506000848460405160200161033a929190610761565b604051602081830303815290604052805190602001209050808251602084016000f592508273ffffffffffffffffffffffffffffffffffffffff1663f399e22e86866040518363ffffffff1660e01b81526004016103999291906107af565b600060405180830381600087803b1580156103b357600080fd5b505af11580156103c7573d6000803e3d6000fd5b50505050826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002085604051610418919061074a565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001839080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508473ffffffffffffffffffffffffffffffffffffffff167f88b91b282690ac7f7f776c9c8ac1986d5a923994d35a501adf869afbf16a739b8585600180549050604051610516939291906107df565b60405180910390a2505092915050565b61075380610abf83390190565b60006105466105418461089d565b610878565b90508281526020810184848401111561055e57600080fd5b610569848285610931565b509392505050565b60008135905061058081610a90565b92915050565b600082601f83011261059757600080fd5b81356105a7848260208601610533565b91505092915050565b6000813590506105bf81610aa7565b92915050565b600080604083850312156105d857600080fd5b60006105e685828601610571565b925050602083013567ffffffffffffffff81111561060357600080fd5b61060f85828601610586565b9150509250929050565b60006020828403121561062b57600080fd5b6000610639848285016105b0565b91505092915050565b61064b816108f5565b82525050565b61066261065d826108f5565b6109a4565b82525050565b6000610673826108ce565b61067d81856108d9565b935061068d818560208601610940565b610696816109f7565b840191505092915050565b60006106ac826108ce565b6106b681856108ea565b93506106c6818560208601610940565b80840191505092915050565b60006106df6016836108d9565b91506106ea82610a15565b602082019050919050565b60006107026017836108d9565b915061070d82610a3e565b602082019050919050565b60006107256009836108ea565b915061073082610a67565b600982019050919050565b61074481610927565b82525050565b600061075682846106a1565b915081905092915050565b600061076c82610718565b91506107788285610651565b60148201915061078882846106a1565b91508190509392505050565b60006020820190506107a96000830184610642565b92915050565b60006040820190506107c46000830185610642565b81810360208301526107d68184610668565b90509392505050565b600060608201905081810360008301526107f98186610668565b90506108086020830185610642565b610815604083018461073b565b949350505050565b60006020820190508181036000830152610836816106d2565b9050919050565b60006020820190508181036000830152610856816106f5565b9050919050565b6000602082019050610872600083018461073b565b92915050565b6000610882610893565b905061088e8282610973565b919050565b6000604051905090565b600067ffffffffffffffff8211156108b8576108b76109c8565b5b6108c1826109f7565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600061090082610907565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561095e578082015181840152602081019050610943565b8381111561096d576000848401525b50505050565b61097c826109f7565b810181811067ffffffffffffffff8211171561099b5761099a6109c8565b5b80604052505050565b60006109af826109b6565b9050919050565b60006109c182610a08565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f47616e64616c6656313a2046554e445f45584953545300000000000000000000600082015250565b7f47616e64616c6656313a205a45524f5f41444452455353000000000000000000600082015250565b7f47616e64616c6656310000000000000000000000000000000000000000000000600082015250565b610a99816108f5565b8114610aa457600080fd5b50565b610ab081610927565b8114610abb57600080fd5b5056fe608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506106f3806100606000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063481c6a75146100515780634a79d50c1461006f578063c45a01551461008d578063f399e22e146100ab575b600080fd5b6100596100c7565b6040516100669190610468565b60405180910390f35b6100776100ed565b6040516100849190610483565b60405180910390f35b61009561017b565b6040516100a29190610468565b60405180910390f35b6100c560048036038101906100c091906103a9565b61019f565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600280546100fa906105ab565b80601f0160208091040260200160405190810160405280929190818152602001828054610126906105ab565b80156101735780601f1061014857610100808354040283529160200191610173565b820191906000526020600020905b81548152906001019060200180831161015657829003601f168201915b505050505081565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461022d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610224906104a5565b60405180910390fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060029080519060200190610284929190610289565b505050565b828054610295906105ab565b90600052602060002090601f0160209004810192826102b757600085556102fe565b82601f106102d057805160ff19168380011785556102fe565b828001600101855582156102fe579182015b828111156102fd5782518255916020019190600101906102e2565b5b50905061030b919061030f565b5090565b5b80821115610328576000816000905550600101610310565b5090565b600061033f61033a846104ea565b6104c5565b90508281526020810184848401111561035757600080fd5b610362848285610569565b509392505050565b600081359050610379816106a6565b92915050565b600082601f83011261039057600080fd5b81356103a084826020860161032c565b91505092915050565b600080604083850312156103bc57600080fd5b60006103ca8582860161036a565b925050602083013567ffffffffffffffff8111156103e757600080fd5b6103f38582860161037f565b9150509250929050565b61040681610537565b82525050565b60006104178261051b565b6104218185610526565b9350610431818560208601610578565b61043a8161066c565b840191505092915050565b6000610452601483610526565b915061045d8261067d565b602082019050919050565b600060208201905061047d60008301846103fd565b92915050565b6000602082019050818103600083015261049d818461040c565b905092915050565b600060208201905081810360008301526104be81610445565b9050919050565b60006104cf6104e0565b90506104db82826105dd565b919050565b6000604051905090565b600067ffffffffffffffff8211156105055761050461063d565b5b61050e8261066c565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600061054282610549565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b82818337600083830152505050565b60005b8381101561059657808201518184015260208101905061057b565b838111156105a5576000848401525b50505050565b600060028204905060018216806105c357607f821691505b602082108114156105d7576105d661060e565b5b50919050565b6105e68261066c565b810181811067ffffffffffffffff821117156106055761060461063d565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f47616e64616c6656313a20464f5242494444454e000000000000000000000000600082015250565b6106af81610537565b81146106ba57600080fd5b5056fea2646970667358221220be315ca1d7d9051cbb9670d04320346113e59dc27ac4bb68bfab381edeb4d33764736f6c63430008040033a26469706673582212203224fa87df497f6f1455369d9be86d241122e376e11a55fb0e962cca00e5a82664736f6c63430008040033"

// DeployFactory deploys a new Ethereum contract, binding an instance of Factory to it.
func DeployFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Factory, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// FundList is a free data retrieval call binding the contract method 0x0b262d02.
//
// Solidity: function fundList(uint256 ) view returns(address)
func (_Factory *FactoryCaller) FundList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "fundList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FundList is a free data retrieval call binding the contract method 0x0b262d02.
//
// Solidity: function fundList(uint256 ) view returns(address)
func (_Factory *FactorySession) FundList(arg0 *big.Int) (common.Address, error) {
	return _Factory.Contract.FundList(&_Factory.CallOpts, arg0)
}

// FundList is a free data retrieval call binding the contract method 0x0b262d02.
//
// Solidity: function fundList(uint256 ) view returns(address)
func (_Factory *FactoryCallerSession) FundList(arg0 *big.Int) (common.Address, error) {
	return _Factory.Contract.FundList(&_Factory.CallOpts, arg0)
}

// FundListLength is a free data retrieval call binding the contract method 0x6bcc4f09.
//
// Solidity: function fundListLength() view returns(uint256)
func (_Factory *FactoryCaller) FundListLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "fundListLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundListLength is a free data retrieval call binding the contract method 0x6bcc4f09.
//
// Solidity: function fundListLength() view returns(uint256)
func (_Factory *FactorySession) FundListLength() (*big.Int, error) {
	return _Factory.Contract.FundListLength(&_Factory.CallOpts)
}

// FundListLength is a free data retrieval call binding the contract method 0x6bcc4f09.
//
// Solidity: function fundListLength() view returns(uint256)
func (_Factory *FactoryCallerSession) FundListLength() (*big.Int, error) {
	return _Factory.Contract.FundListLength(&_Factory.CallOpts)
}

// GetFund is a free data retrieval call binding the contract method 0x50a94e55.
//
// Solidity: function getFund(address , string ) view returns(address)
func (_Factory *FactoryCaller) GetFund(opts *bind.CallOpts, arg0 common.Address, arg1 string) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getFund", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFund is a free data retrieval call binding the contract method 0x50a94e55.
//
// Solidity: function getFund(address , string ) view returns(address)
func (_Factory *FactorySession) GetFund(arg0 common.Address, arg1 string) (common.Address, error) {
	return _Factory.Contract.GetFund(&_Factory.CallOpts, arg0, arg1)
}

// GetFund is a free data retrieval call binding the contract method 0x50a94e55.
//
// Solidity: function getFund(address , string ) view returns(address)
func (_Factory *FactoryCallerSession) GetFund(arg0 common.Address, arg1 string) (common.Address, error) {
	return _Factory.Contract.GetFund(&_Factory.CallOpts, arg0, arg1)
}

// CreateFund is a paid mutator transaction binding the contract method 0xe4b96982.
//
// Solidity: function createFund(address manager, string title) returns(address fund)
func (_Factory *FactoryTransactor) CreateFund(opts *bind.TransactOpts, manager common.Address, title string) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createFund", manager, title)
}

// CreateFund is a paid mutator transaction binding the contract method 0xe4b96982.
//
// Solidity: function createFund(address manager, string title) returns(address fund)
func (_Factory *FactorySession) CreateFund(manager common.Address, title string) (*types.Transaction, error) {
	return _Factory.Contract.CreateFund(&_Factory.TransactOpts, manager, title)
}

// CreateFund is a paid mutator transaction binding the contract method 0xe4b96982.
//
// Solidity: function createFund(address manager, string title) returns(address fund)
func (_Factory *FactoryTransactorSession) CreateFund(manager common.Address, title string) (*types.Transaction, error) {
	return _Factory.Contract.CreateFund(&_Factory.TransactOpts, manager, title)
}

// FactoryFundCreatedIterator is returned from FilterFundCreated and is used to iterate over the raw logs and unpacked data for FundCreated events raised by the Factory contract.
type FactoryFundCreatedIterator struct {
	Event *FactoryFundCreated // Event containing the contract specifics and raw log

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
func (it *FactoryFundCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryFundCreated)
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
		it.Event = new(FactoryFundCreated)
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
func (it *FactoryFundCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryFundCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryFundCreated represents a FundCreated event raised by the Factory contract.
type FactoryFundCreated struct {
	Manager common.Address
	Title   string
	Fund    common.Address
	Arg3    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFundCreated is a free log retrieval operation binding the contract event 0x88b91b282690ac7f7f776c9c8ac1986d5a923994d35a501adf869afbf16a739b.
//
// Solidity: event FundCreated(address indexed manager, string title, address fund, uint256 arg3)
func (_Factory *FactoryFilterer) FilterFundCreated(opts *bind.FilterOpts, manager []common.Address) (*FactoryFundCreatedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "FundCreated", managerRule)
	if err != nil {
		return nil, err
	}
	return &FactoryFundCreatedIterator{contract: _Factory.contract, event: "FundCreated", logs: logs, sub: sub}, nil
}

// WatchFundCreated is a free log subscription operation binding the contract event 0x88b91b282690ac7f7f776c9c8ac1986d5a923994d35a501adf869afbf16a739b.
//
// Solidity: event FundCreated(address indexed manager, string title, address fund, uint256 arg3)
func (_Factory *FactoryFilterer) WatchFundCreated(opts *bind.WatchOpts, sink chan<- *FactoryFundCreated, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "FundCreated", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryFundCreated)
				if err := _Factory.contract.UnpackLog(event, "FundCreated", log); err != nil {
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

// ParseFundCreated is a log parse operation binding the contract event 0x88b91b282690ac7f7f776c9c8ac1986d5a923994d35a501adf869afbf16a739b.
//
// Solidity: event FundCreated(address indexed manager, string title, address fund, uint256 arg3)
func (_Factory *FactoryFilterer) ParseFundCreated(log types.Log) (*FactoryFundCreated, error) {
	event := new(FactoryFundCreated)
	if err := _Factory.contract.UnpackLog(event, "FundCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryFundCreated0Iterator is returned from FilterFundCreated0 and is used to iterate over the raw logs and unpacked data for FundCreated0 events raised by the Factory contract.
type FactoryFundCreated0Iterator struct {
	Event *FactoryFundCreated0 // Event containing the contract specifics and raw log

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
func (it *FactoryFundCreated0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryFundCreated0)
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
		it.Event = new(FactoryFundCreated0)
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
func (it *FactoryFundCreated0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryFundCreated0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryFundCreated0 represents a FundCreated0 event raised by the Factory contract.
type FactoryFundCreated0 struct {
	Manager common.Address
	Fund    common.Address
	Arg2    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFundCreated0 is a free log retrieval operation binding the contract event 0x70890af59a1c61d7489da8e3724463160e14d24e026267fbcce52c5d03b0b03b.
//
// Solidity: event FundCreated(address indexed manager, address fund, uint256 arg2)
func (_Factory *FactoryFilterer) FilterFundCreated0(opts *bind.FilterOpts, manager []common.Address) (*FactoryFundCreated0Iterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "FundCreated0", managerRule)
	if err != nil {
		return nil, err
	}
	return &FactoryFundCreated0Iterator{contract: _Factory.contract, event: "FundCreated0", logs: logs, sub: sub}, nil
}

// WatchFundCreated0 is a free log subscription operation binding the contract event 0x70890af59a1c61d7489da8e3724463160e14d24e026267fbcce52c5d03b0b03b.
//
// Solidity: event FundCreated(address indexed manager, address fund, uint256 arg2)
func (_Factory *FactoryFilterer) WatchFundCreated0(opts *bind.WatchOpts, sink chan<- *FactoryFundCreated0, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "FundCreated0", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryFundCreated0)
				if err := _Factory.contract.UnpackLog(event, "FundCreated0", log); err != nil {
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

// ParseFundCreated0 is a log parse operation binding the contract event 0x70890af59a1c61d7489da8e3724463160e14d24e026267fbcce52c5d03b0b03b.
//
// Solidity: event FundCreated(address indexed manager, address fund, uint256 arg2)
func (_Factory *FactoryFilterer) ParseFundCreated0(log types.Log) (*FactoryFundCreated0, error) {
	event := new(FactoryFundCreated0)
	if err := _Factory.contract.UnpackLog(event, "FundCreated0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
