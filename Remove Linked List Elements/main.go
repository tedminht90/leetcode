package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// func removeElements(head *ListNode, val int) *ListNode {
//     if head == nil {
// 		return nil
// 	}
// 	head.Next = removeElements(head.Next, val)
// 	if head.Val == val {
// 		return head.Next
// 	}
// 	return head
// }

func removeElements(head *ListNode, val int) *ListNode {
	// Kiem tra head dau tien xem co rong hay khong
	if head == nil {
		return nil
	}
	// neu head dau tien bang val
	if head != nil && head.Val == val {
		head = head.Next
	}
	// Duyet qua cac node con lai
	current := head
	for current.Next != nil {
		if current.Next.Val == val{
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func main() {
	values := []int{1, 2, 6, 3, 4, 5, 6}
	head := &ListNode{Val: values[0]}
	cur := head
	for i := 1; i < len(values); i++ {
    	cur.Next = &ListNode{Val: values[i]}
    	cur = cur.Next
	}

	result := removeElements(head, 6)
	for result != nil {
		println(result.Val)
		result = result.Next
	}
}

