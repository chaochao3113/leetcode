package main

import "fmt"

func main() {
	guess := []int{2,2,3}
	answer := []int{3,2,1}
	fmt.Println(game(guess, answer))
}

func game(guess []int, answer []int) int {
	n := 0
	for i := 0; i < 3; i++ {
		if guess[i] == answer[i] {
			n++
		}
	}
	return n
}