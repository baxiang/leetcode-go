package main

import (
	"fmt"
	"sort"
	"strings"
)

func minNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
		return compareNumber(nums[i],nums[j])
	})
    var res strings.Builder
    for i:=0;i<len(nums);i++{
    	res.WriteString(fmt.Sprintf("%d",nums[i]))
	}
	return res.String()
}

func compareNumber(a,b int)bool{
	str1 := fmt.Sprintf("%d%d",a,b)
	str2 := fmt.Sprintf("%d%d",b,a)
	if str1<str2{
		return true
	}
	return false
}




func main() {
	fmt.Println(minNumber([]int{8,2}))
	fmt.Println(minNumber([]int{22,21}))
	fmt.Println(minNumber([]int{2,21}))//212
	fmt.Println(minNumber([]int{3,30,34}))
	fmt.Println(minNumber([]int{3,30,34,5,9}))
	//1399439856075703697382478249389609
	//1399439856075703697382478249389609
	fmt.Println(minNumber([]int{8247,824}))
    fmt.Println(minNumber([]int{824,938,1399,5607,6973,5703,9609,4398,8247}))
}
