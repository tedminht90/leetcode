package main

import "fmt"



type ListNode struct {
	Val int
	Next *ListNode
}


func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}


func main() {

	// 3 -> 2 -> 0 -> -4 -> 2
	head := &ListNode{Val: 3}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 0}
	node3 := &ListNode{Val: -4}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node2
	fmt.Println(hasCycle(head))

}