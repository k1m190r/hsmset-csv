package main

import (
	"encoding/csv"
	"flag"
	"os"
)

func main() {

	flag.Parse()
	files := flag.Args()

	for _, fname := range files {
		f, err := os.Open(fname)
		procFile(f)
	}
}

func procFile() {
}
