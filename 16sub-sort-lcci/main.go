package main

import (
	"fmt"
	"sort"
)

func subSort(array []int) []int {
     tmp := make([]int,len(array))
     copy(tmp,array)
     sort.Slice(tmp,func(i ,j int)bool{
		 return tmp[i]<tmp[j]
	 })
     start :=-1
     end := -1
     for i :=0;i<len(array);i++{
		 if array[i]!=tmp[i] {
			 if start<0 {
				start= i
			 }
			 if end<i {
				 end=i
			 }
		 }
	 }
     return []int{start,end}
}
func main() {
	fmt.Println(subSort([]int{1,2,4,7,10,11,7,12,6,7,16,18,19}))
}
