package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const buf = 2

func main() {
	in := bufio.NewReader(os.Stdin)
	wg := sync.WaitGroup{}
	ch := make(chan int)
	chB := make(chan int, buf)

	// Сколькими способами можно заставить программу ожидать выполнения горутин?
	// Что за способы это будут?

	var x int
	fmt.Println("Введите число теста:\n1 - WaitGroup\n2 - Channel\n3 - Buffered Channel")
	fmt.Fscan(in, &x)

	switch x {
	case 1:
		// В данном примере мы ожидаем завершения горутины так используем WaitGroup
		wg.Add(1)
		go gWaitGroup(&wg)
	case 2:
		// В данном примере для ожидания завершения горутины, мы вызываем чтение данных из канала,
		// пока канал не закрыт, горутина в которой ожидается чтение лочится, пока не прочтёт значение.
		// Если канал закрыли, то будут читаться нулевые значения канала, горутина не будет лочиться.
		go gChanNoBuf(x, ch)
		fmt.Println("Value from channel is: ", <-ch)
		fmt.Println("Value from channel is: ", <-ch) // Если канал закрыт, читать из него можно, просто будем получать нулевое значение.2
	case 3:
		// Записываем значения в канал с буфером 2, когда мы записываем третье значение эта main горутина блокируется
		// и ждёт чтения из других горутин, так как мы вызвали go gBuffedChan(chB) перед тем как записывать значения в канал,
		// main горутина дожидается чтения из gBuffedChan, важно! В самой функции считывается 3 значения, соответсвенно канал снова пуст
		// и мы продолжаем работу main горутины в которой записывается ещё два значения в канал.
		go gBuffedChan(chB)
		for i := 0; i < 5; i++ {
			x = rand.Intn(99) + 1
			fmt.Println("Write to buffed channel: ", x)
			chB <- x
		}

	}

	fmt.Println("Reached the end of main func")
	wg.Wait()
}

// gWaitGroup Использует wg для того, чтобы функция main не завершала свою работу, до завершения этой горутины
func gWaitGroup(wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Hello from goroutine that uses waitGroup!")
	wg.Done()
}

// gChanNoBuf Записывает значение x в канал, а потом закрывает канал.
func gChanNoBuf(x int, ch chan int) {
	// close(ch) - Можно заранее закрыть канал, и убедиться что main func завершится раньше окончания горутины.
	time.Sleep(time.Millisecond * 2500)
	fmt.Println("Hello from goroutine that uses channel")
	ch <- x
	close(ch) // Если канал не закрыть, то программа в которой мы будем считывать значения из канала
	// будет попадать в deadclock, так как программа будет ждать записи в канал, которой
	// никогда не произойдёт.
}

// gBuffedChan
func gBuffedChan(ch chan int) {
	fmt.Println("goroutine is sleep for 2 sec")
	time.Sleep(time.Second * 2)
	for i := 1; i <= 3; i++ {
		fmt.Println("Read from buffered channel: ", <-ch)
	}
}
