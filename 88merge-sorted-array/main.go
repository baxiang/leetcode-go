package main

import "fmt"

func merge1(nums1 []int, m int, nums2 []int, n int)  {
     n1:= m-1
     n2 :=n-1
     l := m+n-1
     for n1>=0&&n2>=0{
     	if nums1[n1]<=nums2[n2]{
     		nums1[l]= nums2[n2]
     		n2--
		}else {
			nums1[l]= nums1[n1]
			n1--
		}
		l--
	 }
	 for n2>=0&&l>=0{
		 nums1[l]= nums2[n2]
		 n2--
		 l--
	 }
}


func merge(nums1 []int, m int, nums2 []int, n int)  {
	n1:= m-1
	n2 :=n-1
	l := m+n-1
	for n2>=0{
		if n1>=0&&nums1[n1]>nums2[n2]{
			nums1[l]= nums1[n1]
			n1--
		}else {
			nums1[l]= nums2[n2]
			n2--
		}
		l--
	}
}

func main() {
	num1 := []int{1,2,3,0,0,0}
	merge(num1,3,[]int{2,5,6},3)
	fmt.Println(num1)
	num2 := []int{0}
	merge(num2,0,[]int{1},1)
	fmt.Println(num2)

	num3 := []int{4,5,6,0,0,0}
	merge(num3,3,[]int{1,2,3},3)
	fmt.Println(num3)
}
