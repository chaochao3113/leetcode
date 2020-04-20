package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{12,345,2,6,7896}
	fmt.Println(findNumbers(nums))
}

func findNumbers(nums []int) int {
	n := 0
	for _, tmp := range nums {
		s := strconv.Itoa(tmp)
		if len(s) % 2 == 0 {
			n++
		}
	}
	return n
}