package main

import (
	"fmt"
	"sort"
	"strings"
)

type Node struct {
	Val byte
	Count int
}


func sortString(s string) string {
	if len(s)<=1{
		return s
	}
	m := make(map[byte]int)
	for i :=range s{
		m[s[i]]++
	}
	var nodeList []*Node
	for k,v:=range m{
		nodeList = append(nodeList,&Node{Val: k, Count: v,})
	}
	sort.Slice(nodeList, func(i, j int) bool {
		   return nodeList[i].Val<nodeList[j].Val
	})

	var res strings.Builder
	for len(res.String())<len(s){
		for i:=0;i<len(nodeList);i++{
			 n := nodeList[i]
			if  n.Count>0{
				res.WriteByte(n.Val)
				n.Count--
			}
		}
		for i:=len(nodeList)-1;i>=0;i--{
			n := nodeList[i]
			if  n.Count>0{
				res.WriteByte(n.Val)
				n.Count--
			}
		}
	}
	return res.String()
}

func main() {
	fmt.Println(sortString("aaaabbbbcccc"))
	fmt.Println(sortString("rat"))
}
