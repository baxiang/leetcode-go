package main

import (
	"fmt"
	"sort"
)

func numSpecialEquivGroups(A []string) int {
	var res int
	m:=make(map[string]struct{})
	for _,word:=range A{
		var even []byte
		var odd []byte
		for i:=range word{
			if i%2==0{
				even=append(even,word[i])
			}else{
				odd=append(odd,word[i])
			}
		}
		sort.Slice(even, func(i, j int) bool {
			return even[i]<even[j]
		})
		sort.Slice(odd, func(i, j int) bool {
			return odd[i]<odd[j]
		})
		s:=string(even)+string(odd)
		if _,ok:=m[s];!ok{
			res++
			m[s]= struct{}{}
		}
	}
	return res
}

func main() {
	fmt.Println(numSpecialEquivGroups([]string{"aa","bb","ab","ba"}))
	fmt.Println(numSpecialEquivGroups([]string{"abcd","cdab","adcb","cbad"}))
}
