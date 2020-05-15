package main

import "unsafe"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	  m := make(map[uintptr]struct{})
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


func main() {
	
}
