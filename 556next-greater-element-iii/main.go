package main

import (
	"fmt"
	"math"
	"sort"
)

func nextGreaterElement(n int) int {
     // 第一步需要拆解成数组
	  var nums []int
	  for n>0{
	  	  r := n %10
	  	  n = n/10
	  	  nums = append(nums,r)
	  }
	  idx :=-1
      //[2,3,4,1]
	  for i:=0;i<len(nums)-1;i++{
		  if nums[i]>nums[i+1] {
		  	  idx = i+1
			  break
		  }
	  }
	 if idx==-1 {
		return -1
	 }
	//[2,3,4,1]
	 for i:=0;i<idx;i++{
	 	 if nums[i]>nums[idx]{
	 	 	nums[i],nums[idx] =nums[idx],nums[i]
	 	 	break
		 }
	 }
     // [1,3,4,2]
	 sort.Slice(nums[:idx], func(i, j int) bool {
		   return  nums[i]>nums[j]
	 })
	 //[4,3,1,2]
	 res :=0
	 for i:=len(nums)-1;i>=0;i--{
	 	  res =res*10+nums[i]
		 if res>math.MaxInt32 {
			 return -1
		 }
	  }
	 return res
}
func main() {
	fmt.Println(nextGreaterElement(1432))
	fmt.Println(nextGreaterElement(1342))
	fmt.Println(nextGreaterElement(123))
	fmt.Println(nextGreaterElement(12))
	fmt.Println(nextGreaterElement(21))
}
