package main

import "fmt"

func add(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}
	return a
}

func main() {
	fmt.Println(add(1,1))
	fmt.Println(add(2,3))
}
