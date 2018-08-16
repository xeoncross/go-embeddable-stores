package main

import (
	"encoding/binary"
	"os"

	tiedot "github.com/HouzuoGuo/tiedot/db"
)

/*
 * This test is super-wonky and only for very rough comparison as tiedot is a
 * document engine, not a key-value store. And we're not testing it correctly
 * anyway.
 */

var tiedotBucketName = "benchmark"

func newTiedot(path string) (kvEngine, error) {
	db, err := tiedot.OpenDB(path)
	if err != nil {
		return nil, err
	}

	// Second run this already exists
	_ = db.Create(tiedotBucketName)

	bucket := db.Use(tiedotBucketName)

	return &tiedotEngine{db: db, bucket: bucket, path: path}, err
}

type tiedotEngine struct {
	path   string
	db     *tiedot.DB
	bucket *tiedot.Col
}

// Not really fair to compare a document store to a key-value store
func (db *tiedotEngine) Put(key []byte, value []byte) error {
	_, err := db.bucket.Insert(map[string]interface{}{
		string(key): value,
	})
	return err
}

func (db *tiedotEngine) Get(key []byte) ([]byte, error) {
	// Overflows for non-32 bit keys, but we don't care
	id := binary.BigEndian.Uint32(key)
	val, _ := db.bucket.Read(int(id))

	// This will never work
	if val, ok := val["whytry"]; ok {
		return val.([]byte), nil
	}

	return []byte("lies, all lies"), nil
}

func (db *tiedotEngine) Close() error {
	return db.db.Close()
}

func (db *tiedotEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}

func (db *tiedotEngine) Cleanup() error {
	return os.RemoveAll(db.path)
}
