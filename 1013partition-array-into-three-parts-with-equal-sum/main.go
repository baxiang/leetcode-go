package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {
	var sum int
    for _,v:=range A{
		sum +=v
	}
	if sum%3!=0{
		return false
	}
	avg :=sum/3

	currSum :=0
	count :=0
	for _,v:=range A{
		currSum +=v
		if currSum==avg{
			count++
			currSum =0
		}
	}
	if count>=3{
		return true
	}
	return false
}


func main() {
	//fmt.Println(canThreePartsEqualSum([]int{0,2,1,-6,6,-7,9,1,2,0,1}))
	//fmt.Println(canThreePartsEqualSum([]int{0,2,1,-6,6,7,9,-1,2,0,1}))
	//fmt.Println(canThreePartsEqualSum([]int{3,3,6,5,-2,2,5,1,-9,4}))
	fmt.Println(canThreePartsEqualSum([]int{10,-10,10,-10,10,-10,10,-10}))
}
