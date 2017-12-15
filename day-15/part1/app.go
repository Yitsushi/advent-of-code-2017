package main

import (
	"fmt"
	"os"
	"strconv"
)

type Generator struct {
	Factor uint64
	State  uint64
}

func (g *Generator) Generate() uint64 {
	g.State = (g.State * g.Factor) % 2147483647

	return g.State
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("program <generator.A> <generator.B>")
		os.Exit(1)
	}

	var param uint64
	param, _ = strconv.ParseUint(args[0], 10, 64)
	genA := Generator{16807, param}
	param, _ = strconv.ParseUint(args[1], 10, 64)
	genB := Generator{48271, param}

	found := 0
	for i := 0; i < 40000000; i++ {
		//for i := 0; i < 10; i++ {
		genA.Generate()
		genB.Generate()

		if (genA.State & 0xffff) == (genB.State & 0xffff) {
			found++
		}
	}

	fmt.Printf("Pairs: %d\n", found)
}
