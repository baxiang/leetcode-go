package main

import "fmt"

func dfs1(grid [][]byte,i,j int){
	if i<0||i>=len(grid) ||j<0||j>=len(grid[i])||grid[i][j]!='1'{
		return
	}
	grid[i][j]= '2'
	dfs(grid,i-1,j)
	dfs(grid,i+1,j)
	dfs(grid,i,j-1)
	dfs(grid,i,j+1)
}

type Point struct {
	i int
	j int
}

func dfs(grid [][]byte,i,j int){
	stack :=make([]*Point,0)
	stack = append(stack,&Point{i: i, j: j})
	for len(stack)>0{
		point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if point.i<0||point.i>=len(grid) ||point.j<0||point.j>=len(grid[point.i])||grid[point.i][point.j]!='1'{
			continue
		}
		grid[point.i][point.j]= '2'
		stack =append(stack,&Point{i: point.i, j: point.j+1})
		stack =append(stack,&Point{i: point.i, j: point.j-1})
		stack =append(stack,&Point{i: point.i-1, j: point.j})
		stack =append(stack,&Point{i: point.i+1, j: point.j})
	}
}



func numIslands(grid [][]byte) int {
	var res int
   for i:=0;i<len(grid);i++{
   	for j:=0;j<len(grid[0]);j++{
   		if grid[i][j]=='0'||grid[i][j]=='2'{
   			continue
		}
		res++
		dfs(grid,i,j)
	}
   }
   return res
}

func main() {
	g :=[][]byte{{'1','1','1','1','0'},
		{'1','1','0','1','0'},
		{'1','1','0','0','0'},
		{'0','0','0','0','0'}}
	fmt.Println(numIslands(g))
}
