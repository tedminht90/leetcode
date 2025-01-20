package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
    Val  int // giá trị số nguyên của node
    Next *ListNode // con trỏ tới node tiếp theo
}

// Hàm tạo danh sách liên kết từ mảng
// Nhận một mảng số nguyên
// Tạo danh sách liên kết từ mảng
// Sử dụng kỹ thuật node dummy tương tự
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
	// &ListNode{Val: num} tạo một node mới với:
	// 	- Val được gán giá trị là num
	// 	- Next mặc định là nil
	// Con trỏ Next của node hiện tại (current.Next) được gán tới node mới
		current.Next = &ListNode{Val: num} 
		current = current.Next // Di chuyển con trỏ current tới node vừa tạo
	}
	return dummy.Next
}

// Merge 2 danh sách liên kết đã sắp xếp
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Tạo node giả làm điểm khởi đầu
    dummy := &ListNode{}
	// current là con trỏ di chuyển, ban đầu trỏ tới dummy
	current := dummy

	// Vòng lặp chính: so sánh và nối các node từ 2 danh sách
	for list1 != nil && list2 != nil { // Chạy cả 2 khi danh sách còn phần tử
		if list1.Val < list2.Val { // So sánh giá trị 2 node đầu tiên
			current.Next = list1 // Nối node nhỏ hơn vào danh sách kết quả
			list1 = list1.Next // Di chuyển con trỏ list1 tới node tiếp theo
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next // Di chuyển con trỏ current tới node vừa nối

	}
	// Nối phần còn lại của list dài hơn
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}
	return dummy.Next // trả về node đầu tiên thực sự
}

// Hàm in danh sách liên kết
func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println()
}

func main() {
	// Nhận một mảng số nguyên
	// Tạo danh sách liên kết từ mảng
	// Sử dụng kỹ thuật node dummy tương tự
	list1 := createList([]int{1, 2, 4})
	list2 := createList([]int{1, 3, 4})

	// Trộn hai danh sách liên kết
	mergedList := mergeTwoLists(list1, list2)

	// In kết quả
	printList(mergedList)
	
}

// Input:
// list1: 1 -> 2 -> 4
// list2: 1 -> 3 -> 4

// Các bước thực hiện:
// 1. So sánh 1 và 1:
//    - Chọn list1.1
//    - current -> 1 (list1)
//    - list1 = 2->4

// 2. So sánh 2 và 1:
//    - Chọn list2.1
//    - current -> 1 (list2)
//    - list2 = 3->4

// 3. So sánh 2 và 3:
//    - Chọn list1.2
//    - current -> 2 (list1)
//    - list1 = 4

// 4. So sánh 4 và 3:
//    - Chọn list2.3
//    - current -> 3 (list2)
//    - list2 = 4

// 5. So sánh 4 và 4:
//    - Chọn list1.4
//    - current -> 4 (list1)
//    - list1 = nil

// 6. list1 rỗng, nối phần còn lại của list2 (4)

// Kết quả: 1 -> 1 -> 2 -> 3 -> 4 -> 4