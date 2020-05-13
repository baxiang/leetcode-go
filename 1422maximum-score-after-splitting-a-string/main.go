package main

import "fmt"

func maxScore(s string) int {
	var max =0
    for i:=1;i<len(s);i++{
    	z := number(s[:i],'0')
    	o := number(s[i:],'1')
    	if max<z+o{
    		max = z+o
		}
	}
	return max

}

func number(s string,t byte)int{
	var res int
	for i:=range s{
		if s[i] == t{
			res++
		}
	}
	return res
}


func main() {
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("00111"))
	fmt.Println(maxScore("1111"))
}
