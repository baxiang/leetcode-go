package main

import "fmt"

func minDeletionSize(A []string) int {
	var res int
	for i := 0; i < len(A[0]); i++ {
		s := make([]byte, len(A))
		for j := 0; j < len(A); j++ {
			s[j] = A[j][i]//
		}
		// 获取每列字符串
		if !isIncrStr(s) {
			res++
		}
	}
	return res
}

func isIncrStr(s []byte) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}


func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
	//"a", "b"
	fmt.Println(minDeletionSize([]string{"a", "b"}))
	//
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))

}
