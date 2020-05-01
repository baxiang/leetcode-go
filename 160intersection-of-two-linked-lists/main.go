package main

import "unsafe"

type ListNode struct {
	     Val int
	     Next *ListNode
 }

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m :=make(map[uintptr]struct{})
	for headA!=nil||headB!=nil{
		if headA!=nil {
			if _,ok:=m[uintptr(unsafe.Pointer(headA))];ok{
				return  headA
			}else {
				m[uintptr(unsafe.Pointer(headA))] = struct{}{}
			}
			headA = headA.Next
		}
		if headB!=nil {
			if _,ok:=m[uintptr(unsafe.Pointer(headB))];ok{
				return  headB
			}else {
				m[uintptr(unsafe.Pointer(headB))] = struct{}{}
			}
			headB= headB.Next
		}
	}
	return nil
}

func main() {
	
}
