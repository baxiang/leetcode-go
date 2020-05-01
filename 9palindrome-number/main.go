package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(123))
}
// 主要难点在如何反转数字? rev * 10 +tmp % 10
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x % 10 == 0) {
		return false
	}
	tmp := x
	rev := 0 // 当前数字的反转数
	for tmp > 0 {
		rev = rev * 10 +tmp % 10 // 反转数公式
		tmp /= 10
	}
	return x==rev
}