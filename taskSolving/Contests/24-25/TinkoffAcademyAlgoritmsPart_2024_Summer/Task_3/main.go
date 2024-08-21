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
	var n, k, a, m int
	fmt.Fscan(in, &n, &k, &a, &m)

	confetKupleno, monets, totalSum := 0, 0, 0

	gen := generator(0, a, m)

	for confetKupleno < n {
		coin := <-gen
		totalSum += coin
		monets++

		if totalSum >= 3*k {
			purchased := totalSum / 3
			if purchased >= k {
				totalSum -= purchased * 3
				confetKupleno += purchased
			}
		}
	}

	result := fmt.Sprintf("%d", monets)
	out.WriteString(result)
}

func absForInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func generator(seed, a, m int) <-chan int {
	ch := make(chan int)
	go func() {
		for {
			seed = lcg(seed, a, m)
			coin := (absForInt(seed%3-1)*5 + absForInt(seed%3)*2) % 8
			ch <- coin
		}
	}()
	return ch
}

func lcg(e, a, m int) int {
	return (a*e + 11) % m
}
