package main

import "fmt"

func distributeCandies(candies int, num_people int) []int {
	ans :=make([]int,num_people)
	count :=0
	index :=0
	for candies>0{
		count++
		i := index%num_people
		if candies>=count{
			ans[i] +=count
		}else {
			ans[i] +=candies
		}
		index++
		candies-=count
	}
	return ans
}


func main() {
	fmt.Println(distributeCandies(7,4))
	fmt.Println(distributeCandies(10,3))
}
