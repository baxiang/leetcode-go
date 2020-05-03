package main

import "fmt"

func isDividingNumbers(i int)bool{
	if i<10{
		return true
	}
	v:=i
	for i!=0{
		r :=i%10
		if r==0||v%r!=0{
			return false
		}
		i =i/10
	}
	return true
}
func selfDividingNumbers(left int, right int) []int {
	 var res []int
     for i:=left;i<=right;i++{
		 if isDividingNumbers(i){
		 	res = append(res,i)
		 }
	 }
	 return res
}

func main() {
	fmt.Println(selfDividingNumbers(1,22))
}
