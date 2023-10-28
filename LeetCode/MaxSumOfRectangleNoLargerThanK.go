package LeetCode

import (
	"fmt"
)

// I will solve this later. LeetCode problem number 363.
func MaxSumSubmatrix(matrix [][]int, k int) int {
	fmt.Println(len(matrix), len(matrix[0]))
	fmt.Println(matrix[0][0])
	prefix2DArray := make([][]int, len(matrix))
	for x := range prefix2DArray {
		prefix2DArray[x] = make([]int, len(matrix[0]))
	}
	result := 0

	fmt.Println("len1", len(prefix2DArray), "len2", len(prefix2DArray[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j == 0 {
				prefix2DArray[i][j] = matrix[i][j]
			} else {
				prefix2DArray[i][j] = prefix2DArray[i][j-1] + matrix[i][j]
			}
			fmt.Print(prefix2DArray[i][j], ",")
		}
		fmt.Println()
	}

	return result
}
