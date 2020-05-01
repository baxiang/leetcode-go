package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	dummy.Next = head
	preNode := dummy // 就是前一个节点
	for head != nil && head.Next!= nil{
		first := head
		second := head.Next
		//交换前两个节点

		first.Next = second.Next
		second.Next = first
		preNode.Next =second

		preNode = first //需要把位置移到前一个节点去
		head = first.Next
	}
	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	preNode := dummy
	for head!=nil&&head.Next!=nil{
		first := head
		second := head.Next

		// 交换数据
		preNode.Next = second
		first.Next = second.Next
		second.Next = first

		// 移动
		preNode = first
		head = first.Next
	}
	return dummy.Next
}


func swapPairs2(head *ListNode) *ListNode {
	//终止条件：链表只剩一个节点或者没节点了，没得交换了。返回的是已经处理好的链表
	if head ==nil||head.Next==nil {
		return head
	}
	//一共三个节点:head, next, swapPairs(next.next)
	//下面的任务便是交换这3个节点中的前两个节点
	next := head.Next
	head.Next= swapPairs2(next.Next)
	next.Next = head
	//根据第二步：返回给上一级的是当前已经完成交换后，即处理好了的链表部分
	return  next
}


func main() {
	a := &ListNode{Val:  1}
	b := &ListNode{Val:  2}
	c := &ListNode{Val:  3}
	d := &ListNode{Val:  4}
	e := &ListNode{Val:  5}
	f := &ListNode{Val:  6}
	a.Next=b
	b.Next =c
	c.Next = d
	d.Next = e
	e.Next = f
	h := swapPairs2(a)
	for h!=nil{
		fmt.Println(h.Val)
		h=h.Next
	}
}
