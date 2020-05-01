package main

import "fmt"

func findSubstring(s string, words []string) []int {
	if len(s)==0{
		return []int{}
	}
	allMap :=make(map[string]int)
	for _,v:=range words{
		allMap[v]++
	}
	wordNum :=len(words)
	wordLen :=len(words[0])
	var res []int
	fmt.Println(len(s)-wordLen*wordNum)
	for i:=0;i<=len(s)-wordLen*wordNum;i++{
		hasMap :=make(map[string]int)
		num :=0
		for num<wordNum{
			word :=s[i+num*wordLen:i+(num+1)*wordLen]
			if v,ok:=allMap[word];ok{
				hasMap[word]++
				if hasMap[word]>v{
					break
				}
			}else {
				break
			}
			num++
		}
		if num==wordNum{
			res = append(res,i)
		}
	}
	return res
}
func findSubstring1(s string, words []string) []int{
	if len(s)==0||len(words)==0 {
		return nil
	}
	allMap :=make(map[string]int)
	for _,v:=range words{
		allMap[v]++
	}
	wordsLen :=len(words)*len(words[0])
	var res []int
	for i:=0;i<=len(s)-wordsLen;i++{
		str :=s[i:i+wordsLen]
		hasMap :=make(map[string]int)
		star:=0
		for ;star<=len(str)-len(words[0]);{
			ss :=str[star:star+len(words[0])]
			if allMap[ss]>0 {
				hasMap[ss]++
				if hasMap[ss]> allMap[ss]{
					break
				}
			}else {
				break
			}
			star = star+len(words[0])
		}
		if star==len(str) {
			res = append(res,i)
		}

	}
	return res
}

func findSubstring2(s string, words []string) []int{
	if len(s)==0||len(words)==0 {
		return nil
	}
	allMap :=make(map[string]int)
	for _,v:=range words{
		allMap[v]++
	}
	wordsLen :=len(words)*len(words[0])
	var res []int
	for i:=0;i<=len(s)-wordsLen;i++{
		str :=s[i:i+wordsLen]
		hasMap :=make(map[string]int)
		start:=0
		for ;start<len(words);start++{
			ss :=str[len(words[0])*start:len(words[0])*(start+1)]
			//fmt.Println(ss)
			if allMap[ss]>0 {
				hasMap[ss]++
				if hasMap[ss]> allMap[ss]{
					break
				}
			}else {
				break
			}
		}
		if start==len(words) {
			res = append(res,i)
		}

	}
	return res
}

func main() {
	fmt.Println(findSubstring2("barfoothefoobarman",[]string{"bar","foo"}))
}
