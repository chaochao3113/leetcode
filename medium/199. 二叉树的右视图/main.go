/**
* @Author: Chao
* @Date: 2020/4/22 12:31
* @Version: 1.0
 */

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	var node TreeNode
	var node2 TreeNode
	var node3 TreeNode
	var node4 TreeNode
	var node5 TreeNode
	node.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5
	node.Left = &node2
	node.Right = &node3
	node2.Left = nil
	node2.Right = &node5
	node3.Left = nil
	node3.Right = &node4

	fmt.Println(rightSideView(&node))
}

func rightSideView(root *TreeNode) []int {
	vision := make([]int, 0)
	dfs(root, 0, &vision)
	return vision
}

func dfs(node *TreeNode, step int, vision *[]int) {
	if node == nil {
		return
	}
	if len(*vision) == step {
		*vision = append(*vision, node.Val)
	}
	dfs(node.Right, step + 1,vision)
	dfs(node.Left, step + 1,vision)
}