package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := []int{}
	leaf2 := []int{}
	dfs(root1, &leaf1)
	dfs(root2, &leaf2)
	return equal(leaf1, leaf2)
}

func dfs(node *TreeNode, leaf *[]int){
	if node == nil{
		return
	}
	if node.Left == nil && node.Right == nil{
		*leaf = append(*leaf, node.Val)
	}
	dfs(node.Left, leaf)
	dfs(node.Right, leaf)
}

func equal(leaf1 []int, leaf2 []int) bool{
	if len(leaf1) != len(leaf2){
		return false
	}
	for i := 0; i < len(leaf1); i++{
		if leaf1[i] != leaf2[i]{
			return false
		}
	}
	return true
}

func main(){
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 5}
	root1.Right = &TreeNode{Val: 1}
	root1.Left.Left = &TreeNode{Val: 6}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Left.Right.Left = &TreeNode{Val: 7}
	root1.Left.Right.Right = &TreeNode{Val: 4}
	root1.Right.Left = &TreeNode{Val: 9}
	root1.Right.Right = &TreeNode{Val: 8}

	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 5}
	root2.Right = &TreeNode{Val: 1}
	root2.Left.Left = &TreeNode{Val: 6}
	root2.Left.Right = &TreeNode{Val: 7}
	root2.Right.Left = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 2}
	root2.Right.Right.Left = &TreeNode{Val: 9}
	root2.Right.Right.Right = &TreeNode{Val: 8}

	fmt.Println(leafSimilar(root1, root2))
}