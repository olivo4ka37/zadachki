package main

import (
	"fmt"
)

func greet(s string, n int) {
	fmt.Printf("Hello %s! From %v worker. \n", s, n)

}

func main() {
	// Ждёт ли программа выполнения обычной горутины?
	s := "Maria"

	for i := 0; i < 5; i++ {
		go greet(s, i)
	}

	fmt.Println("Program end")
	// Ответ - нет. В этом варианте программы каждый раз выполняется различное количество горутин,
	// количество выполненых горутин банально зависит от того, сколько горутин
	// успевают выполниться до завершения главной горутины.
}