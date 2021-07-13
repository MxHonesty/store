package service

import (
	"os"
	"store/pair"
	"store/repository"
	"testing"
)

// Cleans up the Database file.
func cleanUp() {
	_ = os.Remove("test.db")
}

func TestService_bolt_AddPair(t *testing.T) {
	repo, _ := repository.NewBoltRepository("test.db")
	srv := NewService(repo, pair.StringPairFactory{})
	srv.AddPair("a", "a")

	rez := repo.Find("a")
	if !rez {
		t.Errorf("Expected to find value")
	}

	repo.Close()
	cleanUp()
}

func TestService_bolt_Find(t *testing.T) {
	repo, _ := repository.NewBoltRepository("test.db")
	srv := NewService(repo, pair.StringPairFactory{})
	srv.AddPair("a", "a")
	srv.AddPair("b", "b")

	if !srv.Find("a") {
		t.Error("Expected to find key a")
	}

	if srv.Find("c") {
		t.Error("Expected to not find key c")
	}

	repo.Close()
	cleanUp()
}

func TestService_bolt_RemovePair(t *testing.T) {
	repo, _ := repository.NewBoltRepository("test.db")
	srv := NewService(repo, pair.StringPairFactory{})
	srv.AddPair("a", "a")
	srv.AddPair("b", "b")

	removed := srv.RemovePair("a")
	if !removed {
		t.Errorf("Returned false for valid remove")
	}

	removed = srv.RemovePair("a")
	if removed {
		t.Errorf("Removed non existent item")
	}

	repo.Close()
	cleanUp()
}

func TestService_bolt_Search(t *testing.T) {
	repo, _ := repository.NewBoltRepository("test.db")
	srv := NewService(repo, pair.StringPairFactory{})
	srv.AddPair("a", "a")
	srv.AddPair("b", "b")

	el, found := srv.Search("a")
	if !found {
		t.Error("Existing item not found")
	}

	if el.Second() != "a" {
		t.Errorf("Expected Second to be a, got %s", el.Second())
	}

	repo.Close()
	cleanUp()
}
