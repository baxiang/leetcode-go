package main

import "fmt"

type ListNode struct {
	     Val int
	     Next *ListNode
}
func revert(head *ListNode)*ListNode{
	var pre *ListNode
	curr := head
	for curr!= nil{
		pre,curr,curr.Next = curr,curr.Next,pre
	}
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val:  -1,}
	dummy.Next = head
	pre,curr,end := dummy,dummy,dummy
	for end!=nil{
		// 获取当前节点
		curr = pre.Next
		// 获取结束结点
		for i:=0;i<k;i++{
			end = end.Next
			if end == nil{
				return dummy.Next
			}
		}
		// 暂存后节点
		next := end.Next
		end.Next = nil// 至空
		pre.Next = revert(curr)
		// 反转结束后当前子链表的头结点
		pre,end = curr,curr// 都是当前值
		curr.Next = next//需要把链表串起来
	}
	return dummy.Next
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val:  -1,}
	dummy.Next = head
	pre := dummy
	stack := make([]*ListNode,0)
	n := k
	for pre.Next!=nil{
		tmp := pre.Next
		for tmp!= nil&&n>0{
			stack = append(stack,tmp)
			tmp = tmp.Next
			n--
		}
		nextNode := stack[len(stack)-1].Next
		if n==0{
			for len(stack)>0{
				node :=stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				pre.Next = node
				pre = pre.Next
			}
			//需要保证节点的连续行
			pre.Next = nextNode
			n = k
		}else {
			return dummy.Next
		}
	}
	return dummy.Next
}

func main() {
	a := &ListNode{Val:  1}
	b :=&ListNode{Val:  2}
	a.Next =b
	c :=&ListNode{Val:  3}
	b.Next = c
	d :=&ListNode{Val:  4}
	c.Next = d
	e :=&ListNode{Val:  5}
	d.Next = e
	f :=&ListNode{Val:  6}
	e.Next = f
	r := reverseKGroup1(a,3)
	for r!=nil{
		fmt.Println(r.Val)
		r = r.Next
	}
}
