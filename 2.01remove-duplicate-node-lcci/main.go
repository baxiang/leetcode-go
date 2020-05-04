package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes1(head *ListNode) *ListNode {
    m :=make(map[int]struct{})
    dummyHead :=&ListNode{Val:-1}
	pre :=dummyHead
    for head!=nil{
    	if _,ok:=m[head.Val];!ok{
			m[head.Val] = struct{}{}
			pre.Next = head
			pre = head
		}else {
			pre.Next = head.Next
		}
		head = head.Next
	}
	return  dummyHead.Next
}


func removeDuplicateNodes(head *ListNode) *ListNode {
	m :=make(map[int]struct{})
	dummyHead :=&ListNode{Next:head}
	curr :=dummyHead
	for curr!=nil&&curr.Next!=nil{
		val :=curr.Next.Val
		if _,ok:=m[val];ok{
			curr.Next = curr.Next.Next
		}else {
			m[val] = struct{}{}
			curr = curr.Next
		}
	}
	return  dummyHead.Next
}
func main() {
	n1 :=&ListNode{Val:1}
	n2 :=&ListNode{Val:2}
	n1.Next = n2
	n3 :=&ListNode{Val:3}
	n2.Next = n3
	n4 :=&ListNode{Val:3}
	n3.Next = n4
	n5 :=&ListNode{Val:2}
	n4.Next = n5
	n6:=&ListNode{Val:1}
	n5.Next = n6
	h :=removeDuplicateNodes1(n1)
	for h!=nil{
		fmt.Println(h.Val)
		h=h.Next
	}
}
