package main

import (
	"fmt"
	"panbe/blockchain/blockchain"
)

func init() {
	// ...
	fmt.Println("Blockchain Application Initialized");
}

func main() {
	bc := blockchain.NewBlockChain();
	fmt.Println("Chain`s length is : %d", len(bc.chain))
}