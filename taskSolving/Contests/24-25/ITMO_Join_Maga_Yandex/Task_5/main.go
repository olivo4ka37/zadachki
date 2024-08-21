package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func distanceSquared(p1, p2 Point) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return dx*dx + dy*dy
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	// Read the number of points
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	points := make([]Point, n)

	// Read the points
	for i := 0; i < n; i++ {
		scanner.Scan()
		coords := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		points[i] = Point{x, y}
	}

	// Read the number of queries
	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())

	// Read the queries
	queries := make([]Point, q)
	for i := 0; i < q; i++ {
		scanner.Scan()
		coords := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		queries[i] = Point{x, y}
	}

	// Process each query
	for _, query := range queries {
		minDist := math.MaxInt
		var bestPoint Point

		for _, point := range points {
			dist := distanceSquared(point, query)
			if dist < minDist || (dist == minDist && (point.x < bestPoint.x || (point.x == bestPoint.x && point.y < bestPoint.y))) {
				minDist = dist
				bestPoint = point
			}
		}

		fmt.Printf("%d %d\n", bestPoint.x, bestPoint.y)
	}
}
