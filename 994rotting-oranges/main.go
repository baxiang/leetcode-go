package main

import "fmt"

func bfs(grid[][]int,i,j,depth int){
	grid[i][j]= depth
	if i>0&&(grid[i-1][j]==1||grid[i-1][i]-grid[i][j]>1){
		bfs(grid,i-1,j,depth+1)
	}
	if i<len(grid)-1&&(grid[i+1][j]==1||grid[i+1][j]-grid[i][j]>1){
		bfs(grid,i+1,j,depth+1)
	}
	if j>0&&(grid[i][j-1]==1||grid[i][j-1]-grid[i][j]>1){
		bfs(grid,i,j-1,depth+1)
	}
	if j<len(grid[0])-1&&(grid[i][j+1]==1||grid[i][j+1]-grid[i][j]>1){
		bfs(grid,i,j+1,depth+1)
	}
}

func orangesRotting(grid [][]int) int {
     if len(grid)==0{
     	return -1
	 }
	 for i:=0;i<len(grid);i++{
	 	for j:=0;j<len(grid[i]);j++{
	 		 if grid[i][j]==2{
	 		 	 bfs(grid,i,j,2)
			 }
		}
	 }
	 var max =0
	 for i:=0;i<len(grid);i++{
	 	for j:=0;j<len(grid[i]);j++{
	 		if grid[i][j]==1{
	 			return -1
			}
			if max<grid[i][j]{
				max = grid[i][j]
			}
		}
	 }
	 if max ==0{
	 	return 0
	 }
	 return max-2
}

func main() {
	fmt.Println(orangesRotting([][]int{{2,1,1},{1,1,0},{0,1,1}}))
	fmt.Println(orangesRotting([][]int{{2,1,1},{0,1,1},{1,0,1}}))
	fmt.Println(orangesRotting([][]int{{0,2}}))
}
