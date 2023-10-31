package LeetCode

func LongestSubarray(nums []int) int {
	var result, part1, part2, zerosCount, i int

	for i < len(nums) && nums[i] == 0 {
		i++
	}

	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			zerosCount++
			if zerosCount == 2 && nums[i-1] == 0 {
				part1 = 0
				part2 = 0
				zerosCount = 0
			} else if zerosCount == 2 {
				part1 = part2
				part2 = 0
				zerosCount = 1
			}
			continue
		}

		if zerosCount == 0 {
			part1++
		} else {
			part2++
		}

		if result < part1+part2 {
			result = part1 + part2
		}
		//fmt.Println("i is:",i,"zeros is:",zerosCount,"part1 is:",part1,"part2 is:",part2,"result is:",result)
	}

	if result == len(nums) {
		return result - 1
	}

	return result
}
