package main

import (
	"fmt"
	"sort"
)

func sortedSquares1(A []int) []int {
	for i, v := range A {
		A[i] = v * v
	}
	sort.Ints(A)
	return A
}

func sortedSquares(A []int) []int {
	res := make([]int, len(A))
	l := 0
	r := len(A) - 1
	i := len(res) - 1
	var lRes, rRes int
	for l <= r {
		lRes = A[l] * A[l]
		rRes = A[r] * A[r]
		if lRes > rRes {
			res[i] = lRes
			l++
		} else {
			res[i] = rRes
			r--
		}
		i--
	}
	return res
}

func main() {
	//fmt.Println(sortedSquares([]int{-4,-1,0,3,10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
	fmt.Println(sortedSquares([]int{1}))
}
