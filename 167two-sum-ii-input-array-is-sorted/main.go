package main

import "fmt"

func twoSum(numbers []int, target int) []int {
     l :=0
     r :=len(numbers)-1
     for l<=r{
     	res := numbers[l]+numbers[r]
     	if res==target{
     		return []int{l+1,r+1}
		}else if res>target  {
			r--
		}else {
			l++
		}
	 }
	 return nil
}


func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15},9))
}
