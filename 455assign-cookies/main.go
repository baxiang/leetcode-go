package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	if len(g)==0||len(s)==0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	var cookie =0
	var child = 0
	for  child<len(g)&&cookie<len(s){
		if g[child]<=s[cookie] {
			child++
		}
		cookie++
	}
	return child
}

func main() {
	g := []int{1,2,3}
	s := []int{1,1}
	fmt.Println(findContentChildren(g,s))
	g1 := []int{1,2}
	s1 := []int{1,2,3}
	fmt.Println(findContentChildren(g1,s1))
	g2 := []int{1,2,3}
	s2 := []int{3}
	fmt.Println(findContentChildren(g2,s2))
}
