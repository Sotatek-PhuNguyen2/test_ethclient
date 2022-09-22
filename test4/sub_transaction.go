package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func sub_transaction(header *types.Header, client *ethclient.Client) {
	//client1, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/4d7166136e3441beb824859d96bf3e27")
	// block, err := client.BlockByHash(context.Background(), header.Hash())
	block, err := client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		log.Fatal(err)
	}
	sub_transaction1(block)
}

func sub_transaction1(block *types.Block) {
	//fmt.Println(block)
	for _, tx := range block.Transactions() {
		//fmt.Println(tx.Hash().Hex())
		if msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), tx.GasPrice()); err == nil {
			//fmt.Println("From: ", msg.From().Hex())
			if msg.From().Hex() == address {
				fmt.Println("Send Money !!!")
				fmt.Println("Address from: " + msg.From().Hex())
				if tx.To() == nil {
					fmt.Println("Address to: <nil>")
				} else {
					fmt.Println("Address to: " + tx.To().Hex())
				}
				fmt.Println("Tx hash: " + tx.Hash().Hex())
				fmt.Println("Block numvber: " + block.Number().String())
				fmt.Println("Value: " + tx.Value().String())
			} else if tx.To() != nil {
				if tx.To().Hex() == address {
					fmt.Println("Recvice Money !!!")
					fmt.Println("Address from: " + msg.From().Hex())
					fmt.Println("Address to: " + tx.To().Hex())
					fmt.Println("Tx hash: " + tx.Hash().Hex())
					fmt.Println("Block numvber: " + block.Number().String())
					fmt.Println("Value: " + (tx.Value()).String())
				}
			}
		} /* else {
			fmt.Println(nil)
		}
		if tx.To() == nil {
			fmt.Println(nil)
		} else {
			fmt.Println("To: " + tx.To().Hex())
		} */
	}
}
