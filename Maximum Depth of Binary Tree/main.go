package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}

func main(){
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	println(maxDepth(root))
}


// maxDepth(3):
//     left = maxDepth(9):
//         left = maxDepth(nil) = 0
//         right = maxDepth(nil) = 0
//         return max(0,0) + 1 = 1

//     right = maxDepth(20):
//         left = maxDepth(15):
//             left = maxDepth(nil) = 0
//             right = maxDepth(nil) = 0
//             return max(0,0) + 1 = 1
//         right = maxDepth(7):
//             left = maxDepth(nil) = 0
//             right = maxDepth(nil) = 0
//             return max(0,0) + 1 = 1
//         return max(1,1) + 1 = 2

//     return max(1,2) + 1 = 3