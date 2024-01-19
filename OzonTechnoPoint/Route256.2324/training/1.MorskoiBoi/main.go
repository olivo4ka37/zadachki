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

	var n int

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var c1, c2, c3, c4 int
		for z := 0; z < 10; z++ {
			var x int
			fmt.Fscan(in, &x)
			switch x {
			case 1:
				c1++
			case 2:
				c2++
			case 3:
				c3++
			case 4:
				c4++
			}
		}
		if c1 > 4 || c2 > 3 || c3 > 2 || c4 > 1 {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}
