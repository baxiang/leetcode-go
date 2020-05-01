package main

import "fmt"

type ListNode struct {
      Val int
     Next *ListNode
}
//https://leetcode-cn.com/problems/merge-two-sorted-lists/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
     dummy := &ListNode{}
     cur := dummy
     for {
		 if l1==nil&&l2==nil {
			 break
		 }
		 if l1==nil {
			 cur.Next = l2
			 break
		 }
		 if l2==nil {
			 cur.Next = l1
			 break
		 }
		 if l1.Val<l2.Val {
			 cur.Next = l1
			 l1 = l1.Next
			 cur = cur.Next
		 }else {
		 	cur.Next = l2
		 	l2 = l2.Next
		 	cur = cur.Next
		 }
	 }
     return dummy.Next
}
func main() {
	a1 := &ListNode{Val:  1}
	b1 := &ListNode{Val:  2}
	c1 := &ListNode{Val:  4}
	a2 := &ListNode{Val:  1}
	b2 := &ListNode{Val:  3}
	c2 := &ListNode{Val:  4}
	a1.Next=b1
	b1.Next = c1
	a2.Next = b2
	b2.Next = c2
	h := mergeTwoLists(a1,a2)
	for h!=nil{
		fmt.Println(h.Val)
		h=h.Next
	}
}
