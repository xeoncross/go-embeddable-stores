package main

import (
	"github.com/dgraph-io/badger"
)

func newBadgerdb(path string) (kvEngine, error) {
	opts := badger.DefaultOptions
	opts.SyncWrites = false
	opts.Dir = path
	opts.ValueDir = path
	db, err := badger.Open(opts)
	return &badgerdbEngine{db: db, path: path}, err
}

type badgerdbEngine struct {
	path string
	db   *badger.DB
}

func (db *badgerdbEngine) Put(key []byte, value []byte) error {
	return db.db.Update(func(tx *badger.Txn) error {
		return tx.Set(key, value)
	})
}

func (db *badgerdbEngine) Get(key []byte) ([]byte, error) {
	var val []byte
	err := db.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		v, err := item.Value()
		if err != nil {
			return err
		}
		val = v
		return nil
	})
	return val, err
}

func (db *badgerdbEngine) Close() error {
	return db.db.Close()
}

func (db *badgerdbEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}
