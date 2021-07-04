package pair

// A struct that represents the byte string representation of a Pair of data.
// This Module exists in order to make it easy to extend the types of keys and
// values supported by the application.
type Memento struct {
	first []byte
	second []byte
}

// Create a new instance of a Memento.
func NewMemento(first, second []byte) Memento {
	return Memento{first: first, second: second}
}

// Returns the first slice of bytes.
func (m Memento) First() []byte {
	return m.first
}

// Returns the second slice of bytes.
func (m Memento) Second() []byte {
	return m.second
}
