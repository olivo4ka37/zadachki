package LeetCode

func RemoveDuplicatesFromSortedArray26(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	result := 0
	i, j := 0, 0
	for j < len(nums) {
		//fmt.Println("111",result)
		//fmt.Println("111",nums)
		nums[i] = nums[j]
		result++
		for j < len(nums) && nums[i] == nums[j] {
			j++
		}
		i++
		//fmt.Println("222",result)
		//fmt.Println("222",nums)
	}

	return result
}
