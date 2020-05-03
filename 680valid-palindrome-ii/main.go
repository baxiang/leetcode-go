package main

import "fmt"

func validPalindrome(s string) bool {
	m :=make(map[byte]int)
	for i:=range s{
		m[s[i]]++
	}
	var count int
	for _,k:=range m{
		if k%2!=0{
			count++
		}
	}
	l :=len(s)
	if (l%2==0&&count>0)||(l%2!=0&&count>2){
		return false
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("tebbem"))
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("abc"))
	//eeabddddbaeed
	fmt.Println(validPalindrome("cdbeeeabddddbaeedebdc"))
	//aguokepatg
	//bnvfqmgmlc
	//upuufxoohd
	//fpgjdmysgv
	//hmvffcnqxj
	//jxqncffvmh
	//vgsymdjgpf
	//dhooxfuupu
	//culmgmqfvn
	//bgtapekouga
	fmt.Println(validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
}
