package main

import "fmt"

func maxProfit(prices []int) int {
	var res int
	for i:=0;i<len(prices)-1;i++{
        p := prices[i+1]-prices[i]
        if p>0{
        	res+=p
		}
	}
	return res
}


func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
}
