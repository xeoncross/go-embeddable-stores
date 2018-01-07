package main

import (
	"github.com/akrylysov/pogreb"
	"github.com/akrylysov/pogreb/fs"
)

func newPogreb(path string) (kvEngine, error) {
	return pogreb.Open(path, &pogreb.Options{FileSystem: fs.OS})
}
