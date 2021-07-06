package service

import "store/pair"

// The structure of a Service. It encapsulates all the operations of the
// application.
type Service struct {
	repo Repository  // The Repository responsible for storing data.
	pairCreator pair.AbstractFactory
}

// TODO run test suite for a InMemoryRepo-backed Service.

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

// Removes a pair.Pair from the store. Returns true if the operation was done
// successfully.
func (srv *Service) RemovePair(key interface{}) bool {
	return srv.repo.Remove(key)
}

// Returns true if an element with the given key is found.
func (srv *Service) Find(key interface{}) bool {
	return srv.repo.Find(key)
}

// Searches for an element with the given key and returns the element, and a bool
// that is true if the element was found. If the bool is false, then the first
// element will be zero-initialized.
func (srv *Service) Search(key interface{}) (pair.Pair, bool) {
	return srv.repo.Search(key)
}
