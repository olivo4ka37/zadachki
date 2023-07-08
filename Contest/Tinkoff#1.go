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
	var name, b string

	fmt.Fscanln(in, &n)
	fmt.Fscanln(in, &name)
	fmt.Fscanln(in, &b)
	fmt.Println(n, name, b)
}
