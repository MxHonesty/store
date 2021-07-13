package server

// A struct representing a pair of strings as an argument for an rpc call.
type StringPairArgs struct {
	First, Second string
}

// A struct representing a search query result.
type SearchQueryResponse struct {
	First, Second string
	found bool
}
