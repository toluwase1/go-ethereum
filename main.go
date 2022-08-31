package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
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
	fmt.Println("The block number", block.Number())

	/*
		converting an address to string balance and then to an ethereum balance
	*/
	addr := "0x829bd824b016326a401d083b33d092293333a830"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("error getting balance %v\n ", err)
		return
	}
	//1 eth = 10 ^ 18 wei
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	fmt.Println(balance)

	balanceEther := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow(10, 18)))
	fmt.Println("balance in ethereum: ", balanceEther)
}
