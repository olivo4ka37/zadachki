package main

import (
	"fmt"
	"sync"
)

func merge(chs ...<-chan int) <-chan int {
	resultChannel := make(chan int, 1)

	wg := &sync.WaitGroup{}
	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for c := range ch {
				resultChannel <- c
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	return resultChannel
}

// worker - горутина, которая записывает данные в канал
func worker(wg *sync.WaitGroup, ch chan<- int, n int, id int) {
	defer wg.Done() // Сигнализируем о завершении работы горутины
	defer close(ch)

	for i := 0; i < n; i++ { // Отправляем 20 значений
		select {
		case ch <- i: // Отправка данных в канал
			fmt.Printf("Worker %d: Sent %d\n", id, i)
		default:
			fmt.Printf("Worker %d: Channel is closed, exiting\n", id)
			return // Выход из горутины, если канал закрыт
		}
	}
}

// Создается два буферизированных канала с размерами буфера 20 и 30 соответственно.
// Потом вызывается функция которая будет объединять эти два канала в один.
// В конце программа выводит объединенный канал и пишет о своем завершении.
func main() {
	ch1 := make(chan int, 20)
	ch2 := make(chan int, 30)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go worker(&wg, ch1, 20, 1)
	go worker(&wg, ch2, 30, 2)

	wg.Wait()
	resultChannel := merge(ch1, ch2)
	for value := range resultChannel {
		fmt.Println("Received value:", value)
	}

	fmt.Println("Main goroutine finished")
}
