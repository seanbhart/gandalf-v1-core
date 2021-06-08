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
var FactoryBin = "0x608060405234801561001057600080fd5b506120c4806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630b262d021461005157806350a94e55146100815780636bcc4f09146100b1578063e4b96982146100cf575b600080fd5b61006b6004803603810190610066919061061a565b6100ff565b6040516100789190610795565b60405180910390f35b61009b600480360381019061009691906105c6565b61013e565b6040516100a89190610795565b60405180910390f35b6100b9610196565b6040516100c6919061085e565b60405180910390f35b6100e960048036038101906100e491906105c6565b6101a3565b6040516100f69190610795565b60405180910390f35b6001818154811061010f57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000602052816000526040600020818051602081018201805184825260208301602085012081835280955050505050506000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600180549050905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610214576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020b9061083e565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083604051610278919061074b565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146102fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f49061081e565b60405180910390fd5b60006040518060200161030f90610526565b6020820181038252601f19601f8201166040525090506000848460405160200161033a929190610762565b604051602081830303815290604052805190602001209050808251602084016000f592508273ffffffffffffffffffffffffffffffffffffffff1663f399e22e86866040518363ffffffff1660e01b81526004016103999291906107b0565b600060405180830381600087803b1580156103b357600080fd5b505af11580156103c7573d6000803e3d6000fd5b50505050826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002085604051610418919061074b565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001839080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508473ffffffffffffffffffffffffffffffffffffffff167f88b91b282690ac7f7f776c9c8ac1986d5a923994d35a501adf869afbf16a739b8585600180549050604051610516939291906107e0565b60405180910390a2505092915050565b6115cf8062000ac083390190565b60006105476105428461089e565b610879565b90508281526020810184848401111561055f57600080fd5b61056a848285610932565b509392505050565b60008135905061058181610a91565b92915050565b600082601f83011261059857600080fd5b81356105a8848260208601610534565b91505092915050565b6000813590506105c081610aa8565b92915050565b600080604083850312156105d957600080fd5b60006105e785828601610572565b925050602083013567ffffffffffffffff81111561060457600080fd5b61061085828601610587565b9150509250929050565b60006020828403121561062c57600080fd5b600061063a848285016105b1565b91505092915050565b61064c816108f6565b82525050565b61066361065e826108f6565b6109a5565b82525050565b6000610674826108cf565b61067e81856108da565b935061068e818560208601610941565b610697816109f8565b840191505092915050565b60006106ad826108cf565b6106b781856108eb565b93506106c7818560208601610941565b80840191505092915050565b60006106e06016836108da565b91506106eb82610a16565b602082019050919050565b60006107036017836108da565b915061070e82610a3f565b602082019050919050565b60006107266009836108eb565b915061073182610a68565b600982019050919050565b61074581610928565b82525050565b600061075782846106a2565b915081905092915050565b600061076d82610719565b91506107798285610652565b60148201915061078982846106a2565b91508190509392505050565b60006020820190506107aa6000830184610643565b92915050565b60006040820190506107c56000830185610643565b81810360208301526107d78184610669565b90509392505050565b600060608201905081810360008301526107fa8186610669565b90506108096020830185610643565b610816604083018461073c565b949350505050565b60006020820190508181036000830152610837816106d3565b9050919050565b60006020820190508181036000830152610857816106f6565b9050919050565b6000602082019050610873600083018461073c565b92915050565b6000610883610894565b905061088f8282610974565b919050565b6000604051905090565b600067ffffffffffffffff8211156108b9576108b86109c9565b5b6108c2826109f8565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600061090182610908565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561095f578082015181840152602081019050610944565b8381111561096e576000848401525b50505050565b61097d826109f8565b810181811067ffffffffffffffff8211171561099c5761099b6109c9565b5b80604052505050565b60006109b0826109b7565b9050919050565b60006109c282610a09565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f47616e64616c6656313a2046554e445f45584953545300000000000000000000600082015250565b7f47616e64616c6656313a205a45524f5f41444452455353000000000000000000600082015250565b7f47616e64616c6656310000000000000000000000000000000000000000000000600082015250565b610a9a816108f6565b8114610aa557600080fd5b50565b610ab181610928565b8114610abc57600080fd5b5056fe6080604052737a250d5630b4cf539739df2c5dacb4c659f2488d600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507307de306ff27a2b630b1141956844eb1552b956b5600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550731f9840a85d5af5bf1d1762f925bdaddc4201f984600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073d0a1e359811322d97991e03f863a0c30c2cf029c600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561016457600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550739a5811ed25f34834e1451869194180fead119dd2600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611300806102cf6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806373d002241161007157806373d002241461015657806386a19e63146101605780639405de641461016a578063b0514a4614610174578063c45a015514610192578063f399e22e146101b0576100a9565b806316345f18146100ae578063481c6a75146100de5780634a79d50c146100fc57806369deb6591461011a5780636c82532714610138575b600080fd5b6100c860048036038101906100c39190610b96565b6101cc565b6040516100d59190610f12565b60405180910390f35b6100e6610285565b6040516100f39190610dec565b60405180910390f35b6101046102ab565b6040516101119190610f2d565b60405180910390f35b610122610339565b60405161012f9190610ed5565b60405180910390f35b6101406103e5565b60405161014d9190610ef7565b60405180910390f35b61015e61040b565b005b610168610706565b005b6101726107bd565b005b61017c610878565b6040516101899190610f6f565b60405180910390f35b61019a61087e565b6040516101a79190610dec565b60405180910390f35b6101ca60048036038101906101c59190610bbf565b6108a2565b005b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316345f18846040518263ffffffff1660e01b815260040161022a9190610dec565b60206040518083038186803b15801561024257600080fd5b505afa158015610256573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061027a9190610c7d565b905080915050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600280546102b890611173565b80601f01602080910402602001604051908101604052809291908181526020018280546102e490611173565b80156103315780601f1061030657610100808354040283529160200191610331565b820191906000526020600020905b81548152906001019060200180831161031457829003601f168201915b505050505081565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ee8c24b86040518163ffffffff1660e01b815260040160006040518083038186803b1580156103a357600080fd5b505afa1580156103b7573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103e09190610c13565b905090565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073d0a1e359811322d97991e03f863a0c30c2cf029c90506000731f9840a85d5af5bf1d1762f925bdaddc4201f9849050600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1667016345785d8a00006040518363ffffffff1660e01b81526004016104c4929190610e67565b602060405180830381600087803b1580156104de57600080fd5b505af11580156104f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105169190610c54565b50600067016345785d8a00009050600067010a741a4627800090506060848160008151811061056e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816001815181106105e3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600030905060006360bd4f989050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c11d79586868686866040518663ffffffff1660e01b815260040161068e959493929190610f8a565b600060405180830381600087803b1580156106a857600080fd5b505af11580156106bc573d6000803e3d6000fd5b505050507f6d54c56b64987faf49976828c0be87cb21e5ce8dbc256e52b2ad6118931c42fe878688876040516106f59493929190610e90565b60405180910390a150505050505050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330620f42406040518463ffffffff1660e01b815260040161076893929190610e30565b602060405180830381600087803b15801561078257600080fd5b505af1158015610796573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ba9190610c54565b50565b6000600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b815260040161081c929190610e07565b60206040518083038186803b15801561083457600080fd5b505afa158015610848573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086c9190610ca6565b905080600a8190555050565b600a5481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610930576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092790610f4f565b60405180910390fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806002908051906020019061098792919061098c565b505050565b82805461099890611173565b90600052602060002090601f0160209004810192826109ba5760008555610a01565b82601f106109d357805160ff1916838001178555610a01565b82800160010185558215610a01579182015b82811115610a005782518255916020019190600101906109e5565b5b509050610a0e9190610a12565b5090565b5b80821115610a2b576000816000905550600101610a13565b5090565b6000610a42610a3d84611009565b610fe4565b90508083825260208201905082856020860282011115610a6157600080fd5b60005b85811015610a915781610a778882610aee565b845260208401935060208301925050600181019050610a64565b5050509392505050565b6000610aae610aa984611035565b610fe4565b905082815260208101848484011115610ac657600080fd5b610ad1848285611131565b509392505050565b600081359050610ae88161126e565b92915050565b600081519050610afd8161126e565b92915050565b600082601f830112610b1457600080fd5b8151610b24848260208601610a2f565b91505092915050565b600081519050610b3c81611285565b92915050565b600081519050610b518161129c565b92915050565b600082601f830112610b6857600080fd5b8135610b78848260208601610a9b565b91505092915050565b600081519050610b90816112b3565b92915050565b600060208284031215610ba857600080fd5b6000610bb684828501610ad9565b91505092915050565b60008060408385031215610bd257600080fd5b6000610be085828601610ad9565b925050602083013567ffffffffffffffff811115610bfd57600080fd5b610c0985828601610b57565b9150509250929050565b600060208284031215610c2557600080fd5b600082015167ffffffffffffffff811115610c3f57600080fd5b610c4b84828501610b03565b91505092915050565b600060208284031215610c6657600080fd5b6000610c7484828501610b2d565b91505092915050565b600060208284031215610c8f57600080fd5b6000610c9d84828501610b42565b91505092915050565b600060208284031215610cb857600080fd5b6000610cc684828501610b81565b91505092915050565b6000610cdb8383610ce7565b60208301905092915050565b610cf0816110bb565b82525050565b610cff816110bb565b82525050565b6000610d1082611076565b610d1a8185611099565b9350610d2583611066565b8060005b83811015610d56578151610d3d8882610ccf565b9750610d488361108c565b925050600181019050610d29565b5085935050505092915050565b610d6c8161110d565b82525050565b610d7b816110d9565b82525050565b6000610d8c82611081565b610d9681856110aa565b9350610da6818560208601611140565b610daf81611234565b840191505092915050565b6000610dc76014836110aa565b9150610dd282611245565b602082019050919050565b610de681611103565b82525050565b6000602082019050610e016000830184610cf6565b92915050565b6000604082019050610e1c6000830185610cf6565b610e296020830184610cf6565b9392505050565b6000606082019050610e456000830186610cf6565b610e526020830185610cf6565b610e5f6040830184610ddd565b949350505050565b6000604082019050610e7c6000830185610cf6565b610e896020830184610ddd565b9392505050565b6000608082019050610ea56000830187610cf6565b610eb26020830186610ddd565b610ebf6040830185610cf6565b610ecc6060830184610ddd565b95945050505050565b60006020820190508181036000830152610eef8184610d05565b905092915050565b6000602082019050610f0c6000830184610d63565b92915050565b6000602082019050610f276000830184610d72565b92915050565b60006020820190508181036000830152610f478184610d81565b905092915050565b60006020820190508181036000830152610f6881610dba565b9050919050565b6000602082019050610f846000830184610ddd565b92915050565b600060a082019050610f9f6000830188610ddd565b610fac6020830187610ddd565b8181036040830152610fbe8186610d05565b9050610fcd6060830185610cf6565b610fda6080830184610ddd565b9695505050505050565b6000610fee610fff565b9050610ffa82826111a5565b919050565b6000604051905090565b600067ffffffffffffffff82111561102457611023611205565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156110505761104f611205565b5b61105982611234565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60006110c6826110e3565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006111188261111f565b9050919050565b600061112a826110e3565b9050919050565b82818337600083830152505050565b60005b8381101561115e578082015181840152602081019050611143565b8381111561116d576000848401525b50505050565b6000600282049050600182168061118b57607f821691505b6020821081141561119f5761119e6111d6565b5b50919050565b6111ae82611234565b810181811067ffffffffffffffff821117156111cd576111cc611205565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f47616e64616c6656313a20464f5242494444454e000000000000000000000000600082015250565b611277816110bb565b811461128257600080fd5b50565b61128e816110cd565b811461129957600080fd5b50565b6112a5816110d9565b81146112b057600080fd5b50565b6112bc81611103565b81146112c757600080fd5b5056fea264697066735822122068d1dbbec69031b907cff56e4746b0771604e69499268270ae8f535d8e5b85b064736f6c63430008040033a26469706673582212200950f37f75aba96aa68d46108beabfd9d26044c43799a4ebf8f4b0231e37d18c64736f6c63430008040033"

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
