package main

import "fmt"


type Node struct {
	Val int
	Children []*Node
}

func preorder (node *Node) []int {
	if node == nil {
		return []int{}
	}
	result := []int{node.Val}
	for _, child := range node.Children {
		result = append(result, preorder(child)...)
	}
	return result
}

func main() {
	root := &Node{Val: 1}
	root.Children = append(root.Children, &Node{Val: 3})
	root.Children = append(root.Children, &Node{Val: 2})
	root.Children = append(root.Children, &Node{Val: 4})
	root.Children[0].Children = append(root.Children[0].Children, &Node{Val: 5})
	root.Children[0].Children = append(root.Children[0].Children, &Node{Val: 6})


	fmt.Println(preorder(root))
}