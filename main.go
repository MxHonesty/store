package main

import (
	"store/pair"
	"store/repository"
	"store/server"
	"store/service"
)

func main() {
	repo, _ := repository.NewBoltRepository("db.db")
	srv := service.NewService(repo, pair.StringPairFactory{})

	server.RunServer(srv, "3000")
}
