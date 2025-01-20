package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	postorder(root, &result)
	return result
}

func postorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, result)
	postorder(root.Right, result)
	*result = append(*result, root.Val)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	result := postorderTraversal(root)
	fmt.Println(result)
}