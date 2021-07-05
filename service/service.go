package service

import "store/pair"

// The structure of a Service. It encapsulates all the operations of the
// application.
type Service struct {
	repo Repository  // The Repository responsible for storing data.
	pairCreator pair.AbstractFactory
}

// Return a new instance of Service. Must take as argument a Repository and
// a pair.AbstractFactory for creating instances of pair.Pair.
func NewService(repo Repository, pairCreator pair.AbstractFactory) *Service {
	return &Service{repo: repo, pairCreator: pairCreator}
}

// Adds a Pair of elements to the store. The types of these values is arbitrary.
func (srv *Service) AddPair(first, second interface{}) {
	p := srv.pairCreator.CreatePair(first, second)
	srv.repo.Add(p)
}
