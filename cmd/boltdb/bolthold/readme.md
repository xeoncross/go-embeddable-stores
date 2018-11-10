## Bolthold Lookups

https://github.com/timshannon/bolthold is a library to wrap bbolt with encoding/decoding, indexing, and a query wrapper.

This saves us from needing to create the extra boiler plate to get estimated benchmarks for searching through a collection of objects.

## Lookup

Here we create 50k structs and search for about 10% of them.

Searching for ~5k of the small structs takes about `3ms` on my machine as shown below.

## Run

    go test --bench=. --benchmem -cpu 1

    goos: darwin
    goarch: amd64
    pkg: github.com/Xeoncross/go-embeddable-stores/cmd/boltdb/bolthold
    BenchmarkLookups 	     500	   3243333 ns/op	  932139 B/op	   22485 allocs/op
