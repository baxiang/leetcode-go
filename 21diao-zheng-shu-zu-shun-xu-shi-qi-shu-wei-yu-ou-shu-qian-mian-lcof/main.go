package main

import "fmt"

func exchange(nums []int) []int {
     l:=0
     r :=len(nums)-1
     for l<r{
		if nums[l]%2==0&&nums[r]%2!=0{
			nums[l],nums[r]=nums[r],nums[l]
		}
		 for l<len(nums)&&nums[l]%2!=0{
			 l++
		 }
		 for r>=0&&nums[r]%2==0{
			 r--
		 }

	 }
	 return nums
}

func main() {
	fmt.Println(exchange([]int{1,3,5}))
	fmt.Println(exchange([]int{1,2,3,4}))
	fmt.Println(exchange([]int{1,2,3,4,5,6,7}))
}
