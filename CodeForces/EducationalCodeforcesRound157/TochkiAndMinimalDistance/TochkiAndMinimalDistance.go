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
		result1, result2 := 0, 0
		fmt.Fscan(in, &xPoints)
		xyArray := make([]int, 2*xPoints)
		for i := 0; i < len(xyArray); i++ {
			fmt.Fscan(in, &xyArray[i])
		}

		sort.Slice(xyArray, func(i, j int) bool {
			return xyArray[i] > xyArray[j]
		})

		//fmt.Println("xyArray is:", xyArray)
		l := 0
		r := len(xyArray) / 2
		//fmt.Println("r is:", r)
		for r < len(xyArray)-1 {
			result1 += xyArray[l] - xyArray[l+1]
			result2 += xyArray[r] - xyArray[r+1]
			l++
			r++
		}

		if result1 < 0 {
			result1 *= -1
		}

		if result2 < 0 {
			result2 *= -1
		}

		result1 += result2

		fmt.Fprintln(out, result1)
		l = 0
		r = len(xyArray) - 1
		for l < r {
			fmt.Fprintln(out, xyArray[l], xyArray[r])
			l++
			r--
		}
	}
}
