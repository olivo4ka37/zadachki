package main

import (
	"fmt"
	"unsafe"
)

// Основным правилом паддинга(выравнивание) будет, располагать поля в структуре в отсортированном порядке,
// от самых больших к самым маленьким.
func main() {
	var cr Credits
	var cr2 Credits2
	var cr3 EmptyCredits
	fmt.Println(unsafe.Sizeof(cr), unsafe.Sizeof(cr2), unsafe.Sizeof(cr3))
}

// Credits пример хорошего паддинга, занимает 24 байта
type Credits struct {
	Value    float64
	ToPay    float64
	cashBack bool
	user     bool
	id       int8
	id2      int8
	id3      int8
	id4      int8
	id5      int8
	id6      int8
}

// Credits2 пример плохого паддинга, занимает 40 байт
type Credits2 struct {
	cashBack bool    // Требует 1 байт, в строке памяти осталось 7 байт.
	Value    float64 // Требует 8 байт, в текущей строке памяти 7 байт, переходит на следующую строку и полностью её занимает => Как итог 16 байт
	user     bool    // Всё та же схема...
	ToPay    float64 //32
	id       uint8   // Из-за того что предыдущая переменная заняла всю строку, происходит выделение памяти под новую строку в 8 байт.
	id2      uint8   // Из-за того что предыдущая переменная заняла всю строку, происходит выделение памяти под новую строку в 8 байт.
}

// EmptyCredits пример пустой структуры, занимает 0 байт.
type EmptyCredits struct {
}
