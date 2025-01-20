package main

import "fmt"


type ListNode struct {
	Val int
	Next *ListNode
}

func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
}

func main() {
	head := []int{1,1,2}
	list := createList(head)
	
	temp := deleteDuplicates(list)
	printList(temp)


}