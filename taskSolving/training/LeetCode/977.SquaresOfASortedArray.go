package LeetCode

func SortedSquares(nums []int) []int {
	var n, i, x, index int
	n = len(nums)

	result := make([]int, 0, n)
	minusArray := make([]int, 0, n)
	initializedIndex := false
	for i = 0; i < n; i++ {
		x = nums[i]
		if x < 0 {
			x = x * x
			minusArray = append(minusArray, x)
		} else if x == 0 {
			result = append(result, x)
			if !initializedIndex {
				index = len(minusArray) - 1
			}
			initializedIndex = true
		} else {
			if !initializedIndex {
				index = len(minusArray) - 1
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
	if !initializedIndex {
		index = len(minusArray) - 1
	}
	for index >= 0 {
		result = append(result, minusArray[index])
		index--
	}

	return result
}
