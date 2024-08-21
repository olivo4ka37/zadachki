package main

import (
	"bufio"
	"fmt"
	"os"
)

// Автор истории про Петю просто гений!
// Концовка убила :D.
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	nSoldat, mPar := 0, 0
	fmt.Fscan(in, &nSoldat, &mPar)

	parArray, vhojdenii := make([][]int, nSoldat+1), make([]int, nSoldat+1)

	for i := 0; i < mPar; i++ {
		a, b := 0, 0
		fmt.Fscan(in, &a, &b)
		parArray[a] = append(parArray[a], b)
		vhojdenii[b]++
		//fmt.Println(a,b,"i is:", i)
	}

	higherElements := make([]int, 0)
	for i := 1; i <= nSoldat; i++ {
		if vhojdenii[i] == 0 {
			higherElements = append(higherElements, i)
		}
	}
	//	fmt.Println(higherElements)
	//	fmt.Println(vhojdenii)

	result := make([]int, 0)
	//fmt.Println(parArray)

	for len(higherElements) > 0 {
		elementBezVhoda := higherElements[0]
		higherElements = higherElements[1:]
		result = append(result, elementBezVhoda)
		//fmt.Println(elementBezVhoda)
		//fmt.Println(parArray[elementBezVhoda])

		fmt.Println(parArray[elementBezVhoda])
		for _, sosed := range parArray[elementBezVhoda] {
			vhojdenii[sosed]--
			if vhojdenii[sosed] == 0 {
				higherElements = append(higherElements, sosed)
			}
		}
	}

	if len(result) == nSoldat {
		out.WriteString("Yes")
		out.WriteString("\n")
		for _, v := range result {
			x := fmt.Sprintf("%d ", v)
			out.WriteString(x)
		}
	} else {
		out.WriteString("No")
	}
}
