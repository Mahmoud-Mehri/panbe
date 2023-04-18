package main

import (
	"fmt"
	"log"
	"panbe/blockchain"
	"panbe/wallet"
)



func init() {
	// ...
	fmt.Println("Blockchain Application Initialized");
	log.SetPrefix("Panbe - ")
}

func main() {
	w := wallet.NewWallet();
	fmt.Printf("Private Key: %s\n\n", w.PrivateKeyStr())

	bc := blockchain.NewBlockChain();
	bc.AddTransaction("A", "B", 2.53)
	bc.Mine("Miner Address")

	bc.AddTransaction("B", "C", 1.4)
	bc.AddTransaction("C", "A", 1)
	bc.Mine("Miner Address")
	bc.Print()

	fmt.Printf("Balance of Account 'B' is: %.2f\n\n", bc.CalculateAccountBalance("B"))
}