// Điều kiện về chiều cao:

// Tại MỌI NODE trong cây:

// Chênh lệch chiều cao giữa cây con trái và cây con phải ≤ 1
// Tức là: |heightLeft - heightRight| ≤ 1

// Điều kiện về cây con:

// Cả hai cây con (trái và phải) đều phải là cây cân bằng

package main

import "fmt"


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// Hàm height sẽ trả về -1 nếu câu không cân bằng
	return height(root) != -1
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// Kiểm tra chiều cao của cây con trái
	leftHeight := height(node.Left)
	if leftHeight == -1 {
		return -1
	}

	// Kiểm tra chiều cao của cây con phải
	rightHeight := height(node.Right)
	if rightHeight == -1 {
		return -1
	}

	// Kiểm tra điều kiện cây cân bằng
	if abs(leftHeight - rightHeight) > 1 {
		return -1
	}

	// Trả về chiều cao của cây
	return max(leftHeight, rightHeight) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Left.Right = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 8}
	

	fmt.Println(isBalanced(root))
}


// Quá trình kiểm tra theo từng bước:

// Gọi isBalanced(root):

// isBalanced(3):
//     return height(3) != -1

// Tính height(3):

// height(3):
//     // Kiểm tra cây con trái (node 9)
//     leftHeight = height(9):
//         leftHeight = height(nil) = 0
//         rightHeight = height(nil) = 0
//         return max(0,0) + 1 = 1    // Chiều cao node 9
    
//     // Kiểm tra cây con phải (node 20)
//     rightHeight = height(20):
//         leftHeight = height(15):
//             leftHeight = height(nil) = 0
//             rightHeight = height(nil) = 0
//             return 1   // Chiều cao node 15
        
//         rightHeight = height(7):
//             leftHeight = height(nil) = 0
//             rightHeight = height(nil) = 0
//             return 1   // Chiều cao node 7
        
//         return max(1,1) + 1 = 2   // Chiều cao node 20
    
//     // Kiểm tra cân bằng tại node 3
//     |leftHeight - rightHeight| = |1 - 2| = 1 <= 1 (cân bằng)
//     return max(1,2) + 1 = 3   // Chiều cao cây