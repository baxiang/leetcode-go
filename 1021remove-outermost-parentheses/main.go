package main

import (
	"fmt"
	"strings"
)

func removeOuterParentheses(S string) string {
	  if len(S)==0{
	  	return ""
	  }
	  stack := []byte{S[0]}
	  start :=1
	  var res strings.Builder
      for i:=1;i<len(S);i++{
		  if len(stack)>0&&S[i]==')'&&stack[len(stack)-1]=='(' {
		  	  stack = stack[:len(stack)-1]
		  	  if len(stack)==0{
		  	  	  res.WriteString(S[start:i])
				  start = i+2
			  }
		  }else {
		  	 stack = append(stack,S[i])
		  }
	  }
	  return res.String()
}

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println(removeOuterParentheses("()()"))
}
