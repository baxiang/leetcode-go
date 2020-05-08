package main

import "fmt"

func merge(A []int, m int, B []int, n int)  {
       AIndex :=m-1
       BIndex :=n-1
       idx := len(A)-1
       for AIndex>=0&&BIndex>=0{
       	   if A[AIndex]>B[BIndex]{
       	   	  A[idx] = A[AIndex]
       	   	  AIndex--
		   }else if A[AIndex]<=B[BIndex]{
		   	 A[idx] = B[BIndex]
		   	 BIndex--
		   }
		   idx--
	   }
	   for BIndex>=0{
		   A[idx] = B[BIndex]
		   BIndex--
		   idx--
	   }
}

func main() {
	s :=[]int{1,2,3,0,0,0}
	merge(s,3,[]int{2,5,6},3)
	fmt.Println(s)

	s1 :=[]int{0}
	merge(s1,0,[]int{1},1)
	fmt.Println(s1)

	s2 :=[]int{4,5,6,0,0,0}
	merge(s2,3,[]int{1,2,3},3)
	fmt.Println(s2)
}
