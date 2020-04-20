package main

import "fmt"

func main() {
	grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	fmt.Println(countNegatives(grid))
}

func countNegatives(grid [][]int) int {
	n := 0
	for _, row := range grid {
		for _, col := range row {
			if col < 0 {
				n++
			}
		}
	}
	return n
}