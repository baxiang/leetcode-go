package main

import (
	"fmt"
	"math"
)

func Max(a [] int)int{
    var	max = math.MinInt64
	for _,v :=range a{
		if v>max {
			max = v
		}
	}
	return max
}
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums)==0 {
		return nil
	}
	l :=0
	res :=make([]int,0)
	for i:=0;i<=len(nums)-k;i++{
		res =append(res,Max(nums[l:l+k]))
		l++
	}
	return res
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3))
}
