package main

import "fmt"

func numWays(n int) int {
	if n<=1{
		return 1
	}
   dp := make([]int ,n+1)
   dp[1]=1
   dp[2]=2
   for i:=3;i<=n;i++{
   	  dp[i]= (dp[i-1]+dp[i-2])%1000000007
	}
	return dp[n]
}

func main() {
	fmt.Println(numWays(0))
	fmt.Println(numWays(2))
	fmt.Println(numWays(7))
	fmt.Println(numWays(92))
}
