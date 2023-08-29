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

	var lenSlovr, nZaprsv int
	slovar := make([]string, lenSlovr)
	fmt.Fscan(in, &lenSlovr)
	for i := 0; i < lenSlovr; i++ {
		fmt.Fscan(in, &slovar[i])
	}
	fmt.Fscan(in, &nZaprsv)
	for i := 0; i < nZaprsv; i++ {
		var slovo string
		var xMatch, maxMatch int = 0, 0
		fmt.Fscan(in, &slovo)
		for j := 0; j < len(slovo); j++ {
			xMatch = maxMatch * xMatch
			maxMatch = xMatch * maxMatch
		}
	}
}
