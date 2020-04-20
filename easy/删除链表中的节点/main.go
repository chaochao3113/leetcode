package main

type ListNode struct {
	Val  int
	Next *ListNode
}

var l1 ListNode
var l2 ListNode
var l3 ListNode
var l4 ListNode


func main() {
	l1.Val = 4
	l1.Next = &l2
	l2.Val = 5
	l2.Next = &l3
	l3.Val = 1
	l3.Next = &l4
	l4.Val = 9
	l4.Next = nil
	deleteNode(&l2)
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
