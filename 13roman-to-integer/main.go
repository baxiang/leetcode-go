package main

import "fmt"

func romanToInt(s string) int {
   m :=map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
   temp :=0
   var result int
   for i:=len(s)-1;i>=0;i--{
   	   num := m[string(s[i])]
	   //'IV' 4
	   if num<temp&&(num==1||num==10||num==100) {
		   result -=num
	   }else {
	   	   result +=num
	   }
	   temp = num
   }
   return result
}

func main() {
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("MCMXCIV"))
}
