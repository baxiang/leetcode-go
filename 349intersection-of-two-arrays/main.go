package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
      m :=make(map[int]struct{})
      for i:=range nums1{
      	if _,ok:=m[nums1[i]];!ok{
			m[nums1[i]]= struct{}{}
		}
	  }
	  var res []int
	  hasMap :=make(map[int]struct{})
	  for _,v:=range nums2{
		  _,ok:=m[v]
		 _,hasOK:= hasMap[v]
	  	 if ok&&!hasOK{
	  	 	res =append(res,v)
	  	 	hasMap[v]= struct{}{}
		 }
	  }
	  return res
}

func main() {
	fmt.Println(intersection([]int{1,2,2,1},[]int{2,2}))
	fmt.Println(intersection([]int{4,9,5},[]int{9,4,9,8,4}))
}
