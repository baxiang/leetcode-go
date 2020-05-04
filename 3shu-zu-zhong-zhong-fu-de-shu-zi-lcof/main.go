package main

import "fmt"

func findRepeatNumber(nums []int) int {
	m := make([]int,len(nums))
    for _,v:=range nums{
    	if m[v]>0{
    		return v
		}else{
			m[v]++
		}
	}
	return -1
}
func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}
