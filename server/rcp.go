package server

import (
	"fmt"
	"net"
	"net/rpc"
	"store/service"
)

type Server struct {
	service *service.Service
}

// Creates a new Server with a given service.Service.
func NewServer(service *service.Service) *Server {
	return &Server{service: service}
}

// Function that runs the Server on a given port and with a given service.
func RunServer(service *service.Service, port string) {
	server := NewServer(service)
	_ = rpc.Register(server)
	ln, err := net.Listen("tcp", ":" + port)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := ln.Accept()  // Wait for connection.
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)  // Serve connection.
	}
}

// Add a pair of string to the store.
func (srv *Server) Add(i StringPairArgs, reply *bool) error {
	fmt.Println("Received pair " + i.First + " " + i.Second)
	srv.service.AddPair(i.First, i.Second)
	*reply = true
	return nil
}

// Removes a key/value pair from the store by a single key.
func (srv *Server) Remove(key string, reply *bool) error {
	fmt.Println("Received Remove query key: " + key)
	*reply = srv.service.RemovePair(key)
	return nil
}

// Checks if the given key is found inside the store.
func (srv *Server) Find(key string, reply *bool) error {
	fmt.Println("Received Find query key: " + key)
	*reply = srv.service.Find(key)
	return nil
}

// Search query for a key.
func (srv *Server) Search(key string, reply *SearchQueryResponse) error {
	fmt.Println("Received Search query key: " + key)
	pair, found := srv.service.Search(key)
	reply = &SearchQueryResponse{First: pair.First().(string),
		Second: pair.Second().(string), found: found}
	return nil
}

// A Client for testing the Rcp Server.
func MockTestClient(port string) {
	c, err := rpc.Dial("tcp", "127.0.0.1:" + port)
	if err != nil {
		fmt.Println(err)
		return
	}

	var result bool
	args := StringPairArgs{"a", "b"}
	err = c.Call("Server.Add", args, &result)
	if !result {
		fmt.Println("Expected MockTestClient to get True")
	}

	var searchResult SearchQueryResponse
	err = c.Call("Server.Search", "a" ,&searchResult)
	if searchResult.Second != "b" {
		fmt.Println("Expected to find ")
	}

	err = c.Call("Server.Remove", "a", &result)
	if !result {
		fmt.Println("Expected to remove successfully")
	}
}
