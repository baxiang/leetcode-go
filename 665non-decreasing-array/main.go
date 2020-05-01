package main

import "fmt"

func checkPossibility(nums []int) bool {
	count := 0
	for i:=0;i<len(nums)-1;i++{
		if nums[i]<nums[i+1] {
			count++
		}
		if count>1 {
			return false
		}
	}
	return  count==1
}

func main() {
	fmt.Println(checkPossibility([]int{4,2,3}))
	fmt.Println(checkPossibility([]int{4,2,1}))
	fmt.Println(checkPossibility([]int{-1,4,2,3}))
}
