package main

import (
	"bytes"
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

// Slower because of a lack of search
// func BenchmarkByteIndexAdd(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		bi := &ByteIndex{}
// 		for i := 0; i < 100000; i++ {
// 			id := make([]byte, 8)
// 			rand.Read(id)
// 			bi.Add(id)
// 		}
// 	}
// }

// func BenchmarkByteIndexRemove(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		bi := &ByteIndex{}
// 		var ids [][]byte
// 		for i := 0; i < 100000; i++ {
// 			id := make([]byte, 8)
// 			rand.Read(id)
// 			bi.Add(id)
// 			bi.Find(id)
// 			ids = append(ids, id)
// 		}
//
// 		for _, id := range ids {
// 			bi.Remove(id)
// 		}
// 	}
// }

// func BenchmarkKeyListAdd(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		kl := &keyList{}
// 		for i := 0; i < 100000; i++ {
// 			id := make([]byte, 8)
// 			rand.Read(id)
// 			kl.add(id)
// 		}
// 	}
// }

// func BenchmarkKeyListRemove(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		kl := &keyList{}
// 		var ids [][]byte
// 		for i := 0; i < 100000; i++ {
//
// 			id := make([]byte, 8)
// 			rand.Read(id)
// 			kl.add(id)
// 			kl.in(id)
// 			ids = append(ids, id)
// 		}
//
// 		for _, id := range ids {
// 			kl.remove(id)
// 		}
// 	}
// }

func BenchmarkKeyListMarshalByte(b *testing.B) {

	kl := &keyList{}
	for i := 0; i < 100000; i++ {
		id := make([]byte, 8)
		rand.Read(id)
		kl.add(id)
	}

	b.ResetTimer()

	b.Run("Byte", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			data, err := kl.MarshalToByte()
			if err != nil {
				b.Error(err)
			}

			kl = &keyList{}
			err = kl.UnmarshalFromByte(data)
			if err != nil {
				b.Error(err)
			}

			if len(*kl) != 100000 {
				b.Error("Error decoding")
			}
		}
	})

	b.Run("Gob", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			data, err := kl.MarshalToGob()
			if err != nil {
				b.Error(err)
			}

			kl = &keyList{}
			err = kl.UnmarshalFromGob(data)
			if err != nil {
				b.Error(err)
			}

			if len(*kl) != 100000 {
				b.Error("Error decoding")
			}
		}
	})
}

// func BenchmarkKeyListMarshalGob(b *testing.B) {
//
// 	kl := &keyList{}
// 	for i := 0; i < 100000; i++ {
// 		id := make([]byte, 8)
// 		rand.Read(id)
// 		kl.add(id)
// 	}
//
// 	b.ResetTimer()
//
// 	for n := 0; n < b.N; n++ {
// 		data, err := kl.MarshalToGob()
// 		if err != nil {
// 			b.Error(err)
// 		}
//
// 		kl = &keyList{}
// 		err = kl.UnmarshalFromGob(data)
// 		if err != nil {
// 			b.Error(err)
// 		}
//
// 		if len(*kl) != 100000 {
// 			b.Error("Error decoding")
// 		}
// 	}
// }

func TestByteIndex(t *testing.T) {
	bi := &ByteIndex{}

	id := make([]byte, 8)
	rand.Read(id)
	bi.Add(id)

	id2 := make([]byte, 8)
	rand.Read(id2)
	bi.Add(id2)

	id3 := make([]byte, 8)
	rand.Read(id3)
	bi.Add(id3)

	// fmt.Println("id", id)
	// fmt.Println("id2", id2)
	// fmt.Println("id3", id3)
	// fmt.Println("ByteIndex", bi)

	if !bytes.Equal((*bi)[0:8], id) {
		t.Error("ID not saved")
	}

	i := bi.Find(id)
	if i == -1 {
		t.Error("ID not found")
	}

	if i != 0 {
		t.Errorf("ID not in correct location: %d", i)
	}

	bi.Remove(id)
	// fmt.Println("ByteIndex", bi)

	if !bytes.Equal((*bi)[8:16], id3) {
		t.Error("ID not removed")
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
