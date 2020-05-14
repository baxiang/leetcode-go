package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	var oddDummy = &ListNode{}
	var evenDummy = &ListNode{}
	var i = 1
	var preOdd = oddDummy
	var preEven = evenDummy
	for head != nil {
		next := head.Next
		if i%2 != 0 {
			preOdd.Next = head
			preOdd = head
		} else {
			preEven.Next = head
			preEven = head
		}
		i++
		head = next
	}
	preEven.Next = nil
	preOdd.Next = evenDummy.Next
	return oddDummy.Next
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
	//for one!=nil{
	//	fmt.Println(one.Val)
	//	one = one.Next
	//}
	h := oddEvenList(one)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}
