package main

import "fmt"
// leetcode 会超时
func reversePairs(nums []int) int {
	num :=0
    for i:=0;i<len(nums);i++{
    	for j:=i+1;j<len(nums);j++{
			if nums[i]>nums[j] {
				num++
			}
		}
	}
	return num
}
func main() {
	fmt.Println(reversePairs([]int{7,5,6,4}))
}
