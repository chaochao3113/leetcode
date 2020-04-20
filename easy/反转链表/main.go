package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1, l2, l3, l4, l5 ListNode
	l1.Next = &l2; l2.Next = &l3; l3.Next = &l4; l4.Next = &l5
	l1.Val = 1; l2.Val = 2; l3.Val = 3; l4.Val = 4; l5.Val = 5
	node := reverseList(&l1)
	for node != nil  {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	var tem []int
	t := head
	for t != nil {
		tem = append(tem, t.Val)
		t = t.Next
	}
	h := head
	for i := len(tem)-1; i >= 0; i-- {
		h.Val = tem[i]
		h = h.Next
	}

	return head
}
