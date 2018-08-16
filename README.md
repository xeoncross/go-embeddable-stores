pogreb-bench
============

pogreb-bench is a key-value store benchmarking tool. Currently it supports pogreb, goleveldb, bolt and badgerdb.

# Embeddable Go Databases

## boltdb - https://github.com/boltdb/bolt

Solid, well-tested, development is locked. Slowest of the bunch. Largest file size.

## bbolt - https://github.com/coreos/bbolt

Not tested.

## pogreb - https://github.com/akrylysov/pogreb

Optimized store for random lookups. *No support for range/prefix scans*. Slow in this benchmark. [More...](https://artem.krylysov.com/blog/2018/03/24/pogreb-key-value-store/)

## goleveldb - https://github.com/syndtr/goleveldb

Well tested with a redis-like list/hash abstraction above it ([ledisdb](http://ledisdb.com/)) for more complex usage. Smallest storage size.

## badgerdb - https://github.com/dgraph-io/badger

Fastest engine for random lookups and inserts. [More...](https://blog.dgraph.io/post/badger/)

## SQLite

Not tested.

## Other benchmarks

- [badgerdb, goleveldb, boltdb benchmark](https://github.com/zchee/go-benchmarks/blob/master/db/db_bench_test.go) by zchee.
- [badgerdb, goleveldb, boltdb, benchmark](https://github.com/dgraph-io/badger-bench) by dgraph.
