package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	// Tạo mảng 2 chiều để lưu [tổng, số lượng] cho mỗi level
	levels := [][]int{}

	// Gọi DFS để thu thập dữ liệu, bắt đầu từ level 0
	dfs(root, 0, &levels)

	// Tính và trả về kết quả trung bình
	return caculateAverage(levels)
}

// Hàm DFS để thu thập sum và count của mỗi level
func dfs(node *TreeNode, level int, levels *[][]int) {
	// Điều kiện dừng: nếu node rỗng
	if node == nil {
		return
	}

	// Nếu chưa có mảng cho level hiện tại
	if level >= len(*levels) {
		// Thêm mảng mới [sum=0, count=0]
		*levels = append(*levels, []int{0,0})
	}

	// Cập nhật dữ liệu cho level hiện tại:
	(*levels)[level][0] += node.Val // Cộng giá trị vào tổng
	(*levels)[level][1]++ 			// Tăng số lượng node
 
	// Đệ quy xuống các con với level tăng thêm 1
	dfs(node.Left, level + 1, levels)
	dfs(node.Right, level + 1, levels)
}


// Hàm tính trung bình cho mỗi level
func caculateAverage(levels [][]int) []float64 {
	// Tạo mảng kết quả có cùng độ dài với levels
	res := make([]float64, len(levels))

	// Duyệt qua từng level
	for i, level :=  range levels {
		// Tính trung bình: sum/count sau khi ép kiểu về float64
		res[i] = float64(level[0]) / float64(level[1])
	}
	return res
}


func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(averageOfLevels(root))	
}