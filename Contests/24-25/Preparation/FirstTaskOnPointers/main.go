package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Дан массив целых чисел, отсортированный по возрастанию.
Необходимо вернуть массив из квадратов каждого числа, отсортированного по возрастанию.
*/

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, i, x, index int16
	n = 5

	result := make([]int16, 0, n)
	minusArray := make([]int16, 0, n)
	initializedIndex := false
	for i = 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x < 0 {
			x = x * x
			minusArray = append(minusArray, x)
		} else if x == 0 {
			result = append(result, x)
			if !initializedIndex {
				index = int16(len(minusArray)) - 1
			}
			initializedIndex = true
		} else {
			if !initializedIndex {
				index = int16(len(minusArray)) - 1
			}
			initializedIndex = true
			x = x * x
			if index >= 0 {
				for index >= 0 && minusArray[index] <= x {
					result = append(result, minusArray[index])
					index--
				}
			}
			result = append(result, x)
		}
	}
	for index >= 0 {
		result = append(result, minusArray[index])
		index--
	}

	fmt.Println(minusArray)
	fmt.Println(result)
}
