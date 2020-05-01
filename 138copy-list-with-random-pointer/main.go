package main

import "unsafe"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}


func copyRandomList(head *Node) *Node {
    curr := head
    m := make(map[uintptr]*Node)
    for curr!=nil{
    	node :=&Node{Val:curr.Val}
    	m[uintptr(unsafe.Pointer(curr))]=node
    	curr = curr.Next
	}
	curr = head
	for curr!=nil{
		node := m[uintptr(unsafe.Pointer(curr))]
		node.Next = m[uintptr(unsafe.Pointer(curr.Next))]
		node.Random	=m[uintptr(unsafe.Pointer(curr.Random))]
		curr = curr.Next
	}
	return m[uintptr(unsafe.Pointer(head))]
}


func main() {
	
}
