package main

import "fmt"

func main() {
	m := make(map[int]int, 100)
	for i := 0; i < 10; i++ {
		m[i] = i
	}

	fmt.Println(m)
	for i, v := range m {
		fmt.Println(i, v)
	}

	a := []int{1, 2}
	s(a...)
}

func s(a ...int) {
	for _, v := range a {
		fmt.Println(v)
	}
}
