package main

import (
	"fmt"
)

func lengthOfLongestSubstring4(s string) int {
	freq := make([]int,128)
	var res = 0
	start,end := 0,-1
	for start<len(s){
		if end+1<len(s)&&freq[s[end+1]] == 0{
			end++
			freq[s[end]]++
		}else{
			freq[s[start]]--
			start++
		}
		res = max(res,end-start+1)
	}
	return res
}

func max(i,j int)int{
	if i>j{
		return i
	}else{
		return j
	}
}

func lengthOfLongestSubstring2(s string) int{
	m :=make(map[byte]int) //定义map
	max :=0
    start :=0
	for i :=range s{
		v,ok := m[s[i]]
		if ok {
			j := start
			for ;j<=v;j++{
				delete(m,s[j])
			}
			start =j
		}
		m[s[i]] = i
		if max<i-start+1 {
			max = i-start+1
		}
	}
	return max
}
func lengthOfLongestSubstring3(s string) int{
	m := make(map[byte]int)
	var r []byte //窗口数据
	l :=0
	for i :=range s{
		if j,ok:=m[s[i]];ok {
			if i-j>len(r) {
				r =r[i-j:]
			}
		}else {
			r = append(r,s[i])
		}
		m[s[i]] = i
		l = max(l,len(r))
	}
	return l
}


func lengthOfLongestSubstring1(s string) int{
	if len(s)==0 {
		return 0
	}
	ans := 0 // 默认就是1了
	m := map[byte]int{}
	start := -1
	for i:=0 ;i<len(s);i++{// 遍历字符串
		if index,ok := m[s[i]];ok {//存在
			start = max(start,index)//但是要保证起始索引是递增的
		}
		ans = max(ans,i-start)
		m[s[i]]=i
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring(""))//1
	fmt.Println(lengthOfLongestSubstring1("au"))//2
	fmt.Println(lengthOfLongestSubstring1("abcabcbb"))//3
	fmt.Println(lengthOfLongestSubstring1("bbbb"))//1
	fmt.Println(lengthOfLongestSubstring1("pwwkew"))//3
	fmt.Println(lengthOfLongestSubstring1("abba"))//2
	fmt.Println(lengthOfLongestSubstring1("bpfbhmipx"))//7
    fmt.Println(lengthOfLongestSubstring1("eeydgwdykpv"))//7
}
