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
const SingleABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenSwap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"checkAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrowTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracleTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inTokenAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"title\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SingleBin is the compiled bytecode used for deploying new contracts.
var SingleBin = "0x6080604052737a250d5630b4cf539739df2c5dacb4c659f2488d600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507307de306ff27a2b630b1141956844eb1552b956b5600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550731f9840a85d5af5bf1d1762f925bdaddc4201f984600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073d0a1e359811322d97991e03f863a0c30c2cf029c600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561016457600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506113818061027a6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806373d002241161007157806373d00224146101565780637bb7c0d81461016057806386a19e631461017c5780639405de6414610186578063b0514a4614610190578063c45a0155146101ae576100a9565b806316345f18146100ae578063481c6a75146100de5780634a79d50c146100fc57806369deb6591461011a5780636c82532714610138575b600080fd5b6100c860048036038101906100c39190610b9c565b6101cc565b6040516100d59190610f4b565b60405180910390f35b6100e6610285565b6040516100f39190610e25565b60405180910390f35b6101046102ab565b6040516101119190610f66565b60405180910390f35b610122610339565b60405161012f9190610f0e565b60405180910390f35b6101406103e5565b60405161014d9190610f30565b60405180910390f35b61015e61040b565b005b61017a60048036038101906101759190610bc9565b6106ba565b005b6101846107e6565b005b61018e61089d565b005b610198610958565b6040516101a59190610fa8565b60405180910390f35b6101b661095e565b6040516101c39190610e25565b60405180910390f35b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316345f18846040518263ffffffff1660e01b815260040161022a9190610e25565b60206040518083038186803b15801561024257600080fd5b505afa158015610256573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061027a9190610cae565b905080915050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600280546102b8906111ac565b80601f01602080910402602001604051908101604052809291908181526020018280546102e4906111ac565b80156103315780601f1061030657610100808354040283529160200191610331565b820191906000526020600020905b81548152906001019060200180831161031457829003601f168201915b505050505081565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ee8c24b86040518163ffffffff1660e01b815260040160006040518083038186803b1580156103a357600080fd5b505afa1580156103b7573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103e09190610c38565b905090565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073d0a1e359811322d97991e03f863a0c30c2cf029c90506000731f9840a85d5af5bf1d1762f925bdaddc4201f9849050600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1667016345785d8a00006040518363ffffffff1660e01b81526004016104c4929190610ea0565b602060405180830381600087803b1580156104de57600080fd5b505af11580156104f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105169190610c81565b50600067016345785d8a00009050600067010a741a462780009050606084816000815181106105485761054761123e565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816001815181106105975761059661123e565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600030905060006360bd4f989050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c11d79586868686866040518663ffffffff1660e01b8152600401610642959493929190610fc3565b600060405180830381600087803b15801561065c57600080fd5b505af1158015610670573d6000803e3d6000fd5b505050507f6d54c56b64987faf49976828c0be87cb21e5ce8dbc256e52b2ad6118931c42fe878688876040516106a99493929190610ec9565b60405180910390a150505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610748576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073f90610f88565b60405180910390fd5b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816002908051906020019061079f929190610982565b5080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330620f42406040518463ffffffff1660e01b815260040161084893929190610e69565b602060405180830381600087803b15801561086257600080fd5b505af1158015610876573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089a9190610c81565b50565b6000600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b81526004016108fc929190610e40565b60206040518083038186803b15801561091457600080fd5b505afa158015610928573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094c9190610cdb565b905080600a8190555050565b600a5481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b82805461098e906111ac565b90600052602060002090601f0160209004810192826109b057600085556109f7565b82601f106109c957805160ff19168380011785556109f7565b828001600101855582156109f7579182015b828111156109f65782518255916020019190600101906109db565b5b509050610a049190610a08565b5090565b5b80821115610a21576000816000905550600101610a09565b5090565b6000610a38610a3384611042565b61101d565b90508083825260208201905082856020860282011115610a5b57610a5a6112a1565b5b60005b85811015610a8b5781610a718882610aec565b845260208401935060208301925050600181019050610a5e565b5050509392505050565b6000610aa8610aa38461106e565b61101d565b905082815260208101848484011115610ac457610ac36112a6565b5b610acf84828561116a565b509392505050565b600081359050610ae6816112ef565b92915050565b600081519050610afb816112ef565b92915050565b600082601f830112610b1657610b1561129c565b5b8151610b26848260208601610a25565b91505092915050565b600081519050610b3e81611306565b92915050565b600081519050610b538161131d565b92915050565b600082601f830112610b6e57610b6d61129c565b5b8135610b7e848260208601610a95565b91505092915050565b600081519050610b9681611334565b92915050565b600060208284031215610bb257610bb16112b0565b5b6000610bc084828501610ad7565b91505092915050565b600080600060608486031215610be257610be16112b0565b5b6000610bf086828701610ad7565b935050602084013567ffffffffffffffff811115610c1157610c106112ab565b5b610c1d86828701610b59565b9250506040610c2e86828701610ad7565b9150509250925092565b600060208284031215610c4e57610c4d6112b0565b5b600082015167ffffffffffffffff811115610c6c57610c6b6112ab565b5b610c7884828501610b01565b91505092915050565b600060208284031215610c9757610c966112b0565b5b6000610ca584828501610b2f565b91505092915050565b600060208284031215610cc457610cc36112b0565b5b6000610cd284828501610b44565b91505092915050565b600060208284031215610cf157610cf06112b0565b5b6000610cff84828501610b87565b91505092915050565b6000610d148383610d20565b60208301905092915050565b610d29816110f4565b82525050565b610d38816110f4565b82525050565b6000610d49826110af565b610d5381856110d2565b9350610d5e8361109f565b8060005b83811015610d8f578151610d768882610d08565b9750610d81836110c5565b925050600181019050610d62565b5085935050505092915050565b610da581611146565b82525050565b610db481611112565b82525050565b6000610dc5826110ba565b610dcf81856110e3565b9350610ddf818560208601611179565b610de8816112b5565b840191505092915050565b6000610e006014836110e3565b9150610e0b826112c6565b602082019050919050565b610e1f8161113c565b82525050565b6000602082019050610e3a6000830184610d2f565b92915050565b6000604082019050610e556000830185610d2f565b610e626020830184610d2f565b9392505050565b6000606082019050610e7e6000830186610d2f565b610e8b6020830185610d2f565b610e986040830184610e16565b949350505050565b6000604082019050610eb56000830185610d2f565b610ec26020830184610e16565b9392505050565b6000608082019050610ede6000830187610d2f565b610eeb6020830186610e16565b610ef86040830185610d2f565b610f056060830184610e16565b95945050505050565b60006020820190508181036000830152610f288184610d3e565b905092915050565b6000602082019050610f456000830184610d9c565b92915050565b6000602082019050610f606000830184610dab565b92915050565b60006020820190508181036000830152610f808184610dba565b905092915050565b60006020820190508181036000830152610fa181610df3565b9050919050565b6000602082019050610fbd6000830184610e16565b92915050565b600060a082019050610fd86000830188610e16565b610fe56020830187610e16565b8181036040830152610ff78186610d3e565b90506110066060830185610d2f565b6110136080830184610e16565b9695505050505050565b6000611027611038565b905061103382826111de565b919050565b6000604051905090565b600067ffffffffffffffff82111561105d5761105c61126d565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156110895761108861126d565b5b611092826112b5565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60006110ff8261111c565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061115182611158565b9050919050565b60006111638261111c565b9050919050565b82818337600083830152505050565b60005b8381101561119757808201518184015260208101905061117c565b838111156111a6576000848401525b50505050565b600060028204905060018216806111c457607f821691505b602082108114156111d8576111d761120f565b5b50919050565b6111e7826112b5565b810181811067ffffffffffffffff821117156112065761120561126d565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f47616e64616c6656313a20464f5242494444454e000000000000000000000000600082015250565b6112f8816110f4565b811461130357600080fd5b50565b61130f81611106565b811461131a57600080fd5b50565b61132681611112565b811461133157600080fd5b50565b61133d8161113c565b811461134857600080fd5b5056fea2646970667358221220d61a9970298f1db37f6014b5852c223ddac78236ec0e0b42757d32f0a8b7f63464736f6c63430008060033"

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

// Initialize is a paid mutator transaction binding the contract method 0x7bb7c0d8.
//
// Solidity: function initialize(address _manager, string _title, address _oracleAddress) returns()
func (_Single *SingleTransactor) Initialize(opts *bind.TransactOpts, _manager common.Address, _title string, _oracleAddress common.Address) (*types.Transaction, error) {
	return _Single.contract.Transact(opts, "initialize", _manager, _title, _oracleAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x7bb7c0d8.
//
// Solidity: function initialize(address _manager, string _title, address _oracleAddress) returns()
func (_Single *SingleSession) Initialize(_manager common.Address, _title string, _oracleAddress common.Address) (*types.Transaction, error) {
	return _Single.Contract.Initialize(&_Single.TransactOpts, _manager, _title, _oracleAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x7bb7c0d8.
//
// Solidity: function initialize(address _manager, string _title, address _oracleAddress) returns()
func (_Single *SingleTransactorSession) Initialize(_manager common.Address, _title string, _oracleAddress common.Address) (*types.Transaction, error) {
	return _Single.Contract.Initialize(&_Single.TransactOpts, _manager, _title, _oracleAddress)
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
