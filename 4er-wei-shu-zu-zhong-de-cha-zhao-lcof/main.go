package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		 l :=len(matrix[i])
		if l>0 &&matrix[i][0] <= target && target <= matrix[i][l-1] {
			ok := BinarySearch(matrix[i], target)
			if ok {
				return true
			}
		}
	}

	return false
}

func BinarySearch(a []int, target int) bool {
	l := 0
	r := len(a) - 1
	for l <= r {
		mid := l + (r-l)>>1
		if a[mid] == target {
			return true
		} else if a[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

//面试题04. 二维数组中的查找
func main() {
	fmt.Println(findNumberIn2DArray(
		[][]int{{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30}}, 20))
	fmt.Println(findNumberIn2DArray(
		[][]int{{}}, 1))
}
