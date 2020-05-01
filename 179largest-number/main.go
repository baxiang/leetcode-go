package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)
type StrList []string
func(s StrList)Len()int {
	return len(s)
}
func (s StrList)Less(i,j int)bool{
	return  s[i]+s[j]>s[j]+s[i]
}
func(s StrList)Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}

func largestNumber(nums []int) string {
	strs :=make(StrList,len(nums))
	for i,v:=range nums{
       strs[i] = strconv.Itoa(v)
	}
	sort.Slice(strs, func(i ,j int)bool {
		return strs[i]+strs[j]>strs[j]+strs[i]
	})
	s := strings.Join(strs,"")
	if s[0]=='0' {
		return "0"
	}
	return s
}
func main() {
	fmt.Println(largestNumber([]int{3,30,34,5,9}))
	fmt.Println(largestNumber([]int{0,0}))
}
