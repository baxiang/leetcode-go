package main

import "fmt"

func fib(n int) int {
	if n<2 {
		return n
	}
	dp := make([]int,n+1)
	dp[0] =0
	dp[1] =1
	for i:=2;i<=n;i++{
		dp[i] = (dp[i-1]+dp[i-2])%1000000007
	}
   return dp[n]
}


func main() {
	fmt.Println(fib(5))
	fmt.Println(fib(43))
}
