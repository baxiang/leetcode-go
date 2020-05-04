package main

import "unsafe"

type ListNode struct {
	     Val int
	     Next *ListNode
}


func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m :=make(map[uintptr]struct{},0)
	for headA!=nil{
		m[uintptr(unsafe.Pointer(headA))]= struct{}{}
		headA = headA.Next
	}
	for headB!=nil{
		if _,ok:=m[uintptr(unsafe.Pointer(headB))];ok{
			return headB
		}
		headB = headB.Next
	}
	return nil
}

//面试题 02.07. 链表相交
func main() {

}
