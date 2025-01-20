package main


type ListNode struct {
	Val  int
	Next *ListNode
}


// func reverseList(node *ListNode) *ListNode {
// 	if node == nil || node.Next == nil {
// 		return node
// 	}
// 	rest := reverseList(node.Next)

// 	node.Next.Next = node
// 	node.Next = nil

// 	return rest

// }
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// prev = nil
// cur = 1

// Lần lặp 1
// next = 2->6->3->4->5->6
// cur.Next = nil
// prev = 1 -> nil
// cur = 2->6->3->4->5->6

// Lần lặp 2
// next = 6->3->4->5->6
// cur.Next = 1
// prev = 2 -> 1 -> nil
// cur = 6->3->4->5->6

// lần lặp 3
// next = 3->4->5->6
// cur.Next = 2
// prev = 6 -> 2 -> 1 -> nil
// cur = 3->4->5->6

// lần lặp 4
// next = 4->5->6
// cur.Next = 3
// prev = 3->6->2->1->nil
// cur = 4->5->6

// lần lặp 5
// next = 5->6
// cur.Next = 4
// prev = 4->3->6->2->1->nil
// cur = 5->6

// lần lăp 6
// next = 6
// cur.Next = 6
// prev = 5->4->3->6->2->1->nil
// cur = 6

// lần lặp 7
// next = nil
// cur.Next = nil
// prev = 6->5->4->3->2->1->nil
// cur = nil



// --> prev = 6->5->4->3->2->1->nil

func main() {
	values := []int{1, 2, 6, 3, 4, 5, 6}
	head := &ListNode{Val: values[0]}
	cur := head
	for i := 1; i < len(values); i++ {
    	cur.Next = &ListNode{Val: values[i]}
    	cur = cur.Next
	}

	result := reverseList(head)
	for result != nil {
		println(result.Val)
		result = result.Next
	}
}
