package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Размер массива является частью его типа данных

	var a1 [3]int
	fmt.Println(a1)
	fmt.Println(len(a1), cap(a1))

	a2 := [...]int{1, 2, 3}
	fmt.Println(a2[2])

	// Про передачу слайса по указателю в функцию
	sl := make([]int, 0, 3)
	addTenIntsNoPointer(sl)
	fmt.Println(sl)
	addTenInts(&sl)
	fmt.Println(sl)
}

// При нехватки капасити в переданном слайсе, присвоит новою область памяти исходному срезу.
func addTenInts(sl *[]int) {
	for i := 0; i < 10; i++ {
		*sl = append(*sl, rand.Intn(1000))
	}

}

// При нехватки капасити в переданном слайсе, не присвоит новою область памяти исходному срезу, все последующие измененя
// в новой области памяти не будут отображаться на исходном массиве.
func addTenIntsNoPointer(sl []int) {
	for i := 0; i < 10; i++ {
		sl = append(sl, rand.Intn(1000))
	}

}
