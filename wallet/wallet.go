package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"

	"panbe/utils"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey *ecdsa.PublicKey
	address string
}

func NewWallet() *Wallet {
	w := new(Wallet)

	// Step1: Creating ECDSA PrivateKey (32 Bytes) and PublickKey (64 Bytes)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey

	// Step2: Perform SHA-256 hashing on the PublicKey (32 Bytes)
	s2 := sha256.New();
	s2.Write(w.publicKey.X.Bytes())
	s2.Write(w.publicKey.Y.Bytes())
	s2Digest := s2.Sum(nil)

	// Step3: Perform RIPEMD-160 hashing on the result of Step2 (20 Bytes)
	s3 := ripemd160.New()
	s3.Write(s2Digest)
	s3Digest := s3.Sum(nil)

	// Step4: Adding Version Byte at the beginning of RIPEMD-160 hash (0x00 for Main Network)
	s4 := make([]byte, 21)
	s4[0] = 0x00
	copy(s4[1:], s3Digest[:])

	// Step5: Perform SHA-256 hashing on the result of Step4
	s5 := sha256.New()
	s5.Write(s4)
	s5Digest := s5.Sum(nil)

	// Step6: Perform SHA-256 hashing on the result of Step5
	s6 := sha256.New()
	s6.Write(s5Digest)
	s6Digest := s6.Sum(nil)

	// Step7: Take the first 4 bytes of the result as Checksum
	s7 := s6Digest[:4]

	// Step8: Adding Checksum bytesat the end of the RIPEMD-160 value of Step4
	s8 := make([]byte, 25)
	copy(s8[:21], s4[:])
	copy(s8[21:], s7[:])
	
	// Step9: Convert from Bytes to Base58
	s9 := base58.Encode(s8[:])

	// The final result will be the Address value
	w.address = s9;

	return w
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *Wallet) WalletAddress() string {
	return w.address
}

type TransactionInfo struct {
	SenderPrivateKey *ecdsa.PrivateKey
	SenderPublicKey *ecdsa.PublicKey
	SenderAddress string
	RecipientAddress string
	Value float32
}

func NewTransactionInfo(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey,
	sender string, recipient string, value float32) *TransactionInfo {
		return &TransactionInfo{privateKey, publicKey, sender, recipient, value}
}

func (t *TransactionInfo) ToJSON() ([]byte, error) {
	return json.Marshal(struct{
		Sender string `json:"senderAddress"`
		Recipient string `json:"recipientAddress"`
		Value float32 `json:"value"`
	}{
		t.SenderAddress,
		t.RecipientAddress,
		t.Value,
	})
}

func (t *TransactionInfo) SignTransaction() *utils.Signature {
	data, _ := t.ToJSON()
	h := sha256.Sum256(data);
	r, s, _ := ecdsa.Sign(rand.Reader, t.SenderPrivateKey, h[:])
	return &utils.Signature{R: r, S: s}
}
