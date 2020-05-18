package main

import "fmt"

func sumNums(n int) int {
    start :=1
    end :=n
	return ((start+end)*n)>>1
}


//面试题64. 求1+2+…+n
func main() {
	fmt.Println(sumNums(3))
	fmt.Println(sumNums(9))
}
