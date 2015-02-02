package main

import (
	"flag"
	"log"

	"github.com/aodin/spider404/crawl"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 5, "max number of concurrent requests")
	flag.Parse()

	// TODO multiple root args
	if len(flag.Args()) == 0 {
		log.Fatal("Specify a starting URL as the first argument")
	}

	// TODO set the output method here?
	if err := crawl.Site(flag.Args()[0], n); err != nil {
		log.Panic(err)
	}
}
