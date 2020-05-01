package main

import (
	"fmt"
	"math"
	"sort"
)


func threeSumClosest(nums []int, target int) int {
	res :=nums[0]+nums[1]+nums[2]
	sort.Ints(nums)
     for i:=0;i<len(nums)-2;i++{
     	l :=i+1
     	r:=len(nums)-1
     	for l<r{
     		tmp:=nums[i]+nums[l]+nums[r]
			if tmp>target {
				r--
			}else if tmp<target {
				l++
			}else {
				return target
			}
			if int(math.Abs(float64(tmp-target)))<int(math.Abs(float64(res-target))) {
				res = tmp
			}
		}
	 }
	 return res
}
func main() {
	fmt.Println(threeSumClosest([]int{-1,2,1,-4},1))
	fmt.Println(threeSumClosest([]int{0,0,0},1))
}
