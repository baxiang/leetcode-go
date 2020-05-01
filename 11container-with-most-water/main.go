package main

import "fmt"

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func maxArea(height []int) int {
     left :=0
     right := len(height)-1
     area :=0
     h :=0
     for left<right{
     	 w :=right-left
		 if height[left]<=height[right] {
			 h = height[left]
			 left++
		 }else {
			 h = height[right]
			 right--
		 }
		 area = max(area,h*w)
	 }
     return area
}



func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
