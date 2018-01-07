package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type goleveldbEngine struct {
	db   *leveldb.DB
	path string
}

func newGolevelDB(path string) (kvEngine, error) {
	opts := opt.Options{Compression: opt.NoCompression}
	db, err := leveldb.OpenFile(path, &opts)
	if err != nil {
		return nil, err
	}
	err = db.CompactRange(util.Range{})
	return &goleveldbEngine{db: db, path: path}, err
}

func (db *goleveldbEngine) Put(key []byte, value []byte) error {
	return db.db.Put(key, value, nil)
}

func (db *goleveldbEngine) Get(key []byte) ([]byte, error) {
	return db.db.Get(key, nil)
}

func (db *goleveldbEngine) Close() error {
	return db.db.Close()
}

func (db *goleveldbEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}
