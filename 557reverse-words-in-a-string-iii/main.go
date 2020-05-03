package main

import (
	"fmt"
	"strings"
)

func reverseWords1(s string) string {
    l := strings.Split(s," ")
    var str strings.Builder
    for i,v:=range l{
		str.Write(revString(v))
    	if i!=len(l)-1{
			str.WriteString(" ")
		}
	}
	return str.String()
}

func revString(s string)[]byte{
	 l:=0
	 r :=len(s)-1
	 b :=[]byte(s)
	 for l<r{
	 	b[l],b[r] =b[r],b[l]
	 	l++
	 	r--
	 }
	 return b
}
func reverseWords(s string) string {
	b :=[]byte(s)
	l :=0
	for i,v :=range s{
		if v==' '||i==len(s)-1{
			r :=i-1
			if i==len(s)-1{
				r =i
			}
			for l<r{
				b[l],b[r] =b[r],b[l]
				l++
				r--
			}
			l=i+1
		}
	}
	return string(b)
}

func main() {
	                            //s'teL ekat edoCteeL tsetnoc
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
