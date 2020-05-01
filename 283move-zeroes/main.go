package main

import (
	"fmt"
)

func moveZeroes(nums []int)  {
	var l1,l2 []int
      for i:=0;i<len(nums);i++{
		  if nums[i]==0 {
			  l1 = append(l1,nums[i])
		  }else {
			  l2 = append(l2,nums[i])
		  }
	  }
	  l2 = append(l2,l1...)
	  for i:=0;i<len(nums);i++{
	  	nums[i] = l2[i]
	  }
}
func main() {
	 s :=[]int{0,1,0,3,12}
	 moveZeroes(s)
	 fmt.Println(s)
}
