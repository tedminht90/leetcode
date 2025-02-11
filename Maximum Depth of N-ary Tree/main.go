package main

import "fmt"


type Node struct {
	Val int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}
	depth := 0
	for _, child := range root.Children {
		depth = max(depth, maxDepth(child))
	}
	return depth + 1
}


func main() {
	root := &Node{Val: 1}
	root.Children = append(root.Children, &Node{Val: 3})
	root.Children = append(root.Children, &Node{Val: 2})
	root.Children = append(root.Children, &Node{Val: 4})
	root.Children[0].Children = append(root.Children[0].Children, &Node{Val: 5})
	root.Children[0].Children = append(root.Children[0].Children, &Node{Val: 6})

	fmt.Println(maxDepth(root))
	
}