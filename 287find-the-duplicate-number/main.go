package main

import (
	"fmt"
	"sort"
)

func findDuplicate1(nums []int) int {
      flag :=make([]int,len(nums))
      for i:=0;i<len(nums);i++{
      	 if flag[nums[i]]==1{
           return nums[i]
		 }
		 flag[nums[i]]=1
	  }
	  return -1
}
func findDuplicate2(nums []int) int {
   sort.Ints(nums)
   l :=0
   r :=len(nums)
   for l<r{
   	 mid := l+(r-l)>>1
   	 if nums[mid]>mid{
   	 	l =mid+1
	 }else {
	 	r =mid
	 }
   }
   return l
}
func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return 0
}


func main() {
	//1 2,2,3,4
	fmt.Println(findDuplicate([]int{1,3,4,2,2}))
	fmt.Println(findDuplicate([]int{3,1,3,4,2}))
}
