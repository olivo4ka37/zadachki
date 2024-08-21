package LeetCode

func LongestOnes(nums []int, k int) int {
	result, zeroCount, xLen, startPoint := 0, 0, 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && zeroCount < k {
			zeroCount++
			xLen++
		} else if nums[i] == 0 && zeroCount == k {
			if xLen > result {
				result = xLen
			}
			xLen = 0
			zeroCount = 0
			startPoint++
			if i != startPoint-1 {
				i = startPoint
			}
		} else {
			xLen++
		}
	}

	if xLen > result {
		result = xLen
	}

	return result
}
