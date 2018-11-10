Go Embeddable Store
============

Go is a compiled language with a number of high-performance data stores  that can be compiled with Go into a single binary. These databases are called "in-process" or "embedded".

By default, most apps slap MongoDB or MySQL/MariaDB into an application for storing state. Since we left the `.net/node/php/ruby/python` world we can have our database bundled right into our application. For smaller projects, which do not need a globally distributed / sharded database, this is a great devops win.

Embedded databases are propbably not for you if:

- building a serverless app (AWS lambda / Zeit Now)
- dealing with >1 Million HTTP requests a day
- have highly concurrent request patterns

These benchmarks and tests are off for a variety of reasons including differing levels of durability/sync and feature-scope. I recommend people default to `bbolt` for a general store.

# Embeddable Go Databases

### boltdb - https://github.com/boltdb/bolt

Solid, well-tested, development is locked. Each transaction has a consistent view of the data as it existed when the transaction started. Slowest of the bunch. Largest file size. Slowest (due to consistent-data views).

- [Blast](https://github.com/mosuka/blast) - full text search and indexing server cluster via [Raft](https://github.com/hashicorp/raft)
- [Storm](https://github.com/asdine/storm) - ORM-ish layer for database access

### bbolt - https://github.com/etcd-io/bbolt

Fork of boltdb to add new features.

- [BoltHold](https://github.com/timshannon/bolthold/) - simple querying and indexing layer
- [Storm](https://github.com/asdine/storm/) - query-builder, indexes, and struct storage
- [Buckets](https://github.com/joyrexus/buckets) - simple access layer

### pogreb - https://github.com/akrylysov/pogreb

Optimized store for random lookups. *No support for range/prefix scans*. Slow in this benchmark. [More...](https://artem.krylysov.com/blog/2018/03/24/pogreb-key-value-store/)

### goleveldb - https://github.com/syndtr/goleveldb

Well tested. Smallest storage size.

- [ledisdb](http://ledisdb.com/) - redis-like abstraction over it

### badgerdb - https://github.com/dgraph-io/badger

Fastest engine for random lookups and inserts.

- Graph engine [dgraph](https://github.com/dgraph-io/dgraph) built upon it. [More...](https://blog.dgraph.io/post/badger/)
- [Cete document store](https://github.com/1lann/cete)

### keydb - https://github.com/robaho/keydb

Not tested.

# More Complex Engines

### tiedot - https://github.com/HouzuoGuo/tiedot/

Tiedot is a document store, it really can't be compared to these other databases correctly. Included only for loose reference. Super-fast write speeds. Huge storage requirements.

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

    Number of keys: 5000
    Minimum key size: 32, maximum key size: 64
    Minimum value size: 128, maximum value size: 1024
    Concurrency: 3

    Running tiedot benchmark...
    Put: 0.022 sec, 226300 ops/sec
    Get: 0.001 sec, 3552460 ops/sec
    Put + Get time: 0.024 sec
    File size: 512.00MB

    Running pogreb benchmark...
    Put: 0.142 sec, 35293 ops/sec
    Get: 0.002 sec, 2489791 ops/sec
    Put + Get time: 0.144 sec
    File size: 4.25MB

    Running goleveldb benchmark...
    Put: 0.035 sec, 144902 ops/sec
    Get: 0.009 sec, 568586 ops/sec
    Put + Get time: 0.043 sec
    File size: 3.05MB

    Running bolt benchmark...
    Put: 15.868 sec, 315 ops/sec
    Get: 0.004 sec, 1136989 ops/sec
    Put + Get time: 15.873 sec
    File size: 8.00MB

    Running bbolt benchmark...
    Put: 14.710 sec, 339 ops/sec
    Get: 0.006 sec, 874240 ops/sec
    Put + Get time: 14.716 sec
    File size: 8.00MB


## Articles

- https://www.voltdb.com/blog/2015/04/01/foundationdbs-lesson-fast-key-value-store-not-enough/
- https://petewarden.com/2010/10/01/how-i-ended-up-using-s3-as-my-database/
- https://hackernoon.com/what-i-learnt-from-building-3-high-traffic-web-applications-on-an-embedded-key-value-store-68d47249774f
- https://www.cockroachlabs.com/docs/stable/architecture/storage-layer.html
- https://github.com/kval-access-language/kval-language-specification

## Indexing data

- https://github.com/blevesearch/bleve ([Video](https://www.youtube.com/watch?v=OynPw4aOlV0))
- http://roaringbitmap.org/
- https://www.pilosa.com/docs/latest/data-model/
- https://github.com/Xeoncross/keyset

## Encoding/Decoding

If inserting complex data types (structs, maps, slices...) into a key/value store you must encode/decode them first.

1. `protobuf` provides performance, correctness and interoperability. Requires you to generate protobuf files.
2. `gob` is the next fastest format with the smallest data size. More for streaming protocols where you don't want to have to have a pre-shared schema. Poor interoperability.
2. `json` the default. Reasonable performance. Highest interoperability.
3. `xml` is to be avoided.

- https://github.com/alecthomas/go_serialization_benchmarks


# Other benchmarks

- [badgerdb, goleveldb, boltdb benchmark](https://github.com/zchee/go-benchmarks/blob/master/db/db_bench_test.go) by zchee.
- [badgerdb, goleveldb, boltdb, rocksdb benchmark](https://github.com/dgraph-io/badger-bench) by dgraph. (High-Quality)
- [goleveldb, boltdb, pogreb write/put benchmark](https://gist.github.com/mattn/3990033f7bc8a57cd5b86edefb254332) by mattn.

# Guides

- [Expiring boltdb items](http://178.62.97.106/expiring-boltdb-items/)
- [range and prefix scans in boltdb](https://bl.ocks.org/joyrexus/22c3ef0984ed957f54b9)
- [map callbacks to range/prefix scans in Boltdb](https://github.com/joyrexus/buckets/blob/master/rangescan.go)
- [building hash, set, and list using Boltdb](https://github.com/xyproto/simplebolt/blob/master/simplebolt.go)
