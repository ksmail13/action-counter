package main

import (
	"fmt"

	"./config"
	"./server"
)

func main() {
	fmt.Println("hello world")

	conf := &config.Config{}
	serv := &server.Server{}
	serv.Initialize(conf)
	serv.Run(":8080")
}
