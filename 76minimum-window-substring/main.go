package main

import "fmt"

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
	fmt.Println(minWindow("ADOBECODEBANC","ABC"))
}
