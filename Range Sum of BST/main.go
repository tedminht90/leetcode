package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rangeSumBST(node *TreeNode, low, high int) int {
	if node == nil {
		return 0
	}
	if node.Val < low {
		return rangeSumBST(node.Right, low, high)
	}
	if node.Val > high {
		return rangeSumBST(node.Left, low, high)
	}
	return node.Val + rangeSumBST(node.Left, low, high) + rangeSumBST(node.Right, low, high)
}

func main(){
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 15}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 18}
	println(rangeSumBST(root, 7, 15))
}