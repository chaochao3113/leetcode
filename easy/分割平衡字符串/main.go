package main

import "fmt"

func main() {
	s := "RLRRLLRLRL"
	fmt.Println(balancedStringSplit(s))
}

func balancedStringSplit(s string) int {
	n := 0
	v := 0
	for _, t := range s{
		if t == 'L' {
			v++
		} else {
			v--
		}

		if v == 0 {
			n++
		}
	}

	return n
}