package main

import (
	"crypto/rand"
	"encoding/binary"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("leveldb", nil)
	if err != nil {
		log.Fatal(err)
	}

	var key, batchSize uint64
	batchSize = 100000

	for {
		insertKeys(db, key, key+batchSize)
		deleteKeys(db, key, key+batchSize)
		key += batchSize
	}

}

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

func dbKey(index uint64) []byte {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, index)
	return key
}

func insertKeys(db *leveldb.DB, startKey uint64, endKey uint64) {
	for i := startKey; i < endKey; i++ {
		data, _ := randomBytes(256)
		db.Put(dbKey(i), data, nil)
		if i%10000 == 0 {
			log.Printf("Inserted %d records", i)
		}
	}
}

func deleteKeys(db *leveldb.DB, startKey uint64, endKey uint64) {
	for i := startKey; i < endKey; i++ {
		db.Delete(dbKey(i), nil)
		if i%10000 == 0 {
			log.Printf("Deleted %d records", i)
		}
	}
}

// func Int64ToByte(id int64) (b []byte) { //, err error) {
// 	b = make([]byte, 8)
// 	binary.BigEndian.PutUint64(b, uint64(id))
// 	return
// }
//
// func ByteToInt64(b []byte) int64 {
// 	x := binary.BigEndian.Uint64(b)
// 	return int64(x)
// }
