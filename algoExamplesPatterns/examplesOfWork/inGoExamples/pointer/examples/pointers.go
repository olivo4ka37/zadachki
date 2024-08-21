package main

import "fmt"

const year = 2024

func main() {
	//Указатели
	a := 2
	b := &a
	a = 3 // a = 3
	*b = 4
	fmt.Println(b, &b, &a, *b, a)

	// * - Создает новый указатель или получает значение из указателя.
	// & - Получает адрес переменной.
}
