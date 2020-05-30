package main

import "fmt"

func getRow(rowIndex int) []int {
	res :=[]int{1}
	for i:=1;i<=rowIndex;i++{
		res =append([]int{0},res...)
		for j:=0;j<i;j++{
			res[j]= res[j]+res[j+1]
		}
	}
	return res
}

func main() {
	fmt.Println(getRow(0))
}
