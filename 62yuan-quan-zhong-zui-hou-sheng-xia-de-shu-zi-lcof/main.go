package main

import (
	"container/list"
	"fmt"
)


func lastRemaining(n int, m int) int {
	l:= list.New()
	for i:=0;i<n;i++{
		l.PushBack(i)
	}
	i :=0
	for l.Len()>1{
		for i=0;i<m;i++{
			e :=l.Front()
			if i !=m-1 {
				l.MoveToBack(e)
			}else {
				l.Remove(e)
			}
		}
	}
	v := l.Front().Value.(int)
	return v
}
func main() {
	fmt.Println(lastRemaining(5,3))
	fmt.Println(lastRemaining(10,17))
	fmt.Println(lastRemaining(70866,116922))
}
