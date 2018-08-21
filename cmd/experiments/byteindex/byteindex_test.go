package main

import (
	"math/rand"
	"testing"
)

func BenchmarkByte(b *testing.B) {
	rand.Seed(1) // Same every run

	for n := 0; n < b.N; n++ {

		// Generate a bunch of ids
		var ids []uint64
		for i := 0; i < 10000; i++ {
			id := rand.Int63()
			ids = append(ids, uint64(id))
		}

		index, err := MarshalByte(ids)
		if err != nil {
			b.Error(err)
		}

		var decodedIds []uint64
		decodedIds, err = UnmarshalByte(index)
		if err != nil {
			b.Error(err)
		}

		// fmt.Println("ids", ids)
		// fmt.Println("decodedIds", decodedIds)

		if testEq(decodedIds, ids) == false {
			b.Error("mis-match")
		}
	}
}

func BenchmarkGob(b *testing.B) {
	rand.Seed(1) // Same every run

	for n := 0; n < b.N; n++ {

		// Generate a bunch of ids
		var ids []uint64
		for i := 0; i < 10000; i++ {
			id := rand.Int63()
			ids = append(ids, uint64(id))
		}

		index, err := MarshalGob(ids)
		if err != nil {
			b.Error(err)
		}

		var decodedIds []uint64
		decodedIds, err = UnmarshalGob(index)
		if err != nil {
			b.Error(err)
		}

		// fmt.Println("ids", ids)
		// fmt.Println("decodedIds", decodedIds)

		if testEq(decodedIds, ids) == false {
			b.Error("mis-match")
		}
	}
}

func testEq(a, b []uint64) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
