package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	alchemyGoerliKey := os.Getenv("ALCHEMY_GOERLI_KEY")
	client, err := ethclient.Dial("wss://eth-goerli.g.alchemy.com/v2/" + alchemyGoerliKey)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x666B0582d5bb8C5CB5f69AdeF438DFE834F80FAf")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
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
			fmt.Println(vLog)
		}
	}
}
