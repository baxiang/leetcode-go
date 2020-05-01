package main

import "fmt"

func countNegatives1(grid [][]int) int {
	 var res int
     for i:=0;i<len(grid);i++{// 行
     	for j:=0;j<len(grid[i]);j++{// 列
     		if grid[i][j]<0{
     			res++
			}
		}
	 }
	 return res
}


func countNegatives(grid [][]int) int {
	var res int
	for i:=0;i<len(grid);i++{
		 c :=len(grid[i])
		 if grid[i][c-1]>=0{
		 	continue
		 }
		if grid[i][0]<0{
			res +=c
			continue
		}
		 r := binarySearch(grid[i])
		 res +=c-r
	}
	return res
}
func binarySearch(g []int) int {
	l:=0
	r :=len(g)-1
	for l<=r{
		mid :=l+(r-l)>>1
		if g[mid]>=0{
			l = mid+1
		}else {
			r =mid-1
		}
	}
	return l
}



func main() {
	grid := [][]int{{4,3,2,-1},{3,2,1,-1},{1,1,-1,-2},{-1,-1,-2,-3}}
	fmt.Println(countNegatives(grid))
	//fmt.Println(binarySearch([]int{4,3,2,-1}))
	//fmt.Println(binarySearch([]int{3,2,1,-1}))
	//fmt.Println(binarySearch([]int{1,1,-1,-2}))
	//fmt.Println(binarySearch([]int{0,1,2,3}))
}
