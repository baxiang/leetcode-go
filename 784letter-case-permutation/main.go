package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}
func letterCasePermutation(S string) []string {
	res :=make([]string,0)
	dfs1([]byte(S),0,&res)
	return res
}

func dfs(char []byte,index int,res *[]string){
	if index==len(char) {
		*res = append(*res,string(char))
		return
	}
	dfs(char,index+1,res)
	if unicode.IsLetter(rune(char[index])) {
		if unicode.IsLower(rune(char[index])) {
			char[index] = strings.ToUpper(string(char[index]))[0]
		}else {
			char[index] = strings.ToLower(string(char[index]))[0]
		}
		dfs(char,index+1,res)
	}
}

func dfs1(char []byte,index int,res *[]string){
	if index==len(char) {
		*res = append(*res,string(char))
		return
	}
	dfs(char,index+1,res)
	if unicode.IsLetter(rune(char[index])) {
		char[index] = char[index]^1<<5
		dfs(char,index+1,res)
	}
}