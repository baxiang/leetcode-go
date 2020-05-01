package main

import "fmt"

func decompressRLElist(nums []int) []int {
	 var res []int
     for i:=0;i<len(nums);i=i+2{
     	 for j:=0;j<nums[i];j++{
     	 	res = append(res,nums[i+1])
		 }
	 }
	 return res
}


func main() {
	fmt.Println(decompressRLElist([]int{1,2,3,4}))
}
