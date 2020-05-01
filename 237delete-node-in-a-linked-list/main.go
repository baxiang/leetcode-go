package main

import "fmt"




func deleteNodeByHead(head,node *ListNode)*ListNode {
	dummyHead :=&ListNode{
		Val:  -1,
		Next: head,
	}
	var pre = dummyHead
	for head!=nil{
		if head.Val==node.Val {
			next := node.Next
			pre.Next = next
			head = next
		}else {
			pre = head
			head = head.Next
		}
	}
	return dummyHead.Next
}


type ListNode struct {
	Val int
	Next *ListNode
}
// 无法获取前置节点 只能把下一个节点干掉 ，然后把一个节点的值复制给当前节点
func deleteNode(node *ListNode) {
	next := node.Next// 当前的下一个节点
	node.Val = next.Val // 替换值
	node.Next = next.Next //让当前节点指向的是下一个节点的下一个
}


func main() {
	head :=&ListNode{Val:4}
	n1 :=&ListNode{Val:5}
	head.Next = n1
	n2 :=&ListNode{Val:1}
	n1.Next = n2
	n2.Next =&ListNode{Val:9}
	deleteNode(n1)// 去掉5个节点
	curr :=head
	for curr != nil{
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
