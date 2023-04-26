package main

import (
	"server/bcserver"
)

func main() {
	bcServer := bcserver.NewBCServer(5000)
	bcServer.RunServer()
}