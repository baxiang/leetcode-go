package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
    sort.Ints(coins)
    var res = 0
    var list []int
	for i:=len(coins)-1;i>=0;i--{
		for amount>=coins[i]{
			amount-=coins[i]
			list = append(list,coins[i])
			res++
		}
		if amount==0{
			break
		}
	}
	fmt.Println(list)
	if amount>0{
		fmt.Println(amount)
		return -1
	}
	return res
}

func main() {
	//fmt.Println(coinChange([]int{1,2,5},11))
	//fmt.Println(coinChange([]int{2},3))
	fmt.Println(coinChange([]int{186,419,83,408},6249))
}
