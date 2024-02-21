package LeetCode

//You have a long flowerbed in which some of the plots are planted, and some are not.
//However, flowers cannot be planted in adjacent plots.
//Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n,
//return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

//Input: flowerbed = [1,0,0,0,1], n = 1
//Output: true

//Input: flowerbed = [1,0,0,0,1], n = 2
//Output: false

func CanPlaceFlowers(flowerbed []int, n int) bool {
	canPlaced := 0

	if len(flowerbed) == 1 {
		if flowerbed[0] == 1 {
			if n == 0 {
				return true
			} else {
				return false
			}
		} else {
			if n <= 1 {
				return true
			} else {
				return false
			}
		}
	}

	for i := 0; i < len(flowerbed); i++ {
		if i == 0 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				canPlaced++
				flowerbed[i] = 1
			}
		} else if i == len(flowerbed)-1 {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 {
				canPlaced++
				flowerbed[i] = 1
			}
		} else {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				canPlaced++
				flowerbed[i] = 1
			}
		}
	}

	if canPlaced-n >= 0 {
		return true
	}

	return false
}

//#POZOR
