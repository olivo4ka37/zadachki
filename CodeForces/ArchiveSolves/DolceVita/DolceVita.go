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

	var n int

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var shops, bank int

		fmt.Fscan(in, &shops, &bank)

		prices := make([]int, shops)

		for j := 0; j < shops; j++ {
			fmt.Fscan(in, &prices[j])
		}

		sort.Slice(prices, func(i, j int) bool {
			return prices[i] < prices[j]
		})

		var sumPrices, result, xn int
		xn = 1

		for z := 0; z < shops; z++ {
			sumPrices += prices[z]
			if sumPrices > bank {
				break
			}
			result += (bank + xn - sumPrices) / (z + 1)
			xn++
		}
		fmt.Fprintln(out, result)
	}
}
