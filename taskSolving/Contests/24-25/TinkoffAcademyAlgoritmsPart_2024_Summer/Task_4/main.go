package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	startPlatformPositionArray, endPlatformPositionArray := make([]int, m), make([]int, m)
	minDeistvii := 0

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &startPlatformPositionArray[i], &endPlatformPositionArray[i])
	}

	left, right := startPlatformPositionArray[0], endPlatformPositionArray[0]

	for i := 1; i < m; i++ {

		if startPlatformPositionArray[i] > right {
			minDeistvii += startPlatformPositionArray[i] - right
			right = startPlatformPositionArray[i]
			left = startPlatformPositionArray[i] - (right - left)
		} else if endPlatformPositionArray[i] < left {
			minDeistvii += left - endPlatformPositionArray[i]
			left = endPlatformPositionArray[i]
			right = endPlatformPositionArray[i] + (right - left)
		} else {
			left = int(math.Max(float64(left), float64(startPlatformPositionArray[i])))
			right = int(math.Min(float64(right), float64(endPlatformPositionArray[i])))
		}
	}

	result := fmt.Sprintf("%d", minDeistvii)
	out.WriteString(result)
}
