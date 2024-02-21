package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var d, m, y int
		dGran := 28
		fmt.Fscan(in, &d, &m, &y)
		if ((y%4 == 0 && y%100 != 0) || y%400 == 0) && m == 2 {
			dGran++
		}

		switch m {
		case 1:
			dGran += 3
		case 3:
			dGran += 3
		case 4:
			dGran += 2
		case 5:
			dGran += 3
		case 6:
			dGran += 2
		case 7:
			dGran += 3
		case 8:
			dGran += 3
		case 9:
			dGran += 2
		case 10:
			dGran += 3
		case 11:
			dGran += 2
		case 12:
			dGran += 3
		}

		if d > dGran {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}

	}

}
