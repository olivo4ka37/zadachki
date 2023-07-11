package LeetCode

//Там дети с конфетами. Вам дается целочисленный массив candies, где каждая candies[i]
//представляет количество конфет, которые есть у i-го ребенка,
//и целое число extraCandies, обозначающее количество дополнительных конфет, которые есть у вас.
//Возвращает результат логического массива длины n, где результат[i] равен true, если после предоставления
//i-му ребенку всех дополнительных конфет
//у него будет наибольшее количество конфет среди всех детей, или false в противном случае.
//Обратите внимание, что у нескольких детей может быть наибольшее количество конфет.

func KidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	max := candies[0]

	for i := 1; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}
