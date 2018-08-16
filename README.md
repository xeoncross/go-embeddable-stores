pogreb-bench
============

pogreb-bench is a key-value store benchmarking tool. Currently it supports pogreb, goleveldb, bolt and badgerdb.

# Embeddable Go Databases

### boltdb - https://github.com/boltdb/bolt

Solid, well-tested, development is locked. Slowest of the bunch. Largest file size.

### pogreb - https://github.com/akrylysov/pogreb

Optimized store for random lookups. *No support for range/prefix scans*. Slow in this benchmark. [More...](https://artem.krylysov.com/blog/2018/03/24/pogreb-key-value-store/)

### goleveldb - https://github.com/syndtr/goleveldb

Well tested with a redis-like list/hash abstraction above it ([ledisdb](http://ledisdb.com/)) for more complex usage. Smallest storage size.

### badgerdb - https://github.com/dgraph-io/badger

Fastest engine for random lookups and inserts. [More...](https://blog.dgraph.io/post/badger/)

# TODO

### bbolt - https://github.com/coreos/bbolt

Not tested.

### tiedot - https://github.com/HouzuoGuo/tiedot/

Not tested.

### SQLite

Not tested. Multiple versions exist, most wrap the C code.

`database/sql` driver
- https://github.com/mattn/go-sqlite3

no `database/sql` driver
- https://github.com/crawshaw/sqlite
- https://github.com/bvinc/go-sqlite-lite


Comparison between [bvinc/go-sqlite-lite & crawshaw/sqlite](https://www.reddit.com/r/golang/comments/96yd0t/gosqlitelite_a_new_light_weight_sqlite_package/e44eoym/).

# Non-Go Embeddable Databases (Cgo)

- RocksDB
- LMDB
- hyperleveldb


# Other benchmarks

- [badgerdb, goleveldb, boltdb benchmark](https://github.com/zchee/go-benchmarks/blob/master/db/db_bench_test.go) by zchee.
- [badgerdb, goleveldb, boltdb, rocksdb benchmark](https://github.com/dgraph-io/badger-bench) by dgraph. (High-Quality)
- [goleveldb, boltdb, pogreb write/put benchmark](https://gist.github.com/mattn/3990033f7bc8a57cd5b86edefb254332) by mattn.
