package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var n, z, x, c, result int

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &z, &x, &c)
		if z+x+c > 1 {
			result++
		}
	}

	fmt.Fprintln(out, result)

	defer out.Flush()
}
