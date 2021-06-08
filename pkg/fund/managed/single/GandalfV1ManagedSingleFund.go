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
const SingleABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenSwap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"checkAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrowTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracleTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inTokenAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"title\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SingleBin is the compiled bytecode used for deploying new contracts.
var SingleBin = "0x6080604052737a250d5630b4cf539739df2c5dacb4c659f2488d600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507307de306ff27a2b630b1141956844eb1552b956b5600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550731f9840a85d5af5bf1d1762f925bdaddc4201f984600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073d0a1e359811322d97991e03f863a0c30c2cf029c600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561016457600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550739a5811ed25f34834e1451869194180fead119dd2600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611300806102cf6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806373d002241161007157806373d002241461015657806386a19e63146101605780639405de641461016a578063b0514a4614610174578063c45a015514610192578063f399e22e146101b0576100a9565b806316345f18146100ae578063481c6a75146100de5780634a79d50c146100fc57806369deb6591461011a5780636c82532714610138575b600080fd5b6100c860048036038101906100c39190610b96565b6101cc565b6040516100d59190610f12565b60405180910390f35b6100e6610285565b6040516100f39190610dec565b60405180910390f35b6101046102ab565b6040516101119190610f2d565b60405180910390f35b610122610339565b60405161012f9190610ed5565b60405180910390f35b6101406103e5565b60405161014d9190610ef7565b60405180910390f35b61015e61040b565b005b610168610706565b005b6101726107bd565b005b61017c610878565b6040516101899190610f6f565b60405180910390f35b61019a61087e565b6040516101a79190610dec565b60405180910390f35b6101ca60048036038101906101c59190610bbf565b6108a2565b005b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316345f18846040518263ffffffff1660e01b815260040161022a9190610dec565b60206040518083038186803b15801561024257600080fd5b505afa158015610256573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061027a9190610c7d565b905080915050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600280546102b890611173565b80601f01602080910402602001604051908101604052809291908181526020018280546102e490611173565b80156103315780601f1061030657610100808354040283529160200191610331565b820191906000526020600020905b81548152906001019060200180831161031457829003601f168201915b505050505081565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ee8c24b86040518163ffffffff1660e01b815260040160006040518083038186803b1580156103a357600080fd5b505afa1580156103b7573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103e09190610c13565b905090565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073d0a1e359811322d97991e03f863a0c30c2cf029c90506000731f9840a85d5af5bf1d1762f925bdaddc4201f9849050600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1667016345785d8a00006040518363ffffffff1660e01b81526004016104c4929190610e67565b602060405180830381600087803b1580156104de57600080fd5b505af11580156104f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105169190610c54565b50600067016345785d8a00009050600067010a741a4627800090506060848160008151811061056e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816001815181106105e3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600030905060006360bd4f989050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c11d79586868686866040518663ffffffff1660e01b815260040161068e959493929190610f8a565b600060405180830381600087803b1580156106a857600080fd5b505af11580156106bc573d6000803e3d6000fd5b505050507f6d54c56b64987faf49976828c0be87cb21e5ce8dbc256e52b2ad6118931c42fe878688876040516106f59493929190610e90565b60405180910390a150505050505050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330620f42406040518463ffffffff1660e01b815260040161076893929190610e30565b602060405180830381600087803b15801561078257600080fd5b505af1158015610796573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ba9190610c54565b50565b6000600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b815260040161081c929190610e07565b60206040518083038186803b15801561083457600080fd5b505afa158015610848573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086c9190610ca6565b905080600a8190555050565b600a5481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610930576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092790610f4f565b60405180910390fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806002908051906020019061098792919061098c565b505050565b82805461099890611173565b90600052602060002090601f0160209004810192826109ba5760008555610a01565b82601f106109d357805160ff1916838001178555610a01565b82800160010185558215610a01579182015b82811115610a005782518255916020019190600101906109e5565b5b509050610a0e9190610a12565b5090565b5b80821115610a2b576000816000905550600101610a13565b5090565b6000610a42610a3d84611009565b610fe4565b90508083825260208201905082856020860282011115610a6157600080fd5b60005b85811015610a915781610a778882610aee565b845260208401935060208301925050600181019050610a64565b5050509392505050565b6000610aae610aa984611035565b610fe4565b905082815260208101848484011115610ac657600080fd5b610ad1848285611131565b509392505050565b600081359050610ae88161126e565b92915050565b600081519050610afd8161126e565b92915050565b600082601f830112610b1457600080fd5b8151610b24848260208601610a2f565b91505092915050565b600081519050610b3c81611285565b92915050565b600081519050610b518161129c565b92915050565b600082601f830112610b6857600080fd5b8135610b78848260208601610a9b565b91505092915050565b600081519050610b90816112b3565b92915050565b600060208284031215610ba857600080fd5b6000610bb684828501610ad9565b91505092915050565b60008060408385031215610bd257600080fd5b6000610be085828601610ad9565b925050602083013567ffffffffffffffff811115610bfd57600080fd5b610c0985828601610b57565b9150509250929050565b600060208284031215610c2557600080fd5b600082015167ffffffffffffffff811115610c3f57600080fd5b610c4b84828501610b03565b91505092915050565b600060208284031215610c6657600080fd5b6000610c7484828501610b2d565b91505092915050565b600060208284031215610c8f57600080fd5b6000610c9d84828501610b42565b91505092915050565b600060208284031215610cb857600080fd5b6000610cc684828501610b81565b91505092915050565b6000610cdb8383610ce7565b60208301905092915050565b610cf0816110bb565b82525050565b610cff816110bb565b82525050565b6000610d1082611076565b610d1a8185611099565b9350610d2583611066565b8060005b83811015610d56578151610d3d8882610ccf565b9750610d488361108c565b925050600181019050610d29565b5085935050505092915050565b610d6c8161110d565b82525050565b610d7b816110d9565b82525050565b6000610d8c82611081565b610d9681856110aa565b9350610da6818560208601611140565b610daf81611234565b840191505092915050565b6000610dc76014836110aa565b9150610dd282611245565b602082019050919050565b610de681611103565b82525050565b6000602082019050610e016000830184610cf6565b92915050565b6000604082019050610e1c6000830185610cf6565b610e296020830184610cf6565b9392505050565b6000606082019050610e456000830186610cf6565b610e526020830185610cf6565b610e5f6040830184610ddd565b949350505050565b6000604082019050610e7c6000830185610cf6565b610e896020830184610ddd565b9392505050565b6000608082019050610ea56000830187610cf6565b610eb26020830186610ddd565b610ebf6040830185610cf6565b610ecc6060830184610ddd565b95945050505050565b60006020820190508181036000830152610eef8184610d05565b905092915050565b6000602082019050610f0c6000830184610d63565b92915050565b6000602082019050610f276000830184610d72565b92915050565b60006020820190508181036000830152610f478184610d81565b905092915050565b60006020820190508181036000830152610f6881610dba565b9050919050565b6000602082019050610f846000830184610ddd565b92915050565b600060a082019050610f9f6000830188610ddd565b610fac6020830187610ddd565b8181036040830152610fbe8186610d05565b9050610fcd6060830185610cf6565b610fda6080830184610ddd565b9695505050505050565b6000610fee610fff565b9050610ffa82826111a5565b919050565b6000604051905090565b600067ffffffffffffffff82111561102457611023611205565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156110505761104f611205565b5b61105982611234565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60006110c6826110e3565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006111188261111f565b9050919050565b600061112a826110e3565b9050919050565b82818337600083830152505050565b60005b8381101561115e578082015181840152602081019050611143565b8381111561116d576000848401525b50505050565b6000600282049050600182168061118b57607f821691505b6020821081141561119f5761119e6111d6565b5b50919050565b6111ae82611234565b810181811067ffffffffffffffff821117156111cd576111cc611205565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f47616e64616c6656313a20464f5242494444454e000000000000000000000000600082015250565b611277816110bb565b811461128257600080fd5b50565b61128e816110cd565b811461129957600080fd5b50565b6112a5816110d9565b81146112b057600080fd5b50565b6112bc81611103565b81146112c757600080fd5b5056fea264697066735822122068d1dbbec69031b907cff56e4746b0771604e69499268270ae8f535d8e5b85b064736f6c63430008040033"

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

// GetLatestPrice is a free data retrieval call binding the contract method 0x16345f18.
//
// Solidity: function getLatestPrice(address _tokenAddress) view returns(int256)
func (_Single *SingleCaller) GetLatestPrice(opts *bind.CallOpts, _tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Single.contract.Call(opts, &out, "getLatestPrice", _tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0x16345f18.
//
// Solidity: function getLatestPrice(address _tokenAddress) view returns(int256)
func (_Single *SingleSession) GetLatestPrice(_tokenAddress common.Address) (*big.Int, error) {
	return _Single.Contract.GetLatestPrice(&_Single.CallOpts, _tokenAddress)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x16345f18.
//
// Solidity: function getLatestPrice(address _tokenAddress) view returns(int256)
func (_Single *SingleCallerSession) GetLatestPrice(_tokenAddress common.Address) (*big.Int, error) {
	return _Single.Contract.GetLatestPrice(&_Single.CallOpts, _tokenAddress)
}

// GetOracleTokens is a free data retrieval call binding the contract method 0x69deb659.
//
// Solidity: function getOracleTokens() view returns(address[])
func (_Single *SingleCaller) GetOracleTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Single.contract.Call(opts, &out, "getOracleTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOracleTokens is a free data retrieval call binding the contract method 0x69deb659.
//
// Solidity: function getOracleTokens() view returns(address[])
func (_Single *SingleSession) GetOracleTokens() ([]common.Address, error) {
	return _Single.Contract.GetOracleTokens(&_Single.CallOpts)
}

// GetOracleTokens is a free data retrieval call binding the contract method 0x69deb659.
//
// Solidity: function getOracleTokens() view returns(address[])
func (_Single *SingleCallerSession) GetOracleTokens() ([]common.Address, error) {
	return _Single.Contract.GetOracleTokens(&_Single.CallOpts)
}

// InToken is a free data retrieval call binding the contract method 0x6c825327.
//
// Solidity: function inToken() view returns(address)
func (_Single *SingleCaller) InToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Single.contract.Call(opts, &out, "inToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InToken is a free data retrieval call binding the contract method 0x6c825327.
//
// Solidity: function inToken() view returns(address)
func (_Single *SingleSession) InToken() (common.Address, error) {
	return _Single.Contract.InToken(&_Single.CallOpts)
}

// InToken is a free data retrieval call binding the contract method 0x6c825327.
//
// Solidity: function inToken() view returns(address)
func (_Single *SingleCallerSession) InToken() (common.Address, error) {
	return _Single.Contract.InToken(&_Single.CallOpts)
}

// InTokenAllowance is a free data retrieval call binding the contract method 0xb0514a46.
//
// Solidity: function inTokenAllowance() view returns(uint256)
func (_Single *SingleCaller) InTokenAllowance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Single.contract.Call(opts, &out, "inTokenAllowance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InTokenAllowance is a free data retrieval call binding the contract method 0xb0514a46.
//
// Solidity: function inTokenAllowance() view returns(uint256)
func (_Single *SingleSession) InTokenAllowance() (*big.Int, error) {
	return _Single.Contract.InTokenAllowance(&_Single.CallOpts)
}

// InTokenAllowance is a free data retrieval call binding the contract method 0xb0514a46.
//
// Solidity: function inTokenAllowance() view returns(uint256)
func (_Single *SingleCallerSession) InTokenAllowance() (*big.Int, error) {
	return _Single.Contract.InTokenAllowance(&_Single.CallOpts)
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

// CheckAllowance is a paid mutator transaction binding the contract method 0x9405de64.
//
// Solidity: function checkAllowance() returns()
func (_Single *SingleTransactor) CheckAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Single.contract.Transact(opts, "checkAllowance")
}

// CheckAllowance is a paid mutator transaction binding the contract method 0x9405de64.
//
// Solidity: function checkAllowance() returns()
func (_Single *SingleSession) CheckAllowance() (*types.Transaction, error) {
	return _Single.Contract.CheckAllowance(&_Single.TransactOpts)
}

// CheckAllowance is a paid mutator transaction binding the contract method 0x9405de64.
//
// Solidity: function checkAllowance() returns()
func (_Single *SingleTransactorSession) CheckAllowance() (*types.Transaction, error) {
	return _Single.Contract.CheckAllowance(&_Single.TransactOpts)
}

// EscrowTokens is a paid mutator transaction binding the contract method 0x86a19e63.
//
// Solidity: function escrowTokens() returns()
func (_Single *SingleTransactor) EscrowTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Single.contract.Transact(opts, "escrowTokens")
}

// EscrowTokens is a paid mutator transaction binding the contract method 0x86a19e63.
//
// Solidity: function escrowTokens() returns()
func (_Single *SingleSession) EscrowTokens() (*types.Transaction, error) {
	return _Single.Contract.EscrowTokens(&_Single.TransactOpts)
}

// EscrowTokens is a paid mutator transaction binding the contract method 0x86a19e63.
//
// Solidity: function escrowTokens() returns()
func (_Single *SingleTransactorSession) EscrowTokens() (*types.Transaction, error) {
	return _Single.Contract.EscrowTokens(&_Single.TransactOpts)
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

// SwapTokens is a paid mutator transaction binding the contract method 0x73d00224.
//
// Solidity: function swapTokens() returns()
func (_Single *SingleTransactor) SwapTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Single.contract.Transact(opts, "swapTokens")
}

// SwapTokens is a paid mutator transaction binding the contract method 0x73d00224.
//
// Solidity: function swapTokens() returns()
func (_Single *SingleSession) SwapTokens() (*types.Transaction, error) {
	return _Single.Contract.SwapTokens(&_Single.TransactOpts)
}

// SwapTokens is a paid mutator transaction binding the contract method 0x73d00224.
//
// Solidity: function swapTokens() returns()
func (_Single *SingleTransactorSession) SwapTokens() (*types.Transaction, error) {
	return _Single.Contract.SwapTokens(&_Single.TransactOpts)
}

// SingleTokenSwapIterator is returned from FilterTokenSwap and is used to iterate over the raw logs and unpacked data for TokenSwap events raised by the Single contract.
type SingleTokenSwapIterator struct {
	Event *SingleTokenSwap // Event containing the contract specifics and raw log

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
func (it *SingleTokenSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleTokenSwap)
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
		it.Event = new(SingleTokenSwap)
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
func (it *SingleTokenSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleTokenSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleTokenSwap represents a TokenSwap event raised by the Single contract.
type SingleTokenSwap struct {
	TokenIn   common.Address
	AmountIn  *big.Int
	TokenOut  common.Address
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenSwap is a free log retrieval operation binding the contract event 0x6d54c56b64987faf49976828c0be87cb21e5ce8dbc256e52b2ad6118931c42fe.
//
// Solidity: event TokenSwap(address tokenIn, uint256 amountIn, address tokenOut, uint256 amountOut)
func (_Single *SingleFilterer) FilterTokenSwap(opts *bind.FilterOpts) (*SingleTokenSwapIterator, error) {

	logs, sub, err := _Single.contract.FilterLogs(opts, "TokenSwap")
	if err != nil {
		return nil, err
	}
	return &SingleTokenSwapIterator{contract: _Single.contract, event: "TokenSwap", logs: logs, sub: sub}, nil
}

// WatchTokenSwap is a free log subscription operation binding the contract event 0x6d54c56b64987faf49976828c0be87cb21e5ce8dbc256e52b2ad6118931c42fe.
//
// Solidity: event TokenSwap(address tokenIn, uint256 amountIn, address tokenOut, uint256 amountOut)
func (_Single *SingleFilterer) WatchTokenSwap(opts *bind.WatchOpts, sink chan<- *SingleTokenSwap) (event.Subscription, error) {

	logs, sub, err := _Single.contract.WatchLogs(opts, "TokenSwap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleTokenSwap)
				if err := _Single.contract.UnpackLog(event, "TokenSwap", log); err != nil {
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

// ParseTokenSwap is a log parse operation binding the contract event 0x6d54c56b64987faf49976828c0be87cb21e5ce8dbc256e52b2ad6118931c42fe.
//
// Solidity: event TokenSwap(address tokenIn, uint256 amountIn, address tokenOut, uint256 amountOut)
func (_Single *SingleFilterer) ParseTokenSwap(log types.Log) (*SingleTokenSwap, error) {
	event := new(SingleTokenSwap)
	if err := _Single.contract.UnpackLog(event, "TokenSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
