package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "(()())(())(()(()))"
	fmt.Println(removeOuterParentheses(s))
}

func removeOuterParentheses(S string) string {
	flag := 0
	var build strings.Builder
	for _, t := range S {
		if t == '(' {
			flag++
		}
		if flag > 1 {
			build.WriteRune(t)
		}

		if t == ')' {
			flag--
		}
	}

	return build.String()
}