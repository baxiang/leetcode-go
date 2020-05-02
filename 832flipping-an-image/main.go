package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
    for i:=0;i<len(A);i++{
    	l :=0
    	r :=len(A[i])-1
    	for l<=r{
    		if l<r{
				A[i][l],A[i][r]= A[i][r],A[i][l]
				A[i][l]=A[i][l]^1
				A[i][r]=A[i][r]^1
			}else {
				A[i][l]=A[i][l]^1
			}
    		l++
    		r--
		}
	}
	return A
}

func main() {
	//
	fmt.Println(flipAndInvertImage([][]int{{1,1,0},{1,0,1},{0,0,0}}))
}
