package main

import "fmt"

func isUnique(astr string) bool {
      m := make([]byte,128)
      for i :=range astr{
      	  if m[astr[i]]>0{
      	  	return false
		  }
		  m[astr[i]]++
	  }
	  return true
}

//面试题 01.01. 判定字符是否唯一 https://leetcode-cn.com/problems/is-unique-lcci/
func main() {
	fmt.Println(isUnique("leetcode"))
	fmt.Println(isUnique("abc"))
}
