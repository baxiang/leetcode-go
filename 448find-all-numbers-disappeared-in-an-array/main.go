package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	r :=make([]int,len(nums))
	for _,v:=range nums {
		r[v-1] =r[v-1]+1
	}
	var res []int
	for i,v:=range r{
		if v==0 {
			res =append(res,i+1)
		}
	}
	return res
}
func main() {
	fmt.Println(findDisappearedNumbers([]int{4,3,2,7,8,2,3,1}))
}
