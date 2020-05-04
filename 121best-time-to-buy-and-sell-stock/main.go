package main

import "fmt"

func maxProfit(prices []int) int {
	maxVal := 0
	if len(prices) == 0 {
		return maxVal
	}
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if minPrice > prices[i] { // 表示当前买入价格最低
			minPrice = prices[i]
		} else { //表示当前价格适合卖出
			v := prices[i] - minPrice
			maxVal = max(maxVal, v)
		}
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{}))
	fmt.Println(maxProfit([]int{2, 4, 1}))
}
