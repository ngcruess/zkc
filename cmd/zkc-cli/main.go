package main

import (
	"fmt"

	"github.com/zkc/pkg/cui"
)

func main() {
	server, err := cui.NewServer()

	if err != nil {
		fmt.Println(err)
	}

	server.Start()
}
