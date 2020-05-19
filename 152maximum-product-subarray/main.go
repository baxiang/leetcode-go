package main

import (
	"fmt"
	"math"
	"sort"
)

func maxProduct(nums []int) int {
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		var res = 1
		for j := i; j < len(nums); j++ {
			res = res * nums[j]
			if max < res {
				max = res
			}
		}
	}
	return max
}

func maxProduct2(nums []int) int {
	maxDP := make([]int, len(nums))
	minDP := make([]int, len(nums))
	maxDP[0] = nums[0]
	minDP[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		maxDP[i], minDP[i] = numCompare(minDP[i-1]*nums[i], maxDP[i-1]*nums[i], nums[i])
		if max < maxDP[i] {
			max = maxDP[i]
		}
	}
	return max
}

//三个数字排列 获取最大值和最小值
func numCompare(a, b, c int) (max, min int) {
	s := []int{a, b, c,}
	sort.Ints(s)
	return s[2], s[0]
}

func main() {
	fmt.Println(maxProduct2([]int{-2, 0, -1}))
	fmt.Println(maxProduct2([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct2([]int{-2, 3, -4}))
}
