package main

import (
	"fmt"
)

type ListNode struct {
	     Val int
	     Next *ListNode
}
func isPalindrome(head *ListNode) bool {
	if head==nil {
		return  true
	}
	var list []*ListNode
      for head!= nil{
      	list = append(list,head)
      	head = head.Next
	  }
	  i,j:= 0,len(list)-1
	  for i<j{
		  if list[i].Val!=list[j].Val {
			  return false
		  }
		  i++
		  j--
	  }
	  return true
}

func isPalindrome1(head *ListNode) bool {
	if head==nil {
		return  true
	}
	mid := findMid(head)
    rev := revertNode(mid)
	for head!=nil&&rev!= nil{
		if head.Val!= rev.Val {
			return false
		}
		head= head.Next
		rev = rev.Next
	}
	return true
}

func findMid(head *ListNode)*ListNode{
	low,fast := head,head
	for fast!= nil&&fast.Next!=nil{
		low = low.Next
		fast = fast.Next.Next
	}
	return low
}

func revertNode1(head *ListNode)*ListNode{
	var pre *ListNode
	curr := head
	for curr!= nil{
		pre,curr,curr.Next = curr,curr.Next,pre
	}
  return pre
}
func revertNode(head *ListNode)*ListNode{
	var pre *ListNode
	curr := head
	for curr!= nil{
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
		//pre,curr,curr.Next = curr,curr.Next,pre
	}
	return pre
}

func main() {
	a :=&ListNode{
		Val:  1,
	}
	b :=&ListNode{
		Val:  2,
	}
	a.Next = b
	fmt.Println(isPalindrome1(a))

	a1 :=&ListNode{
		Val:  1,
	}
	b1 :=&ListNode{
		Val:  2,
	}
	a1.Next = b1
	c1 :=&ListNode{
		Val:  2,
	}
	b1.Next = c1
	d2 :=&ListNode{
		Val:  1,
	}
	c1.Next = d2
	fmt.Println(isPalindrome1(a1))
}
