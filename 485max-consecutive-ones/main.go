package main

import (
	"math"
	"fmt"
)

func findMaxConsecutiveOnes1(nums []int) int {
	start :=0
	end :=0
	max :=0
	for end<len(nums){
		for start<len(nums)&&nums[start]==0{
			start++
		}
		end =start
		for end<len(nums)&&nums[end]==1{
			end++
		}
		max =int(math.Max(float64(max),float64(end-start)))
		start = end
	}
	return max
}


func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func findMaxConsecutiveOnes2(nums []int) int {
	currCount :=0
	maxVal :=0
	for i:=0;i<len(nums);i++{
		if nums[i]==1{
			currCount ++
		}else {
			maxVal = max(maxVal,currCount)
			currCount =0
		}
	}
	if currCount!=0{// 如果最后几位都是1
		maxVal = max(maxVal,currCount)
	}
	return maxVal
}

func findMaxConsecutiveOnes(nums []int) int {
	left :=0
	right :=0
	maxVal :=0
	for right<len(nums){
		if nums[right]==0{
			maxVal = max(maxVal,right-left)
			left = right+1
		}
		right++
	}
	if left<right{
		maxVal = max(maxVal,right-left)
	}

	return maxVal
}

func main() {
	fmt.Println(findMaxConsecutiveOnes1([]int{1}))
	fmt.Println(findMaxConsecutiveOnes([]int{0}))
	fmt.Println(findMaxConsecutiveOnes1([]int{0,0,1}))
	fmt.Println(findMaxConsecutiveOnes([]int{1,1,0,1,1,1}))
}
