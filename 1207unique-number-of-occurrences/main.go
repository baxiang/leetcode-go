package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
    m :=map[int]int{}
    for _,v :=range arr{
    	m[v]++
	}
    f := map[int]bool{}
    for _,v :=range m{
    	if f[v]{
    		return false
		}else {
			f[v] = true
		}
	}
    return true
}


func main() {
	fmt.Println(uniqueOccurrences([]int{1,2,2,1,1,3}))
	fmt.Println(uniqueOccurrences([]int{1,2}))
	fmt.Println(uniqueOccurrences([]int{-3,0,1,-3,1,1,1,-3,10,0}))
}
