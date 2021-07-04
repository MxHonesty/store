package repository

import (
	"store/pair"
)

// Implementation of Repository that uses bolt as a key/value store.
type BoltRepository struct {

}

func (b *BoltRepository) Add(pair pair.Pair) {
	panic("implement me")
}

func (b *BoltRepository) Remove(key interface{}) bool {
	panic("implement me")
}

func (b *BoltRepository) Find(key interface{}) bool {
	panic("implement me")
}

func (b *BoltRepository) Search(key interface{}) pair.Pair {
	panic("implement me")
}
