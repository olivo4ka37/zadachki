package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string, 4)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go reader(ch, &wg)
	wg.Add(1)
	go sender(ch, &wg)

	wg.Wait()
	fmt.Println("End of program")
}

func sender(ch chan string, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 4)
	ch <- "Message"
	wg.Done()
}

func reader(ch chan string, wg *sync.WaitGroup) {
	fmt.Println(<-ch)
	wg.Done()
}
