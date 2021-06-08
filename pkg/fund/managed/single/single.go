package single

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTitle(client *ethclient.Client, fundAddress common.Address) {
	instance, err := NewSingle(fundAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	title, err := instance.Title(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("fund title: %s\n", title)
}

func GetLatestPrice(client *ethclient.Client, fundAddress common.Address, tokenAddress common.Address) {
	instance, err := NewSingle(fundAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	title, err := instance.Title(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("fund title: %s\n", title)

	fmt.Printf("fund query latest price for token: %s\n", tokenAddress)
	price, err := instance.GetLatestPrice(nil, tokenAddress)
	if err != nil {
		log.Fatalf("get price error: %v\n", err)
	}
	fmt.Printf("price: %d\n", price)
}

func GetTokenAddresses(client *ethclient.Client, fundAddress common.Address) {
	instance, err := NewSingle(fundAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	title, err := instance.Title(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("fund title: %s\n", title)

	addresses, err := instance.GetOracleTokens(nil)
	if err != nil {
		log.Fatalf("get addresses error: %v\n", err)
	}
	fmt.Printf("addresses: %v\n", addresses)
}

func SwapTokens(client *ethclient.Client, fundAddress common.Address) {
	instance, err := NewSingle(fundAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.SwapTokens(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("swap tx: %s\n", tx.Hash().Hex())
}

func FundTest(client *ethclient.Client, fundAddress common.Address) {
	instance, err := NewSingle(fundAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.CheckAllowance(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("swap tx: %s\n", tx.Hash().Hex())
}
