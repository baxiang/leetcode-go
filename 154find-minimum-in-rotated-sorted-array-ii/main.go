package main

import "math"

func findMin(nums []int) int {
     left :=0
     right :=len(nums)-1
     for left<right{
     	mid :=left+(right-left)>>1
     	if nums[mid]>nums[right]{
     		left = mid+1
		}else {
			right= right-1
		}
	 }
	 return nums[left]
}

func main() {
	math.Pow()
}
