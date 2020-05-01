package main

import "fmt"

func mySqrt(x int) int {
	if x==0 ||x==1{
		return x
	}
	left :=0
	right :=x
	for left<=right{
		mid := (left+right)>>1
		r := mid*mid
		if r==x {
			return mid
		}else if r>x  {
			right =mid-1
		}else {
			left = mid+1
		}
	}
	return left-1
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(6))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(10))
}
