package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead :=&ListNode{Val:  -1,}
    curr :=dummyHead
    for l1!=nil||l2!=nil{
    	if l1!=nil&&l2!=nil{
    		if l1.Val<l2.Val{
				curr.Next = l1
				curr = l1
    			l1 =l1.Next
			}else {
				curr.Next = l2
				curr = l2
				l2 =l2.Next
			}
		}else if l1!=nil{
			curr.Next = l1
			break
		}else {
			curr.Next = l2
			break
		}

	}
	return dummyHead.Next
}

//面试题25. 合并两个排序的链表
func main() {
  l1:=&ListNode{Val:  1,}
	two:=&ListNode{Val:  2,}
	l1.Next=two
	four:=&ListNode{Val:  4,}
	two.Next=four


	l2:=&ListNode{Val:  1,}
	three:=&ListNode{Val:  3,}
	l2.Next=three
	f:=&ListNode{Val:  4,}
	three.Next=f

	h := mergeTwoLists(l1,l2)
	for h!=nil{
		fmt.Println(h.Val)
		h =h.Next
	}
}
