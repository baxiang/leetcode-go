package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	m :=make(map[int]struct{})
	curr := head
	var pre *ListNode
	for curr!=nil{
		tmp := curr.Next
		if _,ok:=m[curr.Val];ok{
			pre.Next = tmp
			curr.Next = nil
		}else {
			pre = curr
			m[curr.Val] = struct{}{}
		}
		curr = tmp
	}
	return head
}
func main() {
	a :=ListNode{
		Val:  1,
	}
	b :=&ListNode{
		Val:  1,
	}
	a.Next =b
	c :=&ListNode{
		Val:  2,
	}
	b.Next = c
	d :=&ListNode{
		Val:  3,
	}
	c.Next = d
	f :=&ListNode{
		Val:  3,
	}
	d.Next =f
	r := deleteDuplicates(&a)
	for r!=nil{
		fmt.Println(r.Val)
		r = r.Next
	}
}
