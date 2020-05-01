package main

import (
	"fmt"
	"unsafe"
)

type ListNode struct {
 	Val int
 	Next *ListNode
}
func hasCycle(head *ListNode) bool {
	if head==nil||head.Next==nil {
		return false
	}
    slow :=head
    fast := head.Next
    for fast!= nil&&fast.Next!=nil{
		if slow==fast {
			return true
		}
		slow = slow.Next// 走一步
		fast = fast.Next.Next //走2部
	}
	return false
}

func hasCycle1(head *ListNode) bool {
	if head==nil||head.Next==nil {
		return false
	}
	for head!=nil{
		if head.Val==1>>63 {
			return true
		}
		head.Val = 1>>63
		head = head.Next
	}
	return false
}
func hasCycle2(head *ListNode) bool {
	if head==nil||head.Next==nil {
		return false
	}
	m :=make(map[uintptr]struct{})
	for head!=nil{
		if _,ok:=m[uintptr(unsafe.Pointer(head))];ok {
			return true
		}
		m[uintptr(unsafe.Pointer(head))] = struct{}{}
		head = head.Next
	}
	return false
}

func main() {
	aa := ListNode{Val:  1}
	//b := ListNode{Val:  2}
	//c := ListNode{Val:  0}
	//d := ListNode{Val:  -4}
	//a.Next = &b
	//b.Next = &c
	//c.Next = &d
	//d.Next = &b
	fmt.Println(hasCycle2(&aa))
	a := ListNode{Val:  1}
	b := ListNode{Val:  2}
	c := ListNode{Val:  0}
	d := ListNode{Val:  -4}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &b
	fmt.Println(hasCycle2(&a))
}
