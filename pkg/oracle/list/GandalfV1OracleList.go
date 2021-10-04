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
const OraclelistABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"}],\"name\":\"OracleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"}],\"name\":\"OracleRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"}],\"name\":\"addOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oracleMap\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"removeOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OraclelistBin is the compiled bytecode used for deploying new contracts.
var OraclelistBin = "0x608060405234801561001057600080fd5b50600061002161041b60201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35073da5904bdbfb4ef12a3955aeca103f51dc87c7c3960026000731f9840a85d5af5bf1d1762f925bdaddc4201f98473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001731f9840a85d5af5bf1d1762f925bdaddc4201f9849080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550732ca5a90d34ca333661083f89d831f757a9a50148600260007307de306ff27a2b630b1141956844eb1552b956b573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060017307de306ff27a2b630b1141956844eb1552b956b59080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550739326bfa02add2366b30bacb125260af6410313316002600073d0a1e359811322d97991e03f863a0c30c2cf029c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600173d0a1e359811322d97991e03f863a0c30c2cf029c9080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610423565b600033905090565b611458806104326000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063e5df8b8411610066578063e5df8b8414610120578063ee8c24b814610150578063f0ca4adb1461016e578063f2fde38b1461018a578063fdc85fc4146101a657610093565b806316345f1814610098578063715018a6146100c85780638a45a554146100d25780638da5cb5b14610102575b600080fd5b6100b260048036038101906100ad9190610d65565b6101c2565b6040516100bf9190611066565b60405180910390f35b6100d06102b1565b005b6100ec60048036038101906100e79190610d65565b6103eb565b6040516100f9919061104b565b60405180910390f35b61010a61041e565b6040516101179190610fe5565b60405180910390f35b61013a60048036038101906101359190610dd2565b610447565b6040516101479190610fe5565b60405180910390f35b610158610486565b6040516101659190611029565b60405180910390f35b61018860048036038101906101839190610d92565b610514565b005b6101a4600480360381019061019f9190610d65565b6107d0565b005b6101c060048036038101906101bb9190610d65565b610979565b005b600080600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b15801561026a57600080fd5b505afa15801561027e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102a29190610dff565b50505091505080915050919050565b6102b9610cdf565b73ffffffffffffffffffffffffffffffffffffffff166102d761041e565b73ffffffffffffffffffffffffffffffffffffffff161461032d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610324906110c1565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6001818154811061045757600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600180548060200260200160405190810160405280929190818152602001828054801561050a57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116104c0575b5050505050905090565b61051c610cdf565b73ffffffffffffffffffffffffffffffffffffffff1661053a61041e565b73ffffffffffffffffffffffffffffffffffffffff1614610590576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610587906110c1565b60405180910390fd5b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b15801561063757600080fd5b505afa15801561064b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066f9190610e7a565b60ff16116106b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a990611081565b60405180910390fd5b6001829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f828d2be040dede7698182e08dfa8bfbd663c879aee772509c4a2bd961d0ed43f82826040516107c4929190611000565b60405180910390a15050565b6107d8610cdf565b73ffffffffffffffffffffffffffffffffffffffff166107f661041e565b73ffffffffffffffffffffffffffffffffffffffff161461084c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610843906110c1565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156108bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b3906110a1565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b610981610cdf565b73ffffffffffffffffffffffffffffffffffffffff1661099f61041e565b73ffffffffffffffffffffffffffffffffffffffff16146109f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ec906110c1565b60405180910390fd5b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015610aa157600080fd5b505afa158015610ab5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad99190610e7a565b60ff1611610b1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b13906110e1565b60405180910390fd5b60005b6001805490508160ff161015610c3f576001808080549050610b41919061114b565b81548110610b5257610b51611265565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660018260ff1681548110610b9457610b93611265565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001808080549050610bee919061114b565b81548110610bff57610bfe611265565b5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690558080610c379061120c565b915050610b1f565b50600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690557f6dc84b66cc948d847632b9d829f7cb1cb904fbf2c084554a9bc22ad9d84533408282604051610cd3929190611000565b60405180910390a15050565b600033905090565b600081359050610cf6816113af565b92915050565b600081519050610d0b816113c6565b92915050565b600081359050610d20816113dd565b92915050565b600081519050610d35816113dd565b92915050565b600081519050610d4a8161140b565b92915050565b600081519050610d5f816113f4565b92915050565b600060208284031215610d7b57610d7a611294565b5b6000610d8984828501610ce7565b91505092915050565b60008060408385031215610da957610da8611294565b5b6000610db785828601610ce7565b9250506020610dc885828601610ce7565b9150509250929050565b600060208284031215610de857610de7611294565b5b6000610df684828501610d11565b91505092915050565b600080600080600060a08688031215610e1b57610e1a611294565b5b6000610e2988828901610d3b565b9550506020610e3a88828901610cfc565b9450506040610e4b88828901610d26565b9350506060610e5c88828901610d26565b9250506080610e6d88828901610d3b565b9150509295509295909350565b600060208284031215610e9057610e8f611294565b5b6000610e9e84828501610d50565b91505092915050565b6000610eb38383610ebf565b60208301905092915050565b610ec88161117f565b82525050565b610ed78161117f565b82525050565b6000610ee882611111565b610ef28185611129565b9350610efd83611101565b8060005b83811015610f2e578151610f158882610ea7565b9750610f208361111c565b925050600181019050610f01565b5085935050505092915050565b610f44816111e8565b82525050565b610f5381611191565b82525050565b6000610f6660308361113a565b9150610f7182611299565b604082019050919050565b6000610f8960268361113a565b9150610f94826112e8565b604082019050919050565b6000610fac60208361113a565b9150610fb782611337565b602082019050919050565b6000610fcf60338361113a565b9150610fda82611360565b604082019050919050565b6000602082019050610ffa6000830184610ece565b92915050565b60006040820190506110156000830185610ece565b6110226020830184610ece565b9392505050565b600060208201905081810360008301526110438184610edd565b905092915050565b60006020820190506110606000830184610f3b565b92915050565b600060208201905061107b6000830184610f4a565b92915050565b6000602082019050818103600083015261109a81610f59565b9050919050565b600060208201905081810360008301526110ba81610f7c565b9050919050565b600060208201905081810360008301526110da81610f9f565b9050919050565b600060208201905081810360008301526110fa81610fc2565b9050919050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611156826111bb565b9150611161836111bb565b92508282101561117457611173611236565b5b828203905092915050565b600061118a8261119b565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b600069ffffffffffffffffffff82169050919050565b60006111f3826111fa565b9050919050565b60006112058261119b565b9050919050565b6000611217826111c5565b915060ff82141561122b5761122a611236565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600080fd5b7f6164644f7261636c653a20546f6b656e20616e64204f7261636c65207061697260008201527f20616c7265616479206578697374732e00000000000000000000000000000000602082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f72656d6f76654f7261636c653a20546f6b656e20616e64204f7261636c65207060008201527f61697220646f6573206e6f742065786973742e00000000000000000000000000602082015250565b6113b88161117f565b81146113c357600080fd5b50565b6113cf81611191565b81146113da57600080fd5b50565b6113e6816111bb565b81146113f157600080fd5b50565b6113fd816111c5565b811461140857600080fd5b50565b611414816111d2565b811461141f57600080fd5b5056fea264697066735822122058837cb434bf138a7dc1d8b3a1d4788533bbb50dfeb12f970f308be60ee5860764736f6c63430008060033"

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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oraclelist *OraclelistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oraclelist.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oraclelist *OraclelistSession) Owner() (common.Address, error) {
	return _Oraclelist.Contract.Owner(&_Oraclelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oraclelist *OraclelistCallerSession) Owner() (common.Address, error) {
	return _Oraclelist.Contract.Owner(&_Oraclelist.CallOpts)
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

// AddOracle is a paid mutator transaction binding the contract method 0xf0ca4adb.
//
// Solidity: function addOracle(address _tokenAddress, address _oracleAddress) returns()
func (_Oraclelist *OraclelistTransactor) AddOracle(opts *bind.TransactOpts, _tokenAddress common.Address, _oracleAddress common.Address) (*types.Transaction, error) {
	return _Oraclelist.contract.Transact(opts, "addOracle", _tokenAddress, _oracleAddress)
}

// AddOracle is a paid mutator transaction binding the contract method 0xf0ca4adb.
//
// Solidity: function addOracle(address _tokenAddress, address _oracleAddress) returns()
func (_Oraclelist *OraclelistSession) AddOracle(_tokenAddress common.Address, _oracleAddress common.Address) (*types.Transaction, error) {
	return _Oraclelist.Contract.AddOracle(&_Oraclelist.TransactOpts, _tokenAddress, _oracleAddress)
}

// AddOracle is a paid mutator transaction binding the contract method 0xf0ca4adb.
//
// Solidity: function addOracle(address _tokenAddress, address _oracleAddress) returns()
func (_Oraclelist *OraclelistTransactorSession) AddOracle(_tokenAddress common.Address, _oracleAddress common.Address) (*types.Transaction, error) {
	return _Oraclelist.Contract.AddOracle(&_Oraclelist.TransactOpts, _tokenAddress, _oracleAddress)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xfdc85fc4.
//
// Solidity: function removeOracle(address _tokenAddress) returns()
func (_Oraclelist *OraclelistTransactor) RemoveOracle(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Oraclelist.contract.Transact(opts, "removeOracle", _tokenAddress)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xfdc85fc4.
//
// Solidity: function removeOracle(address _tokenAddress) returns()
func (_Oraclelist *OraclelistSession) RemoveOracle(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Oraclelist.Contract.RemoveOracle(&_Oraclelist.TransactOpts, _tokenAddress)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xfdc85fc4.
//
// Solidity: function removeOracle(address _tokenAddress) returns()
func (_Oraclelist *OraclelistTransactorSession) RemoveOracle(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Oraclelist.Contract.RemoveOracle(&_Oraclelist.TransactOpts, _tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oraclelist *OraclelistTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oraclelist.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oraclelist *OraclelistSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oraclelist.Contract.RenounceOwnership(&_Oraclelist.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oraclelist *OraclelistTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oraclelist.Contract.RenounceOwnership(&_Oraclelist.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oraclelist *OraclelistTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Oraclelist.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oraclelist *OraclelistSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oraclelist.Contract.TransferOwnership(&_Oraclelist.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oraclelist *OraclelistTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oraclelist.Contract.TransferOwnership(&_Oraclelist.TransactOpts, newOwner)
}

// OraclelistOracleAddedIterator is returned from FilterOracleAdded and is used to iterate over the raw logs and unpacked data for OracleAdded events raised by the Oraclelist contract.
type OraclelistOracleAddedIterator struct {
	Event *OraclelistOracleAdded // Event containing the contract specifics and raw log

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
func (it *OraclelistOracleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OraclelistOracleAdded)
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
		it.Event = new(OraclelistOracleAdded)
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
func (it *OraclelistOracleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OraclelistOracleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OraclelistOracleAdded represents a OracleAdded event raised by the Oraclelist contract.
type OraclelistOracleAdded struct {
	TokenAddress  common.Address
	OracleAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOracleAdded is a free log retrieval operation binding the contract event 0x828d2be040dede7698182e08dfa8bfbd663c879aee772509c4a2bd961d0ed43f.
//
// Solidity: event OracleAdded(address _tokenAddress, address _oracleAddress)
func (_Oraclelist *OraclelistFilterer) FilterOracleAdded(opts *bind.FilterOpts) (*OraclelistOracleAddedIterator, error) {

	logs, sub, err := _Oraclelist.contract.FilterLogs(opts, "OracleAdded")
	if err != nil {
		return nil, err
	}
	return &OraclelistOracleAddedIterator{contract: _Oraclelist.contract, event: "OracleAdded", logs: logs, sub: sub}, nil
}

// WatchOracleAdded is a free log subscription operation binding the contract event 0x828d2be040dede7698182e08dfa8bfbd663c879aee772509c4a2bd961d0ed43f.
//
// Solidity: event OracleAdded(address _tokenAddress, address _oracleAddress)
func (_Oraclelist *OraclelistFilterer) WatchOracleAdded(opts *bind.WatchOpts, sink chan<- *OraclelistOracleAdded) (event.Subscription, error) {

	logs, sub, err := _Oraclelist.contract.WatchLogs(opts, "OracleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OraclelistOracleAdded)
				if err := _Oraclelist.contract.UnpackLog(event, "OracleAdded", log); err != nil {
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

// ParseOracleAdded is a log parse operation binding the contract event 0x828d2be040dede7698182e08dfa8bfbd663c879aee772509c4a2bd961d0ed43f.
//
// Solidity: event OracleAdded(address _tokenAddress, address _oracleAddress)
func (_Oraclelist *OraclelistFilterer) ParseOracleAdded(log types.Log) (*OraclelistOracleAdded, error) {
	event := new(OraclelistOracleAdded)
	if err := _Oraclelist.contract.UnpackLog(event, "OracleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OraclelistOracleRemovedIterator is returned from FilterOracleRemoved and is used to iterate over the raw logs and unpacked data for OracleRemoved events raised by the Oraclelist contract.
type OraclelistOracleRemovedIterator struct {
	Event *OraclelistOracleRemoved // Event containing the contract specifics and raw log

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
func (it *OraclelistOracleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OraclelistOracleRemoved)
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
		it.Event = new(OraclelistOracleRemoved)
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
func (it *OraclelistOracleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OraclelistOracleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OraclelistOracleRemoved represents a OracleRemoved event raised by the Oraclelist contract.
type OraclelistOracleRemoved struct {
	TokenAddress  common.Address
	OracleAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOracleRemoved is a free log retrieval operation binding the contract event 0x6dc84b66cc948d847632b9d829f7cb1cb904fbf2c084554a9bc22ad9d8453340.
//
// Solidity: event OracleRemoved(address _tokenAddress, address _oracleAddress)
func (_Oraclelist *OraclelistFilterer) FilterOracleRemoved(opts *bind.FilterOpts) (*OraclelistOracleRemovedIterator, error) {

	logs, sub, err := _Oraclelist.contract.FilterLogs(opts, "OracleRemoved")
	if err != nil {
		return nil, err
	}
	return &OraclelistOracleRemovedIterator{contract: _Oraclelist.contract, event: "OracleRemoved", logs: logs, sub: sub}, nil
}

// WatchOracleRemoved is a free log subscription operation binding the contract event 0x6dc84b66cc948d847632b9d829f7cb1cb904fbf2c084554a9bc22ad9d8453340.
//
// Solidity: event OracleRemoved(address _tokenAddress, address _oracleAddress)
func (_Oraclelist *OraclelistFilterer) WatchOracleRemoved(opts *bind.WatchOpts, sink chan<- *OraclelistOracleRemoved) (event.Subscription, error) {

	logs, sub, err := _Oraclelist.contract.WatchLogs(opts, "OracleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OraclelistOracleRemoved)
				if err := _Oraclelist.contract.UnpackLog(event, "OracleRemoved", log); err != nil {
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

// ParseOracleRemoved is a log parse operation binding the contract event 0x6dc84b66cc948d847632b9d829f7cb1cb904fbf2c084554a9bc22ad9d8453340.
//
// Solidity: event OracleRemoved(address _tokenAddress, address _oracleAddress)
func (_Oraclelist *OraclelistFilterer) ParseOracleRemoved(log types.Log) (*OraclelistOracleRemoved, error) {
	event := new(OraclelistOracleRemoved)
	if err := _Oraclelist.contract.UnpackLog(event, "OracleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OraclelistOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oraclelist contract.
type OraclelistOwnershipTransferredIterator struct {
	Event *OraclelistOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OraclelistOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OraclelistOwnershipTransferred)
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
		it.Event = new(OraclelistOwnershipTransferred)
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
func (it *OraclelistOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OraclelistOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OraclelistOwnershipTransferred represents a OwnershipTransferred event raised by the Oraclelist contract.
type OraclelistOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oraclelist *OraclelistFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OraclelistOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oraclelist.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OraclelistOwnershipTransferredIterator{contract: _Oraclelist.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oraclelist *OraclelistFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OraclelistOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oraclelist.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OraclelistOwnershipTransferred)
				if err := _Oraclelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Oraclelist *OraclelistFilterer) ParseOwnershipTransferred(log types.Log) (*OraclelistOwnershipTransferred, error) {
	event := new(OraclelistOwnershipTransferred)
	if err := _Oraclelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
