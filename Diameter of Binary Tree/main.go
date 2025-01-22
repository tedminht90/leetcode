package main


type TreeNode struct {
	Val   int
	Left *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	depth(root, &res)
	return res
}

func depth(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	left := depth(node.Left, res)
	right := depth(node.Right, res)
	*res = max(*res, left + right)
	return max(left, right) + 1
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	println(diameterOfBinaryTree(root))
}