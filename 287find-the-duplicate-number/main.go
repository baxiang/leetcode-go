package main

import "fmt"

func findDuplicate(nums []int) int {
      flag :=make([]int,len(nums))
      for i:=0;i<len(nums);i++{
      	 if flag[nums[i]]==1{
           return nums[i]
		 }
		 flag[nums[i]]=1
	  }
	  return -1
}
func main() {
	fmt.Println(findDuplicate([]int{1,3,4,2,2}))
	fmt.Println(findDuplicate([]int{3,1,3,4,2}))
}
