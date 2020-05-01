package main

import (
	"fmt"
	"sort"
)

// 三个数字排列 获取最大值和最小值
func search(a,b,c int)(big,small int){
	i :=[]int{a,b,c}
	sort.Ints(i)
	return i[2],i[0]
}
func maxProduct(nums []int) int {
   dpMax :=make([]int,len(nums))
	dpMin:= make([]int,len(nums))
	dpMax[0] =nums[0]
	dpMin[0] = nums[0]
   max :=nums[0]
   for i:=1;i<len(nums);i++{
	 dpMax[i],dpMin[i] = search(dpMax[i-1]*nums[i],dpMin[i-1]*nums[i],nums[i])
	   if dpMax[i]>max {
		   max = dpMax[i]
	   }
	}
	return max
}
func main() {
	fmt.Println(maxProduct([]int{-2,0,-1}))
	fmt.Println(maxProduct([]int{2,3,-2,4}))
	fmt.Println(maxProduct([]int{-2,3,-4}))
}
