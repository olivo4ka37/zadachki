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
	var str string

	fmt.Fscan(in, &n)

	//chetBilets := make([]string, 0, 200000)
	//nechetBilets := make([]string, 0, 200000)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &str)
	}
}
