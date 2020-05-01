package main

import "fmt"

func isValid(s string) bool {
    m :=map[string]string{")":"(","}":"{","]":"["}
    stack := make([]string,0)
    for i :=range s{
		x := string(s[i])
		if x=="("||x=="{"||x=="[" {
			stack = append(stack,x)
		}else if x==")" || x=="}" || x=="]" {
			if len(stack)>0&&stack[len(stack)-1]==m[x] {
				stack = stack[:len(stack)-1]
			}else {
				return false
			}

		}
	}
	return len(stack)==0
}

func main() {
	//fmt.Println(isValid("()"))
	//fmt.Println(isValid("([)]"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("]"))
}
