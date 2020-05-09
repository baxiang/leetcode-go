package main

type MaxQueue struct {
    List []int
    Max  int
}

func Constructor() MaxQueue {
    return MaxQueue{List: make([]int,0),Max: -1}
}

func (this *MaxQueue) Max_value() int {
    return this.Max
}

func (this *MaxQueue) Push_back(value int)  {
     if this.Max<value{
     	this.Max = value
	 }
	 this.List = append(this.List,value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.List)==0{
		return -1
	}
	currVal := this.List[0]
	this.List = this.List[1:]
	if currVal==this.Max{
		this.Max = -1
		for i:=0;i<len(this.List);i++{
			if this.Max<this.List[i]{
				this.Max=this.List[i]
			}
		}
	}
	return currVal
}


//面试题59 - II. 队列的最大值
func main() {
	
}
