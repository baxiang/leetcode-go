package main

import "fmt"

func addDigits(num int) int {
    for num>=10{
    	res:=0
    	for num>0{
    		res +=num%10
    		num=num/10
		}
		num= res
	}
	return num
}


func main() {
	fmt.Println(addDigits(38))
	fmt.Println(addDigits(59))
	fmt.Println(addDigits(199))
}
