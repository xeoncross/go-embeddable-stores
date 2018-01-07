package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	engine       = flag.String("e", "pogreb", "database engine name. pogreb, goleveldb, bolt or badgerdb")
	numKeys      = flag.Int("n", 100000, "number of keys")
	minKeySize   = flag.Int("mink", 16, "minimum key size")
	maxKeySize   = flag.Int("maxk", 64, "maximum key size")
	minValueSize = flag.Int("minv", 128, "minimum value size")
	maxValueSize = flag.Int("maxv", 512, "maximum value size")
	concurrency  = flag.Int("c", 1, "number of concurrent goroutines")
	dir          = flag.String("d", "", "database directory")
	progress     = flag.Bool("p", false, "show progress")
)

func main() {
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	flag.Parse()
	if *dir == "" {
		flag.Usage()
		return
	}
	if err := benchmark(*engine, *dir, *numKeys, *minKeySize, *maxKeySize, *minValueSize, *maxValueSize, *concurrency, *progress); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
