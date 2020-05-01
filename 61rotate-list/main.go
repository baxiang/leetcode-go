package main

import "fmt"
type ListNode struct {
	     Val int
	     Next *ListNode
 }



func rotateRight(head *ListNode, k int) *ListNode {
	if head==nil||head.Next==nil||k==0 {
		return head
	}
	count :=1
	cursor := head
	for cursor.Next!=nil{
		count++
		cursor =cursor.Next
	}
	//// 相等不需要移动
	k =k%count
	if k==0 {
		return head
	}
	//当前是尾结点 // 头尾相连  循环链表
	cursor.Next = head
	// 找到需要分割的节点
	/*
	   1 2 3 4 第0次
	   4 1 2 3 第1次
	   3 4 1 2 第2次
	*/
	l :=count-(k%count)
	pre := head
	for i:=0;i<l-1;i++{
		pre = pre.Next

	}
	newHead := pre.Next
	pre.Next = nil
	return newHead
}
func main() {
	one :=&ListNode{
		Val:    1,
	}
	two :=&ListNode{
		Val:    2,
	}
	three :=&ListNode{
		Val:    3,
	}
	four :=&ListNode{
		Val:    4,
	}
	five :=&ListNode{
		Val:    5,
	}
	one.Next = two
	two.Next = three
	three.Next= four
	four.Next = five
	 h:=rotateRight(one,2)
	 for h!=nil{
	 	fmt.Println(h.Val)
	 	h =h.Next
	 }
}
