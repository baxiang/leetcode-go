package main

import "fmt"

func longestPalindrome(s string) int {
     l := len(s)
     m :=make(map[byte]int)
     for i:=0;i<l;i++{
     	 m[s[i]]++
	 }
	 var res int
	 for _,v:=range m{
	 	 if v/2>0{
	 	 	res +=(v/2)*2
		 }
	 }
	 if l%2!=0||(l%2==0&&res!=l){
	 	res=res+1
	 }
	 return res
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("ccc"))
}
