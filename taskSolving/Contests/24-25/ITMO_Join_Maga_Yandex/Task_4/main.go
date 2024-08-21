package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func distanceBetweenPoints(p1, p2 Point) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return dx*dx + dy*dy
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		x, y := 0, 0
		fmt.Fscan(in, &x, &y)
		points[i] = Point{x, y}
	}

	var q int
	fmt.Fscan(in, &q)

	queries := make([]Point, q)
	for i := 0; i < q; i++ {
		x, y := 0, 0
		fmt.Fscan(in, &x, &y)
		queries[i] = Point{x, y}
	}

	for _, query := range queries {
		closestPoint := points[0]
		minDistance := distanceBetweenPoints(query, points[0])

		for _, point := range points[1:] {
			dist := distanceBetweenPoints(query, point)
			if dist < minDistance || (dist == minDistance && (point.x < closestPoint.x || (point.x == closestPoint.x && point.y < closestPoint.y))) {
				closestPoint = point
				minDistance = dist
			}
		}

		str := fmt.Sprintf("%d %d\n", closestPoint.x, closestPoint.y)
		out.WriteString(str)
	}
}
