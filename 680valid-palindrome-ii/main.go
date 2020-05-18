package main

import "fmt"

func validPalindrome(s string) bool {
	left :=0
	right :=len(s)-1
	for left<right{
		if s[left]!=s[right]{
			return isPalindrome(s[left:right])||isPalindrome(s[left+1:right+1])
		}
		left++
		right--
	}
	return true
}

func isPalindrome(s string)bool{
	left :=0
	right :=len(s)-1
	for left<right{
		if s[left]!=s[right]{
			return false
		}
		left++
		right--
	}
	return true
}

func palindromeHelp(s string,isDelete bool)bool{
	left :=0
	right :=len(s)-1
	for left<right{
		if s[left]!=s[right]{
			if isDelete{
				return palindromeHelp(s[left:right],false)||palindromeHelp(s[left+1:right+1],false)
			}else {
				return false
			}
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("tebbem"))//f
	fmt.Println(validPalindrome("aba")) //t
	fmt.Println(validPalindrome("abca")) //t
	fmt.Println(validPalindrome("abc")) //f
	fmt.Println(validPalindrome("cdbeeeabddddbaeedebdc"))//t
	fmt.Println(validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))//t
}
