package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, lenArr1 int

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &lenArr1)
		slice1 := make([]int, lenArr1)
		for j := 0; j < lenArr1; j++ {
			var iq int
			fmt.Fscan(in, &iq)
			slice1[j] = 1000*(j) + iq
		}
		sort.Ints(slice1)
		for j := 0; j < lenArr1; j++ {
			if slice1[j] > 0 {
				var minRazn int = 101
				var indx int = -1
				var razN int
				for ij := 0; ij < lenArr1; ij++ {
					if j != ij {
						razN = (slice1[j] % 1000) - (slice1[ij] % 1000)
						if razN < 0 {
							razN = -razN
						}
						if minRazn > razN {
							minRazn = razN
							indx = slice1[ij] / 1000
						}
					}
				}
				fmt.Println((slice1[j]/1000)+1, indx+1)
				slice1[j] = -100
				slice1[indx] = -100
			}
		}
		slice1 = nil
		fmt.Println("")
	}
}
