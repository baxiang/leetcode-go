package main

import "fmt"

func sortArrayByParity(A []int) []int {
    l:=0
    r :=len(A)-1
    for l<r{
    	if A[l]%2!=0&&A[r]%2==0{
    		A[l],A[r] = A[r],A[l]
			l++
			r--
		}else if A[l]%2==0{
			l++
		}else {
			r--
		}
	}
	return A
}
func main() {
	fmt.Println(sortArrayByParity([]int{3,1,2,4}))
	fmt.Println(sortArrayByParity([]int{2,3,1,4,5}))
}
