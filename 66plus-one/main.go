package main

import "fmt"

func plusOne(digits []int) []int {
   for i := len(digits)-1;i>=0;i--{
	   digits[i] = (digits[i]+1)%10
	   if digits[i]!=0{// 没有产生进位 直接返回就行
	   	return digits
	   }
   }
   return append([]int{1},digits...)
}


func main() {
	fmt.Println(plusOne([]int{1,2,3}))
	fmt.Println(plusOne([]int{9,9,9,9}))
}
