package main

type CQueue struct {
   list []int
}


func Constructor() CQueue {
     return CQueue{list:make([]int,0)}
}


func (this *CQueue) AppendTail(value int)  {
      this.list = append(this.list,value)
}


func (this *CQueue) DeleteHead() int {
    if len(this.list)==0{
    	return -1
	}
	v :=this.list[0]
	this.list = this.list[1:]
	return v
}


//面试题09. 用两个栈实现队列
func main() {
	
}
