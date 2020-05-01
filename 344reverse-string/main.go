package main

import "fmt"

func reverseString(s []byte)  {
    l :=0
    r :=len(s)-1
    for l<r{
    	s[l],s[r] = s[r],s[l]
    	l++
    	r--
	}
}
func main() {
	s := []byte{'h','e','l','l','o'}
	fmt.Println(string(s))
	reverseString(s)
	fmt.Println(string(s))
}
