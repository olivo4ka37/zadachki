package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Автор истории про Петю просто гений!
// Концовка убила :D.
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	nSoldat, mPar := 0, 0
	fmt.Fscan(in, &nSoldat, &mPar)
	parArray, naSkokChelovNije := make([][]int, nSoldat+1), make([][]int, nSoldat+1)
	for i, _ := range naSkokChelovNije {
		naSkokChelovNije[i] = make([]int, 2)
		naSkokChelovNije[i][1] = i
	}

	for i := 0; i < mPar; i++ {
		a, b := 0, 0
		fmt.Fscan(in, &a, &b)
		parArray[a] = append(parArray[a], b)
		naSkokChelovNije[b][0]++
	}
	//fmt.Println(parArray)
	//fmt.Println(naSkokChelovNije)

	sort.Slice(naSkokChelovNije, func(i, j int) bool {
		return naSkokChelovNije[i][0] < naSkokChelovNije[j][0]
	})

	order := make([]int, 0, nSoldat+1)

	for i := 1; i < len(naSkokChelovNije); i++ {
		if naSkokChelovNije[i][0] == 0 {
			order = append(order, naSkokChelovNije[i][1])
		}
	}
	//fmt.Println(order)

	/*
		order := make([]int, nSoldat+1)
		for i,_ := range order {
			order[i] = naSkokChelovNije[i][1]
		}
	*/

	result := make([]int, 0, nSoldat+1)
	for len(order) > 0 {
		elementBezVhoda := order[0]
		result = append(result, elementBezVhoda)
		order = order[1:]

		//fmt.Println(parArray[elementBezVhoda])
		for _, sosed := range parArray[elementBezVhoda] {
			naSkokChelovNije[sosed][0]--
			if naSkokChelovNije[sosed][0] == 0 {
				order = append(order, sosed)
			}
		}
	}
	//fmt.Println(result)

	//fmt.Println(parArray[1])
	/*
		fmt.Println(naSkokChelovNije)
		fmt.Println(order)
		fmt.Println(parArray)

	*/

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
