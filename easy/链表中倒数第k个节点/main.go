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
	var l4 ListNode
	var l5 ListNode

	l1.Val = 1
	l1.Next = &l2
	l2.Val = 2
	l2.Next  = &l3
	l3.Val = 3
	l3.Next = &l4
	l4.Val = 4
	l4.Next = &l5
	l5.Val = 5
	l5.Next = nil
	k := 2
	fmt.Println(getKthFromEnd(&l1, k).Val, getKthFromEnd(&l1, k).Next.Val)
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var tmp []ListNode
	for head != nil  {
		tmp = append(tmp, *head)
		head = head.Next
	}
	return &tmp[len(tmp)-k]
}
