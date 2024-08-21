package main

import "fmt"

// Описать программу

// Скомпилируется-ли?
// -Да

// Что будет выведено?
// -file error 1
// -file error 2
// -Complete

// Что важно учитывать?
// -Интерфейс хранящий в себе nil значение уже не nil интерйфейс.

// Как сделать так, чтобы foo при значении false возвращало пустую ошибку?
// -Возвращать nil а не переменную c

func main() {
	err1 := foo(true)
	if err1 != nil { // err1 != nil
		fmt.Println("file error 1")
	}

	err2 := foo(false)
	if err2 != nil { // err2 != nil
		fmt.Println("file error 2")
	}

	fmt.Println("Complete")
}

type CustomError struct {
	Text string
}

func (s *CustomError) Error() string {
	return s.Text
}

func foo(fileError bool) error {
	var c *CustomError // У интерфейса есть указа

	if fileError {
		return &CustomError{Text: "Some Error"}
	}

	return c
}

/*
type MyError struct {
	t string
}

func (m *MyError) Error() string {
	return "My Error"
}

func boo() error {
	return &MyError{t: "Niggers!"}
}
*/
