package main

import (
	"os"
	"time"

	"github.com/akrylysov/pogreb"
	"github.com/akrylysov/pogreb/fs"
)

func newPogreb(path string) (kvEngine, error) {
	opt := &pogreb.Options{FileSystem: fs.OS}
	opt.BackgroundSyncInterval = time.Second
	db, err := pogreb.Open(path, opt)
	return &pogrebEngine{db, path}, err
}

type pogrebEngine struct {
	*pogreb.DB
	path string
}

func (db *pogrebEngine) Cleanup() error {
	err := os.RemoveAll(db.path)
	if err == nil {
		err = os.Remove(db.path + ".index")
	}
	return err
}
