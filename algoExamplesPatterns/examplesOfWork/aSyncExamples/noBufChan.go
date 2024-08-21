package main

func main() {
	ch1 := make(chan int)
	ch2 := ch1
	ch1 = nil
	close(ch2)
}
