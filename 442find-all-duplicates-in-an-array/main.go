package main

import (
	"fmt"
	"sort"
)

func findDuplicates(nums []int) []int {
    m := make(map[int]int)
    for _,v:=range nums{
		if i,ok:=m[v];ok {
			m[v]=i+1
		}else {
			m[v]=1
		}
	}
	var res []int
    for k,v:=range m{
		if v==2 {
			res = append(res,k)
		}
	}
	sort.Ints(res)
	return res
}
func main() {
	fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1}))
}
