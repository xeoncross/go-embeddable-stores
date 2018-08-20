package main

import (
	"os"

	"github.com/boltdb/bolt"
)

var boltBucketName = []byte("benchmark")

type boltEngine struct {
	db   *bolt.DB
	path string
	// bucket *bolt.Bucket
}

func newBolt(path string) (kvEngine, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	db.NoSync = true

	// var bucket *bolt.Bucket
	// err = db.Update(func(tx *bolt.Tx) (err error) {
	// 	bucket, err = tx.CreateBucket(boltBucketName)
	// 	return err
	// })
	return &boltEngine{db: db, path: path}, err
}

func (db *boltEngine) Put(key []byte, value []byte) error {
	return db.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketName)
		return b.Put(key, value)
	})
	// return db.bucket.Put(key, value)
}

func (db *boltEngine) Get(key []byte) ([]byte, error) {
	var val []byte
	err := db.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketName)
		val = b.Get(key)
		return nil
	})
	return val, err
	// return db.bucket.Get(key), nil
}

// https://github.com/boltdb/bolt#prefix-scans
// func (db *boltEngine) Search(key []byte) error {
// 	return db.db.View(func(tx *bolt.Tx) error {
// 		c := tx.Bucket(key).Cursor()
//
// 		suffix := []byte("1234")
// 		for k, v := c.Seek(suffix); bytes.HasSuffix(k, suffix); k, v = c.Next() {
// 			fmt.Printf("key=%s, value=%s\n", k, v)
// 		}
//
// 		return nil
// 	})
// }

func (db *boltEngine) Close() error {
	return db.db.Close()
}

func (db *boltEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}

func (db *boltEngine) Cleanup() error {
	return os.RemoveAll(db.path)
}
