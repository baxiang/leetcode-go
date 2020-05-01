package main

import (
	"fmt"
)

func letterCombinations1(digits string) []string {
	if digits == "" || len(digits) == 0 {
		return nil
	}
	nums :=[]string{"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}
	list :=[]string{""}// 默认为空字符
	for i:=0;i<len(digits);i++{
		s :=nums[digits[i]-'0']
		size :=len(list)// 原先的数组长度记录 涉及到底下的入队和出队
		for j:=0;j<size;j++{
			first := list[0]
			list = list[1:]
			for k:=0;k<len(s);k++{
				list = append(list,first+string(s[k]))
			}
		}
	}
	return list
}
func letterCombinations(digits string) []string {
	if digits == "" || len(digits) == 0 {
		return nil
	}
	nums :=[]string{"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}
	var res []string
	res = append(res,"")
	for i :=range digits{
		 s := nums[digits[i]-'0']
		 var temp []string
		 for i:=0;i<len(s);i++{
		 	for j:=0;j<len(res);j++{
		 		temp = append(temp,res[j]+string(s[i]))

			}
		 }
		 res = temp
	}
	return res
}


func main() {
	fmt.Println(letterCombinations("234"))
}
