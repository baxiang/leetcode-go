package main

import "fmt"

type Node struct {
	key int
	value int
	pre *Node
	next *Node
}

type LRUCache struct {
     m map[int] *Node
     head *Node
     tail *Node
	 capacity int
}


func Constructor(capacity int) LRUCache {
    cache := LRUCache{
		  m:        make(map[int]*Node),
		  head:     &Node{},
		  tail:     &Node{},
		  capacity: capacity,
	  }
	  cache.head.next = cache.tail
	  cache.tail.pre = cache.head
	  return cache
}
// 头部添加元素
func (this *LRUCache)AddHead(n *Node){
	 tmp := this.head.next
	 this.head.next = n
	 n.next = tmp
	 n.pre = this.head
	 tmp.pre = n
}
func (this *LRUCache)RemoveNode(n *Node){
	pre := n.pre
	next := n.next
	pre.next = next
	next.pre = pre
	n.pre = nil
	n.next = nil
}


func (this *LRUCache) Get(key int) int {
	node ,ok:= this.m[key]
	if ok {
		this.RemoveNode(node)
		this.AddHead(node)
		return node.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	node ,ok:= this.m[key]
	if ok {// 存在
		if node.value!= value {
			node.value = value
		}
		this.RemoveNode(node)
		this.AddHead(node)
	}else {
		n := &Node{
			key:key,
			value: value,
		}
		if this.capacity==len(this.m) {// 删除
			delete(this.m,this.tail.pre.key)
			this.RemoveNode(this.tail.pre)
		}
		this.AddHead(n)
		this.m[key] = n
	}
}
func main() {
	cache := Constructor(2)
	cache.Put(1,1)
	cache.Put(2,2)
	fmt.Println(cache.Get(1))//1
	cache.Put(3,3)
	fmt.Println(cache.Get(2))//-1
	cache.Put(4,4)
	fmt.Println(cache.Get(1))//-1
	fmt.Println(cache.Get(3))//3
	fmt.Println(cache.Get(4))//4
}
