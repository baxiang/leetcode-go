package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
    stack :=make([]int,0)
    for i:=0;i<len(tokens);i++{
		if tokens[i]=="+" {
			f :=stack[len(stack)-1]
			s:= stack[len(stack)-2]
			stack =stack[:len(stack)-2]
			stack = append(stack,f+s)
		}else if tokens[i]=="-" {
			f :=stack[len(stack)-1]
			s:= stack[len(stack)-2]
			stack =stack[:len(stack)-2]
			stack = append(stack,s-f)
		}else if tokens[i]=="*" {
			f :=stack[len(stack)-1]
			s:= stack[len(stack)-2]
			stack =stack[:len(stack)-2]
			stack = append(stack,s*f)
		}else if tokens[i]=="/" {
			f :=stack[len(stack)-1]
			s:= stack[len(stack)-2]
			stack =stack[:len(stack)-2]
			stack = append(stack,s/f)
		}else {
		    s,_:= strconv.Atoi(tokens[i])
			stack = append(stack,s)
		}
	}
    return stack[0]
}


func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
