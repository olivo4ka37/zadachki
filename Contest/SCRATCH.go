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

	var q, k uint64
	var i, j uint64
	var maxSum uint64 = 0
	var maxLen uint64 = 30002

	fmt.Fscan(in, &q)

	hashMap := make([]bool, maxLen)
	for i = 0; i < maxLen; i++ {
		hashMap[i] = false
	}
	result := make([]uint64, 0)

	for ij := 0; ij < int(q); ij++ {
		fmt.Fscan(in, &k)
		hashMap[k] = true
		for i = 0; i < maxLen; i++ {
			for j = 0; j < maxLen; j++ {
				if i != j && hashMap[i] == true && hashMap[j] == true && i^j > maxSum {
					maxSum = i ^ j
				}
			}
		}
		result = append(result, maxSum)
	}
	for i := range result {
		fmt.Println(result[i])
	}
}
