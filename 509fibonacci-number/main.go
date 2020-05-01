package main

import "fmt"

func fib1(N int) int {
    if N<2{
    	return N
	}
	return fib1(N-1)+fib1(N-2)
}

func fib2(N int) int {
	if N<2{
		return N
	}
	var v1 =0
	var v2 =1
	for i:=2;i<=N;i++{
	 	tmp := v1+v2
		v1 = v2
		v2 = tmp
	}
	return v2
}
func fib3(N int) int {
	if N<2{
		return N
	}
	dp := make([]int,N+1)
	dp[0] = 0
	dp[1]= 1
	for i:=2;i<=N;i++{
		dp[i] = dp[i-1]+dp[i-2]
	}
	return dp[N]
}
func fib(N int) int {
	if N<2{
		return N
	}
	pre := 0
	curr := 1
	for i:=2;i<=N;i++{
		sum :=pre+curr
		pre = curr
		curr = sum
	}
	return curr
}
func main() {
	fmt.Println(fib(5))
}
