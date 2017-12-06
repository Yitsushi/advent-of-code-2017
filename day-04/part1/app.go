package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Unique(m []string) []string {
	n := make([]string, 0)
	ref := make(map[string]bool, len(m))
	for _, v := range m {
		if _, ok := ref[v]; !ok {
			ref[v] = true
			n = append(n, v)
		}
	}
	return n
}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf("First argument must be a file.")
		os.Exit(1)
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	defer file.Close()

	r := bufio.NewReader(file)
	valid := 0
	invalid := 0
	for s, e := Readln(r); e == nil; s, e = Readln(r) {
		words := strings.Fields(s)
		if len(Unique(strings.Fields(s))) == len(words) {
			valid++
		} else {
			invalid++
		}
	}
	fmt.Printf("Valid: %d, Invalid: %d\n", valid, invalid)
}
