package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	aPtr := unsafe.Pointer(&a[0])
	a2 := next(&aPtr, a)
	fmt.Println(&a[0])
	fmt.Println(&a[1])
	fmt.Println(*(*int)(a2))
	fmt.Println((*int)(a2))
}

func next(ptr *unsafe.Pointer, a []int) *int {
	r := (*int)(unsafe.Pointer(uintptr(*ptr) + unsafe.Sizeof(a[0])))

	return r
}
