package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{},0)
	for _,v:=range nums{
		if _,ok:=m[v];ok{
			return true
		}
		m[v]= struct{}{}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1,2,3,1}))
	fmt.Println(containsDuplicate([]int{1,2,3,4}))
}
