package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
    // 排序
	var res [][]int

	sort.Ints(nums)
	for i:=0;i<len(nums);i+=1{
		// 排序 可以去重 重复值没有必须在执行一遍了，也就做到了去重
		if i>0&&nums[i]==nums[i-1] {
			continue
		}
		l := i+1
		r := len(nums)-1
		for l<r{// 移动左右坐标点
			tmp := nums[i]+nums[l]+nums[r]
			if tmp ==0 {
				res = append(res,[]int{nums[i],nums[l],nums[r]})
				l+=1
				r-=1
				// 数字相同 去重
				for l<r&&nums[l]==nums[l-1] {
					l+=1
				}
				for l<r&&nums[r]==nums[r+1] {
					r-=1
				}
			}else if tmp>0 {// 如果和大说明需要取更小的值
				r-=1
			}else {// 如果和小于0 说明需要取更大的值
				l+=1
			}
		}
	}
	return res
}
func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
	fmt.Println(threeSum([]int{0,0,0,0}))
	fmt.Println(threeSum([]int{-2,0,3,-1,4,0,3,4,1,1,1,-3,-5,4,0}))
}
