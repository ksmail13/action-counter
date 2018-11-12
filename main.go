package main

import (
	"fmt"
	"github.com/2DP/action-counter/server"
	"github.com/2DP/action-counter/config"
)

func main() {
	fmt.Println("hello world")
	
	config := &config.Config{}	
	server := &server.Server{}
	server.Initialize(config)
	server.Run(":8080")
}

