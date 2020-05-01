package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	 i := len(a)-1
	 j := len(b)-1
	 var av ,bv ,carry int
	 var res string
	for i>=0||j>=0{
		if i>=0{
			av =int(a[i]-'0')
			i--
		}else {
			av =0
		}
		if j>=0{
			bv =int(b[j]-'0')
			j--
		}else {
			bv =0
		}
		r := av+bv+carry
		carry =r/2
		r = r%2
		res =fmt.Sprintf("%d%s",r,res)
	}
	if carry==1 {
		res =fmt.Sprintf("%d%s",carry,res)
	}
	return res
}


func main() {
	fmt.Println(addBinary("11","1"))//100
	fmt.Println(addBinary("1010","1011"))//10101
	fmt.Println(addBinary("1111","1111"))//10101
}
