package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var x, n, m int
	var result bool
	fmt.Fscan(in, &n, &m)

	arr1 := make([]int, n)
	kartmap := make(map[int]bool, m)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		result = false
		for j := x + 1; j < m+1; j++ {
			if kartmap[j] == false {
				kartmap[j] = true
				result = true
				arr1[i] = j
				break
			}
		}
		if result == false {
			fmt.Fprintln(out, -1)
			defer out.Flush()
			break
		}
	}

	if result == true {
		for i := 0; i < len(arr1); i++ {
			fmt.Fprint(out, arr1[i], " ")
			defer out.Flush()
		}
	}
}
