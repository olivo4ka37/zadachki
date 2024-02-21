package main

import (
	"bufio"
	"fmt"
	"os"
)

func InputFunc() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {

	}

	fmt.Fprintln(out, n)
}
