package main

import "fmt"

func search(nums []int, target int) int {
   left :=0
   right :=len(nums)-1
   for left<=right{
   	 mid :=left+(right-left)>>1
   	 if nums[mid]==target{
   	 	return mid
	 }else if nums[left]<=nums[mid]{//中间值递增序列
		 if nums[left]<=target&&target<nums[mid]{
			 right=mid-1
		 }else {
		 	left= mid+1
		 }
	 }else {// 中间值右边递增序列
		 if nums[mid]<target&&target<=nums[right] {
			 left = mid+1
		 }else {
		 	right = mid-1
		 }
	 }
   }
   return -1
}

func main() {
	fmt.Println(search([]int{4,5,6,7,0,1,2},0))
}
