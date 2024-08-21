package LeetCode

func PivotIndex(nums []int) int {
	result := -1
	arraySum := make([]int, len(nums))
	if len(nums) >= 1 {
		arraySum[0] = nums[0]
	}
	for i := 1; i < len(nums); i++ {
		arraySum[i] = nums[i] + arraySum[i-1]
	}

	for i := 0; i < len(nums); i++ {
		if float64(arraySum[i])-float64(nums[i]) == float64(arraySum[len(arraySum)-1]-nums[i])/float64(2) {
			result = i
			break
		}
	}

	return result
}
