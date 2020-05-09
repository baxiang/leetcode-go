package main

import (
	"fmt"
	"math"
)

func minPathSum(grid [][]int) int {
    return cal(grid,0,0)
}

func cal(grid [][]int,i,j int)int{
	if i==len(grid)||j==len(grid[0]){
		return math.MaxInt64
	}
	if i ==len(grid)-1&&j==len(grid[0])-1{
		return grid[i][j]
	}
	return grid[i][j]+min(cal(grid,i+1,j),cal(grid,i,j+1))
}
func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}

func main() {
	fmt.Println(minPathSum([][]int{{1,3,1}, {1,5,1}, {4,2,1}}))

	//fmt.Println(minPathSum([][]int{
	//	{1,4,8,6,2,2,1,7},
	//	{4,7,3,1,4,5,5,1},
	//	{8,8,2,1,1,8,0,1},
	//	{8,9,2,9,8,0,8,9},
	//	{5,7,5,7,1,8,5,5},
	//	{7,0,9,4,5,6,5,6},
	//	{4,9,9,7,9,1,9,0}},
	//))
}
