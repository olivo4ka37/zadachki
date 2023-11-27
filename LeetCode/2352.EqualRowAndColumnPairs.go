package LeetCode

func EqualPairs(grid [][]int) int {
	result := 0
	flag := true

	mirrorArray := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		mirrorArray[i] = make([]int, len(grid))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			mirrorArray[i][j] = grid[j][i]
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			flag = equal(mirrorArray[i], grid[j])

			if flag == true {
				result++
			} else {
				flag = true
			}
		}
	}
	return result
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
