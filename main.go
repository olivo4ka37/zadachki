package main

import "fmt"

// Задача "Асинхронное слияние каналов"
// Реализовать функцию asyncMerge, которая будет производить асинхронное слияние каналов.

// Исходные данные:
func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4

	close(ch1)
	close(ch2)

	ch3 := asyncMerge[int](ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
}

func asyncMerge[T any](chans ...chan T) chan T {
	return nil
}
