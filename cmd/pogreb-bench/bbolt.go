package main

import (
	"os"

	"github.com/etcd-io/bbolt"
)

var bboltBucketName = []byte("benchmark")

type bboltEngine struct {
	db   *bbolt.DB
	path string
	// bucket *bbolt.Bucket
}

func newBBolt(path string) (kvEngine, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	// db.NoSync = false // default to sync for each write

	err = db.Update(func(tx *bbolt.Tx) (err error) {
		_, err = tx.CreateBucket(bboltBucketName)
		return err
	})
	return &bboltEngine{db: db, path: path}, err
}

func (db *bboltEngine) Put(key []byte, value []byte) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bboltBucketName)
		return b.Put(key, value)
	})
	// return db.bucket.Put(key, value)
}

func (db *bboltEngine) Get(key []byte) ([]byte, error) {
	var val []byte
	err := db.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bboltBucketName)
		val = b.Get(key)
		return nil
	})
	return val, err
	// return db.bucket.Get(key), nil
}

// https://github.com/bboltdb/bbolt#prefix-scans
// func (db *bboltEngine) Search(key []byte) error {
// 	return db.db.View(func(tx *bbolt.Tx) error {
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

func (db *bboltEngine) Close() error {
	return db.db.Close()
}

func (db *bboltEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}

func (db *bboltEngine) Cleanup() error {
	return os.Remove(db.path)
}
