package LeetCode

func MaximumScore(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	result, minimum := nums[k], nums[k]
	i, j := k, k
	jGran := len(nums) - 1
	//fmt.Println(nums)
	for i > 0 || j < jGran {
		if i == 0 {
			j++
			if nums[j] < minimum {
				minimum = nums[j]
			}
		} else if j == jGran {
			i--
			if nums[i] < minimum {
				minimum = nums[i]
			}
		} else if nums[i-1] >= nums[j+1] {
			i--
			if nums[i] < minimum {
				minimum = nums[i]
			}
		} else if nums[i-1] < nums[j+1] {
			j++
			if nums[j] < minimum {
				minimum = nums[j]
			}
		}

		//fmt.Println("minimum is:",minimum,"i is:", i,"j is:", j)
		xResult := minimum * (j - i + 1)
		//fmt.Println(xResult)
		if xResult > result {
			result = xResult
		}
	}

	return result
}

/*
	if nums[i-1] < nums[j+1] {
		j++
		if nums[j] < minimum {
			minimum = nums[j]
		}
	} else if nums[i-1] > nums[j+1] {
		i--
		if nums[i] < minimum {
			minimum = nums[i]
		}
	} else {
		if i != 0 {
			i--
		} else {
			j++
		}
	}
*/
