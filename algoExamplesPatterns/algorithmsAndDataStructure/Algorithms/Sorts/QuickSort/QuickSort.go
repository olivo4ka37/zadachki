package main

import (
	"fmt"
)

// QuickSort выполняет сортировку массива целых чисел методом быстрой сортировки
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]
	left := []int{}
	right := []int{}

	for i := range arr {
		if i == pivotIndex {
			continue
		}
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	QuickSort(left)
	QuickSort(right)

	copy(arr, append(append(left, pivot), right...))
}

func main() {
	// Пример использования QuickSort
	data := []int{-13, 413, 24, 124, 12, 1, 15632, 234, 85, 753, 9, 13, 3415, 864, 1, 97, 315, 15, -12, 21521}
	fmt.Println("Исходный массив:", data)

	QuickSort(data)

	fmt.Println("Отсортированный массив:", data)
}
