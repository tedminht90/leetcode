package main

import "fmt"


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func mergeTrees(node1, node2 *TreeNode) *TreeNode{
	if node1 == nil{
		return node2
	}
	if node2 == nil{
		return node1
	}
	node1.Val += node2.Val
	node1.Left = mergeTrees(node1.Left, node2.Left)
	node1.Right = mergeTrees(node1.Right, node2.Right)
	return node1
}

func printTree(node *TreeNode){
	if node == nil{
		return
	}
	
	printTree(node.Left)
	fmt.Print(node.Val, "")
	printTree(node.Right)
}

func main(){
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 3}
	root1.Right = &TreeNode{Val: 2}
	root1.Left.Left = &TreeNode{Val: 5}

	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Right = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 7}


	printTree(mergeTrees(root1, root2))

	fmt.Println()

}