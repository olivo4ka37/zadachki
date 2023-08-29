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

	var t, rows, cols, x, click, stolbBlyat int

	fmt.Fscan(in, &t)

	for c := 0; c < t; c++ {
		fmt.Fscan(in, &rows, &cols)
		TwoDArray := make([][]int, 0)
		for i := 0; i < rows; i++ {
			var row []int
			for j := 0; j < cols; j++ {
				fmt.Fscan(in, &x)
				row = append(row, x)
			}
			TwoDArray = append(TwoDArray, row)
		}

		fmt.Fscan(in, &click)

		for zxc := 0; zxc < click; zxc++ {
			fmt.Fscan(in, &stolbBlyat)
			sort.SliceStable(TwoDArray, func(i, j int) bool {
				return TwoDArray[i][stolbBlyat-1] < TwoDArray[j][stolbBlyat-1]
			})
		}

		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				fmt.Fprint(out, TwoDArray[i][j], " ")
			}
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out)
	}
	defer out.Flush()

}
