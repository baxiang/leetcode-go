package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	preNode := dummyHead
	for head != nil {
		if head.Val == val {
			next := head.Next
			preNode.Next = next
			head = next
		} else {
			preNode = head
			head = head.Next
		}
	}
	return dummyHead.Next
}

func main() {
	one := &ListNode{Val: 1}
	two := ListNode{Val: 2}
	one.Next = &two
	three := ListNode{Val: 6}
	two.Next = &three
	four := ListNode{Val: 3}
	three.Next = &four
	five := ListNode{Val: 4}
	four.Next = &five
	six := ListNode{Val: 5}
	five.Next = &six
	seven := ListNode{Val: 6}
	six.Next = &seven
	removeElements(one, 6)
	curr := one
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
