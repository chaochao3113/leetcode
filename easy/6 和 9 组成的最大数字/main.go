package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := 9999
	fmt.Println(maximum69Number(num))
}

func maximum69Number (num int) int {
	s := strconv.Itoa(num)

	num, _ = strconv.Atoi(strings.Replace(s, "6", "9", 1))

	return num
}