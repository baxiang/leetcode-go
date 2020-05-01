package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	key int
	value int
}

type LRUCache struct {
    m map[int] *list.Element
    l  *list.List
	capacity  int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
		m:        make(map[int]*list.Element),
		l:        list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v,ok:=this.m[key];ok {
		this.l.MoveToFront(v)
		r := v.Value.(*Node)
		return r.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	if v,ok:=this.m[key];ok {
		 n := v.Value.(*Node)
		 n.value = value
		this.l.MoveToFront(v)
	}else {
		if this.capacity==len(this.m) {
			node := this.l.Back()
			n := node.Value.(*Node)
			delete(this.m,n.key)
			this.l.Remove(node)
		}
		n :=&Node{
			key:   key,
			value: value,
		}
		e := this.l.PushFront(n)
		this.m[key] = e
	}
}


func main() {
	 c := Constructor(2)
	 p :=&c
	 p.Put(2,1)
	 p.Put(2,2)
	 fmt.Println(p.Get(2))
	 p.Put(1,1)
	 p.Put(4,1)
	 fmt.Println(p.Get(2))

}
