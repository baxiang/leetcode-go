package main

import "fmt"

func findComplement1(num int) int {
	i := 1
	for i <= num {
		i = i<<1
	}
	return i - num - 1
}

func findComplement(num int) int {
	tmp := num
	n :=1
	for tmp>0 {
		n=n<<1
		tmp=tmp>>1
	}
	n=n-1
	return num^n
}

func main() {
	fmt.Println(findComplement(5))
}
