package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)

	epsilon := 1e-15
	l, r := 1.0, 100.0

	for r-l > epsilon {
		sr := (l + r) / 2.0
		f := math.Pow(sr, float64(a)) - 1 - math.Pow(sr, float64(a-b)) - math.Pow(sr, float64(a-c))
		if f > 0 {
			r = sr
		} else {
			l = sr
		}
	}

	result := (l + r) / 2.0
	str := fmt.Sprintf("%.10f", result)
	out.WriteString(str)

}
