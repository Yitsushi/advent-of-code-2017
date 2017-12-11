package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Abs_int64(x int64) int64 {
	if x < 0 {
		return -x
	}

	return x
}

func Max_int64(x, y, z int64) int64 {
	var max int64
	if x > y {
		max = x
	} else {
		max = y
	}
	if z > max {
		max = z
	}
	return max
}

type Position struct {
	x int64
	y int64
	z int64
}

func (p *Position) Distance() int64 {
	return (Abs_int64(p.x) + Abs_int64(p.y) + Abs_int64(p.z)) / 2
	//return Max_int64(Abs_int64(p.x), Abs_int64(p.y), Abs_int64(p.z))
}

func (p *Position) Step(d *Position) {
	p.x += d.x
	p.y += d.y
	p.z += d.z
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf("First argument must be a file.")
		os.Exit(1)
	}

	content, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	list := strings.Split(strings.TrimSpace(string(content)), ",")

	directionMap := map[string]*Position{
		"ne": &Position{1, 0, -1}, "se": &Position{1, -1, 0}, // x++
		"n": &Position{0, 1, -1}, "nw": &Position{-1, 1, 0}, // y++
		"s": &Position{0, -1, 1}, "sw": &Position{-1, 0, 1}, // z++
	}

	pos := &Position{}
	var max int64 = 0
	for _, direction := range list {
		pos.Step(directionMap[direction])
		if max < pos.Distance() {
			max = pos.Distance()
		}
	}

	fmt.Printf("Final: %d\n", pos.Distance())
	fmt.Printf("Max: %d\n", max)
}
