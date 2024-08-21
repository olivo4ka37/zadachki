package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// len строки не может изменяться,

	s := "123"
	fmt.Println(s, &s, len(s))
	s = "5454"
	fmt.Println(s, &s, len(s))

}

type stringStruct struct {
	str unsafe.Pointer
	len int
}
