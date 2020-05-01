package main

import "fmt"

func climbStairs(n int) int {
	if n==1 {
		return 1
	}
	dp :=make([]int,n+1)
	dp[1]=1
	dp[2]=2
	for i:=3;i<=n;i++{
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}


//70. 爬楼梯 https://leetcode-cn.com/problems/climbing-stairs/
func main() {
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(6))
	fmt.Println(climbStairs(7))
}
