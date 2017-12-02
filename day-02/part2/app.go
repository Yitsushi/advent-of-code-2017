package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func lineChecksum(line []string) int64 {
	var row []int64 = make([]int64, len(line))

	for i, n := range line {
		x, _ := strconv.ParseInt(n, 10, 64)
		row[i] = x
	}

	for i, n := range row {
		for j, m := range row {
			if i == j {
				continue
			}

			if n%m == 0 {
				return n / m
			}

			if m%n == 0 {
				return m / n
			}
		}
	}

	return 0
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

	var checksum int64 = 0

	defer file.Close()

	r := bufio.NewReader(file)
	for s, e := Readln(r); e == nil; s, e = Readln(r) {
		numbers := strings.Fields(s)
		checksum += lineChecksum(numbers)
		fmt.Println(checksum)
	}
}
