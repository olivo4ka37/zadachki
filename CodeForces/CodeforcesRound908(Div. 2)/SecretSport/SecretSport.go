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

	var n, strLen int
	var str string

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		//aCount := 0
		//bCount := 0
		lessLen := 0
		fmt.Fscan(in, &strLen)
		fmt.Fscan(in, &str)
		for j := len(str) - 1; j > 0; j-- {
			if str[j] != str[j-1] {
				lessLen++
			} else {
				break
			}
		}

	}
}
