package LeetCode

func MaxArea(height []int) int {
	maxSpace, currentSpace, i, j := 0, 0, 0, len(height)-1

	for i < j {
		if height[i] > height[j] {
			currentSpace = height[j] * (j - i)
			if maxSpace < currentSpace {
				maxSpace = currentSpace
			}
			j--
		} else {
			currentSpace = height[i] * (j - i)
			if maxSpace < currentSpace {
				maxSpace = currentSpace
			}
			i++
		}
	}

	return maxSpace
}
