package main

import "fmt"

func tribonacci(n int) int {
	if n<2 {
		return n
	}
    dp:=make([]int,n+1)
    dp[0]=0
    dp[1]=1
    dp[2]=1
    for i:=3;i<=n;i++{
    	dp[i] = dp[i-3]+dp[i-2]+dp[i-1]
	}
	return dp[n]
}

func main() {
	fmt.Println(tribonacci(3))
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(5))
	fmt.Println(tribonacci(25))
}
