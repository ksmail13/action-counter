package main

import (
	"fmt"

	"github.com/ksmail13/action-counter/config"
	"github.com/ksmail13/action-counter/server"
)

func main() {
	fmt.Println("hello world")

	conf := &config.Config{}
	serv := &server.Server{}
	serv.Initialize(conf)
	serv.Run(":8080")
}
