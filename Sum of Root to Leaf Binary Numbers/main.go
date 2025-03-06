package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}
	val = val*2 + node.Val
	if node.Left == nil && node.Right == nil {
		return val
	}
	return dfs(node.Left, val) + dfs(node.Right, val)
}

func main(){
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{0, nil, nil}
	root.Right = &TreeNode{1, nil, nil}
	root.Left.Left = &TreeNode{0, nil, nil}
	root.Left.Right = &TreeNode{1, nil, nil}
	root.Right.Left = &TreeNode{0, nil, nil}
	root.Right.Right = &TreeNode{1, nil, nil}
	println(sumRootToLeaf(root))
}