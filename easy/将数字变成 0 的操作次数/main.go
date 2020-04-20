package main

import "fmt"

func main() {
	num := 123
	fmt.Println(numberOfSteps(num))
}

func numberOfSteps (num int) int {
	n := 0
	for ; num != 0;  {
		if num % 2 == 0 {
			n ++
			num /= 2
		} else {
			n ++
			num -= 1
		}
	}
	return n
}