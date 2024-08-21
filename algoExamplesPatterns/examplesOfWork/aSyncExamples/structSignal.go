package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chSig := make(chan struct{})
	chVal := make(chan int)
	rands := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		rands = append(rands, rand.Intn(1000))
	}
	go readerWorker(chSig, chVal)
	go writerWorker(chSig, chVal, rands)

	time.Sleep(time.Second * 4)

	fmt.Println("Program end")
}

func readerWorker(chQuit chan struct{}, ch chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("Got value %d\n", v)
		case <-chQuit:
			fmt.Printf("Signalled on quit channel. Finishing\n")
			chQuit <- struct{}{}
			close(chQuit)
			return
		}
		time.Sleep(time.Millisecond * 555)
	}
}

func writerWorker(chQuit chan struct{}, ch chan int, arr []int) {
	for _, x := range arr {
		ch <- x
	}

	chQuit <- struct{}{}
}
