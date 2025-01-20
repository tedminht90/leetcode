package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	dfs(root, "", &res)
	return res
}

func dfs(root *TreeNode, path string, res *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path+fmt.Sprintf("%d", root.Val))
		return
	}
	dfs(root.Left, path+fmt.Sprintf("%d->", root.Val), res)
	dfs(root.Right, path+fmt.Sprintf("%d->", root.Val), res)
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println(binaryTreePaths(root))
}