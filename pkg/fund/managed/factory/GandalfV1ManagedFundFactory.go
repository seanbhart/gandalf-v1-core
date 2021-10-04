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
const FactoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fund\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"FundCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fund\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"FundCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"createFund\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"fund\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fundList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundListLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"getFund\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FactoryBin is the compiled bytecode used for deploying new contracts.
var FactoryBin = "0x608060405234801561001057600080fd5b50612251806100206000396000f3fe60806040523480156200001157600080fd5b5060043610620000525760003560e01c80630b262d02146200005757806350a94e55146200008d5780636bcc4f0914620000c35780639549ed9914620000e5575b600080fd5b6200007560048036038101906200006f9190620006f5565b6200011b565b604051620000849190620008a3565b60405180910390f35b620000ab6004803603810190620000a5919062000614565b6200015b565b604051620000ba9190620008a3565b60405180910390f35b620000cd620001b3565b604051620000dc91906200098c565b60405180910390f35b620001036004803603810190620000fd91906200067a565b620001c0565b604051620001129190620008a3565b60405180910390f35b600181815481106200012c57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000602052816000526040600020818051602081018201805184825260208301602085012081835280955050505050506000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600180549050905090565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141562000234576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200022b906200096a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020846040516200029a919062000851565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161462000322576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003199062000948565b60405180910390fd5b60006040518060200162000336906200055a565b6020820181038252601f19601f82011660405250905060008585604051602001620003639291906200086a565b604051602081830303815290604052805190602001209050808251602084016000f592508273ffffffffffffffffffffffffffffffffffffffff16637bb7c0d88787876040518463ffffffff1660e01b8152600401620003c693929190620008c0565b600060405180830381600087803b158015620003e157600080fd5b505af1158015620003f6573d6000803e3d6000fd5b50505050826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208660405162000449919062000851565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001839080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508573ffffffffffffffffffffffffffffffffffffffff167f88b91b282690ac7f7f776c9c8ac1986d5a923994d35a501adf869afbf16a739b8685600180549050604051620005499392919062000904565b60405180910390a250509392505050565b6115fb8062000c2183390190565b60006200057f6200057984620009d2565b620009a9565b9050828152602081018484840111156200059e576200059d62000b44565b5b620005ab84828562000a6d565b509392505050565b600081359050620005c48162000bec565b92915050565b600082601f830112620005e257620005e162000b3f565b5b8135620005f484826020860162000568565b91505092915050565b6000813590506200060e8162000c06565b92915050565b600080604083850312156200062e576200062d62000b4e565b5b60006200063e85828601620005b3565b925050602083013567ffffffffffffffff81111562000662576200066162000b49565b5b6200067085828601620005ca565b9150509250929050565b60008060006060848603121562000696576200069562000b4e565b5b6000620006a686828701620005b3565b935050602084013567ffffffffffffffff811115620006ca57620006c962000b49565b5b620006d886828701620005ca565b9250506040620006eb86828701620005b3565b9150509250925092565b6000602082840312156200070e576200070d62000b4e565b5b60006200071e84828501620005fd565b91505092915050565b620007328162000a2f565b82525050565b6200074d620007478262000a2f565b62000ae8565b82525050565b6000620007608262000a08565b6200076c818562000a13565b93506200077e81856020860162000a7c565b620007898162000b53565b840191505092915050565b6000620007a18262000a08565b620007ad818562000a24565b9350620007bf81856020860162000a7c565b80840191505092915050565b6000620007da60168362000a13565b9150620007e78262000b71565b602082019050919050565b60006200080160178362000a13565b91506200080e8262000b9a565b602082019050919050565b60006200082860098362000a24565b9150620008358262000bc3565b600982019050919050565b6200084b8162000a63565b82525050565b60006200085f828462000794565b915081905092915050565b6000620008778262000819565b915062000885828562000738565b60148201915062000897828462000794565b91508190509392505050565b6000602082019050620008ba600083018462000727565b92915050565b6000606082019050620008d7600083018662000727565b8181036020830152620008eb818562000753565b9050620008fc604083018462000727565b949350505050565b6000606082019050818103600083015262000920818662000753565b905062000931602083018562000727565b62000940604083018462000840565b949350505050565b600060208201905081810360008301526200096381620007cb565b9050919050565b600060208201905081810360008301526200098581620007f2565b9050919050565b6000602082019050620009a3600083018462000840565b92915050565b6000620009b5620009c8565b9050620009c3828262000ab2565b919050565b6000604051905090565b600067ffffffffffffffff821115620009f057620009ef62000b10565b5b620009fb8262000b53565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600062000a3c8262000a43565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101562000a9c57808201518184015260208101905062000a7f565b8381111562000aac576000848401525b50505050565b62000abd8262000b53565b810181811067ffffffffffffffff8211171562000adf5762000ade62000b10565b5b80604052505050565b600062000af58262000afc565b9050919050565b600062000b098262000b64565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f47616e64616c6656313a2046554e445f45584953545300000000000000000000600082015250565b7f47616e64616c6656313a205a45524f5f41444452455353000000000000000000600082015250565b7f47616e64616c6656310000000000000000000000000000000000000000000000600082015250565b62000bf78162000a2f565b811462000c0357600080fd5b50565b62000c118162000a63565b811462000c1d57600080fd5b5056fe6080604052737a250d5630b4cf539739df2c5dacb4c659f2488d600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507307de306ff27a2b630b1141956844eb1552b956b5600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550731f9840a85d5af5bf1d1762f925bdaddc4201f984600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073d0a1e359811322d97991e03f863a0c30c2cf029c600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561016457600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506113818061027a6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806373d002241161007157806373d00224146101565780637bb7c0d81461016057806386a19e631461017c5780639405de6414610186578063b0514a4614610190578063c45a0155146101ae576100a9565b806316345f18146100ae578063481c6a75146100de5780634a79d50c146100fc57806369deb6591461011a5780636c82532714610138575b600080fd5b6100c860048036038101906100c39190610b9c565b6101cc565b6040516100d59190610f4b565b60405180910390f35b6100e6610285565b6040516100f39190610e25565b60405180910390f35b6101046102ab565b6040516101119190610f66565b60405180910390f35b610122610339565b60405161012f9190610f0e565b60405180910390f35b6101406103e5565b60405161014d9190610f30565b60405180910390f35b61015e61040b565b005b61017a60048036038101906101759190610bc9565b6106ba565b005b6101846107e6565b005b61018e61089d565b005b610198610958565b6040516101a59190610fa8565b60405180910390f35b6101b661095e565b6040516101c39190610e25565b60405180910390f35b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316345f18846040518263ffffffff1660e01b815260040161022a9190610e25565b60206040518083038186803b15801561024257600080fd5b505afa158015610256573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061027a9190610cae565b905080915050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600280546102b8906111ac565b80601f01602080910402602001604051908101604052809291908181526020018280546102e4906111ac565b80156103315780601f1061030657610100808354040283529160200191610331565b820191906000526020600020905b81548152906001019060200180831161031457829003601f168201915b505050505081565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ee8c24b86040518163ffffffff1660e01b815260040160006040518083038186803b1580156103a357600080fd5b505afa1580156103b7573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103e09190610c38565b905090565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073d0a1e359811322d97991e03f863a0c30c2cf029c90506000731f9840a85d5af5bf1d1762f925bdaddc4201f9849050600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1667016345785d8a00006040518363ffffffff1660e01b81526004016104c4929190610ea0565b602060405180830381600087803b1580156104de57600080fd5b505af11580156104f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105169190610c81565b50600067016345785d8a00009050600067010a741a462780009050606084816000815181106105485761054761123e565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816001815181106105975761059661123e565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600030905060006360bd4f989050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c11d79586868686866040518663ffffffff1660e01b8152600401610642959493929190610fc3565b600060405180830381600087803b15801561065c57600080fd5b505af1158015610670573d6000803e3d6000fd5b505050507f6d54c56b64987faf49976828c0be87cb21e5ce8dbc256e52b2ad6118931c42fe878688876040516106a99493929190610ec9565b60405180910390a150505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610748576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073f90610f88565b60405180910390fd5b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816002908051906020019061079f929190610982565b5080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330620f42406040518463ffffffff1660e01b815260040161084893929190610e69565b602060405180830381600087803b15801561086257600080fd5b505af1158015610876573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089a9190610c81565b50565b6000600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b81526004016108fc929190610e40565b60206040518083038186803b15801561091457600080fd5b505afa158015610928573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094c9190610cdb565b905080600a8190555050565b600a5481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b82805461098e906111ac565b90600052602060002090601f0160209004810192826109b057600085556109f7565b82601f106109c957805160ff19168380011785556109f7565b828001600101855582156109f7579182015b828111156109f65782518255916020019190600101906109db565b5b509050610a049190610a08565b5090565b5b80821115610a21576000816000905550600101610a09565b5090565b6000610a38610a3384611042565b61101d565b90508083825260208201905082856020860282011115610a5b57610a5a6112a1565b5b60005b85811015610a8b5781610a718882610aec565b845260208401935060208301925050600181019050610a5e565b5050509392505050565b6000610aa8610aa38461106e565b61101d565b905082815260208101848484011115610ac457610ac36112a6565b5b610acf84828561116a565b509392505050565b600081359050610ae6816112ef565b92915050565b600081519050610afb816112ef565b92915050565b600082601f830112610b1657610b1561129c565b5b8151610b26848260208601610a25565b91505092915050565b600081519050610b3e81611306565b92915050565b600081519050610b538161131d565b92915050565b600082601f830112610b6e57610b6d61129c565b5b8135610b7e848260208601610a95565b91505092915050565b600081519050610b9681611334565b92915050565b600060208284031215610bb257610bb16112b0565b5b6000610bc084828501610ad7565b91505092915050565b600080600060608486031215610be257610be16112b0565b5b6000610bf086828701610ad7565b935050602084013567ffffffffffffffff811115610c1157610c106112ab565b5b610c1d86828701610b59565b9250506040610c2e86828701610ad7565b9150509250925092565b600060208284031215610c4e57610c4d6112b0565b5b600082015167ffffffffffffffff811115610c6c57610c6b6112ab565b5b610c7884828501610b01565b91505092915050565b600060208284031215610c9757610c966112b0565b5b6000610ca584828501610b2f565b91505092915050565b600060208284031215610cc457610cc36112b0565b5b6000610cd284828501610b44565b91505092915050565b600060208284031215610cf157610cf06112b0565b5b6000610cff84828501610b87565b91505092915050565b6000610d148383610d20565b60208301905092915050565b610d29816110f4565b82525050565b610d38816110f4565b82525050565b6000610d49826110af565b610d5381856110d2565b9350610d5e8361109f565b8060005b83811015610d8f578151610d768882610d08565b9750610d81836110c5565b925050600181019050610d62565b5085935050505092915050565b610da581611146565b82525050565b610db481611112565b82525050565b6000610dc5826110ba565b610dcf81856110e3565b9350610ddf818560208601611179565b610de8816112b5565b840191505092915050565b6000610e006014836110e3565b9150610e0b826112c6565b602082019050919050565b610e1f8161113c565b82525050565b6000602082019050610e3a6000830184610d2f565b92915050565b6000604082019050610e556000830185610d2f565b610e626020830184610d2f565b9392505050565b6000606082019050610e7e6000830186610d2f565b610e8b6020830185610d2f565b610e986040830184610e16565b949350505050565b6000604082019050610eb56000830185610d2f565b610ec26020830184610e16565b9392505050565b6000608082019050610ede6000830187610d2f565b610eeb6020830186610e16565b610ef86040830185610d2f565b610f056060830184610e16565b95945050505050565b60006020820190508181036000830152610f288184610d3e565b905092915050565b6000602082019050610f456000830184610d9c565b92915050565b6000602082019050610f606000830184610dab565b92915050565b60006020820190508181036000830152610f808184610dba565b905092915050565b60006020820190508181036000830152610fa181610df3565b9050919050565b6000602082019050610fbd6000830184610e16565b92915050565b600060a082019050610fd86000830188610e16565b610fe56020830187610e16565b8181036040830152610ff78186610d3e565b90506110066060830185610d2f565b6110136080830184610e16565b9695505050505050565b6000611027611038565b905061103382826111de565b919050565b6000604051905090565b600067ffffffffffffffff82111561105d5761105c61126d565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156110895761108861126d565b5b611092826112b5565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60006110ff8261111c565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061115182611158565b9050919050565b60006111638261111c565b9050919050565b82818337600083830152505050565b60005b8381101561119757808201518184015260208101905061117c565b838111156111a6576000848401525b50505050565b600060028204905060018216806111c457607f821691505b602082108114156111d8576111d761120f565b5b50919050565b6111e7826112b5565b810181811067ffffffffffffffff821117156112065761120561126d565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f47616e64616c6656313a20464f5242494444454e000000000000000000000000600082015250565b6112f8816110f4565b811461130357600080fd5b50565b61130f81611106565b811461131a57600080fd5b50565b61132681611112565b811461133157600080fd5b50565b61133d8161113c565b811461134857600080fd5b5056fea2646970667358221220d61a9970298f1db37f6014b5852c223ddac78236ec0e0b42757d32f0a8b7f63464736f6c63430008060033a2646970667358221220610e5d37cc54bc555932384f89b8181e25b63b14c1f82b8ff8109a81cbc6613964736f6c63430008060033"

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

// CreateFund is a paid mutator transaction binding the contract method 0x9549ed99.
//
// Solidity: function createFund(address manager, string title, address oracleAddress) returns(address fund)
func (_Factory *FactoryTransactor) CreateFund(opts *bind.TransactOpts, manager common.Address, title string, oracleAddress common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createFund", manager, title, oracleAddress)
}

// CreateFund is a paid mutator transaction binding the contract method 0x9549ed99.
//
// Solidity: function createFund(address manager, string title, address oracleAddress) returns(address fund)
func (_Factory *FactorySession) CreateFund(manager common.Address, title string, oracleAddress common.Address) (*types.Transaction, error) {
	return _Factory.Contract.CreateFund(&_Factory.TransactOpts, manager, title, oracleAddress)
}

// CreateFund is a paid mutator transaction binding the contract method 0x9549ed99.
//
// Solidity: function createFund(address manager, string title, address oracleAddress) returns(address fund)
func (_Factory *FactoryTransactorSession) CreateFund(manager common.Address, title string, oracleAddress common.Address) (*types.Transaction, error) {
	return _Factory.Contract.CreateFund(&_Factory.TransactOpts, manager, title, oracleAddress)
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
