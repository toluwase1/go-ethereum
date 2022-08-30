package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var infuraURL string = "https://mainnet.infura.io/v3/753ced8554424d2bb4cae5b3f70d48c9"
var ganacheUrl = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("error creating an ether client: %v", err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("error getting a block: %v", err)
	}
	fmt.Println(block.Number())
}
