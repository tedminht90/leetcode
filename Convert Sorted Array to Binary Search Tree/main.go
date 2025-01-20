package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
		return nil
	}
	// Find middle element as root
	mid := len(nums) / 2

	// Create root node with middle element
	root := &TreeNode{Val: nums[mid]}

	// Recursively create left and right subtree
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root

}

func main(){
	nums := []int{-10,-3,0,5,9}
	sortedArrayToBST(nums)
}

// Input array: [-10,-3,0,5,9]

// 1. nums = [-10,-3,0,5,9]
//    len = 5, mid = 5/2 = 2
//    root = 0

// 2. Left subtree: [-10,-3]
//    len = 2, mid = 2/2 = 1
//    node = -3
//    - Left: [-10] -> node = -10
//    - Right: [] -> nil

// 3. Right subtree: [5,9]
//    len = 2, mid = 2/2 = 1
//    node = 9
//    - Left: [5] -> node = 5
//    - Right: [] -> nil

// Cây kết quả:
//       0
//      / \
//    -3   9
//    /   /
// -10   5