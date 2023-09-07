package main

import "fmt"

func main() {
	// Задание размерности n
	n := 10

	// Создание и инициализация двумерного массива размерности n x 2
	matrix := make([][]int, n)
	fmt.Println(matrix)
	for i := range matrix {
		matrix[i] = make([]int, 2)
		matrix[i][0] = i * 2
		matrix[i][1] = i*2 + 1
	}
	fmt.Println(matrix)

	// Вывод двумерного массива
	for i := 0; i < n; i++ {
		fmt.Printf("%d: %d %d\n", i, matrix[i][0], matrix[i][1])
	}
}
