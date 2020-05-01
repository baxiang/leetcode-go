package main

import "fmt"
type ListNode struct {
     Val int
     Next *ListNode
 }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
      //dump node
      d := &ListNode{Next:head}
      len :=0
      first := head
      for first!=nil{
      	first = first.Next
      	len++
	  }
      len =len-n// revert index
      first = d
      for len >0{
      	len--
      	first = first.Next
	  }
      first.Next = first.Next.Next
      return d.Next
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	//dump node
	dummyNode := &ListNode{}// 创建哨兵节点
	dummyNode.Next = head // 将哨兵节点指向头结点
	curr := dummyNode // 从哨兵节点开始存储 可以保证n=len(list)正数第一个节点
	var list []*ListNode // 遍历数组存储
	for curr!=nil{
		list = append(list,curr)
		curr = curr.Next
	}
	pre := list[len(list)-n-1]
	curr = list[len(list)-n]
	pre.Next = curr.Next // 去掉当前节点
	return dummyNode.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	//dump node
	dummyNode := &ListNode{}// 创建哨兵节点
	dummyNode.Next = head // 将哨兵节点指向头结点
	var fast = head
	for i:=0;i<n;i++{
		fast =fast.Next
	}
	var slow = head
	var pro = dummyNode
	for fast!=nil{
		pro = slow
		slow= slow.Next
		fast = fast.Next
	}
	pro.Next = slow.Next
	return dummyNode.Next
}





func main() {
	head := &ListNode{Val:1}
	head.Next = &ListNode{Val:2}
	node := head.Next
	node.Next = &ListNode{Val:3}
	node = node.Next
	node.Next = &ListNode{Val:4}
	node = node.Next
	node.Next = &ListNode{Val:5}
	d :=removeNthFromEnd2(head,2)
	for  d!=nil {
		fmt.Println(d.Val)
		d = d.Next
	}
}
