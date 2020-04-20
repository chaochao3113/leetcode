package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//var l1 ListNode
	//var l2 ListNode
	//var l3 ListNode
	//var l4 ListNode
	//var l5 ListNode
	//var l6 ListNode
	//l1.Next = &l2; l2.Next = &l3; l3.Next = nil; l4.Next = &l5; l5.Next = &l6; l6.Next = nil
	//l1.Val = 1; l2.Val = 2; l3.Val = 4; l4.Val = 1; l5.Val = 3; l6.Val = 4
	//node := mergeTwoLists(&l1, &l4)
	node := mergeTwoLists(nil, nil)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var t ListNode
	head := &t

	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	for {
		if l1.Val >= l2.Val {
			head.Val = l2.Val
			l2 = l2.Next
			if l2 == nil {
				head.Next = l1
				break
			}
		} else if l1.Val < l2.Val {
			head.Val = l1.Val
			l1 = l1.Next
			if l1 == nil {
				head.Next = l2
				break
			}
		}

		head.Next = new(ListNode)
		head = head.Next

	}

	return &t
}
