package main

import "fmt"

//2n
func removeElement(nums []int, val int) int {
	var l []int
    for _,v :=range nums{
		if val!=v{
			l = append(l,v)
		}
	}
	for i,v:=range l{
		nums[i]=v
	}
	return len(l)
}

func removeElement2(nums []int, val int) int {
	var index = 0
	for _,v :=range nums{
		if val!=v{
			nums[index]= v
			index++
		}
	}
	return index
}


func removeElement3(nums []int, val int) int {
	slow := 0
	quick :=0
	for quick<len(nums){
		if nums[quick]!=val {
			nums[slow]=nums[quick]
			slow++
		}
		quick++
	}
	return slow
}


func main() {
	fmt.Println(removeElement3([]int{3,2,2,3},3))//2
	fmt.Println(removeElement3([]int{0,1,2,2,3,0,4,2},2))//5
	fmt.Println(removeElement3([]int{2},3))//1
	fmt.Println(removeElement3([]int{2},2))//0
}
