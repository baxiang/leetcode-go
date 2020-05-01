package main

import (
	"container/heap"
	"fmt"
)


type Node struct {
	Val int
	Num int
}

type NodeHeap []*Node

func(h NodeHeap)Len()int{
	return len(h)
}
func(h NodeHeap)Swap(i,j int){
	h[i],h[j] = h[j],h[i]
}
func(h NodeHeap)Less(i,j int)bool{
	return h[i].Num<h[j].Num
}
func(h *NodeHeap)Push(x interface{}){
	*h = append(*h,x.(*Node))
}
func (h *NodeHeap)Pop()interface{}  {
	x :=(*h)[len(*h)-1]
	*h =(*h)[:len(*h)-1]
	return x
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}

func topKFrequent(nums []int, k int) []int {
	m :=make(map[int]int)
	for _,v :=range nums{
		m[v]++
	}
	topK :=min(k,len(m))
	size :=0
	h :=&NodeHeap{}
	for k,v :=range m{
		n := &Node{
			Val: k,
			Num: v,
		}
		if size<topK{
			heap.Push(h,n)
			size++
		}else {
			if v >(*h)[0].Num{
				heap.Pop(h)
				heap.Push(h,n)
			}
		}
	}
	var res []int
	for i:=0;i<topK;i++{
		res = append(res,heap.Pop(h).(*Node).Val)
	}
	return res
}
func main() {
	fmt.Println(topKFrequent([]int{1,1,1,2,2,3},2)) //1,2
	fmt.Println(topKFrequent([]int{1},1))
}
