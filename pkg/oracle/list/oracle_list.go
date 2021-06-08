package oraclelist

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/seanbhart/gandalf-v1-core/pkg/tx"
)

func Deploy(client *ethclient.Client, privateKeyHex string, gasLimit uint64) (common.Address, *Oraclelist) {
	auth, err := tx.GetAuth(client, privateKeyHex, gasLimit)
	if err != nil {
		log.Fatal(err)
	}

	address, tx, instance, err := DeployOraclelist(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	return address, instance
}

func GetTokenAddresses(client *ethclient.Client, oracleListAddress common.Address) {
	instance, err := NewOraclelist(oracleListAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	addresses, err := instance.GetTokenAddresses(nil)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(addresses); i++ {
		fmt.Printf("address: %s\n", addresses[i].Hash().Hex())
	}
}

func GetLatestPrice(client *ethclient.Client, oracleListAddress common.Address, tokenAddress common.Address) {
	instance, err := NewOraclelist(oracleListAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	px, err := instance.GetLatestPrice(nil, tokenAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("latest price: %d\n", px)
}
