package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var kg int

	fmt.Fscan(in, &kg)

	if kg%2 == 0 && kg != 2 {
		fmt.Fprint(out, "YES")
	} else {
		fmt.Fprint(out, "NO")
	}

	defer out.Flush()
}
