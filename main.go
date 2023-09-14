package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	//out := bufio.NewWriter(os.Stdout)

	var n, m int
	//var xSymbol byte

	fmt.Fscan(in, &n, &m)
	var arrASCII [][]byte
	for z := 0; z <= n; z++ {
		row, _ := in.ReadBytes('\n')
		//fmt.Println(string(row))
		arrASCII = append(arrASCII, row)
	}

	for i := 0; i <= n; i++ {
		fmt.Print(string(arrASCII[i]))
	}
}
