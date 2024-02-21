package main

import "fmt"

func main() {
	var a, b, c, x, y, z int

	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Scanf("%d %d %d", &x, &y, &z)

	nx := x/a + y/b + z/c + 2
	kff := 2
	fmt.Println(factr(nx) / factr(nx-kff) / factr(kff))
}

func factr(n int) uint64 {
	var Value uint64 = 1
	if n < 0 {
		return 0
	} else {
		for i := 1; i <= n; i++ {
			Value *= uint64(i)
		}

	}
	return Value
}
