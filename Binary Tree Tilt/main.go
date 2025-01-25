package main

import "fmt"


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	tilt := 0
	dfs(root, &tilt)
	return tilt
}

func dfs(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, tilt)
	right := dfs(root.Right, tilt)
	*tilt += abs(left - right)
	return left + right + root.Val
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}


func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	fmt.Println(findTilt(root))
}
