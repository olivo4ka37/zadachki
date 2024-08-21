package LeetCode

func SingleNumber(nums []int) int {
	mapByte := make([]bool, 60001)
	result := 0

	for _, i := range nums {
		if mapByte[i+30000] == false {
			mapByte[i+30000] = true
		} else {
			mapByte[i+30000] = false
		}
		//fmt.Println("i is:",i,"mapByte[i+30000] is:",mapByte[i+30000])
	}

	for i := 0; i < len(mapByte); i++ {
		if mapByte[i] == true {
			result = i - 30000
		}
		//fmt.Println("i is:",i,"result is:",result)
	}

	return result
}
