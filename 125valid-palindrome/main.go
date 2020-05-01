package main

import (
	"fmt"
	"unicode"
)

func isValid(i rune)bool{
	return unicode.IsDigit(i)||unicode.IsLetter(i)
}

func isPalindrome1(s string) bool {
	if len(s)<2{
		return true
	}
	l :=0
	r :=len(s)-1
	for l<r{
		vl, vr := rune(s[l]), rune(s[r])
		if !isValid(vl)&&!isValid(vr){
			l++
			r--
		}else if !isValid(vl){
			l++
		}else if !isValid(vr){
			r--
		}else if unicode.ToUpper(vl) != unicode.ToUpper(vr){
			return false
		}else{
			l++
			r--
		}
	}
	return true
}
func isPalindrome(s string) bool {
	// 判断是否是需要比较的合法回文字符,
	// 题目说空格是合法的回文字符很误导人, 其实这个地方需要排除空格
	isValid := func(v rune) bool {
		return unicode.IsDigit(v) || unicode.IsLetter(v)
	}
	i, j := 0, len(s)-1
	for i < j {
         for i<len(s)&&!isValid(rune(s[i])){
         	i++
		 }
		for j>=0&&!isValid(rune(s[j])){
			j--
		}
	     if i<len(s)&&j>0&&(unicode.ToUpper(rune(s[i])) != unicode.ToUpper(rune(s[j]))) {
			return false // 如果都合法如果不相等则直接返回
		} else { // 如果相等则移动两个指针
			i++
			j--
		}
	}
	return true
}

func main() {
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	//fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("a."))
	fmt.Println(isPalindrome(".,"))
}
