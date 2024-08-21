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

	var shkafCount int
	var socksCount int
	var chetnoe = true

	fmt.Fscan(in, &shkafCount)
	for i := 0; i < shkafCount; i++ {
		fmt.Fscan(in, &socksCount)
		if socksCount%2 == 1 {
			chetnoe = !chetnoe
		}
	}

	if chetnoe {
		out.WriteString("YES")
	} else {
		out.WriteString("NO")
	}

}
