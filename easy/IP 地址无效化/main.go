package main

import (
	"fmt"
	"strings"
)

func main() {
	address := "1.1.1.1"
	fmt.Println(defangIPaddr(address))
}

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}