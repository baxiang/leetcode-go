package main

import "fmt"

func isHappy(n int) bool {
	  n = digitResult(n)
	  m :=make(map[int]struct{},0)
	  for n!=1{
	  	if _,ok:=m[n];!ok{
			m[n]= struct{}{}

			n =digitResult(n)
			fmt.Println(n)
		}else {
			return false
		}
	  }
	  return true
}
func digitResult(n int) int {
	var res int
	for n>0{
		r :=n%10
		res +=r*r
		n =n/10
	}
	return res
}

func main() {
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(19))
}
