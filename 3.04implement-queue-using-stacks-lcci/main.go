package main
type MyQueue struct {
     list []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
   return MyQueue{list:make([]int,0)}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.list = append(this.list,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	v :=this.list[0]
	this.list = this.list[1:]
	return v
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.list[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.list)==0
}
func main() {
	
}
