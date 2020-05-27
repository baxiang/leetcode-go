package main

import "fmt"

func countSubstrings(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) {
				res++
			}
		}
	}
	return res
}

func countSubstrings1(s string) int {

	var res int
	flag := make([][]bool,len(s))
	for i:=range flag{
		flag[i]=  make([]bool,len(s))
	}
	for j := 0; j < len(s); j++ {
		for i := j ; i >=0; i-- {
			if s[i] == s[j] && (j - i < 2 || flag[i + 1][j - 1]) {
				flag[i][j] = true
				res++
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	fmt.Println(s)
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	fmt.Println(countSubstrings1("abba"))
	//fmt.Println(countSubstrings1("aaa"))
}
