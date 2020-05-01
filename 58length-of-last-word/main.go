package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
     s =strings.TrimSpace(s)
	i:=len(s)-1
    for ;i>=0;i--{
		if string(s[i]) == " "{
			break
		}
	}
	return len(s)-i-1
}

func main() {
	fmt.Println(lengthOfLastWord("a "))
	fmt.Println(lengthOfLastWord("Hello World"))
}
