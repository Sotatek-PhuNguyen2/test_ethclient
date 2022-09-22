package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/4d7166136e3441beb824859d96bf3e27")
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().String())
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())
		fmt.Println(tx.Value().String())    // 10000000000000000
		fmt.Println(tx.Gas())               // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce())             // 110644
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if msg, err := tx.AsMessage(types.LatestSignerForChainID(big.NewInt(int64(1))), chainID); err == nil {
			fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}
	}
}

/* 	// Grab block by hash then iterate over transactions by index
   	blockHash := common.HexToHash("0x2a875a424a5236d5ae8f3524c86f158abe63499ee6089ca07abd01e5b1257cb1")
   	count, err := client.TransactionCount(context.Background(), blockHash)
   	if err != nil {
   		log.Fatal(err)
   	}

   	for idx := uint(0); idx < count; idx++ {
   		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
   		if err != nil {
   			log.Fatal(err)
   		}

   		fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
   	} */

// Grab a transaction by it's individual hash
/* 	txHash := common.HexToHash("0xae9c3776de9ed6bf0e025704bbeced567b428c78e00330b59c25fe45e7ef87a9")
   	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
   	if err != nil {
   		log.Fatal(err)
   	}

   	fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
   	fmt.Printf("Pending?: %v\n", isPending) */
/* 	account := common.HexToAddress("0xc956eb4869D6212ed9f4419bF35140E4Bc392F9B")
bigInt, err := client.BalanceAt(context.Background(), account, nil)
if err != nil {
	log.Fatalf(err.Error())
}
fmt.Println(bigInt) */
