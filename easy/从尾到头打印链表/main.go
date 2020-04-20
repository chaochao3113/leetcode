package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 ListNode
	var l2 ListNode
	var l3 ListNode
	l1.Val = 1
	l1.Next = &l2
	l2.Val = 3
	l2.Next = &l3
	l3.Val =  2
	l3.Next = nil
	fmt.Println(reversePrint(&l1))
}

func reversePrint(head *ListNode) []int {
	var n []int
	for head != nil {
		n = append(n, head.Val)
		head = head.Next
	}

	for i := 0; i < len(n)/2; i++ {
		a := n[len(n)-i-1]
		n[len(n)-1-i] = n[i]
		n[i] = a
	}

	return n
}
