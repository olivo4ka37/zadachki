package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var minValue1 int = 0
	var minValue2 int = 0

	previous := Lottery(n, 1)
	for i := 1; i < n/2+1; i++ {

		lot := Lottery(i, n-i)

		if lot < previous {
			minValue1 = i
			minValue2 = n - i
			previous = lot
		}
	}
	fmt.Println(minValue1, " ", minValue2)
}

func Lottery(x, y int, integers ...int) int {

	itog := x * y / GCD(x, y)

	for i := 0; i < len(integers); i++ {
		itog = Lottery(itog, integers[i])
	}

	return itog
}

func GCD(x, y int) int {

	for y != 0 {
		t := y
		y = x % y
		x = t
	}

	return x
}
