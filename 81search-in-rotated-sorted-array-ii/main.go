package main

import "fmt"

func search(nums []int, target int) bool {
	left :=0
	right :=len(nums)-1
	for left<=right{
		mid :=left+(right-left)>>1
		if nums[mid]==target{
			return true
		}else if nums[mid]==nums[left]{
			left++
		} else if nums[left]<nums[mid]{
			if nums[left]<=target&&target<nums[mid]{
				right =mid-1
			}else {
				left = mid+1
			}
		}else {
			if nums[mid]<target&&target<=nums[right]{
				left = mid+1
			}else {
				right = mid-1
			}
		}
	}
	return false
}

func main() {
	fmt.Println(search([]int{2,5,6,0,0,1,2},0))
	fmt.Println(search([]int{2,5,6,0,0,1,2},3))
	fmt.Println(search([]int{3,1},1))
}
