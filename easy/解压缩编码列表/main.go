package main

import "fmt"

func main() {
	nums := []int{1,2,3,4}
	fmt.Println(decompressRLElist(nums))
}

func decompressRLElist(nums []int) []int {
	var result []int
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < nums[i]; j++ {
			result = append(result, nums[i+1])
		}
		i++
	}
	return result
}