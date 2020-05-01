package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
    stack :=make([]*ListNode,0)
    for head!=nil{
    	stack = append(stack,head)
    	head =head.Next
	}
	var res []int
	for len(stack)>0{
		node :=stack[len(stack)-1]
		res = append(res,node.Val)
		stack = stack[:len(stack)-1]
	}
	return res
}

func main() {
	one :=&ListNode{Val:1}
	three :=&ListNode{Val:3}
	one.Next = three
	two :=&ListNode{Val:2}
	three.Next=two
	fmt.Println(reversePrint(one))
}
