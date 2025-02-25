package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isUnivalTree(node *TreeNode) bool{
	return isUnivalTreeHelper(node, node.Val)
}

func isUnivalTreeHelper(node *TreeNode, val int) bool{
	if node == nil{
		return true
	}
	if node.Val != val{
		return false
	}
	return isUnivalTreeHelper(node.Left, val) && isUnivalTreeHelper(node.Right, val)
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 1}


	fmt.Println(isUnivalTree(root))
}