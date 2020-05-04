package main

import "fmt"

func sortArrayByParityII(A []int) []int {
	   j :=1
	  for i:=0;i<len(A);i=i+2{
		  if (A[i]&1)!=0 {// 偶数位查找不是偶数的
			  for A[j]%2!=0{// 奇数查找不是奇数的
				  j=j+2
			  }
			  A[i],A[j]=A[j],A[i]
		  }
	  }
	  return A
}


func main() {
	fmt.Println(sortArrayByParityII([]int{2,3,1,1,4,0,0,4,3,3}))
	fmt.Println(sortArrayByParityII([]int{4,2,5,7}))
}
