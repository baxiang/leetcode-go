package main

import "fmt"

func subtractProductAndSum(n int) int {
     var sum int =0
     var res int =1
     for n!=0{
     	r := n%10
     	n =n/10
     	sum +=r
     	res *=r
	 }
     return res-sum
}
func main() {
	fmt.Println(subtractProductAndSum(234))
	fmt.Println(subtractProductAndSum(4421))
}
