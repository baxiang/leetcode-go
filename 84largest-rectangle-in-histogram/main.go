package main

import "fmt"

func largestRectangleArea(heights []int) int {
	var res int
	for i := 0; i < len(heights); i++ {
		var min = heights[i]
		for j := i; j < len(heights); j++ {
			if min > heights[j] {
				min = heights[j]
			}
			width := j - i + 1
			area := min * width
			if area > res {
				res = area
			}
		}
	}
	return res
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
