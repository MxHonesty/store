package main

import (
	"fmt"
	"store/pair"
	"store/repository"
	"store/server"
	"store/service"
)

func main() {
	repo, _ := repository.NewBoltRepository("db.db")
	srv := service.NewService(repo, pair.StringPairFactory{})

	go server.RunServer(srv, "9999")
	go server.MockTestClient("9999")

	var input string
	_, _ = fmt.Scanln(&input)
}
