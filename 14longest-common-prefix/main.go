package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs)==0 {
		return ""
	}
	if len(strs)==1 {
		return strs[0]
	}
	preStr :=strs[0]
	for i:=1;i<len(strs);i++{
		for !strings.HasPrefix(strs[i],preStr){
			if len(preStr)==0 {
				return ""
			}
			preStr =preStr[0:len(preStr)-1]
		}
	}
	return preStr
}


func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
}
