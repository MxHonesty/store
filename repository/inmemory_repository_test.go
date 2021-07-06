package repository

import (
	"store/pair"
	"testing"
)

// Test case for initializing the InMemoryRepository.
func TestNewInMemoryRepository(t *testing.T) {
	repo := NewInMemoryRepository()
	if repo.Count() != 0 {
		t.Errorf("Expected size 0, got %d", repo.Count())
	}
}

// Test case for adding an element in the slice.
func TestInMemoryRepository_Add(t *testing.T) {
	repo := NewInMemoryRepository()
	repo.Add(pair.NewStringPair("a", "a"))

	if repo.Count() != 1 {
		t.Errorf("Expected size to be 1, got %d", repo.Count())
	}

	if repo.elems[0].First() != "a" {  // Make sure the pair is correct.
		t.Errorf("Expected First element of pair to be a, got %s",
			repo.elems[0].First())
	}

	if repo.elems[0].Second() != "a" {
		t.Errorf("Expected Second element of pair to be a, got %s",
			repo.elems[0].Second())
	}
}

func TestInMemoryRepository_Find(t *testing.T) {
	repo := NewInMemoryRepository()
	repo.Add(pair.NewStringPair("a", "a"))
	repo.Add(pair.NewStringPair("b", "a"))
	repo.Add(pair.NewStringPair("c", "a"))
	repo.Add(pair.NewStringPair("d", "a"))

	if !repo.Find("a") {
		t.Error("Didn't find existing item")
	}

	if !repo.Find("b") {
		t.Error("Didn't find existing item")
	}

	if !repo.Find("c") {
		t.Error("Didn't find existing item")
	}

	if !repo.Find("d") {
		t.Error("Didn't find existing item")
	}

	if repo.Find("abc") {
		t.Error("Found non existing item")
	}
}

func TestInMemoryRepository_Remove(t *testing.T) {
	repo := NewInMemoryRepository()
	repo.Add(pair.NewStringPair("a", "a"))
	repo.Add(pair.NewStringPair("b", "a"))
	repo.Add(pair.NewStringPair("c", "a"))
	repo.Add(pair.NewStringPair("d", "a"))

	rez := repo.Remove("b")

	if rez != true {
		t.Errorf("Expected to return true, got false")
	}

	// Check item count.
	if repo.Count() != 3 {
		t.Errorf("Expected to have 3 items, we have %d", repo.Count())
	}

	rez = repo.Remove("c")
	if rez != true {
		t.Errorf("Expected to return true, got false")
	}

	// Check that the correct items were removed.
	if repo.elems[0].First() != "a" {
		t.Errorf("Expected First key to be a, got %s",
			repo.elems[0].First())
	}

	if repo.elems[1].First() != "d" {
		t.Errorf("Expected First key to be d, got %s",
			repo.elems[1].First())
	}

	// Try to delete a non existing item.
	rez = repo.Remove("b")
	if rez != false {
		t.Errorf("Repo reported deleting non existing item")
	}

	if repo.Count() != 2 {
		t.Errorf("Expected size 2, got %d", repo.Count())
	}
}

func TestInMemoryRepository_Search(t *testing.T) {
	repo := NewInMemoryRepository()
	repo.Add(pair.NewStringPair("a", "a"))
	repo.Add(pair.NewStringPair("b", "a"))
	repo.Add(pair.NewStringPair("c", "a"))
	repo.Add(pair.NewStringPair("d", "a"))

	el, rez := repo.Search("a")
	if !rez {
		t.Error("Didn't find existing item")
	}

	if el.First() != "a" {
		t.Errorf("Expected returned item First to be a, got %s", el.First())
	}
}
