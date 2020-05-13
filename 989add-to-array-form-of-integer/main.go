package main

import "fmt"

func addToArrayForm(A []int, K int) []int {
	var B []int
	for K>0{
		B = append([]int{K%10},B...)
		K=K/10
	}
	a1 :=len(A)-1
	b1 :=len(B)-1
	carry :=0
	var res []int
	for a1>=0||b1>=0{
		n1 :=0
		if a1>=0{
			n1=A[a1]
			a1--
		}
		n2 :=0
		if b1>=0{
			n2=B[b1]
			b1--
		}
		r := n1+n2+carry
		carry=r/10
		res = append([]int{r%10},res...)
	}
	if carry>0 {
		res = append([]int{carry},res...)
	}
	return res
}

func main() {
	fmt.Println(addToArrayForm([]int{1,2,0,0},34))
	fmt.Println(addToArrayForm([]int{2,7,4},181))
}
