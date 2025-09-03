package util

import (
	"fmt"

	badger "github.com/dgraph-io/badger/v4"
)

func BadgerData() error {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		txn.Set([]byte("a"), []byte("b"))
		return nil
	})
	if err != nil {
		return err
	}

	return db.View(func(txn *badger.Txn) error {
		item, rerr := txn.Get([]byte("a"))
		if rerr != nil {
			return rerr
		}
		rerr = item.Value(func(val []byte) error {
			fmt.Printf("The val is: %s\n", val)
			return nil
		})
		if rerr != nil {
			return rerr
		}
		return nil
	})
}
