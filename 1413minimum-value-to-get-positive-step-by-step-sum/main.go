package main

import (
	"fmt"
	"math"
)
// 需要一个数 和所有数组相加和都保持不小于1
func minStartValue(nums []int) int {
	 var sum int
	 var res =1
     for _,v:=range nums{
     	sum +=v
     	if sum<1{
     		res = int(math.Max(float64(res),float64(1-sum)))
		}
	 }
	 return res
}

func main() {
	fmt.Println(minStartValue([]int{-3,2,-3,4,2}))
	fmt.Println(minStartValue([]int{1,2}))
	fmt.Println(minStartValue([]int{1,-2,-3}))
}
