package LeetCode

func FindMaxAverage(nums []int, k int) float64 {
	var result, currentSum int = 0, 0
	for i := 0; i < len(nums); i++ {
		if i < k {
			currentSum += nums[i]
			result = currentSum
		} else {
			currentSum += -nums[i-k] + nums[i]
			if currentSum > result {
				result = currentSum
			}
		}
	}

	return float64(result) / float64(k)
}
