package main

import "fmt"

func sumZero(n int) []int {
	if n==1{
		return []int{0}
	}
	var sum =n
	res :=[]int{n}
    for i:=n-2;i>0;i--{
   	    res=append(res,i)
   	    sum +=i
    }
	res=append(res,0-sum)
	return res
}

func main() {
	fmt.Println(sumZero(5))
	fmt.Println(sumZero(3))
	fmt.Println(sumZero(2))
	fmt.Println(sumZero(1))
}
