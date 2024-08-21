package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	x, y int
}

type KDNode struct {
	point   Point
	left    *KDNode
	right   *KDNode
	isXNode bool
}

func buildKDTree(points []Point, depth int) *KDNode {
	if len(points) == 0 {
		return nil
	}

	axis := depth % 2
	sort.Slice(points, func(i, j int) bool {
		if axis == 0 {
			return points[i].x < points[j].x
		}
		return points[i].y < points[j].y
	})

	mid := len(points) / 2
	node := &KDNode{
		point:   points[mid],
		isXNode: axis == 0,
	}

	node.left = buildKDTree(points[:mid], depth+1)
	node.right = buildKDTree(points[mid+1:], depth+1)

	return node
}

func squaredDistance(p1, p2 Point) int {
	return (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)
}

func findClosest(root *KDNode, target Point, depth int, best *Point, bestDist *int) {
	if root == nil {
		return
	}

	dist := squaredDistance(root.point, target)
	if dist < *bestDist || (*bestDist == dist && (root.point.x < best.x || (root.point.x == best.x && root.point.y < best.y))) {
		*best = root.point
		*bestDist = dist
	}

	axis := depth % 2
	var nextNode, otherNode *KDNode
	if (axis == 0 && target.x < root.point.x) || (axis == 1 && target.y < root.point.y) {
		nextNode = root.left
		otherNode = root.right
	} else {
		nextNode = root.right
		otherNode = root.left
	}

	findClosest(nextNode, target, depth+1, best, bestDist)

	if axis == 0 && (target.x-root.point.x)*(target.x-root.point.x) < *bestDist ||
		axis == 1 && (target.y-root.point.y)*(target.y-root.point.y) < *bestDist {
		findClosest(otherNode, target, depth+1, best, bestDist)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	n := 0
	fmt.Sscanf(scanner.Text(), "%d", &n)

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &points[i].x, &points[i].y)
	}

	root := buildKDTree(points, 0)

	scanner.Scan()
	q := 0
	fmt.Sscanf(scanner.Text(), "%d", &q)

	queries := make([]Point, q)
	for i := 0; i < q; i++ {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &queries[i].x, &queries[i].y)
	}

	for _, query := range queries {
		best := points[0]
		bestDist := squaredDistance(query, best)
		findClosest(root, query, 0, &best, &bestDist)
		fmt.Fprintf(writer, "%d %d\n", best.x, best.y)
	}

	writer.Flush()
}
