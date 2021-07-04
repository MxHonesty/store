package pair

// Common interface for all Pair instances.
// The database can store any kind of key/value pairs
// as long as they can be translated to a bytes array.
type Pair interface {
	First() interface{}
	Second() interface{}
	SetMemento(m Memento)
	GetMemento() Memento
}
