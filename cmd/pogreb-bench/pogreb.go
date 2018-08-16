package main

import (
	"os"

	"github.com/akrylysov/pogreb"
	"github.com/akrylysov/pogreb/fs"
)

func newPogreb(path string) (kvEngine, error) {
	db, err := pogreb.Open(path, &pogreb.Options{FileSystem: fs.OS})
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
