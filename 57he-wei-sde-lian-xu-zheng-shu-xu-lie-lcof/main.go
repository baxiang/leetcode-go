package main

import "fmt"

func findContinuousSequence(target int) [][]int {

	left :=1
	right :=2
	var res [][]int
	for left<right{
		sum :=(left+right)*(right-left+1)/2
		if sum==target{
			var list []int
			for i:=left;i<=right;i++{
				list = append(list,i)
			}
			res = append(res,list)
			left++
		}else if sum<target{
			right++
		}else {
			left++
		}
	}
	return res
}

func main() {
	fmt.Println(findContinuousSequence(9))
	fmt.Println(findContinuousSequence(15))
	fmt.Println(findContinuousSequence(1))
}
