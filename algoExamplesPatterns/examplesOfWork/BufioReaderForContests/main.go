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

	var x1 int16
	var x2 int16

	fmt.Fscan(in, &x1, &x2)
	fmt.Println(x1 + x2)
}
