package main

import (
	"fmt"
)

// 判断需要移动的次数
func maxSlidingWindow(nums []int, k int) []int {
	//len-k+1
	p :=len(nums)-k+1
	slide:=make([]int,p)
	for i:=0;i<p;i++{
		max := nums[i]
		for j:=i+1;j<i+k;j++{
			if max<nums[j] {
				max =nums[j]
			}
		}
		slide[i] = max
	}
	return slide
}

type Dequeue struct {
	arr []int
}
func NewDequeue (size int)*Dequeue{
	return &Dequeue{arr:make([]int,size)}
}
//
func(d *Dequeue)AddFirst(v int){
	d.arr = append([]int{v},d.arr...)
}

func(d *Dequeue)RemoveFirst(){
	if len(d.arr)==0 {
		return
	}
	d.arr =d.arr[1:]
}
func(d *Dequeue)isEmpty()bool{
	return len(d.arr)==0
}

func(d *Dequeue)GetFirst()int{
	if len(d.arr)==0 {
		return 0
	}
	return d.arr[0]
}
func(d *Dequeue)AddLast(v int){
	d.arr = append(d.arr,v)
}
func(d *Dequeue)RemoveLast(){
	if len(d.arr)==0 {
		return
	}
	d.arr =d.arr[:len(d.arr)-1]
}
func(d *Dequeue)GetLast()int{
	if len(d.arr)==0 {
		return 0
	}
	return d.arr[len(d.arr)-1]
}
func maxSlidingWindow1(nums []int, k int )[]int{
	numLens :=len(nums)
	if numLens<=0{
		return nil
	}
	queues := NewDequeue(numLens)
	var result []int
	for i,v :=range nums{
		for !queues.isEmpty()&&v>=queues.GetLast(){
			queues.RemoveLast()
		}
		queues.AddLast(nums[i])
		if i>=k&&nums[i-k]==queues.GetFirst() {// 窗口最大值已经不在超过了窗口里 需要移开
			queues.RemoveFirst()
		}
		if i>=k-1 {//需要大于等于窗口才把最大元素放入
			// 放入结果数组
			result = append(result,queues.GetFirst())
		}

	}
   return result
}

func maxSlidingWindow2(nums []int, k int )[]int{
	numLens :=len(nums)
	if numLens<=0{
		return nil
	}
	var window []int //存储坐标点
	var result []int // 窗口最大值
	for i :=range nums{
		if len(window)>0&&window[0]==i-k {//需要保证最大值的索引值在当前窗口访问内
			window = window[1:]
		}
		for len(window)>0&&(nums[window[len(window)-1]]<=nums[i] ) {// 最后值小于所求当前值
			window = window[:len(window)-1] // 保证最左边最大值
		}
		window =append(window,i)
		if i>=k-1{
			result = append(result,nums[window[0]])
		}
	}
	return result
}


// 滑动窗口最大值
func main() {
	// [3 3 5 5 6 7]
	fmt.Println(maxSlidingWindow2([]int{1,3,-1,-3,5,3,6,7},3))
	// [1,-1]
	//fmt.Println(maxSlidingWindow2([]int{1,-1},1))
}
