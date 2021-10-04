package factory

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/seanbhart/gandalf-v1-core/pkg/tx"
)

func Deploy(client *ethclient.Client, privateKeyHex string, gasLimit uint64) (common.Address, *Factory) {
	auth, err := tx.GetAuth(client, privateKeyHex, gasLimit)
	if err != nil {
		log.Fatal(err)
	}

	address, tx, instance, err := DeployFactory(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	return address, instance
}

func CreateFund(client *ethclient.Client, factoryAddress common.Address, oracleAddress common.Address, gasLimit uint64, privateKeyHex string, manager common.Address, title string) {
	// EventsListen(client, factoryAddress)

	auth, err := tx.GetAuth(client, privateKeyHex, gasLimit)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := NewFactory(factoryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.CreateFund(auth, manager, title, oracleAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func FundListCount(client *ethclient.Client, factoryAddress common.Address) {
	instance, err := NewFactory(factoryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	count, err := instance.FundListLength(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("fund count: %s\n", count)
}

func QueryFundIndex(client *ethclient.Client, factoryAddress common.Address, index *big.Int) {
	instance, err := NewFactory(factoryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	count, err := instance.FundList(nil, index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("fund @ index: %s\n", count)
}

func QueryFund(client *ethclient.Client, factoryAddress common.Address, manager common.Address, title string) {
	instance, err := NewFactory(factoryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.GetFund(nil, manager, title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx response: %v\n", tx)
	fmt.Printf("tx hex: %s\n", tx.Hash().Hex())
}
