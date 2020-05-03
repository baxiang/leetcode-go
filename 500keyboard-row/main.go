package main

import "fmt"

func findWords(words []string) []string {
     m :=make([]int,123)
     for _,v :=range "qwertyuiopQWERTYUIOP"{
     	m[v]=0
	 }
	for _,v :=range "asdfghjklASDFGHJKL"{
		m[v]=1
	}
	for _,v :=range "zxcvbnmZXCVBNM"{
		m[v]=2
	}
	var res []string
	for _,word:=range words{
		level :=m[word[0]]
		isLevel :=true
		for i:=1;i<len(word);i++{
			if level!=m[word[i]]{
				isLevel = false
				break
			}
		}
		if isLevel{
			res = append(res,word)
		}
	}
	return res
}


func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
