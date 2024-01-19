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
		var str string
		needToWrite := false
		bykvaCount := 3
		ciphraCount := 0
		fmt.Fscan(in, &str)
		for z := 0; z < len(str); z++ {
			if z == 0 && (str[z] < 65 || str[z] > 90) {
				break
			} else if z == 0 {
				continue
			} else if z == len(str)-1 {
				if str[len(str)-1] >= 48 && str[len(str)-1] <= 57 {
					ciphraCount++
				} else {
					bykvaCount++
				}
				if bykvaCount == 2 && ciphraCount == 0 {
					needToWrite = true
				}
				break
			}

			if str[z] >= 48 && str[z] <= 57 {
				ciphraCount++
			} else if str[z] >= 65 && str[z] <= 90 {
				bykvaCount++
			}

			if (str[z+1] >= 65 && str[z+1] <= 90) && (bykvaCount == 3 && ciphraCount > 0 && ciphraCount <= 2) {
				bykvaCount = 0
				ciphraCount = 0
			}

			if (ciphraCount != 0 && bykvaCount != 3) || (ciphraCount != 0 && bykvaCount == 0) || (ciphraCount == 2 && bykvaCount != 3) {
				break
			}
		}

		if needToWrite == true {
			bykvaCount = 0
			for _, x := range str {
				if bykvaCount == 3 {
					fmt.Fprint(out, " ")
					bykvaCount = 0
				}
				fmt.Fprint(out, string(x))
				if x >= 65 && x <= 90 {
					bykvaCount++
				}
			}
			fmt.Fprintln(out, "")
		} else {
			fmt.Fprint(out, "-")
			fmt.Fprintln(out, "")
		}
	}
}
