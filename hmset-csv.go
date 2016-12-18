package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var buf bytes.Buffer
var cs int
var cf *os.File

func checkErr(err error) {
	if err != nil {
		log.Fatal()
	}
}

func main() {

	flag.Parse()
	files := flag.Args()

	for _, fname := range files {
		f, err := os.Open(fname)
		checkErr(err)

		cf = f
		cs = 0
		buf.Reset()

		procFile(f)
		f.Close()
	}
}

func procFile(f *os.File) {
	r := csv.NewReader(bufio.NewReader(f))

	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal()
		}

		addRec(rec)
	}
	send2redis()
}

func redStr(s string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
}

// build redis proto string and send to write buffer
func addRec(rec []string) {
	if len(rec) != 2 {
		return
	}

	for i, s := range rec {
		rec[i] = strings.TrimSpace(s)
	}

	if rec[0] == "" || rec[1] == "" {
		return
	}

	for _, s := range rec {
		buf.WriteString(redStr(s))
		cs += 1
	}

}

func hmset(s string) string {
	var ret bytes.Buffer

	ret.WriteString(fmt.Sprintf("*%d\r\n", cs+2)) // +2 for HMSET KEY
	ret.WriteString(redStr("HMSET"))
	fn := strings.Replace(filepath.Base(cf.Name()), ",", ":", -1)
	name := strings.TrimSuffix(fn, filepath.Ext(fn))
	ret.WriteString(redStr(name))
	ret.WriteString(s)

	rs := ret.String()

	return rs
}

// count the strings, add hmset key and send to stdout
func send2redis() {
	vals := buf.String()
	fmt.Printf(hmset(vals))
}
