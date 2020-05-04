package main

import (
	"fmt"
	"sort"
)

func majorityElement1(nums []int) int {
     m :=make(map[int]int)
     t :=len(nums)/2
     for _,v:=range nums{
     	if m[v]>=t{
     		return v
		}
		m[v]++
	 }
	 return -1
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}


func main() {
	fmt.Println(majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
