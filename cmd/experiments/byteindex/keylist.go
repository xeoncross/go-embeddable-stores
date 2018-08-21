package main

import (
	"bytes"
	"encoding/gob"
	"sort"
)

/*
BenchmarkByte-8        	    3000	    380887 ns/op	 1155563 B/op	      64 allocs/op
BenchmarkGob-8         	    2000	    624693 ns/op	 1055519 B/op	     228 allocs/op
BenchmarkByteIndex-8   	   20000	    127662 ns/op	     164 B/op	       1 allocs/op
BenchmarkKeyList-8     	  500000	    626237 ns/op	     290 B/op	       1 allocs/op
BenchmarkKeyListMarshalByte/Byte-8         	     100	  13880275 ns/op	19249168 B/op	      66 allocs/op
BenchmarkKeyListMarshalByte/Gob-8          	     100	  17191247 ns/op	12861329 B/op	  200221 allocs/op
*/

// Because keyList is sorted the lookup times are much faster than ByteIndex.
// Helps speed up adding and removing indexes

// https://github.com/timshannon/bolthold/blob/master/index.go#L108
// https://play.golang.org/p/U1zVINjzOJf
// https://gist.github.com/schmohlio/615ab4d47bc01020786ef58aec622fdf

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

func (v *keyList) MarshalToByte() (index []byte, err error) {
	// Create a byte array large enough, then fill it with each value
	index = make([]byte, len((*v))*8)
	for i, id := range *v {
		copy(id[0:], index[i:i+8])
		// index[i] = id[0]
		// index[i+1] = id[1]
		// index[i+2] = id[2]
		// index[i+3] = id[3]
		// index[i+4] = id[4]
		// index[i+5] = id[5]
		// index[i+6] = id[6]
		// index[i+7] = id[7]
	}
	return
}

func (v *keyList) MarshalToByte2() (index []byte, err error) {
	// return append([]byte{}, ((*v)...)), nil
	for _, id := range *v {
		index = append(index, id...)
	}
	return
}

func (v *keyList) UnmarshalFromByte(index []byte) error {
	(*v) = make([][]byte, len(index)/8)
	for i := 0; i < len(index); i += 8 {
		// (*v) = append((*v), index[i:i+8])
		(*v)[i/8] = index[i : i+8]
	}
	return nil
}

func (v *keyList) UnmarshalFromByte2(index []byte) error {
	(*v) = make([][]byte, len(index)/8)
	for i := 0; i < len(index); i += 8 {
		(*v) = append((*v), index[i:i+8])
	}
	return nil
}

func (v *keyList) MarshalToGob() (index []byte, err error) {
	b := new(bytes.Buffer)
	err = gob.NewEncoder(b).Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (v *keyList) UnmarshalFromGob(index []byte) error {
	b := bytes.NewBuffer(index)
	return gob.NewDecoder(b).Decode(&v)
}
