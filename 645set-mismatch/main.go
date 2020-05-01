package main

import "fmt"

func findErrorNums(nums []int) []int {
   m :=make(map[int]int)
   var n,j int
   for i,v:=range nums{
	   if k,ok:=m[v];ok {
		   n = v
		   j = nums[i]+(i-k)
	   }else {
	   	m[v]=i
	   }
   }
   return []int{n,j}
}

func main() {
	fmt.Println(findErrorNums([]int{1,2,2,4}))
}
