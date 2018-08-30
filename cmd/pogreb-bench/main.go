package main

import (
	_ "expvar"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	engine       = flag.String("e", "pogreb", "database engine name. pogreb, goleveldb, bolt or badgerdb")
	numKeys      = flag.Int("n", 5000, "number of keys")
	minKeySize   = flag.Int("mink", 32, "minimum key size")
	maxKeySize   = flag.Int("maxk", 64, "maximum key size")
	minValueSize = flag.Int("minv", 128, "minimum value size")
	maxValueSize = flag.Int("maxv", 1024, "maximum value size")
	concurrency  = flag.Int("c", 3, "number of concurrent goroutines")
	dir          = flag.String("d", "data", "database directory")
	progress     = flag.Bool("p", false, "show progress")
)

func main() {
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	flag.Parse()
	if *dir == "" {
		flag.Usage()
		return
	}

	// Setup monitoring
	go http.ListenAndServe(":1234", http.DefaultServeMux)
	time.Sleep(time.Second)

	fmt.Printf("Number of keys: %d\n", *numKeys)
	fmt.Printf("Minimum key size: %d, maximum key size: %d\n", *minKeySize, *maxKeySize)
	fmt.Printf("Minimum value size: %d, maximum value size: %d\n", *minValueSize, *maxValueSize)
	fmt.Printf("Concurrency: %d\n", *concurrency)
	fmt.Println()

	for engine := range engines {
		if err := benchmark(engine, *dir, *numKeys, *minKeySize, *maxKeySize, *minValueSize, *maxValueSize, *concurrency, *progress); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Println()
		time.Sleep(time.Second * 3)
	}
}
