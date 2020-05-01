package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	 res :=make([]int,len(nums))
	 flag :=false
     for i:=0;i<len(nums);i++{
		 flag = false
		 for j:=1;j<len(nums);j++{
		 	 idx :=(j+i)%len(nums)
			 if nums[i]<nums[idx] {
			 	 res[i] = nums[idx]
			 	 flag = true
				break
			 }
		 }
		 if !flag {
			 res[i] = -1
		 }
	 }
	 return res
}

func main() {
	fmt.Println(nextGreaterElements([]int{1,2,1}))
	fmt.Println(nextGreaterElements([]int{5,4,3,2,1}))
}
