package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(node *TreeNode) int {
	if node == nil {
		return -1
	}

	if node.Left == nil && node.Right == nil {
		return -1
	}

	left := node.Left.Val
	right := node.Right.Val

	if left == node.Val {
		left = findSecondMinimumValue(node.Left)
	}

	if right == node.Val {
		right = findSecondMinimumValue(node.Right)
	}

	if left != -1 && right != -1 {
		return min(left, right)
	}

	if left != -1 {
		return left
	}

	return right
}


func main(){
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}


	fmt.Println(findSecondMinimumValue(root))
}