package main

import "strings"

func wordBreak(s string, wordDict []string) bool {
    for _,v:=range wordDict{
    	if strings.Contains(s,v){
    		return true
		}
	}
	return false
}

func main() {
	
}
