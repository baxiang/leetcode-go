package main

import "fmt"

func generate(numRows int) [][]int {
	 var res [][]int
	 if numRows==0{
	 	return res
	 }
	 res = append(res,[]int{1})
     for i:=1;i<numRows;i++{
     	m :=[]int{0}
     	m =append(m,res[i-1]...)
     	for j:=0;j<len(m)-1;j++{
     		m[j]=m[j]+m[j+1]
		}
		res = append(res,m)
	 }
	 return res
}

func main() {
	fmt.Println(generate(5))
}
