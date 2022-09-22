package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/4d7166136e3441beb824859d96bf3e27"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	}
	defer client.Close()
	hash := common.HexToHash("0xa9f2ecc010e6c85b9aba55a86669a01e02a3231a4019e5638a5e106a0d02ae68")
	block, _, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	}

	fmt.Println(block)
}
