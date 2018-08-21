package main

import (
	"bytes"
	"sort"
)

/*
BenchmarkByte-8        	    3000	    380887 ns/op	 1155563 B/op	      64 allocs/op
BenchmarkGob-8         	    2000	    624693 ns/op	 1055519 B/op	     228 allocs/op
BenchmarkByteIndex-8   	   20000	    127662 ns/op	     164 B/op	       1 allocs/op
BenchmarkKeyList-8     	  500000	    626237 ns/op	     290 B/op	       1 allocs/op
*/

// https://github.com/timshannon/bolthold/blob/master/index.go#L108

// keyList is a slice of unique, sorted keys([]byte) such as what an index points to
type keyList [][]byte

func (v *keyList) add(key []byte) {
	i := sort.Search(len(*v), func(i int) bool {
		return bytes.Compare((*v)[i], key) >= 0
	})

	if i < len(*v) && bytes.Equal((*v)[i], key) {
		// already added
		return
	}

	*v = append(*v, nil)
	copy((*v)[i+1:], (*v)[i:])
	(*v)[i] = key
}

func (v *keyList) remove(key []byte) {
	i := sort.Search(len(*v), func(i int) bool {
		return bytes.Compare((*v)[i], key) >= 0
	})

	if i < len(*v) {
		copy((*v)[i:], (*v)[i+1:])
		(*v)[len(*v)-1] = nil
		*v = (*v)[:len(*v)-1]
	}
}

func (v *keyList) in(key []byte) bool {
	i := sort.Search(len(*v), func(i int) bool {
		return bytes.Compare((*v)[i], key) >= 0
	})

	return (i < len(*v) && bytes.Equal((*v)[i], key))
}
