package main

import (
	"fmt"
)

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}
func trap(height []int)int{
	l :=len(height)
	res :=0
	// 起始位置是第一个，而不是0
	// 结束位置是倒数第二个，而不是倒数第一个
	for i:=1;i<l-1;i++{
		leftMax :=0
		for k:=0;k<=i;k++{
			leftMax =max(height[k],leftMax)
		}
		rightMax :=0
		for j:=i;j<l;j++{
			rightMax =max(height[j],rightMax)
		}
		minVal := min(leftMax,rightMax)
        water := minVal-height[i]
        res =res+water
	}
	return res
}





func trap1(height []int) int {
	if len(height)==0{
		return 0
	}
	left :=0
	right :=len(height)-1
	leftMax :=0
	rightMax :=0
	res :=0
	for left< right{
		if height[left]<=height[right]{
			if height[left]>=leftMax{
				leftMax = height[left]
			}else {
				res +=leftMax-height[left]
			}
			left++
		}else {
			if height[right]>=rightMax{
				rightMax = height[right]
			}else {
				res +=rightMax-height[right]
			}
			right--
		}
	}
	return res
}

func trap2(height []int)int{
	l :=len(height)
	res :=0
	// 起始位置是第一个，而不是0
	// 结束位置是倒数第二个，而不是倒数第一个
	for i:=1;i<l-1;i++{
		leftMax :=height[i]
		for k:=i-1;k>=0;k--{
			if leftMax<height[k]{
				leftMax = height[k]
			}
		}
		if leftMax<=height[i]{
			continue
		}
		rightMax :=height[i]
		for j:=i+1;j<l;j++{
			if rightMax<height[j]{
				rightMax = height[j]
			}
		}
		if rightMax<=height[i]{
			continue
		}
		minVal := min(leftMax,rightMax)
		water := minVal-height[i]
		res =res+water
	}
	return res
}

// 动态规划
func trap3(height []int)int{
	if len(height)==0{
		return 0
	}
	size :=len(height)
	res :=0
	leftDP :=make([]int,size)
	rightDP :=make([]int,size)
	leftDP[0]= height[0]
	rightDP[size-1] = height[size-1]

	for i:=1;i<size;i++{
		leftDP[i] = max(leftDP[i-1],height[i])
	}
	for i:=size-2;i>=0;i--{
		rightDP[i] = max(rightDP[i+1],height[i])
	}
	for i:=1;i<size-1;i++{
		minHeight := min(leftDP[i],rightDP[i])
		waterVal :=minHeight-height[i]
		res +=waterVal
	}
	return res
}

func main() {
	fmt.Println(trap3([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}
