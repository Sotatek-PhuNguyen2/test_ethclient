package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var address string = "0xc956eb4869D6212ed9f4419bF35140E4Bc392F9B"

func main() {
	client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/4d7166136e3441beb824859d96bf3e27")
	if err != nil {
		log.Fatal(err)
	}
	var headers = make(chan *types.Header, 10000)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			go sub_transaction(header, client)
		}
	}
}
