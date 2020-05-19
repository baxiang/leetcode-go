package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome1(head *ListNode) bool {
	var str string
	var revertStr string
	for head != nil {
		str = fmt.Sprintf("%s%d", str, head.Val)
		revertStr = fmt.Sprintf("%d%s", head.Val, revertStr)
		head = head.Next
	}
	return str == revertStr
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 查找中间节点
	midNode := findMidNode(head)
	// 翻转以中间节点的链表
	currTail := revertLinkedNode(midNode)
	// 判断头节点和尾节点值是否一直
	for head != nil && currTail != nil {
		if head.Val != currTail.Val {
			return false
		}
		head = head.Next
		currTail = currTail.Next
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}
	var res []*ListNode
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	left := 0
	right := len(res) - 1
	for left < right {
		if res[left].Val != res[right].Val {
			return false
		}
		left++
		right--
	}
	return true
}

func findMidNode(head *ListNode) *ListNode {

	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func revertLinkedNode(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	for head != nil {
		next := head.Next
		proNode := dummyHead.Next
		dummyHead.Next = head
		head.Next = proNode
		head = next
	}
	return dummyHead.Next
}

func revertLinkedNode1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	node := revertLinkedNode1(next)
	next.Next = head
	head.Next = nil
	return node
}

func main() {
	t1 := ListNode{Val: 1}
	t2 := ListNode{Val: 2}
	t1.Next = &t2
	t3 := ListNode{Val: 1}
	t2.Next = &t3
	t4 := ListNode{Val: 4}
	t3.Next = &t4
	t5 := ListNode{Val: 5}
	t4.Next = &t5

	fmt.Println(isPalindrome(&t1))

}
