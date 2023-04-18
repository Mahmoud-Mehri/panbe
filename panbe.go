package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
	// "panbe/blockchain/blockchain"
)

const (
	NETWORK_DIFFICULTY = 2
	BLOCKCHAIN_ADDRESS = "FFFFFF";
	TRANSACTION_FEE = 0.01;
	MINING_REWARD = 0.01;
)


type Block struct {
	nonce int
	timestamp int64
	previousHash [32]byte
	transactions []*Transaction
}

func NewBlock(nonce int, prevHash [32]byte, trans []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = prevHash
	b.transactions = trans
	return b;
}

func (b *Block) ToJSON() ([]byte, error) {
	return json.Marshal(struct{
		// Timestamp int64 `json:"timestamp"`
		Nonce int `json:"nonce"`
		PreviousHash [32]byte `json:"previousHash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		// Timestamp: b.timestamp,
		Nonce: b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

func (b *Block) Hash() [32]byte {
	data, _ := b.ToJSON()
	// fmt.Printf("Block Object: %s\n", string(data))
	return sha256.Sum256(data)
}

func (b *Block) Print() {
	fmt.Printf("-Nonce          %d\n", b.nonce)
	fmt.Printf("-Timestamp      %d\n", b.timestamp)
	fmt.Printf("-PreviousHash   %x\n", b.previousHash)
	fmt.Printf("-Hash           %x\n", b.Hash())
	fmt.Printf("-Transactions:\n")
	for _, t := range b.transactions {
		t.Print()
	}
}

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlockChain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.chain = append(bc.chain, b);
	return bc
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain) - 1]
}

func (bc *Blockchain) CreateBlock(Nonce int) *Block {
	b := NewBlock(Nonce, bc.LastBlock().Hash(), bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{};
	return b
}

func (bc *Blockchain) Print() {
	for i, b := range bc.chain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25),  i, strings.Repeat("=", 25))
		b.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 50))
}

func (bc *Blockchain) AddTransaction(sender, recipient string, value float32) *Transaction {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
	return t
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	tranPool := make([]*Transaction, 0);
	for _, t := range bc.transactionPool {
		tranPool = append(tranPool, NewTransaction(t.senderAddress, t.recipientAddress, t.value))
	}
	return tranPool;
}

func (bc *Blockchain) Validate(nonce int, prevHash [32]byte, trans []*Transaction, difficulty int) bool {
	sign := strings.Repeat("0", difficulty)
	newBlock := &Block{nonce, 0, prevHash, trans}
	newHash := fmt.Sprintf("%s", newBlock.Hash());
	return newHash[:difficulty] == sign
}

func (bc *Blockchain) ProofOfWork() int {
	nonce := 0;
	prevHash := bc.LastBlock().Hash()
	trans := bc.CopyTransactionPool()
	for !bc.Validate(nonce, prevHash, trans, NETWORK_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) Mine(minerAddress string) {
	bc.AddTransaction(BLOCKCHAIN_ADDRESS, minerAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	bc.CreateBlock(nonce)

}

type Transaction struct {
	senderAddress string
	recipientAddress string
	value float32
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{
		sender,
		recipient,
		value,
	}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	fmt.Printf("   Sender Address     %s \n", t.senderAddress)
	fmt.Printf("   Recipient Address  %s \n", t.recipientAddress)
	fmt.Printf("   Value              %.2f\n", t.value)
}

func (t *Transaction) ToJSON() ([]byte, error) {
	return json.Marshal(struct{
		SenderAddress string `json:"senderAddress"`
		RecipientAddress string `json:"recipientAddress"`
		Value float32 `json:"value"`
	}{
		t.senderAddress,
		t.recipientAddress,
		t.value,
	})
}

func init() {
	// ...
	fmt.Println("Blockchain Application Initialized");
	log.SetPrefix("Panbe - ")
}

func main() {
	bc := NewBlockChain();
	bc.AddTransaction("A", "B", 2.53)
	nonce := bc.ProofOfWork();
	bc.CreateBlock(nonce)

	bc.AddTransaction("B", "C", 6.4)
	bc.AddTransaction("C", "A", 1)
	nonce = bc.ProofOfWork()
	bc.CreateBlock(nonce)
	log.Printf("Chain`s length is : %d\n", len(bc.chain))
	bc.Print()
}