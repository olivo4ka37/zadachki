package main

import (
	"fmt"
	"unsafe"
)

func main() {
	c := 0          // Просто переменная
	var e Pntr = &c // Созданный тип Pntr типа *int, то есть указателя на int
	var d PPn = &e  // Созданный тип PPn типа Pntr, то есть указатель на Pntr
	var f PPn3 = &d
	fmt.Println(&c, &*e, &**d) // &c Сначала выводим адрес c. Потом &*e через разименовыватель, переходим к ячейке c,
	// через & получаем адрес текущей ячейки, то есть ячейки c. Потом &**d,
	// через разименовыватель переходим к значению ячейки e, это указатель,
	// мы еще раз его разименовываем, переходим к ячейке c, пишем её адрес.
	// Вот почему пишет три раза один и тот же адрес.
	fmt.Println(&e, &*d) //&e пишет адрес ячейки e, &*d тоже пишет адрес ячейки e, так как сначала разименовали указатель,
	// а потом только получили адрес ячейки.
	fmt.Println(e, &e, *e) // 1. Значение переменной e, 2. Адрес переменной e, 3. Разименование указателя, который хранит сама переменная e.
	fmt.Println(&***f)     // Можно разименовывать 3 и более раз если есть указатели.
}

type _string struct {
	elements unsafe.Pointer // underlying bytes
	len      int            // number of bytes
}

// Ptr is a named pointer type whose base type is int.
type Pntr *int

// PP is a named pointer type whose base type is Ptr.
type PPn *Pntr

type PPn3 *PPn
