package main

import (
	"fmt"
)

func singleNumbers(nums []int) []int {
	var res int
   for _,v:=range nums{
   	   res= res^v
   	   fmt.Println(v,res)
   }
   return nil
}

func main() {
	fmt.Println(singleNumbers([]int{4,1,4,6}))
	fmt.Println(singleNumbers([]int{1,2,10,4,1,4,3,3}))
}
