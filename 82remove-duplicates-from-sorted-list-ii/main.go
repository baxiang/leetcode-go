package main

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead:=&ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dummyHead
	curr := head
	count :=0
	for curr!=nil&&curr.Next!= nil{
		// 有重复一直循环
		for curr.Next!=nil&&curr.Next.Val==curr.Val{
			count ++
			curr = curr.Next
		}
		if count>0 {
			pre.Next= curr.Next
			curr = curr.Next
			count = 0
		}else {
			pre = curr
			curr = curr.Next
		}
	}
	return dummyHead.Next
}

func main() {
	
}
