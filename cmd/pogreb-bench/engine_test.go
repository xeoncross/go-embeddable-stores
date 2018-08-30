package main

import (
	"math/rand"
	"path"
	"testing"
)

// RandomBytes up to X length
func randomBytes(min, max int) (b []byte) {
	b = make([]byte, min+rand.Intn(max-min))
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func BenchmarkEngines100(b *testing.B) {

	for engineName := range engines {
		ctr, err := getEngineCtr(engineName)
		if err != nil {
			b.Error(err)
		}

		b.Run(engineName, func(b2 *testing.B) {
			dbpath := path.Join("./data", "bench_"+engineName)
			db, err := ctr(dbpath)
			if err != nil {
				b.Error(err)
			}

			for n := 0; n < b2.N; n++ {
				db.Put(randomBytes(8, 64), randomBytes(8, 64))
			}

			err = db.Close()
			if err != nil {
				b.Error(err)
			}

			err = db.Cleanup()
			if err != nil {
				b.Error(err)
			}
		})

	}
}
