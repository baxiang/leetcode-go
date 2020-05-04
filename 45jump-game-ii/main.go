package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {

	//记录能够到达的最远距离
	maxPosition :=0
	steps :=0
	//记录上一步 step 跳的最远距离也就是边界
	end :=0
	//        //这里有个小细节，因为是起跳的时候就 + 1 了，如果最后一次跳跃刚好到达了最后一个位置，那么遍历到最后一个位置的时候就会再次起跳，这是不允许的，因此不能遍历最后一个位置
	for i:=0;i<len(nums)-1;i++{
		//更新最大距离
		maxPosition = int(math.Max(float64(maxPosition),float64(nums[i]+i)))
		/*
		   我们什么时候需要步数 + 1？
		   当到达上一步的最远距离的时候，那么意味着我们需要进行一次新的起跳，那么步数 + 1
		   并且更新最远距离
		*/
		////第一次起跳 或 到达跳跃的边界
		if i==end{
			end =maxPosition
			steps++
		}
	}
	return steps
}


func main() {
	fmt.Println(jump([]int{1,2}))
	fmt.Println(jump([]int{1,2,3}))
	fmt.Println(jump([]int{2,3,1,1,4}))
}
