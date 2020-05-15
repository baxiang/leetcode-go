package main

import "fmt"

func subarraySum1(nums []int, k int) int {
	var count =0
	for i :=0;i<len(nums);i++{
		sum :=0
		for j:=i;j<len(nums);j++{
			sum +=nums[j]
			if sum==k{
				count++
			}
		}
	}
	return count
}

func subarraySum(nums []int, k int) int {
	m := map[int]int{0:1}
	var sum int
	var count int
	for i :=0;i<len(nums);i++{
		sum +=nums[i]
		if v,ok:=m[sum-k];ok{
			count+=v
		}
		m[sum]++
	}
	return count
}



func main() {
	fmt.Println(subarraySum([]int{1,1,1},2))
	fmt.Println(subarraySum([]int{3,4,7,2,-3,1,4,2},7))
}
