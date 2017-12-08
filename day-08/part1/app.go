package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var Memory map[string]int64 = map[string]int64{}

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

func test(register, op, value string) bool {
	var r int64 = 0
	var ok bool = false
	if r, ok = Memory[register]; !ok {
		Memory[register] = 0
		r = 0
	}

	c, _ := strconv.ParseInt(value, 10, 64)

	switch op {
	case ">":
		return r > c
	case "<":
		return r < c
	case "==":
		return r == c
	case "<=":
		return r <= c
	case ">=":
		return r >= c
	case "!=":
		return r != c
	}

	return false
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

	var max int64 = int64(math.Inf(-64))

	r := bufio.NewReader(file)
	for s, e := Readln(r); e == nil; s, e = Readln(r) {
		words := strings.Fields(s)
		if _, ok := Memory[words[0]]; !ok {
			Memory[words[0]] = 0
		}

		if test(words[4], words[5], words[6]) {
			c, _ := strconv.ParseInt(words[2], 10, 64)
			if words[1] == "dec" {
				c = -c
			}
			Memory[words[0]] += c
		}
	}

	for _, v := range Memory {
		if v > max {
			max = v
		}
	}

	fmt.Printf("Maximum value: %d\n", max)
}
