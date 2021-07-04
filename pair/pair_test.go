package pair

import (
	"reflect"
	"testing"
)

// Test case for creating a new StringPair.
func TestNewStringPair(t *testing.T) {
	pair := NewStringPair("a", "b")
	if pair.First() != "a" {
		t.Errorf("Expected to find a, found %s", pair.First())
	} else if pair.Second() != "b" {
		t.Errorf("Expected to find b, found %s", pair.Second())
	}
}

// Test case for creating a StringPair from Memento.
func TestNewStringPairFromMemento(t *testing.T) {
	m := NewMemento([]byte("abc"), []byte("cba"))
	pair := NewStringPairFromMemento(m)
	if pair.First() != "abc" {
		t.Errorf("Expected to find abc, found %s", pair.First())
	} else if pair.Second() != "cba" {
		t.Errorf("Expected to find cba, found %s", pair.Second())
	}
}

// Test case for creating a Memento from a StringPair.
func TestStringPair_GetMemento(t *testing.T) {
	m := NewMemento([]byte("abc"), []byte("cba"))
	pair := NewStringPair("abc", "cba")
	pairM := pair.GetMemento()

	if !reflect.DeepEqual(pairM.First(), m.First()) {
		t.Errorf("Expected %s, be equal to %s", string(pairM.First()), string(m.First()))
	} else if !reflect.DeepEqual(pairM.Second(), m.Second()) {
		t.Errorf("Expected %s, be equal to %s", string(pairM.Second()), string(m.Second()))
	}
}
