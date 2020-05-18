package main

import "fmt"

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	m := map[int]int{}
	sum := 0
	for i, v := range A {
		if v%2 == 0 {
			m[i] = v
			sum += v
		}

	}
	var res []int
	for i := 0; i < len(queries); i++ {
		val := queries[i][0]
		idx := queries[i][1]
		preVal := A[idx]
		currVal := A[idx] + val
		A[idx] = currVal
		if preVal%2 != 0 && currVal%2 == 0 {
			sum += currVal
		} else if preVal%2 == 0 && currVal%2 != 0 {
			sum -= preVal
		} else if preVal%2 == 0 && currVal%2 == 0 {
			sum += val
		}
		res = append(res, sum)
	}
	return res
}
func main() {
	fmt.Println(sumEvenAfterQueries([]int{1,2,3,4},[][]int{{1,0},{-3,1},{-4,0},{2,3}}))
}
