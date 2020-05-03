package main

import "fmt"

func luckyNumbers (matrix [][]int) []int {
	//// 为了减少重复计算 k 是列坐标 v 行坐标 表示当前列最大值的行坐标
    m :=make(map[int]int)
	for i:=0;i<len(matrix);i++{
		var min =matrix[i][0]
		var c = 0
		// 计算每行最小的列坐标点
		for j:=1;j<len(matrix[i]);j++{
            if min>matrix[i][j]{
            	min = matrix[i][j]
            	c = j
			}
		}
		//如果已经计算过当前列， 看看是否相等
		if v,ok:=m[c];ok{//
			if v==i{
				 return []int{matrix[i][c]}
			}else {
				continue
			}
		}
		//  已经知道当前行最小值的列坐标，计算当前列最大值行坐标信息
		var max =matrix[0][c]
		var r = 0
		//求列最大的行坐标值
		for k:=1;k<len(matrix);k++{
			if max<matrix[k][c]{
				max = matrix[k][c]
				r = k
			}
		}
		if i==r{// 判断是否相等
			return []int{matrix[i][c]}
		}else {
			m[c]=r //记录当前列最大的行坐标
		}
	}
	return nil
}

func main() {
	fmt.Println(luckyNumbers([][]int{{3,7,8},{9,11,13},{15,16,17}}))
	fmt.Println(luckyNumbers([][]int{{1,10,4,2},{9,3,8,7},{15,16,17,12}}))
	fmt.Println(luckyNumbers([][]int{{7,8},{1,2}}))

}
