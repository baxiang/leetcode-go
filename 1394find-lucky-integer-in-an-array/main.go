package main

import (
	"fmt"
)

func findLucky(arr []int) int {
      m :=make(map[int]int)
      for _,v:=range arr{
      	   m[v]++
	  }
	  max := -1
	  for k,v:=range m{
	  	  if k==v&&max<k{
	  	  	 max =k
		  }
	  }
	  return max
}

func main() {
	fmt.Println(findLucky([]int{2,2,3,4}))
	fmt.Println(findLucky([]int{1,2,2,3,3,3}))
	fmt.Println(findLucky([]int{2,2,2,3,3}))
	fmt.Println(findLucky([]int{5}))
	fmt.Println(findLucky([]int{7,7,7,7,7,7,7}))
}
