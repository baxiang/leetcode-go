package main

import "fmt"

type ListNode struct {
	 Val int
	 Next *ListNode
 }


func reverseList1(head *ListNode) *ListNode {
	if head == nil||head.Next==nil{
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil{
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil||head.Next==nil{
		return head
	}
	dummyNode := &ListNode{}// 哑结点
	cur := head
	for cur != nil{
		next := cur.Next // 为了遍历，临时变量存储当前节点的下一个节点
		head :=  dummyNode.Next // 临时变量存储当前哑结点的指向的头结点
		dummyNode.Next = cur // 往头部插入数据,当前节点变成了头结点了
		cur.Next = head // 为了维持链条 需要把之前的头结点变成下一个节点
		cur = next // 遍历移动
	}
	return dummyNode.Next // 返回头结点
}


func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr!= nil{//当为空表示迭代结束
		//前一个值 ，当前值，后一个值
		pre,curr,curr.Next = curr,curr.Next,pre
	}
	return pre
}

func makeLinkedList(list []int)*ListNode{
	dummyNode := &ListNode{} // 创建哑结点 指向头节点
	curr := dummyNode
	for i:=0;i<len(list);i++{
		curr.Next= &ListNode{Val:list[i]}
		curr = curr.Next
	}
	return dummyNode.Next
}



func main() {
	head := makeLinkedList([]int{1,2,3,4,5})
	head = reverseList(head)
	for head!=nil{
		fmt.Println(head.Val)
		head = head.Next
	}
}
