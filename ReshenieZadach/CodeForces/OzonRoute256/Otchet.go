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

	var n, len int
	var itog bool
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &len)
		slice1 := make([]bool, 50002)
		var x, predX int
		fmt.Fscan(in, &predX)
		for j := 1; j < len; j++ {
			fmt.Fscan(in, &x)
			if slice1[x] == true {
				itog = true
			}
			if predX != x {
				slice1[predX] = true
				predX = x
			}
		}
		if itog == true {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
		slice1 = nil
		itog = false
	}
}
