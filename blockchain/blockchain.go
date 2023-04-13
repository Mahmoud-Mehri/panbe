package blockchain

import (
	"fmt"
	"time"
)

type Block struct {
	nonce int
	previousHash string
	timestamp int64
	transactions []string
}

func NewBlock(nonce int, prevHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = prevHash
	return b;
}

func (b *Block) Print() {
	fmt.Printf("-Nonce          %d\n", b.nonce)
	fmt.Printf("-Timestamp      %d\n", b.timestamp)
	fmt.Printf("-PreviousHash   %s\n", b.previousHash)
	fmt.Printf("-Transactions   %s\n", b.transactions)
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockChain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "First Hash Value");
	return bc
}

func (bc *Blockchain) CreateBlock(Nonce int, PrevHash string) *Block {
	b := NewBlock(Nonce, PrevHash)
	bc.chain = append(bc.chain, b)
	return b
}