package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
     var res [][]int
     sort.Ints(nums)
     for i:=0;i<len(nums)-3;i++{
		 if i>0&&nums[i]==nums[i-1] {
			continue
		 }
     	for j:=i+1;j<len(nums)-2;j++{
			if j>i+1&&nums[j]==nums[j-1] {
				continue
			}
     		l :=j+1
     		r :=len(nums)-1
			for l<r {

				temp :=nums[i]+nums[j]+nums[l]+nums[r]
				if temp<target {
					l++
				}else if temp>target {
					r--
				}else {
					res = append(res,[]int{nums[i],nums[j],nums[l],nums[r]})
					l++
					r--
					for l>j+1&&nums[l]==nums[l-1] {
						l++
					}
					for l<len(nums)-2&&nums[r]==nums[r+1]{
						r--
					}
				}

			}
		}
	 }
	 return res
}


func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2},0))
	fmt.Println(fourSum([]int{-3,-2,-1,0,0,1,2,3},0))
	fmt.Println(fourSum([]int{-1,0,-5,-2,-2,-4,0,1,-2},-9))

}
