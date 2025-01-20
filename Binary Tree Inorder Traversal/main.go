package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}


// Recursive solution

// Đây là hàm wrapper (hàm bọc ngoài)
// Tạo slice result để lưu các giá trị trong quá trình duyệt
// Truyền địa chỉ của result vào hàm đệ quy để có thể thay đổi giá trị của nó
func inorderTraversal(root *TreeNode) []int {
	var result []int // Tạo slice rỗng để lưu kết quả
	inorder(root, &result) // Gọi hàm đệ quy, truyền con trỏ của result
	return result // Trả về kết quả cuối cùng
}

func inorder(node *TreeNode, result *[]int){
	if node == nil{ // Điều kiện dựng đệ quy
		return
	}
	inorder(node.Left, result) // Bước 1: Duyệt cây con trái
	*result = append(*result, node.Val) // Bước 2: Thêm giá trị của node hiện tại vào result
	inorder(node.Right, result) // Bước 3: Duyệt cây con phải
}


func main(){
	// Tạo cây nhị phân [1,null,2,3]
   	/*
       1
        \
         2
        /
       3
   	*/
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Left.Right = &TreeNode{Val: 4}
	fmt.Println(inorderTraversal(root))
}

// 1. Gọi inorder(node1, &result): // node1 là gốc, giá trị 1
//    - node = 1
//    - Gọi inorder(node1.Left, &result) // null → return
//    - Thêm 1 vào result: [1]
//    - Gọi inorder(node1.Right, &result) // node2

// 2. Tại node2 (giá trị 2):
//    - node = 2
//    - Gọi inorder(node2.Left, &result) // node3

// 3. Tại node3 (giá trị 3):
//    - node = 3
//    - Gọi inorder(node3.Left, &result) // null → return
//    - Thêm 3 vào result: [1,3]
//    - Gọi inorder(node3.Right, &result) // null → return

// 4. Quay lại node2:
//    - Thêm 2 vào result: [1,3,2]
//    - Gọi inorder(node2.Right, &result) // null → return

// Kết quả cuối cùng: [1,3,2]