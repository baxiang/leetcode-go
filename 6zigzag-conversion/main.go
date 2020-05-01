package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows==1||numRows>=len(s) {
		return s
	}
	lookup := map[int][]byte{}
	for i:=0;i<numRows;i++{
		lookup[i]=[]byte{}
	}
	idx :=0
	step :=1
	for i:=0;i<len(s);i++{
		lookup[idx] = append(lookup[idx],s[i])
		if idx==0  {
			step =1
		}else if idx == numRows-1 {
			step =-1
		}
		idx+=step
	}
	var r []byte
	for i:=0;i<numRows;i++{
       r = append(r,lookup[i]...)
	}
   return string(r)
}

func main() {
	fmt.Println(convert("LEETCODEISHIRING",3))
}
