package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(node *ListNode) *ListNode {
	slow, fast := node, node
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main(){
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println(middleNode(root))
}