package main

import "fmt"

func missingNumber1(nums []int) int {
    for i:=1;i<len(nums);i++{
        if nums[i]-nums[i-1]!=1{
        	return nums[i-1]+1
		}
	}
	if nums[0]!=0{
		return 0
	}
	return nums[len(nums)-1]+1
}
func missingNumber(nums []int) int {
	left :=0
	right :=len(nums)
	for left<right{
		mid := left+(right-left)>>1
		if nums[mid]==mid{
			left = mid+1
		}else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(missingNumber([]int{0}))//1
	fmt.Println(missingNumber([]int{1}))//0
	fmt.Println(missingNumber([]int{0,1}))//2
	fmt.Println(missingNumber([]int{0,1,3}))//2
	fmt.Println(missingNumber([]int{0,1,2,3,4,5,6,7,9}))//8
}
