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

	var n, sum int

	fmt.Fscan(in, &n)
	tovars := [10001]int{}
	for i := 0; i < n; i++ {
		var mincost int = 10002
		var maxcost int = 0
		var n1 int
		fmt.Fscan(in, &n1)
		for j := 0; j < n1; j++ {
			var cost int
			fmt.Fscan(in, &cost)
			tovars[cost]++
			if tovars[cost]%3 == 0 {
				sum += cost * 2
				tovars[cost] = 0
			}
			if mincost > cost {
				mincost = cost
			}
			if maxcost < cost {
				maxcost = cost
			}
		}
		for j := mincost; j <= maxcost; j++ {
			sum += tovars[j] * j
			tovars[j] = 0
		}
		fmt.Println(sum)
		sum = 0
		mincost = 10002
		maxcost = 0
	}
}
