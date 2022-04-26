package main

import (
	"log"

	"github.com/zkc/pkg/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Panic("sdfoksadfj")
	}
	server.Start()
}
