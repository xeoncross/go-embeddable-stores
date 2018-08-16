# Transactions & Batches

BoltDB / bbolt are fast when spawning 1k goroutines inserting records as shown
in this benchmark. Memory usage stays around 20MB (Alloc) for me while only
taking 10 seconds to insert 1m key/values.
