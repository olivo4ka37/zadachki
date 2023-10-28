package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var j, a, b int

	fmt.Fscan(in, &j)

	for i := 0; i < j; i++ {
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(out, a+b)
		a, b = 0, 0
	}

	defer out.Flush()
}
