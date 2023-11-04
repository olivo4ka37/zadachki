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

	var n, xPoints int

	fmt.Fscan(in, &n)
	for c := 0; c < n; c++ {
		result1 := 0
		fmt.Fscan(in, &xPoints)
		xyArray := make([]int, 2*xPoints)
		for i := 0; i < xPoints; i++ {
			fmt.Fscan(in, xyArray[i], xyArray[i+1])
			i++
		}

		sort.Slice(xyArray, func(i, j int) bool {
			return xyArray[i] > xyArray[j]
		})

		for i := 0; i < len(xyArray); i++ {
			result1 += xyArray[i] - xyArray[i+1]
			i++
		}

		if result1 < 0 {
			result1 *= -1
		}

		fmt.Fprintln(out, result1)
		l := 0
		r := len(xyArray) - 1
		for l < r {
			fmt.Fprintln(out, xyArray[l], xyArray[r])
			l++
			r--
		}
	}
}
