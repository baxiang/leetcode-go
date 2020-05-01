package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	m1 :=make(map[int]int)
	for _,v:=range nums1{
		m1[v]++
	}
	m2 :=make(map[int]int)
	for _,v:=range nums2{
		m2[v]++
	}
	var res[]int
	for k,v :=range m1{
		r := m2[k]
		r = min(v,r)
		for i:=0;i<r;i++{
			res =append(res,k)
		}
	}
	return res
}

func min(a,b int)int  {
	if a<b{
		return a
	}
	return b
}

func main() {
	fmt.Println(intersect([]int{1,2,2,1},[]int{2,2}))
	fmt.Println(intersect([]int{4,9,5},[]int{9,4,9,8,4}))
}
