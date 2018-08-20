package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"sync"
	"time"

	bolt "github.com/coreos/bbolt"
)

func main() {
	path := "test.db"

	os.RemoveAll(path)
	db, _ := bolt.Open(path, 0600, nil)
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket([]byte("MyBucket"))
		return nil
	})

	var wg sync.WaitGroup
	key := []byte("keyhere")

	start := time.Now()

	var i uint64
	for i = 0; i < 10000; i++ {
		wg.Add(1)
		go func(i uint64) {
			_, err := incr(db, key)
			if err != nil {
				fmt.Println(err)
			}
			// fmt.Println("goroutine", i, id)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(time.Since(start))

	fmt.Println(i)
	err := db.View(func(tx *bolt.Tx) (err error) {
		bucket := tx.Bucket([]byte("MyBucket"))
		b := bucket.Get(key)
		var id int64
		if len(b) != 0 {
			id = ByteToInt64(b)
		}
		fmt.Println(b, "=", id)
		return
	})

	if err != nil {
		fmt.Println(err)
	}

}

func incr(db *bolt.DB, key []byte) (id int64, err error) {
	// mu.Lock()
	// defer mu.Unlock()

	err = db.Update(func(tx *bolt.Tx) (err error) {
		bucket := tx.Bucket([]byte("MyBucket"))
		b := bucket.Get(key)
		if len(b) != 0 {
			id = ByteToInt64(b)
		}
		id += 1
		err = bucket.Put(key, Int64ToByte(id))
		return
	})

	return
}

func Int64ToByte(id int64) (b []byte) { //, err error) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(id))
	return
}

func ByteToInt64(b []byte) int64 {
	x := binary.BigEndian.Uint64(b)
	return int64(x)
}
