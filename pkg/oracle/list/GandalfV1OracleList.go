// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oraclelist

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

// OraclelistABI is the input ABI used to generate the binding from.
const OraclelistABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oracleMap\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OraclelistBin is the compiled bytecode used for deploying new contracts.
var OraclelistBin = "0x608060405234801561001057600080fd5b506000739b6ff80ff8348852d5281de45e66b7ed36e7b8a99080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073da5904bdbfb4ef12a3955aeca103f51dc87c7c3960016000739b6ff80ff8348852d5281de45e66b7ed36e7b8a973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506106908061013d6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806316345f18146100515780638a45a55414610081578063e5df8b84146100b1578063ee8c24b8146100e1575b600080fd5b61006b60048036038101906100669190610357565b6100ff565b604051610078919061052a565b60405180910390f35b61009b60048036038101906100969190610357565b6101ee565b6040516100a8919061050f565b60405180910390f35b6100cb60048036038101906100c69190610380565b610221565b6040516100d891906104d2565b60405180910390f35b6100e9610260565b6040516100f691906104ed565b60405180910390f35b600080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b1580156101a757600080fd5b505afa1580156101bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101df91906103a9565b50505091505080915050919050565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000818154811061023157600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060008054806020026020016040519081016040528092919081815260200182805480156102e457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161029a575b5050505050905090565b6000813590506102fd816105fe565b92915050565b60008151905061031281610615565b92915050565b6000813590506103278161062c565b92915050565b60008151905061033c8161062c565b92915050565b60008151905061035181610643565b92915050565b60006020828403121561036957600080fd5b6000610377848285016102ee565b91505092915050565b60006020828403121561039257600080fd5b60006103a084828501610318565b91505092915050565b600080600080600060a086880312156103c157600080fd5b60006103cf88828901610342565b95505060206103e088828901610303565b94505060406103f18882890161032d565b93505060606104028882890161032d565b925050608061041388828901610342565b9150509295509295909350565b600061042c8383610438565b60208301905092915050565b6104418161057e565b82525050565b6104508161057e565b82525050565b600061046182610555565b61046b818561056d565b935061047683610545565b8060005b838110156104a757815161048e8882610420565b975061049983610560565b92505060018101905061047a565b5085935050505092915050565b6104bd816105da565b82525050565b6104cc81610590565b82525050565b60006020820190506104e76000830184610447565b92915050565b600060208201905081810360008301526105078184610456565b905092915050565b600060208201905061052460008301846104b4565b92915050565b600060208201905061053f60008301846104c3565b92915050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b60006105898261059a565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600069ffffffffffffffffffff82169050919050565b60006105e5826105ec565b9050919050565b60006105f78261059a565b9050919050565b6106078161057e565b811461061257600080fd5b50565b61061e81610590565b811461062957600080fd5b50565b610635816105ba565b811461064057600080fd5b50565b61064c816105c4565b811461065757600080fd5b5056fea2646970667358221220370c303a6dd70450a200a2b27e76c690c368dcc745eccab2fcaa873b8498561f64736f6c63430008040033"

// DeployOraclelist deploys a new Ethereum contract, binding an instance of Oraclelist to it.
func DeployOraclelist(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Oraclelist, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclelistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OraclelistBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oraclelist{OraclelistCaller: OraclelistCaller{contract: contract}, OraclelistTransactor: OraclelistTransactor{contract: contract}, OraclelistFilterer: OraclelistFilterer{contract: contract}}, nil
}

// Oraclelist is an auto generated Go binding around an Ethereum contract.
type Oraclelist struct {
	OraclelistCaller     // Read-only binding to the contract
	OraclelistTransactor // Write-only binding to the contract
	OraclelistFilterer   // Log filterer for contract events
}

// OraclelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OraclelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclelistSession struct {
	Contract     *Oraclelist       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OraclelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclelistCallerSession struct {
	Contract *OraclelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OraclelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclelistTransactorSession struct {
	Contract     *OraclelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OraclelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclelistRaw struct {
	Contract *Oraclelist // Generic contract binding to access the raw methods on
}

// OraclelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclelistCallerRaw struct {
	Contract *OraclelistCaller // Generic read-only contract binding to access the raw methods on
}

// OraclelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclelistTransactorRaw struct {
	Contract *OraclelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclelist creates a new instance of Oraclelist, bound to a specific deployed contract.
func NewOraclelist(address common.Address, backend bind.ContractBackend) (*Oraclelist, error) {
	contract, err := bindOraclelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oraclelist{OraclelistCaller: OraclelistCaller{contract: contract}, OraclelistTransactor: OraclelistTransactor{contract: contract}, OraclelistFilterer: OraclelistFilterer{contract: contract}}, nil
}

// NewOraclelistCaller creates a new read-only instance of Oraclelist, bound to a specific deployed contract.
func NewOraclelistCaller(address common.Address, caller bind.ContractCaller) (*OraclelistCaller, error) {
	contract, err := bindOraclelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OraclelistCaller{contract: contract}, nil
}

// NewOraclelistTransactor creates a new write-only instance of Oraclelist, bound to a specific deployed contract.
func NewOraclelistTransactor(address common.Address, transactor bind.ContractTransactor) (*OraclelistTransactor, error) {
	contract, err := bindOraclelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OraclelistTransactor{contract: contract}, nil
}

// NewOraclelistFilterer creates a new log filterer instance of Oraclelist, bound to a specific deployed contract.
func NewOraclelistFilterer(address common.Address, filterer bind.ContractFilterer) (*OraclelistFilterer, error) {
	contract, err := bindOraclelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OraclelistFilterer{contract: contract}, nil
}

// bindOraclelist binds a generic wrapper to an already deployed contract.
func bindOraclelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oraclelist *OraclelistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oraclelist.Contract.OraclelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oraclelist *OraclelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oraclelist.Contract.OraclelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oraclelist *OraclelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oraclelist.Contract.OraclelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oraclelist *OraclelistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oraclelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oraclelist *OraclelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oraclelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oraclelist *OraclelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oraclelist.Contract.contract.Transact(opts, method, params...)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x16345f18.
//
// Solidity: function getLatestPrice(address _tokenAddress) view returns(int256)
func (_Oraclelist *OraclelistCaller) GetLatestPrice(opts *bind.CallOpts, _tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oraclelist.contract.Call(opts, &out, "getLatestPrice", _tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0x16345f18.
//
// Solidity: function getLatestPrice(address _tokenAddress) view returns(int256)
func (_Oraclelist *OraclelistSession) GetLatestPrice(_tokenAddress common.Address) (*big.Int, error) {
	return _Oraclelist.Contract.GetLatestPrice(&_Oraclelist.CallOpts, _tokenAddress)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x16345f18.
//
// Solidity: function getLatestPrice(address _tokenAddress) view returns(int256)
func (_Oraclelist *OraclelistCallerSession) GetLatestPrice(_tokenAddress common.Address) (*big.Int, error) {
	return _Oraclelist.Contract.GetLatestPrice(&_Oraclelist.CallOpts, _tokenAddress)
}

// GetTokenAddresses is a free data retrieval call binding the contract method 0xee8c24b8.
//
// Solidity: function getTokenAddresses() view returns(address[])
func (_Oraclelist *OraclelistCaller) GetTokenAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Oraclelist.contract.Call(opts, &out, "getTokenAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenAddresses is a free data retrieval call binding the contract method 0xee8c24b8.
//
// Solidity: function getTokenAddresses() view returns(address[])
func (_Oraclelist *OraclelistSession) GetTokenAddresses() ([]common.Address, error) {
	return _Oraclelist.Contract.GetTokenAddresses(&_Oraclelist.CallOpts)
}

// GetTokenAddresses is a free data retrieval call binding the contract method 0xee8c24b8.
//
// Solidity: function getTokenAddresses() view returns(address[])
func (_Oraclelist *OraclelistCallerSession) GetTokenAddresses() ([]common.Address, error) {
	return _Oraclelist.Contract.GetTokenAddresses(&_Oraclelist.CallOpts)
}

// OracleMap is a free data retrieval call binding the contract method 0x8a45a554.
//
// Solidity: function oracleMap(address ) view returns(address)
func (_Oraclelist *OraclelistCaller) OracleMap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Oraclelist.contract.Call(opts, &out, "oracleMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleMap is a free data retrieval call binding the contract method 0x8a45a554.
//
// Solidity: function oracleMap(address ) view returns(address)
func (_Oraclelist *OraclelistSession) OracleMap(arg0 common.Address) (common.Address, error) {
	return _Oraclelist.Contract.OracleMap(&_Oraclelist.CallOpts, arg0)
}

// OracleMap is a free data retrieval call binding the contract method 0x8a45a554.
//
// Solidity: function oracleMap(address ) view returns(address)
func (_Oraclelist *OraclelistCallerSession) OracleMap(arg0 common.Address) (common.Address, error) {
	return _Oraclelist.Contract.OracleMap(&_Oraclelist.CallOpts, arg0)
}

// TokenAddresses is a free data retrieval call binding the contract method 0xe5df8b84.
//
// Solidity: function tokenAddresses(uint256 ) view returns(address)
func (_Oraclelist *OraclelistCaller) TokenAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Oraclelist.contract.Call(opts, &out, "tokenAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddresses is a free data retrieval call binding the contract method 0xe5df8b84.
//
// Solidity: function tokenAddresses(uint256 ) view returns(address)
func (_Oraclelist *OraclelistSession) TokenAddresses(arg0 *big.Int) (common.Address, error) {
	return _Oraclelist.Contract.TokenAddresses(&_Oraclelist.CallOpts, arg0)
}

// TokenAddresses is a free data retrieval call binding the contract method 0xe5df8b84.
//
// Solidity: function tokenAddresses(uint256 ) view returns(address)
func (_Oraclelist *OraclelistCallerSession) TokenAddresses(arg0 *big.Int) (common.Address, error) {
	return _Oraclelist.Contract.TokenAddresses(&_Oraclelist.CallOpts, arg0)
}
