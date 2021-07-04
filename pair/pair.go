package pair

// The representation of a Key/Value Pair.
type StringPair struct {
	first  string
	second string
}

// Instantiate a new StringPair from two string values.
func NewStringPair(first, second string) *StringPair {
	return &StringPair{first: first, second: second}
}

// Instantiate a new StringPair from a Memento instance.
func NewStringPairFromMemento(m Memento) *StringPair {
	pair := &StringPair{}
	pair.SetMemento(m)
	return pair
}

// Reconstruct the state of the Pair from a Memento.
func (p *StringPair) SetMemento(m Memento) {
	first := string(m.First())
	second := string(m.Second())
	p.first = first
	p.second = second
}

// Create a snapshot of the current state of the StringPair.
// This snapshot is represented as a Memento instance.
func (p *StringPair) GetMemento() Memento {
	return NewMemento([]byte(p.first), []byte(p.second))
}

// Return the first element of the Pair.
func (p *StringPair) First() interface{} {
	return p.first
}

// Return the second element of the Pair.
func (p *StringPair) Second() interface{} {
	return p.second
}
