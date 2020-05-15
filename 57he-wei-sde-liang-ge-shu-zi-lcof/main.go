package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]struct{})
    for _,v :=range nums{
    	if _,ok:=m[target-v];ok{
    		return []int{target-v,v}
		}else {
			m[v]= struct{}{}
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	left :=0
	right :=len(nums)-1
	for left<right{
		 sum := nums[left]+nums[right]
		if sum==target{
			return []int{nums[left],nums[right]}
		}else if sum>target{
			right--
		}else {
			left++
		}
	}
	return nil
}



func main() {
	fmt.Println(twoSum([]int{2,7,11,15},9))
	fmt.Println(twoSum([]int{16,16,18,24,30,32},48))
}
