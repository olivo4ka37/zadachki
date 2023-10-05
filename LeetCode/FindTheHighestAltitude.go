package LeetCode

func LargestAltitude(gain []int) int {
	result, currentAltitude := 0, 0
	for i := 0; i < len(gain); i++ {
		currentAltitude += gain[i]
		if currentAltitude > result {
			result = currentAltitude
		}
	}

	return result
}
