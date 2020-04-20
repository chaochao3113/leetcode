package main

import "fmt"

func main() {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseLeftWords(s, k))
}

func reverseLeftWords(s string, n int) string {
	return s[n:]+s[:n]
}