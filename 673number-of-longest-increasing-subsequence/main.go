package main

import (
	"fmt"
)

func findNumberOfLIS(nums []int) int {
     dp := make([]int,len(nums))
	 count := make([]int,len(nums))
	maxLength := 1
     for i :=0;i<len(nums);i++{
     	dp[i] = 1
     	count[i] = 1
     	for j:=0;j<i;j++{
			if nums[j]<nums[i] {
				if dp[j]+1>dp[i] {
					dp[i] = dp[j]+1
					count[i] = count[j]
				}else if dp[j]+1==dp[i] {
					count[i] += count[j]
				}
				
			}
		}
		 if maxLength < dp[i]{
			 maxLength = dp[i]
		 }
	 }
	var maxNum int
	for i := 0; i <= len(dp) - 1; i += 1{
		if dp[i] == maxLength{
			maxNum += count[i]
		}
	}
	return maxNum
}

func main() {
	fmt.Println(findNumberOfLIS([]int{1,3,5,4,7}))
	fmt.Println(findNumberOfLIS([]int{2,2,2,2,2}))
}
