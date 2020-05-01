package main

import "fmt"
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
    left :=0
    right :=len(matrix[0])-1
    top :=0
    bottom :=len(matrix)-1
    grid := make([]int,(right+1)*(bottom+1))
    index :=0
    for left<=right&&top<=bottom{
    	// 上层 边界判断
    	for i:=left;i<=right&&top<=bottom;i++{
    		grid[index] = matrix[top][i]
    		index++
		}
		//fmt.Println("top","=",grid)
		top++
		// 右侧
		for i:=top;i<=bottom&&left<=right;i++{
			grid[index]= matrix[i][right]
			index++
		}
		//fmt.Println("right","=",grid)
		right--

		//下
		for i:=right;i>=left&&top<=bottom;i--{
			grid[index]= matrix[bottom][i]
			index++
		}
		//fmt.Println("bottom","=",grid)
		bottom--
		//左
		for i:=bottom;i>=top&&left<=right;i--{
			grid[index]= matrix[i][left]
			index++
		}
		//fmt.Println("left","=",grid)
		left++
		//fmt.Println(grid)
	}
	return grid
}


func main() {
	matrix :=[][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	fmt.Println("1","=",spiralOrder(matrix))

}
