package main

import (
	"github.com/boltdb/bolt"
)

var boltBucketName = []byte("benchmark")

type boltEngine struct {
	db   *bolt.DB
	path string
}

func newBolt(path string) (kvEngine, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	db.NoSync = true
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket(boltBucketName)
		return err
	})
	return &boltEngine{db: db, path: path}, err
}

func (db *boltEngine) Put(key []byte, value []byte) error {
	return db.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketName)
		return b.Put(key, value)
	})
}

func (db *boltEngine) Get(key []byte) ([]byte, error) {
	var val []byte
	err := db.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketName)
		val = b.Get(key)
		return nil
	})
	return val, err
}

func (db *boltEngine) Close() error {
	return db.db.Close()
}

func (db *boltEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}
