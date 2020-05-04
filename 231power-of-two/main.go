package main

import "fmt"

func isPowerOfTwo1(n int) bool {
	if n<=0{
		return false
	}
    for n>1{
    	if n%2!=0{
    		return false
		}
		n=n/2
	}
	return true
}


func isPowerOfTwo(n int) bool {
	return n>0&&n&(n-1)==0
}


func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(218))
}
