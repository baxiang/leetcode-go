package main

import (
	"fmt"
	"math"
	"strings"
)

func strToInt(str string) int {
	str = strings.TrimSpace(str)
	var singn = 1
	var result int
	for i :=range str{
		if i==0&&str[i]=='-' {
			singn = -1
		}else if i==0&&str[i]=='+' {
			singn = 1
		} else if i==0&&(str[i]>'9'||str[i]<'0'){
			return  0
		}else if str[i]<='9'&&str[i]>='0' {
			n := int(str[i]-'0')
			result =result*10+n
		}else {
			break
		}
	}
	 result =result*singn
	if result>math.MaxInt32 {
		return math.MaxInt32
	}
	if result<math.MinInt32 {
		return math.MinInt32
	}
	return result
}
func main() {
	fmt.Println(strToInt("42"))
	fmt.Println(strToInt("   -42"))
	fmt.Println(strToInt("4193 with words"))
	fmt.Println(strToInt("words and 987"))
	fmt.Println(strToInt("-91283472332"))
	fmt.Println(strToInt("3.1415926"))
}
