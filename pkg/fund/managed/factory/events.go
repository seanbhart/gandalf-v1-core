package factory

import (
	"context"
	"fmt"
	"log"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func EventsListen(client *ethclient.Client, factoryAddress common.Address) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{factoryAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}

func EventsQuery(client *ethclient.Client, factoryAddress common.Address) {
	query := ethereum.FilterQuery{
		// FromBlock: big.NewInt(2394201),
		// ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			factoryAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(FactoryABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Printf("Blockhash: %v\n", vLog.BlockHash.Hex())
		fmt.Printf("BlockNumber: %v\n", vLog.BlockNumber)
		fmt.Printf("TxHash: %v\n", vLog.TxHash.Hex())

		event, err := contractAbi.Unpack("FundCreated", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("event: %v\n", event)

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Printf("topics: %v\n", topics)
	}

	eventSignature := []byte("FundCreated(address,string,address,uint256)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Printf("signature: %v\n", hash.Hex())
}
