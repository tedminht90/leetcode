package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow := head // slow pointer bắt đầu từ head
	fast := head // fast pointer bắt đầu từ head
	for fast != nil && fast.Next != nil { 
		slow = slow.Next // slow pointer tăng 1 bước
		fast = fast.Next.Next // fast pointer tăng 2 bước
	}
	if fast != nil {
		slow = slow.Next // nếu fast != nil thì slow tăng thêm 1 bước
	}
	slow = reverse(slow) // reverse list từ slow pointer trở đi
	fast = head // fast pointer trở về head
	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	list := []int{1, 2}
	var head *ListNode
	head = &ListNode{Val: list[0]}
	current := head
	for i := 1; i < len(list); i++ {
		current = &ListNode{Val: list[i]}
		current = current.Next
	}

	fmt.Println(isPalindrome(head))
}