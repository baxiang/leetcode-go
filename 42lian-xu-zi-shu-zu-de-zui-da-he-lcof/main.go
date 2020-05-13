package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	max :=math.MinInt64
	for i:=0;i<len(nums);i++{
		 for j:=i+1;j<=len(nums);j++{
			r := SubArraySum(nums[i:j])
			if max<r {
				 max = r
			}
		 }
	 }
	 return max
}
func SubArraySum(nums[]int)int{
	var res int
	for i:=0;i<len(nums);i++{
		res +=nums[i]
	}
	return res
}


func maxSubArray1(nums []int) int {
    if len(nums)==0{
    	return 0
	}
	max :=nums[0]
	res :=make([][]int,len(nums))
	for i:=0;i<len(nums);i++{
		res[i]= make([]int,len(nums)+1)
		res[i][i]=nums[i]
		if max<res[i][i] {
			max = res[i][i]
		}
		for j:=i+1;j<len(nums);j++{
			res[i][j]= res[i][j-1]+nums[j]
			if max<res[i][j] {
				max = res[i][j]
			}
		}
	}
	return max
}





func main() {
	//fmt.Println(maxSubArray1([]int{-2,1,-3,4,-1,2,1,-5,4}))
	//fmt.Println(maxSubArray1([]int{1}))
	fmt.Println(maxSubArray1([]int{-2,1}))
	//fmt.Println(maxSubArray([]int{-2,1}))
}
