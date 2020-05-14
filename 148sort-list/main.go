package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val int
	Next *ListNode
}

//归并排序链表
func sortList(head *ListNode) *ListNode {
	var res []*ListNode
	for head !=nil{
		res = append(res,head)
		head = head.Next
	}
	sort.Slice(res, func(i, j int) bool {
		 return res[i].Val<res[j].Val
	})
    dummyHead :=&ListNode{}
    preNode := dummyHead
	for i:=0;i<len(res);i++{
		preNode.Next = res[i]
		preNode = res[i]
	}
    preNode.Next = nil
	return dummyHead.Next
}


func main() {
	four :=&ListNode{Val:4}
	two :=&ListNode{Val:2}
	four.Next = two
	one :=&ListNode{Val:1}
	two.Next = one
	three :=&ListNode{Val:3}
	one.Next = three

	r := sortList(four)
	for r!=nil{
		fmt.Println(r.Val)
		r = r.Next
	}


}
