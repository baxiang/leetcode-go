package main

import (
	"fmt"
)

func countSegments(s string) int {
	segmentCount  := 0
	for  i := 0; i < len(s); i++ {
		if ((i == 0 || s[i-1] == ' ') && s[i] != ' ') {
			segmentCount++
		}
	}
	return segmentCount
}

func main() {
	fmt.Println(countSegments("Hello, my name is John"))
	fmt.Println(countSegments("                "))
}
