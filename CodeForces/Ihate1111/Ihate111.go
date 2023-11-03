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

	//Чикне макнагет теорем Любое число больше (m*n - (m+n)) может состоять из некого набора m и n (x*m + y*n) x и y некие значения.

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x > 1099 {
			fmt.Fprintln(out, "YES")
		} else {
			for x > 0 {
				if x%11 == 0 {
					fmt.Fprintln(out, "YES")
					x = 0
				} else {
					x -= 111
					if x == 0 {
						fmt.Fprintln(out, "YES")
					}
				}
			}
			if x != 0 {
				fmt.Fprintln(out, "NO")
			}
		}
	}
}
