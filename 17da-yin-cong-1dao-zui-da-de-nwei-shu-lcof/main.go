package main

import (
	"math"
	"fmt"
)

func printNumbers(n int) []int {
    var res []int
    var max int
    for n>0{
    	max =max*10+9
    	n--
	}
    for i:=1;i<=max;i++{
    	res = append(res,i)
	}
	return res
}

func printNumbers1(n int) []int {
	var res []int
	for i:=1;i<int(math.Pow10(n));i++{
		res = append(res,i)
	}
	return res
}

func main() {
	//fmt.Println(printNumbers1(1))
	fmt.Println(printNumbers1(2))
	//fmt.Println(printNumbers(3))
}
