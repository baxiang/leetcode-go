package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
    var list []int
	for head!=nil{
		list = append(list,head.Val)
		head = head.Next
	}
	var res int
	l := len(list)
	for i:=0;i<l;i++{
		if list[i]==1{
			r := int(math.Pow(2,float64(l-i-1)))
			res +=r
		}
	}
	return res
}

func main() {
	head :=&ListNode{Val:1}
	one :=&ListNode{Val:0}
	head.Next= one
	two :=&ListNode{Val:1}
	one.Next = two

	fmt.Println(getDecimalValue(head))
}
