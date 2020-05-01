package main

import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    m := make(map[int]int)
	for i,v:=range nums2{
		m[v]= i
	}
     var res = make([]int,len(nums1))
	 for i:=0;i<len(nums1);i++{
	 	isExist :=false
		for j:=m[nums1[i]];j<len(nums2);j++{
			if nums1[i]<nums2[j] {
				res[i] = nums2[j]
				isExist = true
				break
			}
		}
		 if !isExist {
			 res[i]= -1
		 }
	}
	return res
}
func main() {
	//fmt.Println(nextGreaterElement([]int{4,1,2},[]int{1,3,4,2}))
	fmt.Println(nextGreaterElement([]int{3,1,5,7,9,2,6},[]int{1,2,3,5,6,7,9,11}))
}
