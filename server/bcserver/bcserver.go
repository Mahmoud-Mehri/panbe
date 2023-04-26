package bcserver

import (
	"fmt"
	"log"
	"net/http"
	"panbe/blockchain"
	"panbe/wallet"
)

var blockchains map[string]*blockchain.Blockchain = make(map[string]*blockchain.Blockchain)


type BCServer struct {
	port int
}

func NewBCServer(port int) *BCServer {
	return &BCServer{
		port: port,
	}
}

func (bs *BCServer) Port() int {
	return bs.port;
}

func (bs BCServer) GetBlockchain() *blockchain.Blockchain {
	bc, ok := blockchains["blockchain"];
	if !ok {
		w := wallet.NewWallet();
		bc = blockchain.NewBlockChain(w.WalletAddress())
		blockchains["blockchain"] = bc;
	}
	return bc

}

func HandleRequest(wr http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

	}
}

func (bs *BCServer) RunServer() {
	http.HandleFunc("/", HandleRequest)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", bs.port), nil))
}



