package main

import "fmt"

func rotate(nums []int, k int)  {
     tmp :=make([]int,len(nums))
     copy(tmp,nums)
     for i,v:=range tmp{
     	 index :=(i+k)%len(tmp)
     	 nums[index]= v
	 }
}


func main() {
	s :=[]int{1,2,3,4,5,6,7}
	rotate(s,3)
	fmt.Println(s)
	s1 :=[]int{-1,-100,3,99}
	rotate(s1,2)
	fmt.Println(s1)
}
