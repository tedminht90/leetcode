package main

import "fmt"


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

func printNode(node *TreeNode){
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	printNode(node.Left)
	printNode(node.Right)
}

func main(){
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	val := 2

	printNode(searchBST(root, val))
}