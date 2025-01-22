package main

type TreeNode struct {
	Val   int
	Left *TreeNode
	Right *TreeNode
}

func getMinimumDifference(node *TreeNode) int {
	var res, pre int
	res = 1 << 31 // Khởi tạo res là giá trị lớn nhất của int
	pre = -1 // Khởi tạo pre là -1
	inOrder(node, &res, &pre) // Duyệt cây theo thứ tự inorder
	return res
}

func inOrder(node *TreeNode, res, pre *int) {
	if node == nil {
		return
	}
	// Duyệt cây con trái, sau đó so sánh giá trị của node hiện tại với giá trị của node trước đó
	inOrder(node.Left, res, pre)

	// Xử lý node hiện tại
	if *pre != -1 { // Nếu node trước đó không phải là nil
		if node.Val - *pre < *res { // So sánh giá trị hiện tại với giá trị trước đó
			*res = node.Val - *pre // Cập nhật res
		}
	}
	*pre = node.Val // Cập nhật giá trị của node trước đó

	// Duyeệt cây con phải
	inOrder(node.Right, res, pre)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	println(getMinimumDifference(root))
}
