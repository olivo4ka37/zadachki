package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	scan := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	scan.Scan()
	str := strings.Split(scan.Text(), " ")

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(str[i])
	}

	l := make([]int, n)
	r := make([]int, n)

	// Найти ближайший меньший элемент слева
	arr := []int{}
	for i := 0; i < n; i++ {
		for len(arr) > 0 && a[arr[len(arr)-1]] >= a[i] {
			arr = arr[:len(arr)-1]
		}
		if len(arr) == 0 {
			l[i] = 0
		} else {
			l[i] = a[arr[len(arr)-1]]
		}
		arr = append(arr, i)
	}

	// Найти ближайший меньший элемент справа
	arr = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(arr) > 0 && a[arr[len(arr)-1]] >= a[i] {
			arr = arr[:len(arr)-1]
		}
		if len(arr) == 0 {
			r[i] = 0
		} else {
			r[i] = a[arr[len(arr)-1]]
		}
		arr = append(arr, i)
	}

	for i := 0; i < n; i++ {
		answer := fmt.Sprintf("%d %d\n", l[i], r[i])
		out.WriteString(answer)
	}
}
