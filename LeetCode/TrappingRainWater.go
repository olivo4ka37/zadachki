package LeetCode

func Trap(height []int) int {
	result, left, right, leftMax, rightmax := 0, 0, len(height)-1, 0, 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				result += (leftMax - height[left])
			}
			left++
		} else {
			if height[right] >= rightmax {
				rightmax = height[right]
			} else {
				result += (rightmax - height[right])
			}
			right--
		}

	}

	return result
}
