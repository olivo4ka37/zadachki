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
		var lenN, ix int
		var compressionON bool = false
		var differenceX int
		result := make([]int, lenN)
		fmt.Fscan(in, &lenN)

		for j := 0; j < lenN; j++ {
			var x int
			fmt.Fscan(in, &x)
			if !compressionON {
				result[ix] = x
				compressionON = true
			} else {
				differenceX = result[ix] - x
				if differenceX == 0 || differenceX*differenceX > 1 {
					ix++
					result[ix] = 0
					ix++
					compressionON = false
				} else if differenceX == -1 {

				} else if differenceX == +1 {

				}
			}
		}

		fmt.Println(len(result))
		fmt.Println(result)
	}

}
