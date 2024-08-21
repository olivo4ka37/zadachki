package LeetCode

import (
	"fmt"
	"sort"
)

func MaxOperations(nums []int, k int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	fmt.Println("nums is:", nums)

	i, j, result := 0, len(nums)-1, 0

	for i < j {
		if nums[i]+nums[j] == k {
			i++
			j--
			result++
		} else if nums[i]+nums[j] > k {
			i++
		} else {
			j--
		}
		fmt.Println("nums[i] is:", nums[i], "nums[j] is:", nums[j], "result is:", result)
	}

	return result
}
