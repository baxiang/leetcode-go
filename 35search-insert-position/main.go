package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left:=0
	right:= len(nums)-1
	for left<=right{
		mid := (left+(right-right)) >>1
		if nums[mid]==target {
			return mid
		} else if nums[mid] <target{
			left=mid+1
		}else {
			right= mid-1
		}
	}
	return left
}

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6},5))
	fmt.Println(searchInsert([]int{1,3,5,6},7))
	fmt.Println(searchInsert([]int{1},1))
}
