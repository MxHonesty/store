package service

import "store/pair"

// Common interface for all Repository types.
type Repository interface {
	Add(pair pair.Pair)  // Add a pair.Pair to the Repository.
	Remove(key string) bool  // Remove a pair.Pair with a key from Repository.
	Find(key string) bool  // Find an item in the Repository.
	Search(key string) (pair.Pair, bool)  // Search an item in the Repository.
}
