package main

import "fmt"

func findNumbers(nums []int) int {
	var res int
    for _,v :=range nums{
    	if findDigit(v){
    		res++
		}
	}
	return res
}

func findDigit(num int)bool{
	var n int
	for num!=0{
		num=num/10
		n++
	}
	if n%2==0{
		return true
	}
	return false
}

func main() {
	fmt.Println(findNumbers([]int{12,345,2,6,7896}))
	fmt.Println(findNumbers([]int{555,901,482,1771}))
}
