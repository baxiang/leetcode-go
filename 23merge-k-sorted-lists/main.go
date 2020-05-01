package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type NodeSlice []*ListNode

func (p NodeSlice) Len() int           { return len(p) }
func (p NodeSlice) Less(i, j int) bool { return p[i].Val < p[j].Val }
func (p NodeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *NodeSlice) Push(x interface{}) {
	node := x.(*ListNode)
	*p = append(*p, node)
}
func (p *NodeSlice) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}
func mergeKLists(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	curr := dummyHead
	nodeList := make(NodeSlice, 0)
	for i := range lists {
		if lists[i] != nil {
			nodeList = append(nodeList, lists[i])
		}
	}
	heap.Init(&nodeList)
	for len(nodeList) > 0 {
		item := heap.Pop(&nodeList).(*ListNode)// 最小元素
		next := item.Next
		item.Next = nil
		curr.Next = item
		curr = item
		if next != nil {
			heap.Push(&nodeList, next)
		}
	}
	return dummyHead.Next
}

func mergeKLists1(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	nodeList := make(NodeSlice, 0)
	for i := range lists {
		h :=lists[i]
		for h!= nil {
			nodeList = append(nodeList, h)
			h=h.Next
		}
	}
	sort.Sort(nodeList)
	curr := dummyHead
	for _,v :=range nodeList{
		curr.Next = v
		curr = v
	}
	return dummyHead.Next
}



func main() {
    a1 :=ListNode{Val:  1,}
	b1 :=&ListNode{Val: 4,}
	a1.Next = b1
	c1 :=ListNode{Val:  5,}
	b1.Next =&c1
	a2 :=ListNode{Val:  1,}
	b2 :=ListNode{Val:  3,}
	a2.Next = &b2
	c2 :=ListNode{Val:  4,}
	b2.Next =&c2
	a3 :=ListNode{Val:  2,}
	b3 :=ListNode{Val:  6,}
	a3.Next =&b3
	r := mergeKLists([]*ListNode{&a1,&a2,&a3,nil})
	for r!= nil{
		fmt.Println(r.Val)
		r = r.Next
	}
}
