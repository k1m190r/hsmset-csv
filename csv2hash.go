package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func procRec(rec string) {
	// build up the redis command and feed it to the red.
}

func main() {
	flag.Parse()
	files := flag.Args()

	for x, fname := range files {
		fmt.Printf("%v, %v\n", x, fname)

		f, err := os.Open(fname)
		check(err)

		r := csv.NewReader(bufio.NewReader(f))

		for {
			record, err := r.Read()

			switch err {
			case io.EOF:
				return
			case nil:
				break
			default:
				log.Fatal(err)
			}

			procRec(record)
		}
	}

}
