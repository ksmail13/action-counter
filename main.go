package main

import (
	"fmt"

	"github.com/ksmail13/action-counter/config"
	"github.com/ksmail13/action-counter/repository"
	"github.com/ksmail13/action-counter/server"
)

func main() {
	fmt.Println("hello world")

	conf := &config.Config{RedisAddr: "dev.redis:10379", RedisPassword: "", RedisDB: 0}
	serv := &server.Server{Repo: repository.Redis(conf)}

	serv.Initialize(conf)
	serv.Run(":8080")
}
