package LeetCode

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	//fmt.Println("nums is:", nums)

	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//fmt.Println("i is:", i, "j is:", j, "k is:", k)
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[j]+nums[k]+nums[i] <= 0 {
				k--
			} else {
				j++
			}
			//fmt.Println("j is:", j, "k is:", k, "RESULT IS:", result)
		}
	}

	return result
}
