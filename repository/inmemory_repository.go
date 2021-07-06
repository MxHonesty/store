package repository

import (
	"store/pair"
)

// TODO Implement

// In-Memory implementation of the Repository. The intent of this module is
// mainly for simplifying testing by storing all the key/value pairs in a slice.
type InMemoryRepository struct {
	elems []pair.Pair
}

func NewInMemoryRepository() *InMemoryRepository {
	elems := make([]pair.Pair, 0, 10)  // Create a new empty slice.
	return &InMemoryRepository{elems: elems}
}

// Add a pair.Pair to the service.Repository.
func (i *InMemoryRepository) Add(pair pair.Pair) {
	i.elems = append(i.elems, pair)
}

// Removes a pair.Pair with a given key from the service.Repository. Returns true
// if a pair.Pair was removed, returns false if none was found.
func (i *InMemoryRepository) Remove(key interface{}) bool {
	rez := make([]pair.Pair, 0)
	index := -1

	for i, el := range i.elems {  // Search for the element with the given key.
		if el.First() == key {
			index = i
			break
		}
	}

	if index != -1 {  // If element with given key is found.
		rez = append(rez, i.elems[:index]...)
		rez = append(rez, i.elems[index+1:]...)
		i.elems = rez
		return true
	} else {
		return false
	}
}

// Returns true if a pair.Pair with the given key exists.
func (i *InMemoryRepository) Find(key interface{}) bool {
	for _, el := range i.elems {
		if el.First() == key {
			return true
		}
	}
	return false
}

// Searches for a pair.Pair with the given key. IF such a pair is found. return
// it. The bool return value determines if such a pair could be found.
func (i *InMemoryRepository) Search(key interface{}) (pair.Pair, bool) {
	for _, el := range i.elems {
		if el.First() == key {
			return el, true
		}
	}
	return nil, false
}

// Returns the number of elements in the Repository.
func (i *InMemoryRepository) Count() int {
	return len(i.elems)
}
