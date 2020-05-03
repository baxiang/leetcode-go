package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	sum :=0
	for _,v:=range nums{
		sum+=v
	}
	currSum :=0
	var res []int
	for i:=len(nums)-1;currSum<=sum/2;i--{
		currSum +=nums[i]
		res = append(res,nums[i])
	}
	return res
}


func main() {
	fmt.Println(minSubsequence([]int{4,4,7,6,7}))
}
