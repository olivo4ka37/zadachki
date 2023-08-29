package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var n, start, end int
	var str, nakleika string

	fmt.Fscan(in, &str)
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &start, &end, &nakleika)
		str = nakleivanie(str, nakleika, start, end)
	}

	fmt.Fprint(out, str)
	defer out.Flush()
}

func nakleivanie(str, nakleika string, start, end int) string {
	startStr := str[:start-1]
	endStr := str[end:]

	return startStr + nakleika + endStr
}
