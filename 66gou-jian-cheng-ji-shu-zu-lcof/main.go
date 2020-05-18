package main

import "fmt"

func constructArr1(a []int) []int {
	var res []int
    for i :=0;i<len(a);i++{
    	v :=1
    	for j:=0;j<len(a);j++{
    		if i==j{
    			continue
			}
			v*=a[j]
		}
		res = append(res,v)
	}
	return res
}
func constructArr2(a []int) []int {
	if len(a)<=1{
		return a
	}
	size := len(a)
	leftRes :=make([]int,size)
	leftRes[0]= 1
	for i :=1;i<size;i++{
		leftRes[i] = leftRes[i-1]*a[i-1]
	}
	rightRes :=make([]int,size)
	rightRes[size-1] =1
	for i :=size-2;i>=0;i--{
		rightRes[i] = rightRes[i+1]*a[i+1]
	}
	res := make([]int,size)
	for i := 0; i < size; i++{
		res[i] = leftRes[i] * rightRes[i]
	}
	return res
}
func constructArr(a []int) []int {
	if len(a)<=1{
		return a
	}
	size := len(a)
	res :=make([]int,size)
	mutVal :=1
	for i :=0;i<size;i++{
		res[i] = mutVal // 先乘左边的数(不包括自己)
		mutVal *=a[i]
	}
	mutVal =1
	for i :=size-1;i>=0;i--{
		res[i] *= mutVal //再乘右边的数(不包括自己)
		mutVal *=a[i]
	}
	return res
}

func main() {
	fmt.Println(constructArr([]int{1,2,3,4,5}))
	fmt.Println(constructArr([]int{1,2,0,4,5}))
}
