package main

import "fmt"

func game(guess []int, answer []int) int {
	var res int
     for i,v:=range guess{
     	 if answer[i]==v{
     	 	res++
		 }
	 }
	 return res
}

func main() {
	fmt.Println(game([]int{1,2,3},[]int{1,2,3}))
}
