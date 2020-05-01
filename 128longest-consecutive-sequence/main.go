package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums)==0{
		return 0
	}
    sort.Ints(nums)
    count :=1
    max := count
    for i:=1;i<len(nums);i++{
    	if nums[i]==nums[i-1]{
			continue
		}
    	if nums[i-1]+1==nums[i]{
    		count++
		}else {
			if max<count{
				max = count
			}
			count=1
		}
	}
	if max<count{
		max = count
	}
	return max
}
func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0,-1}))
	fmt.Println(longestConsecutive([]int{1,2,0,1}))
}
