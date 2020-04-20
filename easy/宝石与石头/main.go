package main

import "fmt"

func main() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))
}

func numJewelsInStones(J string, S string) int {
	n := 0
	var count [256]int
	for _, v := range S {
		count[v]++
	}
	for _, v := range J {
		n += count[v]
	}
	return n
}