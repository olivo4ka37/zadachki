package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var t, k, n, m int
	var xSymbol rune

	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(in, &k, &n, &m)
		var arrASCII [][]byte
		for z := 0; z <= n; z++ {
			row, _ := in.ReadBytes('\n')
			arrASCII = append(arrASCII, row)
		}

		for x := 1; x < k; x++ {

		}

	}
	defer out.Flush()
}
