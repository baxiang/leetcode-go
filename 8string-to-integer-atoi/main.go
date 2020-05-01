package main

import (
	"math"
	"fmt"
	"strings"
)

func myAtoi(str string) int {
	//去掉收尾空格
	str = strings.TrimSpace(str)
	result := 0
	sign := 1

	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}
       // 数值最大检测
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return sign * result
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi(   " -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
}
