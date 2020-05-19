package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	greatDummy := &ListNode{}
	lessDummy := &ListNode{}
	currGreat := greatDummy
	currLess := lessDummy
	for head != nil {
		if head.Val < x {
			currLess.Next = head
			currLess = head
		} else {
			currGreat.Next = head
			currGreat = head
		}
		head = head.Next
	}
	currGreat.Next = nil
	currLess.Next = greatDummy.Next
	return lessDummy.Next
}

func main() {

}
