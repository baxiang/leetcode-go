package main

import "fmt"

func majorityElement(nums []int) int {
    m := make(map[int]int)
    max :=0
    num :=nums[0]
    for _,v:=range nums{
    	m[v]++
		if max<m[v] {
			max = m[v]
			num= v
		}
	}
	fmt.Println(m)
	return num
}

func main() {
	//fmt.Println(majorityElement([]int{3,2,3}))
	//fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
	fmt.Println(majorityElement([]int{3,3,4}))
}
