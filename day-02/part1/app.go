package main

import (
	"bufio"
	"fmt"
	"math"
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
	var min int64 = math.MaxInt64
	var max int64 = math.MinInt64
	for _, n := range line {
		x, _ := strconv.ParseInt(n, 10, 64)
		if min > x {
			min = x
		}
		if max < x {
			max = x
		}
	}

	return max - min
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
