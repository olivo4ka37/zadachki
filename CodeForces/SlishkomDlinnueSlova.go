package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var n int
	var slovo string

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &slovo)

		if len(slovo) > 10 {
			slovo = string(slovo[0]) + strconv.Itoa(len(slovo)-2) + string(slovo[len(slovo)-1])
			fmt.Fprintln(out, slovo)
		} else {
			fmt.Fprintln(out, slovo)
		}

		defer out.Flush()
	}

}
