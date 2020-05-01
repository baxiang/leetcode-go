package main

import "fmt"

func divingBoard(shorter int, longer int, k int) []int {
	res := make([]int,0)
	if k==0 {
		return res
	}
	if shorter==longer{
		return append(res,k*shorter)
	}
	for i:=0;i<=k;i++{
		res = append(res,(k-i)*shorter+i*longer)
	}
	return res
}
func main() {
	fmt.Println(divingBoard(1,1,0))//[0]
	fmt.Println(divingBoard(1,1,10000))
	fmt.Println(divingBoard(1,2,3))
}
