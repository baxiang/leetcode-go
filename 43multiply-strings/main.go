package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
     var n1 int=0
     for i:=0;i<len(num1);i++{
     	 n :=int (num1[i]-'0')
     	 n1 =n1*10+n
	 }
	 var res string
	for i:=len(num2)-1;i>=0;i--{
		//n := int (num2[i]-'0')

	}
	return fmt.Sprintf("%f",res)
}

func addString(num1 string, num2 string) string {

	n1 :=len(num1)-1
	n2 :=len(num2)-2
	carry :=0
	var res string
	for n1>=0||n2>=0{
		a :=0
		if n1>=0 {
			a = int(num1[n1]-'0')
			n1--
		}
		b :=0
		if n2>=0 {
			b = int(num2[n2]-'0')
			n2--
		}
		r := a+b+carry
		carry =r/10
		r = r%10
		res = fmt.Sprintf("%d%s",r,res)
	}
	if carry==1 {
		res = fmt.Sprintf("%d%s",carry,res)
	}
	return res
}


func main() {
	//fmt.Println(multiply("2","3"))
	//fmt.Println(multiply("123","456"))
	fmt.Println(multiply("498828660196","840477629533"))
}
