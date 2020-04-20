package main

import "fmt"

func main() {
	n := 234
	fmt.Println(subtractProductAndSum(n))
}

func subtractProductAndSum(n int) int {
	var num []int
	for ; n != 0; {
		num = append(num, n % 10)
		n /= 10
	}

	mult := 1
	sum := 0
	for i := 0; i < len(num); i++ {
		mult *= num[i]
		sum += num[i]
	}

	return mult - sum
}