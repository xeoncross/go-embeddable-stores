pogreb-bench
============

pogreb-bench is a benchmarking tool for databases that can be compiled with Go into a single binary. These databases are called "in-process" or "embedded".

These benchmarks are off for a variety of reasons including differing levels of durability/sync and feature-scope. I recommend people default to goleveldb for a general store. 

# Embeddable Go Databases

### boltdb - https://github.com/boltdb/bolt

Solid, well-tested, development is locked. Each transaction has a consistent view of the data as it existed when the transaction started. Slowest of the bunch. Has [Storm ORM](https://github.com/asdine/storm). Largest file size. Slowest (due to consistent-data views).

### pogreb - https://github.com/akrylysov/pogreb

Optimized store for random lookups. *No support for range/prefix scans*. Slow in this benchmark. [More...](https://artem.krylysov.com/blog/2018/03/24/pogreb-key-value-store/)

### goleveldb - https://github.com/syndtr/goleveldb

Well tested. A redis-like list/hash abstraction ([ledisdb](http://ledisdb.com/)) built with it for more complex usage. Smallest storage size.

### badgerdb - https://github.com/dgraph-io/badger

Fastest engine for random lookups and inserts. Graph engine [dgraph](https://github.com/dgraph-io/dgraph) built upon it. [More...](https://blog.dgraph.io/post/badger/)

### tiedot - https://github.com/HouzuoGuo/tiedot/

Tiedot is a document store, it really can't be compared to these other databases correctly. Included only for loose reference. Super-fast write speeds. Huge storage requirements.

# TODO

### bbolt - https://github.com/coreos/bbolt

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


# Sample Results

    Number of keys: 500000
    Minimum key size: 32, maximum key size: 64
    Minimum value size: 128, maximum value size: 1024
    Concurrency: 3

    Running tiedot benchmark...
    Put: 7.935 sec, 63012 ops/sec
    Get: 0.116 sec, 4298637 ops/sec
    Put + Get time: 8.051 sec
    File size: 1.25GB

    Running badgerdb benchmark...
    Put: 9.713 sec, 51477 ops/sec
    Get: 1.436 sec, 348216 ops/sec
    Put + Get time: 11.149 sec
    File size: 373MB

    Running goleveldb benchmark...
    Put: 22.387 sec, 22334 ops/sec
    Get: 2.158 sec, 231742 ops/sec
    Put + Get time: 24.545 sec
    File size: 306MB

    Running pogreb benchmark...
    Put: 58.528 sec, 8542 ops/sec
    Get: 0.224 sec, 2234631 ops/sec
    Put + Get time: 58.751 sec
    File size: 424MB

# Other benchmarks

- [badgerdb, goleveldb, boltdb benchmark](https://github.com/zchee/go-benchmarks/blob/master/db/db_bench_test.go) by zchee.
- [badgerdb, goleveldb, boltdb, rocksdb benchmark](https://github.com/dgraph-io/badger-bench) by dgraph. (High-Quality)
- [goleveldb, boltdb, pogreb write/put benchmark](https://gist.github.com/mattn/3990033f7bc8a57cd5b86edefb254332) by mattn.
