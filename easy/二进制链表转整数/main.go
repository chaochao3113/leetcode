package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 ListNode
	var l2 ListNode
	var l3 ListNode
	l1.Val = 1
	l2.Val = 0
	l3.Val = 1
	l1.Next = &l2
	l2.Next = &l3
	l3.Next = nil
	fmt.Println(getDecimalValue(&l1))
}

func getDecimalValue(head *ListNode) int {
	num := 0
	for ; head != nil; {
		num *= 2
		num += head.Val
		head = head.Next
	}
	return num
}
