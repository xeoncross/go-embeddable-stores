package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var mu sync.Mutex

func main() {
	db, err := leveldb.OpenFile("database", nil)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	key := []byte("keyhere")

	var i uint64
	for i = 0; i < 10; i++ {
		wg.Add(1)
		go func(i uint64) {
			id, err := incr(db, key)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("goroutine", i, id)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(i)
	id, err := db.Get(key, &opt.ReadOptions{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id, "=", ByteToInt64(id))

}

func incr(db *leveldb.DB, key []byte) (id int64, err error) {
	mu.Lock()
	defer mu.Unlock()

	// May or may not exist
	b, err := db.Get(key, &opt.ReadOptions{DontFillCache: true})
	if err == nil {
		return
	}
	if len(b) != 0 {
		id = ByteToInt64(b)
	}

	id += 1
	err = db.Put(key, Int64ToByte(id), &opt.WriteOptions{Sync: true})
	if err != nil {
		return
	}

	return
}

// func incr(db *leveldb.DB, key []byte) (id int64, err error) {
// 	mu.Lock()
// 	defer mu.Unlock()
//
// 	tx, err := db.OpenTransaction()
// 	if err != nil {
// 		return
// 	}
//
// 	// May or may not exist
// 	b, err := tx.Get(key, &opt.ReadOptions{DontFillCache: true})
// 	if err == nil {
// 		tx.Discard()
// 		return
// 	}
// 	if len(b) != 0 {
// 		id = ByteToInt64(b)
// 	}
//
// 	id += 1
// 	err = tx.Put(key, Int64ToByte(id), &opt.WriteOptions{Sync: true})
// 	if err != nil {
// 		return
// 	}
//
// 	err = tx.Commit()
// 	return
// }

func Int64ToByte(id int64) (b []byte) { //, err error) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(id))
	return
}

func ByteToInt64(b []byte) int64 {
	x := binary.BigEndian.Uint64(b)
	return int64(x)
}
