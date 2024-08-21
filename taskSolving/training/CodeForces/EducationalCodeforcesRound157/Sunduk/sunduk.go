package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, x, y, result, k int

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &y, &k)
		if y < x {
			result = x
			fmt.Fprintln(out, result)
		} else {
			if x+k >= y {
				result = y
				fmt.Fprintln(out, result)
			} else {
				result = 2*y - x - k
				fmt.Fprintln(out, result)
			}
		}
	}
}
