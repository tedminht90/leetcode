package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	// Gọi hàm đệ quy để duyệt cây
	preorder(root, &result)
	return result
}

func preorder(root *TreeNode, result *[]int) {
	// Điều kiện dừng nếu node là nil
	if root == nil {
		return
	}
	// Thêm giá trị của node vào mảng kết quả
	*result = append(*result, root.Val)

	// Duyệt đệ quy ở nhánh trái
	preorder(root.Left, result)
	// Duyệt đệ quy ở nhánh phải
	preorder(root.Right, result)
}


func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	result := preorderTraversal(root)

	fmt.Println(result)
}