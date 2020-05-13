package main

import "fmt"

func Abs(a, b int) int {
	r := a - b
	if r < 0 {
		return -r
	}
	return r
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if Abs(nums[j], nums[i]) <= t && Abs(j, i) <= k {
				return true
			}
		}
	}
	return false
}
func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))

}
