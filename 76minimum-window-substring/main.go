package main

import (
	"fmt"
)

func minWindow1(s string, t string) string {
	need := make(map[byte]int)
	for i :=range t{
		need[t[i]]++
	}

	window :=make(map[byte]int)
	var left,right,match int
	var minLen = 0
	var resL int//窗口最小左边索引
	var resR int//窗口最小右边索引
	for right<len(s){
		if _,ok:=need[s[right]];ok{
			window[s[right]]++
			if window[s[right]]==need[s[right]] {
				match++
			}
		}
		for match ==len(need){
			l := right-left+1
			//fmt.Println(s[left:right+1])
			if minLen==0||l<minLen {
				resL = left
				resR = right
				minLen = l
			}
			if v,ok:=need[s[left]];ok{
				window[s[left]]--
				if window[s[left]]<v{
					match--
				}
			}
			left++
		}
		right++
	}
	if minLen==0 {
		return ""
	}
	return s[resL:resR+1]
}

func isContainer(strMap map[byte]int,m map[byte]int)bool{
	for k,v:=range m{
		if strMap[k]<v{
			return false
		}
	}
	return true
}



func minWindow(s string, t string) string {
	need :=make(map[byte]int)
	for i:=0;i<len(t);i++{
		need[t[i]]++
	}
	var left,right,match,minLen,resL,resR int
	window := make(map[byte]int)
	for right<=len(s)-1{
		if _,ok:=need[s[right]];ok{
			window[s[right]]++
			if window[s[right]]==need[s[right]] {
				match++
			}
		}
		for match ==len(need){//窗口匹配
			l := right-left+1
			if minLen==0||l<minLen {
				resL = left
				resR = right
                minLen = l
			}
			if v,ok:=need[s[left]];ok{
				window[s[left]]--
				if window[s[left]]<v{
					match--
				}
			}
			left++
		}
		right++
	}
	if minLen==0 {
		return ""
	}
   return s[resL:resR+1]
}
func main() {
	//fmt.Println(minWindow1("BANC","ABC"))
	fmt.Println(minWindow1("ADOBECODEBANC","ABC"))
}
