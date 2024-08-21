package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func findNeighbors(n int, names []string, exits []int, out *bufio.Writer) {
	// Создаем двусвязный список для имен детей
	circle := list.New()
	nameMap := make(map[int]*list.Element)
	for i := 0; i < n; i++ {
		element := circle.PushBack(names[i])
		nameMap[i+1] = element
	}

	// Итерация по детям, которые выходят из круга
	for _, exit := range exits {
		// Получаем элемент, который будет удален
		elem := nameMap[exit]

		// Получаем соседей
		var left, right string
		if elem.Prev() != nil {
			left = elem.Prev().Value.(string)
		} else {
			left = circle.Back().Value.(string)
		}
		if elem.Next() != nil {
			right = elem.Next().Value.(string)
		} else {
			right = circle.Front().Value.(string)
		}

		// Печатаем имена слева и справа
		str := fmt.Sprintf("%s %s\n", left, right)
		out.WriteString(str)

		// Удаляем элемент из списка
		circle.Remove(elem)
		delete(nameMap, exit)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	names := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &names[i])
	}

	exits := make([]int, n-3)
	for i := 0; i < n-3; i++ {
		fmt.Fscan(in, &exits[i])
	}

	// Вызов функции для нахождения соседей
	findNeighbors(n, names, exits, out)
}
