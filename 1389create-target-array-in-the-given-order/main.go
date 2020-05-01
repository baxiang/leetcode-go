package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(index))
	sortedNum := 0 //当前已经排列总数
	for i := 0; i < len(nums); i++ {
		v := index[i]
		if i == v {
			res[i] = nums[i]
		} else {
			for n := sortedNum; n > v; n-- { //需要把数字往后移动
				res[n] = res[n-1]
			}
			res[v] = nums[i]
		}
		sortedNum++
	}
	return res
}

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
	fmt.Println(createTargetArray([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}))
	fmt.Println(createTargetArray([]int{1}, []int{0}))
}
