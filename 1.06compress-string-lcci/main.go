package main

import (
	"fmt"
	"strings"
)

func compressString(S string) string {
	if len(S) == 0 {
		return S
	}
	preByte := S[0]
	count := 1
	var res strings.Builder
	for i := 1; i < len(S); i++ {
		if preByte == S[i] {
			count++
		} else {
			res.WriteByte(preByte)
			res.WriteString(fmt.Sprintf("%d", count))
			preByte = S[i]
			count = 1
		}
	}

	res.WriteByte(preByte)
	res.WriteString(fmt.Sprintf("%d", count))
	str := res.String()
	if len(str) >= len(S) {
		return S
	}
	return str
}

// 面试题 1.06 字符串压缩
func main() {
	fmt.Println(compressString("aabcccccaaa"))
	fmt.Println(compressString("abbccd"))
	fmt.Println(compressString("bb"))
}
