package main

import (
	"flag"
	"log"
)

func main() {
	var test string
	flag.StringVar(&test, "test", "1", "set test string")
	flag.Parse()
	log.Print(test)
}