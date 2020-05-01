package main


type ListNode struct {
	Val int
	Next *ListNode
}


func reverseList(head *ListNode) *ListNode {
     var pre *ListNode
     curr := head
     for curr!= nil{
     	pre,curr,curr.Next = curr,curr.Next,pre
	 }
	 return pre
}
func reverseList1(head *ListNode) *ListNode {
	dummyHead :=&ListNode{
		Val:  -1,
	}
	for head!= nil{
		pre := dummyHead.Next
		next :=head.Next
		dummyHead.Next = head
		head.Next = pre
		head = next
	}
	return dummyHead.Next
}


func main() {
	
}
