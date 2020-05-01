package main

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent1(nums []int) []int {
	var res []int

	for i:=0;i<len(nums);i++{
		var list []int
		list =append(list,nums[:i]...)
		list =append(list,nums[i+1:]...)
		fmt.Println(list)
		r := findSmallerNumbers(list,nums[i])
		res = append(res,r)
	}
	return res
}


func smallerNumbersThanCurrent(nums []int) []int {
	copyNum :=make([]int,len(nums))
	copy(copyNum,nums)
	sort.Ints(copyNum)
	m := make(map[int]int)
	for i,v :=range copyNum{
		if _,ok :=m[v];!ok{
			m[v]=i
		}
	}
	for i,v :=range nums{
		copyNum[i] = m[v]
	}
	return copyNum
}

func findSmallerNumbers(nums []int,t int)int{
	var res int
	for i:=0;i<len(nums);i++{
		if nums[i]<t{
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8,1,2,2,3}))
}