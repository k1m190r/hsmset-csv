package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var args = make([]string, 0, 100)

func makeRESP(args []string) (ret string) {
	grow := []string{"*"}
	grow = append(grow, strconv.Itoa(len(args)))
	// TODO grow $len(args[i])args[i]
	return ret
}

func newKey(key string, args []string) {
	switch len(args) {
	case 0: // args is empty
	default:
		makeRESP(args)
		args = args[0:1] // zero out an array
		args[0] = key
	}
}

func procRec(rec []string, out io.Writer) {
	if out == nil {
		out = os.Stdout
	}

	// trim spaces
	nrec := make([]string, len(rec))
	for i, s := range rec {
		nrec[i] = strings.TrimSpace(s)
	}

	switch nrec[0] {
	case "":
		if (nrec[1] == "") || (nrec[2] == "") {
			break
		}
		args = append(args, nrec[1:3]...)
	default:
		newKey(nrec[0], args)
	}
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

			procRec(record, nil)
		}
	}

}
