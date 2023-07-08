package AlgorithmsAndDataStructure

import "fmt"

//Простой способ реализовать структуру данных временной очереди в Go - это использовать срез:
//
//    чтобы добавить в очередь, вы используете встроенную функцию append
//    чтобы удалить из очереди - срезать первый элемент

func Ochered() {
	var queue []string

	queue = append(queue, "Hello ") // Добавление в очередь
	queue = append(queue, "world!")

	for len(queue) > 0 {
		fmt.Print(queue[0]) // Первый элемент
		queue = queue[1:]   // Удаление из очереди
	}
}
