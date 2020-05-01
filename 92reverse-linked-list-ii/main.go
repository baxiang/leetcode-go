package main

import "fmt"

type ListNode struct {
	     Val int
	     Next *ListNode
}

func reverseBetween1(head *ListNode, m int, n int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	begin := dummyNode
	for i:=0;i<m-1;i++{
		begin = begin.Next
	}
	var pre *ListNode
	curr := begin.Next
	for i:=0;i<n-m+1;i++{
		tmp := curr.Next
		curr.Next = pre
		pre = curr
		curr = tmp
	}
	begin.Next.Next = curr// 链接后节点
	begin.Next = pre // 开始反转的最后节点
	return dummyNode.Next
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummyHead :=&ListNode{
		Val:  -1,
		Next: head,
	}
	curr :=head
	i := 1
	pre :=dummyHead
	var next *ListNode
	var start *ListNode
	for curr!=nil{
		if i==m {
			start = curr
		}
		if i==n{
			next = curr.Next
			curr.Next = nil
			pre.Next = revertNode(start)
			start.Next = next
			return dummyHead.Next
		}
		if i<=m-1{
			pre = curr //获取前一个节点
		}
		curr = curr.Next
		i++
	}

	return dummyHead.Next
}

func revertNode(head*ListNode)*ListNode{
	curr :=head
	var pre *ListNode
	for curr!=nil{
		pre,curr,curr.Next = curr,curr.Next,pre
	}
	return pre
}
func revertNode1(head*ListNode)*ListNode{
	  var pre *ListNode
	  curr := head
	  for curr!=nil{
		  tmp :=curr.Next
		  curr.Next = pre
		  pre = curr
	  	  curr= tmp
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
    // head = revertNode(head)
     head = reverseBetween(head,1,2)
     for head!=nil{
     	fmt.Println(head.Val)
     	head = head.Next
	 }
}
