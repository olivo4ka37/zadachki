package CodeWars

import "sort"

func Solution(ar []int) int {
	xArray := make([]int, 0, 1000000)

	for i := 0; i < len(ar); i++ {
		xArray = append(xArray, ar[i])
		if ar[i] == 1 {
			return 1 * len(ar)
		}
	}

	sort.Slice(xArray, func(i, j int) bool { // Сортировка массива по убыванию
		return xArray[i] > xArray[j]
	})

	var minOstatok int = int(^uint(0) >> 1)
	for i := 0; i < len(xArray)-1; i++ {
		for j := i + 1; j < len(xArray); j++ {
			if xArray[i]%xArray[j] != 0 && xArray[i]%xArray[j] < minOstatok { //Нахождение минимального остатка от целочисленного деления
				minOstatok = xArray[i] % xArray[j]
			}
		}
	}

	// Искать от меньшего к большему и добавить переменную по которой будет производиться не с нулевым элементом, а со всеми элементами больше нашего элемента

	if minOstatok == int(^uint(0)>>1) { // Если минимальная Разность осталось неизменной, то это значит что массив равный, и в нем нечего менять
		return xArray[0] * len(xArray)
	}

	result := len(xArray) * minOstatok
	return result
}
