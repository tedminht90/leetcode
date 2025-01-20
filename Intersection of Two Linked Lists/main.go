package main


type ListNode struct {
	Val int
	Next *ListNode
}

func createList(value []int) *ListNode {
	if value == nil || len(value) == 0 {
		return nil
	}

	head := &ListNode{Val: value[0]}
	current := head
	for i := 1; i < len(value); i++ {
		current.Next = &ListNode{Val: value[i]}
		current = current.Next
	}
	return head
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// Kiểm tra nếu một trong hai danh sách rỗng thì trả về nil
	if headA == nil || headB == nil {
        return nil // Không thể có giao điểm nếu một trong hai danh sách rỗng
    }

    // Khởi tạo hai con trỏ, mỗi con trỏ sẽ duyệt qua cả hai danh sách
    ptrA := headA // Con trỏ ptrA bắt đầu từ đầu danh sách A
    ptrB := headB // Con trỏ ptrB bắt đầu từ đầu danh sách B
    
	// Vòng lặp chạy cho đến khi hai con trỏ gặp nhau
    for ptrA != ptrB {
		// Di chuyển cả 2 con trỏ đến node tiếp theo
        ptrA = ptrA.Next
        ptrB = ptrB.Next
        
		// Nếu cả 2 con trỏ đều là nil, tức là đã duyệt hết cả hai danh sách
		// mà không tìm thấy giao điểm, trả về nil
        if ptrA == nil && ptrB == nil {
            return nil
        }
        // Nếu con trỏ A đến cuối danh sách A, 
		// thì chuyển con trỏ A sang danh sách B
        if ptrA == nil {
            ptrA = headB
        }
        
		// Nếu con trỏ B đến cuối danh sách B,
		// thì chuyển con trỏ B sang danh sách A
        if ptrB == nil {
            ptrB = headA
        }
    }
    // Khi vòng lặp kết thúc, hai con trỏ đã gặp nhau
	// Trả về node tại thời điểm giao nhau
    return ptrA
}

func main() {
	intersection := createList([]int{8, 4, 5})


	// Create List A: 4 -> 1 -> 8 -> 4 -> 5
	listA := createList([]int{4, 1})
	tailA := listA
	for tailA.Next != nil {
		tailA = tailA.Next
	}
	tailA.Next = intersection


	// Create List B: 5 -> 0 -> 1 -> 8 -> 4 -> 5
	listB := createList([]int{5, 0, 1})
	tailB := listB
	for tailB.Next != nil {
		tailB = tailB.Next
	}
	tailB.Next = intersection

	// Find the intersection of two linked lists
	result := getIntersectionNode(listA, listB)
	if result != nil {
		println(result.Val)
	} else {
		println("nil")
	}
}