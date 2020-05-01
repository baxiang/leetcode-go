package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	mid := len(nums1) / 2
	if len(nums1)%2 != 0 {
		return float64(nums1[mid])
	}
	return float64(nums1[mid]+nums1[mid-1]) / 2.0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	mid := len(nums1) / 2
	if len(nums1)%2 != 0 {
		return float64(nums1[mid])
	}
	return float64(nums1[mid]+nums1[mid-1]) / 2.0
}




func main() {
	//fmt.Println(findMedianSortedArrays([]int{1,3},[]int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
