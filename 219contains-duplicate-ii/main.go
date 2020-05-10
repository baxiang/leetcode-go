package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if r, ok := m[v]; ok {
			if i-r <= k {
				return true
			}
		}
		m[v] = i
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1},3))
	fmt.Println(containsNearbyDuplicate([]int{1,0,1,1},1))
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1,2,3},2))
}
