package main

import "fmt"


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool)
	return find(root, k, m)
}
func find(root *TreeNode, k int, m map[int]bool) bool {
	if root == nil {
		return false
	}
	if m[k-root.Val] {
		return true
	}
	m[root.Val] = true
	return find(root.Left, k, m) || find(root.Right, k, m)
}


func main(){
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 7}

	k := 9

	fmt.Println(findTarget(root, k))
}