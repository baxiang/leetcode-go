package main

func maxValue(grid [][]int) int {
    row :=len(grid)// 行
    column := len(grid[0])//列
    dp := make([][]int,row)
    for i:=range dp{
    	dp[i] = make([]int,column)
	}
	dp[0][0]= grid[0][0]
	for i :=1;i<row;i++{//第一列行
		dp[i][0]= dp[i-1][0]+grid[i][0]
	}
	for j:=1;j<column;j++{//第一行的列
		dp[0][j] = dp[0][j-1]+grid[0][j]
	}
	for i:=1;i<row;i++{
		for j:=1;j<column;j++{
			dp[i][j]= grid[i][j]+Max(dp[i][j-1],dp[i-1][j])
		}
	}
	return dp[row-1][column-1]
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//面试题47. 礼物的最大价值
func main() {
	
}
