package main

import "fmt"

type Direction int

func Iota() {
	type Direction int

	const (
		North Direction = iota
		East
		South
		West
	)

	var d Direction = North
	fmt.Print(d)
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
}

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}
