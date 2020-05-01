package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
     dp :=make([]int,len(nums))//定义规则 dp[i] 到i表达当前最大的升序
     maxLen := 0
     for i:=0;i<len(nums);i++{
     	 dp[i] =1
     	 for j:=0;j<i;j++{
			 if nums[j]<nums[i] { // 遍历i的子串 如果存在升序
				 dp[i] = max(dp[j]+1,dp[i])
			 }
		 }
		 maxLen = max(maxLen,dp[i])
	 }
	 return maxLen
}







//300. 最长上升子序列 https://leetcode-cn.com/problems/longest-increasing-subsequence/
func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
