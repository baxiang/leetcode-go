package main

import (
	"fmt"
)

func exchangeBits1(num int) int {
	//0xaaaaaaaa 10101010101010101010101010101010 (偶数位为1，奇数位为0）
	 even := (num & 0xaaaaaaaa) >> 1 // 获取当前数字的偶数位，然后向右移动一位
	 //0x55555555  1010101010101010101010101010101 (偶数位为0，奇数位为1）
	 odd := (num & 0x55555555) << 1 // 获取当前位数的奇数 然后向左移动一位
	return even + odd
}


func exchangeBits(num int) int {
	//0xaaaaaaaa 10101010101010101010101010101010 (偶数位为1，奇数位为0）
	even := (num & 0b10101010_10101010_10101010_10101010) >> 1 // 获取当前数字的偶数位，然后向右移动一位
	//0x55555555  1010101010101010101010101010101 (偶数位为0，奇数位为1）
	odd := (num & 0b01010101_01010101_01010101_01010101) << 1 // 获取当前位数的奇数 然后向左移动一位
	return even | odd
}


func main() {
	fmt.Println(exchangeBits(1))
	fmt.Println(exchangeBits(2))
	fmt.Println(exchangeBits(3))
	fmt.Println(exchangeBits(4))
}
