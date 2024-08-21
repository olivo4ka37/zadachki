package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//Переменные
	hello := "Андрей"
	fmt.Println(hello)
	fmt.Println(len(hello))                    // Пишет количество байтов в строке
	fmt.Println(utf8.RuneCountInString(hello)) // Пишет количество символов в строке

	const year = 2024 // Адаптируется под тип в котором будет использоваться, только если это возможно. (Не дженерик)
	//	var str string = "s"
	//fmt.Println(str+year) - Будет выдавать ошибку компиляции

	var nextYear uint = 2025
	fmt.Println(nextYear + year)
	var month int16 = 04
	fmt.Println(month + year)

	//Указатели
	a := 2
	b := &a
	*b = 3  // a = 3
	c := &a // новый указатель на переменную

	d := new(int)
	*d = 12
	*c = *d // c = 12 -> a = 12
	*d = 13
	fmt.Println(a, *b, *c, *d)

	c = d
	*c = 14
	fmt.Println(a, *b, *c, *d)

	ar := make([]int, 1, 1)
	ar[0] = 100
	x := &ar[0]
	//fmt.Println(x, &ar[0])
	for i := 0; i < 100; i++ {
		ar = append(ar, 16)
	}
	ar[0] = 1337
	//fmt.Println(ar)
	fmt.Println(x, &ar[0])
}
