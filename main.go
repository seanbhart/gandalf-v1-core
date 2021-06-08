package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/seanbhart/gandalf-v1-core/pkg/tx"
	"github.com/seanbhart/gandalf-v1-core/pkg/utils"

	mfactory "github.com/seanbhart/gandalf-v1-core/pkg/fund/managed/factory"
	msingle "github.com/seanbhart/gandalf-v1-core/pkg/fund/managed/single"
	olist "github.com/seanbhart/gandalf-v1-core/pkg/oracle/list"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "COMMAND", "[INPUT]")
		return
	}
	command := os.Args[1]

	client, err := ethclient.Dial(utils.GetEnv("NETWORK"))
	if err != nil {
		log.Fatal(err)
	}

	senderKey := utils.GetEnv("ACCOUNT_KEY_PRIV")
	senderAddress := utils.GetAddress(senderKey)
	factoryAddress := common.HexToAddress(utils.GetEnv("FACTORY_ADDRESS"))
	oracleAddress := common.HexToAddress(utils.GetEnv("ORACLE_ADDRESS"))

	switch command {
	case "factorydeploy":
		if len(os.Args) < 3 {
			fmt.Println("Usage:", os.Args[0], "factorydeploy", "[gas limit]")
			return
		}
		gasLimit, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		mfactory.Deploy(client, senderKey, uint64(gasLimit))

	case "factorygetevents":
		mfactory.EventsQuery(client, factoryAddress)

	case "factoryeventslisten":
		mfactory.EventsListen(client, factoryAddress)

	case "fundcreate":
		if len(os.Args) < 4 {
			fmt.Println("Usage:", os.Args[0], "fundcreate", "[fund title]", "[gas limit]")
			return
		}
		title := os.Args[2]
		gasLimit, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		mfactory.CreateFund(client, factoryAddress, uint64(gasLimit), senderKey, senderAddress, title)

	case "fundcount":
		mfactory.FundListCount(client, factoryAddress)

	case "fundqueryindex":
		if len(os.Args) < 3 {
			fmt.Println("Usage:", os.Args[0], "fundqueryindex", "[fund list index]")
			return
		}
		index := new(big.Int)
		index, ok := index.SetString(os.Args[2], 10)
		if !ok {
			fmt.Println("SetString: error with [fund list index]")
			return
		}
		mfactory.QueryFundIndex(client, factoryAddress, index)

	case "fundqueryname":
		if len(os.Args) < 3 {
			fmt.Println("Usage:", os.Args[0], "fundqueryname", "[fund name]")
			return
		}
		name := os.Args[2]
		mfactory.QueryFund(client, factoryAddress, senderAddress, name)

	case "fundgettitle":
		if len(os.Args) < 3 {
			fmt.Println("Usage:", os.Args[0], "fundgettitle", "[fund address]")
			return
		}
		fundAddress := common.HexToAddress(os.Args[2])
		msingle.GetTitle(client, fundAddress)

	case "fundgetlatestprice":
		if len(os.Args) < 4 {
			fmt.Println("Usage:", os.Args[0], "fundgetlatestprice", "[fund address]", "[token address]")
			return
		}
		fundAddress := common.HexToAddress(os.Args[2])
		tokenAddress := common.HexToAddress(os.Args[3])
		msingle.GetLatestPrice(client, fundAddress, tokenAddress)

	case "fundgetoracleaddresses":
		if len(os.Args) < 3 {
			fmt.Println("Usage:", os.Args[0], "fundgetoracleaddresses", "[fund address]")
			return
		}
		fundAddress := common.HexToAddress(os.Args[2])
		msingle.GetTokenAddresses(client, fundAddress)

	case "fundswap":
		if len(os.Args) < 3 {
			fmt.Println("Usage:", os.Args[0], "fundswap", "[fund address]")
			return
		}
		fundAddress := common.HexToAddress(os.Args[2])
		msingle.SwapTokens(client, fundAddress)

	case "fundtest":
		if len(os.Args) < 3 {
			fmt.Println("Usage:", os.Args[0], "fundtest", "[fund address]")
			return
		}
		fundAddress := common.HexToAddress(os.Args[2])
		msingle.FundTest(client, fundAddress)

	case "oracledeploy":
		if len(os.Args) < 3 {
			fmt.Println("Usage:", os.Args[0], "oracledeploy", "[gas limit]")
			return
		}
		gasLimit, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		olist.Deploy(client, senderKey, uint64(gasLimit))

	case "oraclelatestprice":
		if len(os.Args) < 3 {
			fmt.Println("Usage:", os.Args[0], "oraclelatestprice", "[token address]")
			return
		}
		tokenAddress := common.HexToAddress(os.Args[2])
		olist.GetLatestPrice(client, oracleAddress, tokenAddress)

	case "sendeth":
		if len(os.Args) < 5 {
			fmt.Println("Usage:", os.Args[0], "sendeth", "[to address]", "[value]", "[gas limit]")
			return
		}
		toAddress := common.HexToAddress(os.Args[2])
		value := new(big.Int)
		value, ok := value.SetString(os.Args[3], 10)
		if !ok {
			fmt.Println("SetString: error with [value]")
			return
		}
		gasLimit, err := strconv.Atoi(os.Args[4])
		if err != nil {
			log.Fatal(err)
		}
		tx.SendEth(client, toAddress, value, senderKey, uint64(gasLimit))
	}
}
