package main

import "fmt"

func Max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
func maxProfit(prices []int) int {
    //dp:= make([][]int,len(prices))
    //for i:=range prices{
    //	dp[i] = make([]int,len(prices))
	//}
	max :=0
	for i:=0;i<len(prices);i++{
		for j:=i+1;j<len(prices);j++{
			if prices[j]>prices[i]  {
				max = Max(max,prices[j]-prices[i])
			}
		}
	}
   return max
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}
