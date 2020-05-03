package main

import "fmt"

type ListNode struct {
	     Val int
	     Next *ListNode
	 }

func middleNode1(head *ListNode) *ListNode {
	var res []*ListNode
    for head!=nil{
    	res =append(res,head)
    	head = head.Next
	}
	return res[len(res)/2]
}

func middleNode(head *ListNode) *ListNode {
	slow :=head
	fast :=head
	for fast.Next!=nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}


func main() {
	one :=&ListNode{Val:  1}
	two :=&ListNode{Val:  2}
	one.Next=two
	three :=&ListNode{Val:  3}
	two.Next=three
	four :=&ListNode{Val:  4}
	three.Next=four
	five :=&ListNode{Val:  5}
	four.Next=five
	//six :=&ListNode{Val:  6}
	//five.Next=six
    fmt.Println(middleNode(one).Val)
}
