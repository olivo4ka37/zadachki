package main

import "fmt"

func main() {
	deferTest()
}

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic happend first time:", err)
		}
	}()

	fmt.Println("Some useful work")
	panic("Something bad happen")
	return
}
