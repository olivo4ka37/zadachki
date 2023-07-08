package main

import (
	"bufio"
	"fmt"
	"os"
)

func InputFunc() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	var testInt int

	fmt.Fscan(in, &testCount, &testInt)
	fmt.Println(testCount &^ testInt)
}
