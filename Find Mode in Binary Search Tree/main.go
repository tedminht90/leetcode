package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func findMode(root *TreeNode) []int {
    // Tạo map để đếm số lần xuất hiện của các node
    count := make(map[int]int)

    // Hàm đệ quy để duyệt cây và đếm số lần xuất hiện của các node
    dfs(root, count)

    // Tìm ra số lần xuất hiện nhiều nhất
    maxCount := 0
    for _, v := range count {
        if v > maxCount {
            maxCount = v
        }
    }

    // Tìm các giá trị có tần suất xuất hiện nhiều nhất
    var result []int
    for k, v := range count {
        if v == maxCount {
            result = append(result, k)
        }
    }
    return result
}

func dfs(root *TreeNode, count map[int]int) {
    if root == nil {
        return
    }
    // Tăng tần suất xuất hiện của node hiện tại
    count[root.Val]++

    // Đệ quy sang 2 node con
    dfs(root.Left, count)
    dfs(root.Right, count)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 2}
	root.Right.Right.Right = &TreeNode{Val: 3}
	root.Right.Right.Left = &TreeNode{Val: 3}
	root.Right.Right.Left.Left = &TreeNode{Val: 3}
	fmt.Println(findMode(root))
}