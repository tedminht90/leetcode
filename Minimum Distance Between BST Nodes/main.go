package main

import "fmt"


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func minDiffInBST(root *TreeNode) int {
	var res int = 1<<63 - 1 // Khởi tạo res là MAX_INT
	var prev *TreeNode = nil // Con trỏ đến nút trước đó
	inOrder(root, &res, &prev)
	return res
}

// Cần thay đổi res và prev nên truyền vào con trỏ
// prev là con trỏ kép để thay đổi giá trị của nút trước đó,
// nếu đơn thì thay đổi prev sẽ không được bảo toàn
// Con trỏ kép cho phép thay đổi được con trỏ prev gốc
func inOrder(root *TreeNode, res *int, prev **TreeNode){
	if root == nil {
		return
	}
	// Duyệt trái
	inOrder(root.Left, res, prev)

	// Xử lý nút hiện tại
	if *prev != nil {
        diff := root.Val - (*prev).Val
        if diff < *res {
            *res = diff
        }
    }
	*prev = root
	inOrder(root.Right, res, prev)
}

func main(){
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	fmt.Println(minDiffInBST(root))
}