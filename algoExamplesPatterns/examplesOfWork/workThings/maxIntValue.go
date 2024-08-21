package main

import "fmt"

func main() {
	x := 9223372036854775807
	x++
	fmt.Println(x)
	x2 := 9223372036854775807
	x2 += 202
	fmt.Println(x2)
}
