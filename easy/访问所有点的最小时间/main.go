package main

import (
	"fmt"
)

func main() {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	fmt.Println(minTimeToVisitAllPoints(points))
}

func minTimeToVisitAllPoints(points [][]int) int {
	sum := 0
	for i := 0; i < len(points)-1; i++ {
		sum += max(points[i], points[i + 1])
	}
	return sum
}

func max(t1 []int, t2 []int) int {
	if abs(t1[0] - t2[0]) > abs(t1[1] - t2[1]){
		return abs(t1[0] - t2[0])
	}

	return abs(t1[1] - t2[1])
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}