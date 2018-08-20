package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"

	bolt "github.com/coreos/bbolt"
)

func main() {
	path := "test.db"

	os.RemoveAll(path)
	db, _ := bolt.Open(path, 0600, nil)
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket([]byte("test"))
		return nil
	})

	valueSize := 16

	for i := 0; i < 1000000; i++ {
		go func(i int) {
			db.Batch(func(tx *bolt.Tx) error {
				key := strconv.Itoa(i)
				b := tx.Bucket([]byte("test"))
				b.Put([]byte(key), getBytes(valueSize))
				if i%100000 == 0 {
					alloc, sys := getMemUsage()
					log.Printf("Key = '%v', Alloc = %4vM, Sys = %4vM, Goroutines = %d", key, alloc, sys, runtime.NumGoroutine())
				}
				return nil
			})
		}(i)
		time.Sleep(time.Microsecond)
	}
}

func getBytes(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic("Can't read from rand.Read")
	}
	return b
}

func getMemUsage() (uint64, uint64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return bToMb(m.Alloc), bToMb(m.Sys)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
