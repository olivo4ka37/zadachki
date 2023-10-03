package LeetCode

func Trap(height []int) int {
	result, j, i, currentWater, leftMax := 0, 0, 0, 0, 0

	for j < len(height) {
		if leftMax < height[j] {
			leftMax = height[j]
		}
		if i == len(height) {
			currentWater = 0
			j++
			i = j
		} else if height[j] <= height[i] && height[j] >= leftMax {
			result += currentWater
			currentWater = 0
			j = i
			i++
		} else {
			currentWater += height[j] - height[i]
			i++
		}
		//fmt.Println("j is:",j,"i is:",i,"result is:", result)
	}

	//fmt.Println("RESULT IS:", result)

	rightMax, secondResult := 0, 0
	j, i = len(height)-1, len(height)-1

	for j > 0 {
		if rightMax < height[j] {
			rightMax = height[j]
		}
		if i == -1 {
			currentWater = 0
			j--
			i = j
		} else if height[j] < height[i] && height[j] >= rightMax {
			secondResult += currentWater
			currentWater = 0
			j = i
			i--
		} else {
			currentWater += height[j] - height[i]
			i--
		}
		//fmt.Println("j is:",j,"i is:",i,"result is:", result)
	}

	return result + secondResult
}
