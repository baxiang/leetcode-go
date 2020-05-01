package main

import "fmt"

func numberOfSteps (num int) int {
	var res int
	for num !=0{
		r := num%2
		if r ==0{//偶数
			num=num/2
		}else {
			num= num-1
		}
		res++
	}
	return res
}


func main() {
	fmt.Println(numberOfSteps(14))
}
