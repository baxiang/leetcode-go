package main

import (
	"fmt"
)

func findPairs(nums []int, k int) int {
	numsHas := make(map[int]bool)
	diffHas := make(map[int]bool)
   for _,num :=range nums{
	   if  numsHas[num-k]{
		   diffHas[num - k] = true
	   }
	   if numsHas[num+k] {
		   diffHas[num] = true
	   }
	   numsHas[num] = true
   }
   return len(diffHas)
}

func main() {
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5},2))
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5},1))
	fmt.Println(findPairs([]int{1,2,3,4,5},3))
}
