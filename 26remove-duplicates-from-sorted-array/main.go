package main

import "fmt"

func removeDuplicates2(nums []int) int {
	 m := make(map[int]struct{})
	 index :=0
     for _,v:=range nums{
     	if _,ok:=m[v];!ok{
     		m[v] = struct{}{}
			nums[index]=v
     		index++
		 }
	 }
	 nums = nums[:index]
	 return len(m)
}
//快慢指针
func removeDuplicates1(nums []int) int {
	 slow := 0//慢指针
	 //快指针变 块指针便利数组
	for quick :=1; quick<len(nums) ;quick++  {
		if nums[slow]!=nums[quick] {
			slow++// 只有当前值和快指针不相同
			nums[slow]=nums[quick]
		}
	}
	return slow +1
}
func removeDuplicates(nums []int) int {
	for i:=len(nums)-1;i>0;i--{
		if nums[i]==nums[i-1]{
			nums = append(nums[:i],nums[i+1:]...)
		}
	}

	return len(nums)
}

func main() {
	l := []int{1,1,2}
	r := removeDuplicates(l)
	for i:=0;i<r;i++{
		fmt.Println(l[i])
	}
}
