package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := l1
	q := l2
	currentHead := dummy
	var sum,carry int // carry 进位值
	for true  {
		if q != nil || p != nil {
			var x,y int
			if p!=nil {
				x=p.Val
				p = p.Next
			}
			if q!=nil {
				y =q.Val
				q = q.Next
			}

			sum = carry + x + y
			carry = sum / 10
		} else {
			break
		}
		currentHead.Val = sum % 10
		if p!=nil||q!=nil {
			currentHead.Next = &ListNode{}
			currentHead = currentHead.Next
		}

	}
	if carry != 0 {
		currentHead.Next = &ListNode{carry, nil}
	}
	return dummy
}
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	p :=l1
	q := l2
	curr := dummyHead
	carry :=0
	for p!=nil||q!=nil{
		var x,y int
		if p!=nil {
			x = p.Val
			p = p.Next
		}
		if q!= nil {
			y = q.Val
			q = q.Next
		}
		sum := carry+x+y
		carry = sum/10
		curr.Next =&ListNode{
			Val:  sum%10,
			Next: nil,
		}
		curr = curr.Next
	}
	if carry>0{
		curr.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return dummyHead.Next
}
func main() {
	l1:= &ListNode{Val:2,Next:&ListNode{Val:4,Next:&ListNode{Val:3}}}
	l2:= &ListNode{Val:5,Next:&ListNode{Val:6,Next:&ListNode{Val:4}}}
	r := addTwoNumbers1(l1,l2)
	curr := r
	var str strings.Builder
	for curr!=nil{
		str.WriteString(fmt.Sprintf("%d -> ",curr.Val))
		curr = curr.Next
	}
	fmt.Println(strings.TrimRight(str.String()," -> "))
	ll1:= &ListNode{Val:0,Next:nil}
	ll2:= &ListNode{Val:0,Next:nil}
	r = addTwoNumbers1(ll1,ll2)
	curr = r
	var str1 strings.Builder
	for curr!=nil{
		str1.WriteString(fmt.Sprintf("%d -> ",curr.Val))
		curr = curr.Next
	}
	fmt.Println(strings.TrimRight(str1.String()," -> "))
}

