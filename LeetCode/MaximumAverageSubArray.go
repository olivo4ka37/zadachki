package LeetCode

import "fmt"

func FindMaxAverage(nums []int, k int) float64 {
	var result, currentSum int = 0, 0

	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}

	lastInt := nums[0]
	result = currentSum

	fmt.Println("result is:", result, "lastInt is:", lastInt)

	for i := k; i < len(nums); i++ {
		if result-lastInt+nums[i] > result {
			result += -lastInt + nums[i]
		}
		fmt.Println("i is:", i, "result is:", result, "lastInt is:", lastInt, "nums[i] is:", nums[i])
		lastInt = nums[i-k+1]
	}

	return float64(result) / float64(k)
}
