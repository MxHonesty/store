package pair

// The interface for all Pair factories.
type AbstractFactory interface {
	CreatePair(first, second interface{}) Pair
}

// Factory for creating a StringPair.
type StringPairFactory struct {}

// Takes two string arguments and creates a StringPair.
func (s StringPairFactory) CreatePair(first, second interface{}) Pair {
	created := NewStringPair(first.(string), second.(string))
	return created
}
