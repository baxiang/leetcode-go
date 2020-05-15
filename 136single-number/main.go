package main

import "fmt"

//func singleNumber(nums []int) int {
//        n :=map[int]int{}
//        for _,v :=range nums{
//        	n[v]++
//		}
//        for k,v :=range n{
//        	if v ==1{
//        		return k
//			}
//		}
//        return 0
//}

// 异或运算
//任何数和 00 做异或运算，结果仍然是原来的数
// 任何数和其自身做异或运算，结果是 00
func singleNumber(nums []int) int {
	var ans = 0
	for _,n :=range nums {
		ans ^= n
	}
	return ans
}

func main() {
	fmt.Println(singleNumber([]int{2,2,1}))
	fmt.Println(singleNumber([]int{4,1,2,1,2}))
	fmt.Println(singleNumber([]int{-1}))
}
