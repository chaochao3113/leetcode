package main

import "fmt"

func main() {
	n := 1
	fmt.Println(printNumbers(n))
}

func printNumbers(n int) []int {
	max := 0
	for i := 0; i < n; i++ {
		max = max * 10 + 9
	}
	num := make([]int, max)
	for i := 1; i <= max; i++ {
		num[i - 1] = i
	}
	return num
}