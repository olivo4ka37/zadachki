package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	fmt.Println("Program started")
	m := make(map[int]int, 10)
	wg := sync.WaitGroup{}
	mut := sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeToMap(m, rand.Intn(100), &wg, mut)
	}

	wg.Wait()

	fmt.Println("Program done")
}

func writeToMap(m map[int]int, n int, wg *sync.WaitGroup, mut sync.Mutex) {

	m[n] = n
	wg.Done()
	return
}
