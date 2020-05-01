package main

import (
	"container/heap"
	"fmt"

)

type IntList []int

func (l IntList)Len()int{
	return len(l)
}
func(l IntList)Swap(i,j int){
	l[i],l[j] = l[j],l[i]
}

func(l IntList) Less(i, j int) bool{
	return l[i]<l[j]
}
func (l *IntList)Push(x interface{}){
	*l = append(*l,x.(int))
}

func (l *IntList)Pop()interface{}{
	x := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return x
}



func thirdMax(nums []int) int {
	//sort.Ints(nums)
	//list := make(IntList,0)
	m :=make(map[int]struct{})
	//heap.Init(&list)
	var list []int
	for i :=0;i<len(nums);i++{
		if _,ok:=m[nums[i]];!ok{
			list =append(list,nums[i])
			m[nums[i]]= struct {}{}
		}

	}
	heapSort(list)
	if len(list)<3 {
		return list[0]
	}
    return list[2]
}

func heapSort(input []int){
	inputLen := len(input)
	if inputLen == 0 {
		return
	}
	for i:=0; i<inputLen; i++{
		minAjust(input[i:])
	}
}

func minAjust(input []int){
	inputLen := len(input)
	if inputLen <= 1{
		return
	}
	for i:= inputLen/2 -1; i>=0; i--{
		if (2*i+1 <= inputLen-1) && (input[i] <= input[2*i+1]){
			input[i], input[2*i+1] = input[2*i+1], input[i]
		}
		if (2*i+2<= inputLen-1) && (input[i] <= input[2*i+2]){
			input[i], input[2*i+2] = input[2*i+2], input[i]
		}
	}
}

func sift(list []int, left, right int) {
	fIdx := left
	sIdx := fIdx*2 + 1
	// 构造大根堆
	for sIdx <= right {
		//判断左孩子:sIdx 右孩子:sIdx+1
		if sIdx < right && list[sIdx] < list[sIdx+1] {
			sIdx++
		}
		// 最终和大的儿子比较
		if list[fIdx] < list[sIdx] {
			// 交换
			list[fIdx], list[sIdx] = list[sIdx], list[fIdx]
			// 交换后重新检查被修改的子节点为大根堆
			fIdx = sIdx
			sIdx = 2*fIdx + 1
		} else {
			// 已经是大根堆
			break
		}
	}
}

func HeapSort(list []int) {
	length := len(list)
	//建立初始堆
	sift(list, 0, length-1)
	for idx := length / 2; idx >= 0; idx-- {
		// 从后往前调整
		sift(list, idx, length-1)
	}
	// 将大根堆的根节点和堆最后一个元素交换，重新对前面的 length - 1 调整堆
	for idx := length - 1; idx >= 1; idx-- {
		list[0], list[idx] = list[idx], list[0]
		sift(list, 0, idx-1)
	}
	// 结果就是逆序输出大根堆
}

func buidHeap(a []int, n int) {

	//heapify from the last parent node
	for i := n / 2; i >= 1; i-- {
		heapifyUpToDown(a, i, n)
	}

}

//sort by ascend, a index begin from 1, has n elements
func sort(a []int, n int) {
	buidHeap(a, n)

	k := n
	for k >= 1 {
		swap(a, 1, k)
		heapifyUpToDown(a, 1, k-1)
		k--
	}
}

//heapify from up to down , node index = top
func heapifyUpToDown(a []int, top int, count int) {

	for i := top; i <= count/2; {

		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		swap(a, i, maxIndex)
		i = maxIndex
	}

}

//swap two elements
func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}


func main() {
	//fmt.Println(thirdMax([]int{3,2,1}))//1
	//fmt.Println(thirdMax([]int{1,2}))//2
	//fmt.Println(thirdMax([]int{2,2,3,1}))//1
	//
	//fmt.Println(thirdMax([]int{1,2,2,5,3,5}))//2
	//fmt.Println(thirdMax([]int{1,-2147483648,2}))//-2147483648
	//fmt.Println(thirdMax([]int{1,2,-2147483648}))//-2147483648

	nums1 :=[]int{1,-2147483648,2}
	var list IntList
	for i :=0;i<len(nums1);i++{
		list = append(list,nums1[i])
	}
	heap.Init(&list)
	fmt.Println(list)
	//HeapSort(nums1)
	nums2 :=[]int{1,2,-2147483648}
	//HeapSort(nums2)
	//fmt.Println(nums2)
	var list2 IntList
	for i :=0;i<len(nums2);i++{
		list2 = append(list2,nums2[i])
	}
	heap.Init(&list2)
	fmt.Println(list2)

	//
	//nums3 :=[]int{3,28,1}
	//HeapSort(nums3)
	//fmt.Println(nums3)
}
