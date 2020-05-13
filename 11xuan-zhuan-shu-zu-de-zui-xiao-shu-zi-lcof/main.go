package main

import "fmt"

func minArray(numbers []int) int {
    left :=0
    right :=len(numbers)-1
    for left<right{
    	mid :=left+(right-left)>>1
    	if numbers[mid]>numbers[right]{
    		left = mid+1
		}else if numbers[mid]<=numbers[right] {
			right = right-1
		}
	}
	return numbers[left]
}

func main() {
	fmt.Println(minArray([]int{5,1,2,3,4}))//1
	fmt.Println(minArray([]int{3,4,5,6,7,1,2}))//1
	fmt.Println(minArray([]int{2,3,4,5,6,1}))//1
	fmt.Println(minArray([]int{1,3,5}))
}
