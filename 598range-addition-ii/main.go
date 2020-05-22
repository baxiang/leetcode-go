package main

import "fmt"

func maxCount1(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < len(ops); i++ {
		for j := 0; j < ops[i][0]; j++ {
			for k := 0; k < ops[i][1]; k++ {
				matrix[j][k]++
			}
		}
	}
	max := matrix[0][0]
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if max == matrix[i][j] {
				res++
			}
		}
	}
	return res
}

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	r := ops[0][0]
	c := ops[0][1]
	for i := 1; i < len(ops); i++ {
		if r > ops[i][0] {
			r = ops[i][0]
		}
		if c > ops[i][1] {
			c = ops[i][1]
		}
	}
	return r * c
}

func main() {
	fmt.Println(maxCount1(3, 3, [][]int{{2, 2}, {3, 3}}))
}
