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
