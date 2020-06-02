package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	var anchor = 0
	var write = 0
	for read := 0; read < len(chars); read++ {
		if read+1 == len(chars) || chars[read+1] != chars[read] {
			chars[write] = chars[anchor]
			write++
			if read > anchor {
				char := strconv.Itoa(read - anchor + 1)
				for i := 0; i < len(char); i++ {
					chars[write] = char[i]
					write++
				}
			}
			anchor = read + 1
		}
	}
	return write
}

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'a', 'b', 'b'}))
	fmt.Println(compress([]byte{'a','a','b','b','c','c','c'}))
	fmt.Println(compress([]byte{'a'}))
	fmt.Println(compress([]byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'}))
}
