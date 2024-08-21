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
		var lenN int
		result := make([]int, lenN)
		fmt.Fscan(in, &lenN)

		var beforeX, compressCode int
		var plusFlag bool = false
		//var equalFlag bool = false
		var minusFlag bool = false

		fmt.Fscan(in, &beforeX)

		for j := 1; j < lenN; j++ {
			var x int
			fmt.Fscan(in, &x)

			if beforeX == x && plusFlag != true && minusFlag != true {
				result = append(result, x)
				result = append(result, 0)
			} else if beforeX == x+1 && minusFlag != true {
				plusFlag = true
				compressCode++
			} else if beforeX == x-1 && plusFlag != true {
				minusFlag = true
				compressCode--
			} else if plusFlag == true && beforeX != x+1 {

			}

		}
	}

	fmt.Fprintln(out, n)
}
