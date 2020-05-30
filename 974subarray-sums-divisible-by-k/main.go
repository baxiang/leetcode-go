package main

import "fmt"

func subarraysDivByK1(A []int, K int) int {
	var res int
	for i := 0; i < len(A); i++ {
		var sum =0
		for j := i ; j <len(A); j++ {
			sum +=A[j]
			if sum%K == 0 {
				res++
			}
		}
	}
	return res
}
func subarraysDivByK(A []int, K int) int {
	record := map[int]int{0: 1} // 如果A中有元素能整除K， 那么相关元素单独能构成符合题意的子数组，需要统计
	prefixSum, result := 0, 0
	for i := 0; i < len(A); i++ {
		prefixSum += A[i]
		mod := (prefixSum % K+K)%K
		//if mod < 0 { // prefixSum可能是负数，导致mod为负数， 取余的结果mod为负，跟取余结果为mod+K是等价的
		//	mod += K
		//}
		fmt.Println(mod)
		result += record[mod]
		record[mod]++
	}
	return result
}


func main() {
	subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)
	fmt.Println()
}
