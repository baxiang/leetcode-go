package main

import (
	"fmt"
	"math"
)

func reverse1(x int) int {
   var stack []int
   for x!=0{
   	   r :=x%10
   	   stack = append(stack,r)
   	   x =x/10
   }
   var r int
   index :=0
   for len(stack)>0{
   	   top :=stack[len(stack)-1]
   	   stack = stack[:len(stack)-1]
	   r = int(math.Pow10(index))*top+r
   	   index++
   }
   // 边界
	if !( -(1<<31) <= r && r <= (1<<31)-1) {
		return 0
	}
   return r
}

func reverse(x int) int {

	var y =0
	for x!=0{
		y =y*10+x%10
		if !( -(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x =x/10
	}
	return y
}
func main() {
	fmt.Println(reverse(-1234))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
}
