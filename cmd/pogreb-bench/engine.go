package main

import (
	"errors"
	"os"
	"path/filepath"
)

type kvEngine interface {
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Close() error
	FileSize() (int64, error)
	Cleanup() error
}

type engineCtr func(string) (kvEngine, error)

var engines = map[string]engineCtr{
	"pogreb":    newPogreb,
	"goleveldb": newGolevelDB,
	"bolt":      newBolt,
	"bbolt":     newBBolt,
	// "badgerdb": newBadgerdb,
	"tiedot": newTiedot,
}

func getEngineCtr(name string) (engineCtr, error) {
	if ctr, ok := engines[name]; ok {
		return ctr, nil
	}
	return nil, errors.New("unknown engine")
}

func dirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
