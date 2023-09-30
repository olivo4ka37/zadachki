package LeetCode

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	result, xAnswers := make([][]int, 0), 0

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	fmt.Println("nums is:", nums)

	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		fmt.Println("i is:", i, "j is:", j, "k is:", k)
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, nil)
				result[xAnswers] = make([]int, 3)
				result[xAnswers][0] = nums[i]
				result[xAnswers][1] = nums[j]
				result[xAnswers][2] = nums[k]
				j++
				k--
				xAnswers++
			} else if nums[j]+nums[k]+nums[i] < 0 {
				k--
			} else {
				j++
			}
			fmt.Println("j is:", j, "k is:", k, "RESULT IS:", result)
		}
	}

	fmt.Println("RESULT IS111:", result)

	sort.Slice(result, func(i, j int) bool {
		return result[i][1] > result[j][1]
	})

	sort.SliceStable(result, func(i, j int) bool {
		return result[i][0] > result[j][0]
	})

	fmt.Println("RESULT IS222:", result)

	for i := 0; i < len(result)-1; i++ {
		if result[i][0] == result[i+1][0] && result[i][1] == result[i+1][1] && result[i][2] == result[i+1][2] {
			result = append(result[:i], result[i+1:]...)
			i--
		}
	}

	return result
}
