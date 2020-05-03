package main

import (
	"fmt"
	"math"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
     var res int
	for i:=0;i<len(arr1);i++{
		isDistanceValue :=true
		for j:=0;j<len(arr2);j++{
			if int(math.Abs(float64(arr1[i]-arr2[j])))<=d {
				isDistanceValue = false
				break
			}
		}
		if isDistanceValue{
			res++
		}
	}
	return res
}


func main() {
	fmt.Println(findTheDistanceValue([]int{4,5,8},[]int{10,9,1,8},2))
	fmt.Println(findTheDistanceValue([]int{1,4,2,3},[]int{-4,-3,6,10,20,30},3))
	fmt.Println(findTheDistanceValue([]int{12,1,100,3},[]int{-5,-2,10,-3,7},6))


}
