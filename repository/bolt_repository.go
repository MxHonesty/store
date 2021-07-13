package repository

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"store/pair"
)

// Implementation of Repository that uses bolt as a key/value store.
type BoltRepository struct {
	database *bolt.DB
	dbName string
}

// Creates a new instance of BoltRepository. If the error returned is nil, the
// Repository was correctly created.
func NewBoltRepository(dbName string) (*BoltRepository, error) {
	db, err := bolt.Open(dbName, 0666, nil)
	store := &BoltRepository{database: db, dbName: dbName}
	store.createBucket()  // init Main bucket.
	return store, err
}

// Closes the db.
func (b	*BoltRepository) Close() {
	_ = b.database.Close()
}

// Creates the main Bucket for the Store.
func (b *BoltRepository) createBucket() {
	err := b.database.Update(func(tx *bolt.Tx) error {
		_ , err := tx.CreateBucketIfNotExists([]byte("Main"))
		if err != nil {
			return fmt.Errorf("create Main bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		println(err.Error())
	}
}

func (b *BoltRepository) Add(pair pair.Pair) {
	memento := pair.GetMemento()
	key := memento.First()
	value := memento.Second()

	err := b.database.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Main"))
		err := bucket.Put(key, value)
		return err
	})
	if err != nil {
		println(err.Error())
	}
}

func (b *BoltRepository) Remove(key string) bool {
	removed := false

	exists := b.Find(key)
	if !exists {
		return false  // Non existent can't be removed.
	}

	err := b.database.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Main"))
		err := bucket.Delete([]byte(key))
		if err != nil {
			return err
		}
		removed = true  // if no error, item removed.
		return nil

	})
	if err != nil {
		println(err.Error())
	}

	return removed
}

// Returns true if the given key is found inside the store. It checks the type of
// key to be []byte.
func (b *BoltRepository) Find(key string) bool {
	found := false
	err := b.database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Main"))
		rez := bucket.Get([]byte(key))
		if rez != nil {
			found = true  // If found
		}
		return nil
	})
	if err != nil {
		println(err.Error())
	}

	return found
}

func (b *BoltRepository) Search(key string) (pair.Pair, bool) {
	found := false
	var foundPair pair.Pair

	err := b.database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Main"))
		rez := bucket.Get([]byte(key))
		if rez != nil {
			m := pair.NewMemento([]byte(key), rez)
			foundPair = pair.NewStringPairFromMemento(m)
			found = true  // If found
		}
		return nil
	})
	if err != nil {
		println(err.Error())
	}

	return foundPair, found
}
