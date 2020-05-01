package main

import "fmt"

type MinStack struct {
    s []Node
}
type Node struct {
	d int //data
	m int //current min
}


/** initialize your data structure here. */
func Constructor() MinStack {
   return MinStack{}
}


func (this *MinStack) Push(x int)  {
	d := Node{d:x,m:x}
	if len(this.s)>0 &&this.s[len(this.s)-1].m< x{
		d.m=this.s[len(this.s)-1].m
	}
     this.s = append(this.s,d)

}


func (this *MinStack) Pop()  {
	this.s = this.s[:len(this.s)-1]
}


func (this *MinStack) Top() int {
	return this.s[len(this.s)-1].d
}


func (this *MinStack) GetMin() int {
	return this.s[len(this.s)-1].m
}

func main() {
	m := MinStack{}
	m.Push(-10)
	m.Push(14)
	fmt.Println(m.GetMin())
	fmt.Println(m.GetMin())
	m.Push(-20)
	fmt.Println(m.GetMin())
	fmt.Println(m.GetMin())
	fmt.Println(m.Top())
	fmt.Println(m.GetMin())
	m.Pop()
	m.Push(10)
	m.Push(-7)
	fmt.Println(m.GetMin())
}
