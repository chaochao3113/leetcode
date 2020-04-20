package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "We are happy"
	fmt.Println(replaceSpace(s))
}

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}