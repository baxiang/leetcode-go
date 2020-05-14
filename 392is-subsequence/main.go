package main

import "fmt"

func isSubsequence(s string, t string) bool {
    s1 :=0
    t1 :=0
    for t1<len(t)&&s1<len(s){
    	if s[s1]==t[t1]{
    		s1++
		}
		t1++
	}
	if s1==len(s){
		return true
	}
	return false
}


func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("acb", "ahbgdc"))
}
