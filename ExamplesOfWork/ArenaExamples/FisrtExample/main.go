//go:build goexperiment.arenas
// +build goexperiment.arenas

package main

import (
	"arena"
	"fmt"
)

type Person struct {
	Lastname  string
	Firstname string
}

func main() {
	mem := arena.NewArena()
	defer mem.Free()

	for i := 0; i < 10; i++ {
		obj := arena.New[Person](mem)
		print(obj)
	}

	fmt.Println("123")
}
