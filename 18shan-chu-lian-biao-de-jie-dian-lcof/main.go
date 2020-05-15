package main

import "fmt"

type ListNode struct {
	     Val int
	     Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	preNode := dummyHead
	for head!= nil{
		next := head.Next
		if head.Val==val{
			preNode.Next = next
			break
		}
		preNode = head
		head = next
	}
	return dummyHead.Next
}


func main() {
	four :=&ListNode{Val: 4}
	five :=&ListNode{Val: 5}
	four.Next = five
	one :=&ListNode{Val: 1}
	five.Next = one
	nine :=&ListNode{Val: 9}
	one.Next = nine


	r := deleteNode(four,4)
	for r!=nil{
		fmt.Println(r.Val)
		r = r.Next
	}
}
