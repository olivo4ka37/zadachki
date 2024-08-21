package main

import (
	"fmt"
	"time"
)

func main() {
	bCh := make(chan int, 2)

	go readCh(bCh)

	for i := 1; i <= 4; i++ {
		fmt.Println("Write to channel: ", i)
		bCh <- i
	}
}

func readCh(ch chan int) {
	time.Sleep(time.Second * 2)
	for i := 1; i <= 2; i++ {
		fmt.Println("Read from channel: ", <-ch)
	}
}
