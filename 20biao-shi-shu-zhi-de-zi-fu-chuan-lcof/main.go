package main

import (
	"fmt"
	"unicode"
)

func isNumber(s string) bool {
	point :=0
	for i:=0;i<len(s);i++{
		if (s[i]=='+'||s[i]=='-')&&i!=0{
			return false
		} else if s[i]=='.'&&i==0 {
			return false
		} else if s[i]=='e'&&(i+1>=len(s)||!unicode.IsDigit(rune(s[i+1]))){
			return false
		} else if s[i]=='e'&&(i-1<0||!unicode.IsDigit(rune(s[i-1]))){
			return false
		} else if s[i]=='.'&&point>0{
			return false
		}else if s[i]=='.'&&point==0{
			point++
		} else if !unicode.IsDigit(rune(s[i]))&&!(
			s[i]=='e'||s[i]=='+'||s[i]=='-'){
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isNumber("."))
	fmt.Println(isNumber("e9"))
	//fmt.Println(isNumber("0123"))
	//fmt.Println(isNumber("3.1415"))
	//fmt.Println(isNumber("+100"))
	//fmt.Println(isNumber("12e"))
	//fmt.Println(isNumber("1a3.14"))
	//fmt.Println(isNumber("1.2.3"))
	//fmt.Println(isNumber("+-5"))
	//fmt.Println(isNumber("-1E-16"))
	//fmt.Println(isNumber("12e+5.4"))
}
