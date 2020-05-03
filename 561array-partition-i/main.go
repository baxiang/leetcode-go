package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    var res int
    for i:=0;i<len(nums);i=i+2{
    	 res +=nums[i]
	}
	return res
}


func main() {
	fmt.Println(arrayPairSum([]int{1,4,3,2}))
}
