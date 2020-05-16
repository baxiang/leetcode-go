package main

import "fmt"

func nthUglyNumber(n int) int {
	count := 0
	number :=1
	uglyNumber :=1
    for count<n{
    	if isUglyNum(number){
    		count++
			uglyNumber = number
		}
		number++
	}
	return uglyNumber
}

func isUglyNum(n int)bool{
	for n%2==0{
		n=n/2
	}
	for n%3==0{
		n=n/3
	}
	for n%5==0{
		n=n/5
	}
	return n<=1
}

//面试题49. 丑数
func main() {
    //fmt.Println(isUglyNum(2))
	fmt.Println(nthUglyNumber(10))
	fmt.Println(nthUglyNumber(11))
	//for i:=1;i<=20;i++{
	//	if isUglyNum(i){
	//		fmt.Println(i)
	//	}
	//}
}
