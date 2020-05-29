package main

import (
	"fmt"
	"strconv"
)

func decodeString1(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			var index int
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j] == '[' {
					index = j
					break
				}
			}
			list := make([]byte, len(stack)-index+1)
			copy(list, stack[index-1:])
			var countStr = ""
			for index = index - 1; index >= 0; index-- {
				if stack[index] >= '0' && stack[index] <= '9' {
					countStr = fmt.Sprintf("%d%s", stack[index]-'0', countStr)
				} else {
					break
				}
			}
			count, _ := strconv.Atoi(countStr)
			if index < 0 {
				stack = stack[:0]
			} else {
				stack = stack[:index+1]
			}
			list = list[2:]
			for n := 0; n < count; n++ {
				for m := 0; m < len(list); m++ {
					stack = append(stack, list[m])
				}

			}
		}
	}
	return string(stack)
}

func decodeString(s string) string {
	alphaStack := make([]byte, 0)
	countStack := make([]int,0)
	alphaIndex := make([]int,0)
	var count int
	for i:=0 ;i < len(s);i++{
		if s[i] >='0'&&s[i]<='9' {
			count =count*10 +int(s[i]-'0')
		}else  if s[i]=='['{ // [ 后面肯定是字母
			// 可以把当前的数字放入次数数组
			countStack = append(countStack,count)
			count =0

			alphaIndex = append(alphaIndex,len(alphaStack))
		}else if s[i] ==']'{ // 对当前字母段出站
			// 获取当前字母段的次数 然后次数出栈
			c :=countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			// 获取当前字母段开始的索引 然后索引出栈
			index :=alphaIndex[len(alphaIndex)-1]
			alphaIndex = alphaIndex[:len(alphaIndex)-1]

			//获取当前的字母段
			str := string(alphaStack[index:])
			// 字母数组出站
			alphaStack = alphaStack[:index]

			for j:=0;j<c;j++{
				alphaStack = append(alphaStack,[]byte(str)...)
			}
		}else {
			alphaStack = append(alphaStack,s[i])
		}
	}
	return string(alphaStack)
}



func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("100[leetcode]"))
	fmt.Println(decodeString("3[a]2[b4[F]c]"))
}
