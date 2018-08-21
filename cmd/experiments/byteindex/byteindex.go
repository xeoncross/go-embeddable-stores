package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/gob"
	"fmt"
)

func main() {

	// The actual index is a N * 8 []byte long slice of values
	var index []byte

	var generatedIds []uint64

	for i := 0; i < 10; i++ {
		// Generate a bunch of random 64bit numbers (in []byte form)
		// id := rand.Int63()
		ib := make([]byte, 8)
		rand.Read(ib)

		id := binary.BigEndian.Uint64(ib)
		generatedIds = append(generatedIds, id)

		index = append(index, ib...)
	}

	fmt.Println("index", index)
	fmt.Println("generatedIds", generatedIds)

	var ids []uint64
	for i := 0; i < len(index); i += 8 {
		id := binary.BigEndian.Uint64(index[i : i+8])
		ids = append(ids, id)
	}

	fmt.Println("ids", ids)
}

func MarshalByte(v []uint64) (index []byte, err error) {
	for _, id := range v {
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, id)
		index = append(index, b...)
	}
	return
}

func UnmarshalByte(index []byte) (ids []uint64, err error) {
	for i := 0; i < len(index); i += 8 {
		id := binary.BigEndian.Uint64(index[i : i+8])
		ids = append(ids, id)
	}
	return
}

func MarshalGob(v []uint64) ([]byte, error) {
	b := new(bytes.Buffer)
	err := gob.NewEncoder(b).Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func UnmarshalGob(data []byte) (ids []uint64, err error) {
	b := bytes.NewBuffer(data)
	err = gob.NewDecoder(b).Decode(&ids)
	return
}
