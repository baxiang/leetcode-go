package main

import (
	"fmt"
)
//func getLeastNumbers0(arr []int, k int) []int {
//	sort.Sort(Array(arr))
//	return arr[:k]
//}
//
//type Array []int
//
//func (a Array) Len() int           { return len(a) }
//func (a Array) Less(i, j int) bool { return a[i] < a[j] }
//func (a Array) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }


//func getLeastNumbers(arr []int, k int) []int {
//    sort.Ints(arr)
//    return arr[:k]
//}
func getLeastNumbers1(arr []int, k int) []int {
	//sort.Ints(arr)
	 sort(arr,0,len(arr)-1)
	return arr
}

//func quickSort1(arr []int){
//
//}
func sort(arr [] int ,left,right int){
	if left<right {
         p := partition(arr,left,right)
         sort(arr,left,p-1)
         sort(arr,p+1,right)
	}
}

func partition(arr [] int ,start,end int)int{
	i,j,temp :=start,end,arr[start]
	for i<j{
		// 从左往右找比基准值大的数字
		for arr[j]>temp&&i<j{
			j--
		}
		// 从右往左找比基准值小的数字
		for arr[i]<=temp&&i<j{
			i++
		}
		arr[i],arr[j] = arr[j],arr[i]

	}
	arr[i],arr[start] = arr[start],arr[i]
	return i
}

func swap(arr [] int ,left,right int){
	temp := arr[left]
	arr[left]=arr[right]
	arr[right]= temp
}

func quickSort(arr []int)[]int{
	if len(arr) <2{// 一个元素不需要排序 直接返回
		return arr
	}
	pivot :=arr[0]
	var left []int //小于基准值
	var mid []int // 等于基准值
	var right []int // 大于基准值
	for _,v :=range arr{
		if v<pivot {
			left = append(left,v)
		}else if v>pivot{
			right = append(right,v)
		}else {
			mid = append(mid,v)
		}
	}
	 l :=quickSort(left)
	 l = append(l,mid...)
	 r :=quickSort(right)
	return append(l,r...)
}


func buildHeap(arr []int){
	// 最后一个节点
	last := len(arr)-1
	start :=(last-1)/2
    for start>=0{
       heapfix(arr,start)
       start--
	}
}

func heapfix(arr []int,i int){
	len := len(arr)
	if i>len {
		return
	}
	left :=i<<2+1
	right :=i<<2+2
	max := i
	if left<len&&arr[left]>arr[max]{
		max = left
	}
	if right<len&&arr[right]>arr[max]{
		max = right
	}
	if max!=i {
		arr[max],arr[i] = arr[i],arr[max]
		heapfix(arr,max)
	}
}

func quickSort1(arr []int)[]int{
	if len(arr)<2 {
		return arr
	}
	var min []int
	var max []int
	var mid []int
	m :=arr[0]
	for i:=0;i<len(arr);i++{
		if m<arr[i] {// big data
			max = append(max,arr[i])
		}
		if m>arr[i] { // small
			min = append(min,arr[i])
		}
		if m==arr[i] {// equal
			mid = append(mid,arr[i])
		}
	}
	r := quickSort1(min)
	r = append(r,mid...)
	r =  append(r,quickSort1(max)...)
   return r
}


//面试题40. 最小的k个数 https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
func main() {
	//fmt.Println(getLeastNumbers1([]int{3,2,1,4,7,9},4))
	l :=[]int{10,3,2,1,0,4,7,9}
	buildHeap(l)
	fmt.Println(quickSort1(l))
}
