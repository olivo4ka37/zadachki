package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var n, k, temp int
	var uslovie string

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {

		fmt.Fscan(in, &k)
		var arrayTemp = [2]int{15, 30}
		for x := 0; x < k; x++ {
			fmt.Fscan(in, &uslovie, &temp)
			if uslovie == ">=" {
				if temp > arrayTemp[0] {
					arrayTemp[0] = temp
				}
			} else if uslovie == "<=" {
				if temp < arrayTemp[1] {
					arrayTemp[1] = temp
				}
			}

			if arrayTemp[0] > arrayTemp[1] {
				fmt.Fprintln(out, -1)
				defer out.Flush()
			} else {
				fmt.Fprintln(out, arrayTemp[0])
				defer out.Flush()
			}
		}
		fmt.Fprintln(out, "")
		defer out.Flush()
	}

}
