package LeetCode

import "fmt"

func MaxSumSubmatrix(matrix [][]int, k int) int {
	fmt.Println(len(matrix), len(matrix[0]))
	fmt.Println(matrix[0][0])
	prefix2DArray := make([][]int, len(matrix), len(matrix[0]))
	result := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 && j == 0 {
				prefix2DArray[i][j] = matrix[i][j]
			} else if i == 0 && j != 0 {
				prefix2DArray[i][j] = prefix2DArray[i][j-1] + matrix[i][j]
			} else {
				prefix2DArray[i][j] = prefix2DArray[i-1][j-1] + matrix[i][j]
			}

		}
	}

	return result
}
