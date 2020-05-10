package main

import "fmt"

func spiralOrder1(matrix [][]int) []int {
	var res []int
	if len(matrix)==0{
		return res
	}
	row :=len(matrix)
	col := len(matrix[0])
	up :=0//上界
	down :=row-1//下界
	left :=0//左界
	right :=col-1// 右界
	total := row*col//总数
	for {
		// 从左到右
		for i:=left;i<=right;i++{
			res = append(res,matrix[up][i])
		}
		if len(res)==total{
			break
		}
		up++// 上边界+1 向下移动一位
		for i:=up;i<=down;i++{
			res = append(res,matrix[i][right])
		}
		if len(res)==total{
			break
		}
		right--
		for i:=right;i>=left;i--{
			res = append(res,matrix[down][i])
		}
		if len(res)==total{
			break
		}
		down--
		for i:=down;i>=up;i--{
			res = append(res,matrix[i][left])
		}
		if len(res)==total{
			break
		}
		left++
	}
	return res
}


func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix)==0{
		return res
	}
	row :=len(matrix)
	col := len(matrix[0])
    up :=0//上界
    down :=row-1//下界
    left :=0//左界
    right :=col-1// 右界
    for {
    	// 从左到右
    	for i:=left;i<=right;i++{
    		res = append(res,matrix[up][i])
		}
		up++// 上边界+1 向下移动一位
		if up>down{
			break
		}
		// 从上到下
		for i:=up;i<=down;i++{
			res = append(res,matrix[i][right])
		}
		// 右边界移动
		right--
		if right<left{
			break
		}
		//从右到左
		for i:=right;i>=left;i--{
			res = append(res,matrix[down][i])
		}
		//下边界向上移动一位
		down--
		if down<up{
			break
		}
        // 从下到上
		for i:=down;i>=up;i--{
			res = append(res,matrix[i][left])
		}
        // 左边界向右移动一位
		left++
		if left>right{
			break
		}
	}
   return res
}


func main() {
	fmt.Println(spiralOrder1([][]int{
		{1, 2, 3 },
	{ 4, 5, 6 },
	{ 7, 8, 9 },
	}))

	fmt.Println(spiralOrder1([][]int{
		{1, 2, 3, 4},
	{5, 6, 7, 8},
	{9, 10, 11, 12},
	}))
}
