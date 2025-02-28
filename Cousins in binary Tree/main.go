package main


type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

// Hàm chính kiểm tra hai nút có phải là anh em họ không
func isCousins(root *TreeNode, x int, y int) bool {
    // Tìm thông tin về nút x và y
    xParent, xDepth := findNodeInfo(root, nil, 0, x)
    yParent, yDepth := findNodeInfo(root, nil, 0, y)
    
    // Kiểm tra điều kiện anh em họ: khác cha, cùng độ sâu
    return xParent != yParent && xDepth == yDepth
}

// Hàm phụ tìm thông tin về một nút trong cây
func findNodeInfo(node, parent *TreeNode, depth int, target int) (*TreeNode, int) {
    // Trường hợp cơ sở: nút rỗng
    if node == nil {
        return nil, -1
    }
    
    // Nếu tìm thấy nút cần tìm
    if node.Val == target {
        return parent, depth
    }
    
    // Tìm kiếm trong cây con trái
    leftParent, leftDepth := findNodeInfo(node.Left, node, depth+1, target)
    if leftParent != nil {
        return leftParent, leftDepth
    }
    
    // Tìm kiếm trong cây con phải
    rightParent, rightDepth := findNodeInfo(node.Right, node, depth+1, target)
    return rightParent, rightDepth
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	x := 5
	y := 4
	println(isCousins(root, x, y))

}