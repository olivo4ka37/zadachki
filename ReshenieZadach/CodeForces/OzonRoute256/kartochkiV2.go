package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var x, n, m int
	result := true

	//n - количество друзей которым нужно раздать карточки, m - количество карточек в целом
	fmt.Fscan(in, &n, &m)

	arr1 := make([][]int, n) // Инициализация двухмерного массива для хранения количества карточек (x) у друга и его индекс номера (i)

	for i := 0; i < n; i++ {
		arr1[i] = make([]int, 2)
		fmt.Fscan(in, &x)
		arr1[i][0] = x
		arr1[i][1] = i
	}

	// Сортируем двухмерный массив по убыванию элементов первой строки
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i][0] > arr1[j][0]
	})

	// проверяем возможность раздать все карточки друзьям, начиная раздачу с самой большой карты, если раздача возможна
	// выдаем этому другу карточку, если нет, заканчиваем цикл пишем -1, значение result ставим на false
	for j := 0; j < n; j++ {
		if m <= arr1[j][0] {
			result = false
			fmt.Fprintln(out, -1)
			defer out.Flush()
			break
		} else {
			arr1[j][0] = m
			m--
		}
	}

	// Если result == true значит все карточки можно выдать, сортируем массив по возрастанию индексов друзей и выводим значение карточки по
	// индексу друга.
	if result == true {
		sort.Slice(arr1, func(i, j int) bool {
			return arr1[i][1] < arr1[j][1]
		})
		for i := 0; i < len(arr1); i++ {
			fmt.Fprint(out, arr1[i][0], " ")
			defer out.Flush()
		}
	}

}
