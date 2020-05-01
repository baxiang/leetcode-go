package main

import "fmt"

func removeDuplicates(nums []int) int {
    slow :=2
    for fast :=2;fast<len(nums);fast++{
    	if nums[fast]!=nums[slow-2]{
			nums[slow]=nums[fast]
    		slow++
		}
	}
	return slow
}
func main() {
	fmt.Println(removeDuplicates([]int{1,1,1,2,2,3}))
}
