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
	fmt.Printf("Private Key: %s\n", w.PrivateKeyStr())
	fmt.Printf("Public Key: %s\n", w.PublicKeyStr())
	fmt.Printf("Address: %s\n", w.WalletAddress())

	t := wallet.NewTransactionInfo(w.PrivateKey(), w.PublicKey(), w.WalletAddress(), "RecipientAddr", 2.5)
	fmt.Printf("Signature: %s \n", t.SignTransaction())

	bc := blockchain.NewBlockChain("")
	if bc.AddTransaction(t, t.SignTransaction()) {
		bc.Print()
	}

	// bc.AddTransaction("A", "B", 2.53)
	// bc.Mine("Miner Address")

	// bc.AddTransaction("B", "C", 1.4)
	// bc.AddTransaction("C", "A", 1)
	// bc.Mine("Miner Address")

	// fmt.Printf("Balance of Account 'B' is: %.2f\n\n", bc.CalculateAccountBalance("B"))
}