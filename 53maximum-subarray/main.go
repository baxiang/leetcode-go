package main

import "fmt"



func Max(a,b int)int{
	if a>b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
    dp := make([]int ,len(nums))
    dp[0]=nums[0]
    res := dp[0]
    for i:=1;i<len(nums);i++{
		dp[i]=Max(nums[i],dp[i-1]+nums[i])
		if res<dp[i] {
			res = dp[i]
		}
	}
	return res
}
//53. 最大子序和 https://leetcode-cn.com/problems/maximum-subarray/
func main() {
	fmt.Println(maxSubArray([]int{-2}))
	fmt.Println(maxSubArray([]int{-2,1}))
	fmt.Println(maxSubArray([]int{-2,1,-3}))
	fmt.Println(maxSubArray([]int{2,1,-3,4,-1,2,1,-5,4}))
}
