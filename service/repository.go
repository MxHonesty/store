package service

import "store/pair"

// Common interface for all Repository types.
type Repository interface {
	Add(pair pair.Pair)  // Add a pair.Pair to the Repository.
	Remove(key interface{}) bool  // Remove a pair.Pair with a key from Repository.
	Find(key interface{}) bool  // Find an item in the Repository.
	Search(key interface{}) pair.Pair  // Search an item in the Repository.
}
