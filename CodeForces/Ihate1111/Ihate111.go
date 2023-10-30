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

	var n, x int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		x1 := 111111111
		fmt.Fscan(in, &x)
		for x1 > 1 {
			if x > x1 {
				x -= x1
			} else if x == x1 {
				fmt.Fprintln(out, "YES")
				x = 0
				break
			} else {
				x1 /= 10
			}
		}
		if x != 0 {
			fmt.Fprintln(out, "NO")
		}
	}

}
