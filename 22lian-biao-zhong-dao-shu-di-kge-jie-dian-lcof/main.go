package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd1(head *ListNode, k int) *ListNode {
	var res []*ListNode
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	l := len(res)
	if l >= k {
		return res[l-k]
	}
	return nil
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var fast = head
	for k > 0 && fast != nil {
		fast = fast.Next
		k--
	}
	var slow = head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	one := &ListNode{Val: 1}
	two := &ListNode{Val: 2}
	one.Next = two
	three := &ListNode{Val: 3}
	two.Next = three
	four := &ListNode{Val: 4}
	three.Next = four
	five := &ListNode{Val: 5}
	four.Next = five
	r := getKthFromEnd(one, 2)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
