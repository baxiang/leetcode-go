package main

import "fmt"

func findSubsequences(nums []int) [][]int {
    for i:=0;i<len(nums);i++{
    	 l :=[]int {nums[i]}
    	for j:=i+1;j<len(nums);j++{
			if nums[j]>=nums[i] {

				l = append(l,nums[j])
			}else {
				l =[]int {nums[i]}
			}
			if len(l)>=2{
				fmt.Println(l)
			}

		}
	}
	return nil
}

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}
