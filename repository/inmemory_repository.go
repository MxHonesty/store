package repository

import (
	"store/pair"
)

// In-Memory implementation of the Repository. The intent of this module is
// mainly for simplifying testing by storing all the key/value pairs in a slice.
type InMemoryRepository struct {

}

func (i *InMemoryRepository) Add(pair pair.Pair) {
	panic("implement me")
}

func (i *InMemoryRepository) Remove(key interface{}) bool {
	panic("implement me")
}

func (i *InMemoryRepository) Find(key interface{}) bool {
	panic("implement me")
}

func (i *InMemoryRepository) Search(key interface{}) (pair.Pair, bool) {
	panic("implement me")
}
