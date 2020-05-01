package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
    dummyHead :=&ListNode{Val:  -1}
    var pre *ListNode
    curr := head
	for curr!=nil{
		pre = dummyHead
		next := curr.Next
		// 查找插入位置，从头结点开始遍历，知道当前节点的数据大于插入值停止
		for pre.Next!= nil&&pre.Next.Val<curr.Val{
			pre = pre.Next
		}
		// pre 节点就是需要插入节点位置的前节点
		// 插入一个新的节点
		curr.Next = pre.Next
		pre.Next = curr
		curr = next
	}
	return dummyHead.Next
}


func main() {
	four :=&ListNode{Val:  4}
	two :=&ListNode{Val:  2}
	four.Next = two
	one :=&ListNode{Val:  1}
	two.Next = one
	three :=&ListNode{Val:  3}
	one.Next= three
	h:= insertionSortList(four)
	for h!=nil{
		fmt.Println(h.Val)
		h =h.Next
	}
	fmt.Println("------")

	 head :=&ListNode{Val:  -1}
	 five:=&ListNode{Val:  5}
	 head.Next = five
	 five.Next = three
	 three.Next= four
	 zero :=&ListNode{Val:  0}
	 four.Next= zero
	 h = insertionSortList(head)
	 for h!=nil{
		fmt.Println(h.Val)
		h =h.Next
	 }
}
